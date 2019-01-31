package langserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/url"

	"github.com/influxdata/flux/complete"
	"github.com/sourcegraph/go-lsp"
	"github.com/sourcegraph/jsonrpc2"
	"go.uber.org/zap"
)

type Server struct {
	handler  *Handler
	logger   *zap.Logger
	listener net.Listener
	closed   bool
	shutdown bool

	workspace string
}

func New(h Handler, l *zap.Logger) *Server {
	return &Server{
		handler: &h,
		logger:  l,
	}
}

func (s *Server) Serve(l net.Listener) error {
	if s.listener != nil {
		return errors.New("already listening")
	}
	s.listener = l
	for {
		conn, err := l.Accept()
		if err != nil {
			if s.closed {
				return nil
			}
			return err
		}
		go s.serve(conn)
	}
}

func (s *Server) serve(rw io.ReadWriteCloser) error {
	stream := jsonrpc2.NewBufferedStream(rw, jsonrpc2.VSCodeObjectCodec{})
	handler := jsonrpc2.HandlerWithError(s.handle)
	conn := jsonrpc2.NewConn(context.TODO(), stream, handler)
	<-conn.DisconnectNotify()
	return nil
}

func (s *Server) handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (result interface{}, err error) {
	if s.shutdown && req.Method != "exit" {
		return nil, &jsonrpc2.Error{
			Code:    jsonrpc2.CodeInvalidRequest,
			Message: "server is shutdown",
		}
	}
	logger := s.logger
	if req.ID.IsString {
		logger = logger.With(zap.String("id", req.ID.Str))
	} else {
		logger = logger.With(zap.Uint64("id", req.ID.Num))
	}
	logger = logger.With(zap.String("method", req.Method))

	switch req.Method {
	case "initialize":
		// TODO(jsternberg): Keep track if the server was already initialized.
		if req.Params == nil {
			return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
		}
		var params lsp.InitializeParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, err
		}

		logger.Info("Initialize", zap.Int("processId", params.ProcessID), zap.String("path", string(params.RootURI)))
		if err := s.reset(params); err != nil {
			return nil, err
		}
		return lsp.InitializeResult{
			Capabilities: lsp.ServerCapabilities{
				CompletionProvider: &lsp.CompletionOptions{
					TriggerCharacters: []string{"."},
				},
			},
		}, nil
	case "shutdown":
		s.shutdown = true
		return nil, nil
	case "exit":
		s.closed = true
		if err := s.listener.Close(); err != nil {
			return nil, err
		}
		return nil, nil
	case "textDocument/completion":
		if req.Params == nil {
			return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
		}
		var params lsp.CompletionParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, err
		}
		return s.completions(params.TextDocument.URI)
	}

	return nil, &jsonrpc2.Error{
		Code:    jsonrpc2.CodeMethodNotFound,
		Message: fmt.Sprintf("method not supported: %s", req.Method),
	}
}

func (s *Server) reset(params lsp.InitializeParams) error {
	if params.RootURI != "" {
		u, err := url.Parse(string(params.RootURI))
		if err != nil {
			return err
		}

		if u.Scheme != "file" {
			return fmt.Errorf("invalid uri scheme: %s", u.Scheme)
		}
		s.workspace = u.Path
	}
	return nil
}

func (s *Server) completions(uri lsp.DocumentURI) (lsp.CompletionList, error) {
	text, err := s.getText(uri)
	if err != nil {
		return lsp.CompletionList{}, err
	}
	list, err := complete.StaticComplete(text)
	if err != nil {
		return lsp.CompletionList{}, err
	}
	items := make([]lsp.CompletionItem, 0, len(list))
	for _, item := range list {
		items = append(items, lsp.CompletionItem{
			Label: item,
		})
	}
	return lsp.CompletionList{
		Items: items,
	}, nil
}

func (s *Server) getText(uri lsp.DocumentURI) (string, error) {
	u, err := url.Parse(string(uri))
	if err != nil {
		return "", err
	}

	if u.Scheme != "file" {
		return "", fmt.Errorf("invalid uri scheme: %s", u.Scheme)
	}

	data, err := ioutil.ReadFile(u.Path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}