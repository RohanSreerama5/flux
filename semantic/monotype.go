package semantic

import (
	"fmt"
	"strings"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/influxdata/flux/codes"
	"github.com/influxdata/flux/internal/errors"
	"github.com/influxdata/flux/semantic/internal/fbsemantic"
)

type fbTabler interface {
	Init(buf []byte, i flatbuffers.UOffsetT)
	Table() flatbuffers.Table
}

// MonoType represents a monotype.  This struct is a thin wrapper around
// Go code generated by the FlatBuffers compiler.
type MonoType struct {
	mt  fbsemantic.MonoType
	tbl fbTabler
}

// NewMonoType constructs a new monotype from a FlatBuffers table and the given kind of monotype.
func NewMonoType(tbl flatbuffers.Table, t fbsemantic.MonoType) (MonoType, error) {
	var tbler fbTabler
	switch t {
	case fbsemantic.MonoTypeNONE:
		return MonoType{}, errors.Newf(codes.Internal, "missing type, got type: %v", fbsemantic.EnumNamesMonoType[t])
	case fbsemantic.MonoTypeBasic:
		tbler = new(fbsemantic.Basic)
	case fbsemantic.MonoTypeVar:
		tbler = new(fbsemantic.Var)
	case fbsemantic.MonoTypeArr:
		tbler = new(fbsemantic.Arr)
	case fbsemantic.MonoTypeRow:
		tbler = new(fbsemantic.Row)
	case fbsemantic.MonoTypeFun:
		tbler = new(fbsemantic.Fun)
	default:
		return MonoType{}, errors.Newf(codes.Internal, "unknown type (%v)", t)
	}
	tbler.Init(tbl.Bytes, tbl.Pos)
	return MonoType{mt: t, tbl: tbler}, nil
}

func (mt MonoType) Nature() Nature {
	switch mt.mt {
	case fbsemantic.MonoTypeBasic:
		t, _ := mt.Basic()
		switch t {
		case fbsemantic.TypeBool:
			return Bool
		case fbsemantic.TypeInt:
			return Int
		case fbsemantic.TypeUint:
			return UInt
		case fbsemantic.TypeFloat:
			return Float
		case fbsemantic.TypeString:
			return String
		case fbsemantic.TypeDuration:
			return Duration
		case fbsemantic.TypeTime:
			return Time
		case fbsemantic.TypeRegexp:
			return Regexp
		case fbsemantic.TypeBytes:
			return Bytes
		default:
			return Invalid
		}
	case fbsemantic.MonoTypeArr:
		return Array
	case fbsemantic.MonoTypeRow:
		return Object
	case fbsemantic.MonoTypeFun:
		return Function
	case fbsemantic.MonoTypeNONE,
		fbsemantic.MonoTypeVar:
		fallthrough
	default:
		return Invalid
	}
}

// Kind specifies a particular kind of monotype.
type Kind fbsemantic.MonoType

const (
	Unknown = Kind(fbsemantic.MonoTypeNONE)
	Basic   = Kind(fbsemantic.MonoTypeBasic)
	Var     = Kind(fbsemantic.MonoTypeVar)
	Arr     = Kind(fbsemantic.MonoTypeArr)
	Row     = Kind(fbsemantic.MonoTypeRow)
	Fun     = Kind(fbsemantic.MonoTypeFun)
)

// Kind returns what kind of monotype the receiver is.
func (mt MonoType) Kind() Kind {
	return Kind(mt.mt)
}

var (
	BasicBool     MonoType
	BasicInt      MonoType
	BasicUint     MonoType
	BasicFloat    MonoType
	BasicString   MonoType
	BasicDuration MonoType
	BasicTime     MonoType
	BasicRegexp   MonoType
	BasicBytes    MonoType
)

func init() {
	builder := flatbuffers.NewBuilder(1024)
	fbsemantic.BasicStart(builder)
	fbsemantic.BasicAddT(builder, fbsemantic.TypeBool)
	basicBool := fbsemantic.BasicEnd(builder)

	// TODO (algow): initial all basic values

	// TODO (algow): this probably doesn't work...
	builder.Finish(basicBool)
	basicTable := flatbuffers.Table{
		Bytes: builder.FinishedBytes(),
		Pos:   basicBool,
	}
	BasicBool, _ = NewMonoType(basicTable, fbsemantic.MonoTypeBasic)
}

func getBasic(tbl fbTabler) (*fbsemantic.Basic, error) {
	b, ok := tbl.(*fbsemantic.Basic)
	if !ok {
		return nil, errors.New(codes.Internal, "MonoType is not a basic type")
	}
	return b, nil
}

// Basic returns the basic type for this monotype if it is a basic type,
// and an error otherwise.
func (mt MonoType) Basic() (fbsemantic.Type, error) {
	b, err := getBasic(mt.tbl)
	if err != nil {
		return fbsemantic.TypeBool, err
	}
	return b.T(), nil
}

func getVar(tbl fbTabler) (*fbsemantic.Var, error) {
	v, ok := tbl.(*fbsemantic.Var)
	if !ok {
		return nil, errors.New(codes.Internal, "MonoType is not a type var")
	}
	return v, nil

}

// VarNum returns the type variable number if this monotype is a type variable,
// and an error otherwise.
func (mt MonoType) VarNum() (uint64, error) {
	v, err := getVar(mt.tbl)
	if err != nil {
		return 0, err
	}
	return v.I(), nil
}

func monoTypeFromVar(v *fbsemantic.Var) MonoType {
	return MonoType{
		mt:  fbsemantic.MonoTypeVar,
		tbl: v,
	}
}

func getFun(tbl fbTabler) (*fbsemantic.Fun, error) {
	f, ok := tbl.(*fbsemantic.Fun)
	if !ok {
		return nil, errors.New(codes.Internal, "MonoType is not a function")
	}
	return f, nil
}

// NumArguments returns the number of arguments if this monotype is a function,
// and an error otherwise.
func (mt MonoType) NumArguments() (int, error) {
	f, err := getFun(mt.tbl)
	if err != nil {
		return 0, err
	}
	return f.ArgsLength(), nil
}

// Argument returns the argument give an ordinal position if this monotype is a function,
// and an error otherwise.
func (mt MonoType) Argument(i int) (*Argument, error) {
	f, err := getFun(mt.tbl)
	if err != nil {
		return nil, err
	}
	if i < 0 || i >= f.ArgsLength() {
		return nil, errors.Newf(codes.Internal, "request for out-of-bounds argument: %v of %v", i, f.ArgsLength())
	}
	a := new(fbsemantic.Argument)
	if !f.Args(a, i) {
		return nil, errors.New(codes.Internal, "missing argument")
	}
	return newArgument(a)
}

func (mt MonoType) ReturnType() (MonoType, error) {
	f, ok := mt.tbl.(*fbsemantic.Fun)
	if !ok {
		return MonoType{}, errors.New(codes.Internal, "ReturnType() called on non-function MonoType")
	}
	var tbl flatbuffers.Table
	if !f.Retn(&tbl) {
		return MonoType{}, errors.New(codes.Internal, "missing return type")
	}
	return NewMonoType(tbl, f.RetnType())
}

func getArr(tbl fbTabler) (*fbsemantic.Arr, error) {
	arr, ok := tbl.(*fbsemantic.Arr)
	if !ok {
		return nil, errors.New(codes.Internal, "MonoType is not an array")
	}
	return arr, nil
}

// ElemType returns the element type if this monotype is an array, and an error otherise.
func (mt MonoType) ElemType() (MonoType, error) {
	arr, err := getArr(mt.tbl)
	if err != nil {
		return MonoType{}, err
	}
	var tbl flatbuffers.Table
	if !arr.T(&tbl) {
		return MonoType{}, errors.New(codes.Internal, "missing array type")
	}
	return NewMonoType(tbl, arr.TType())
}

func getRow(tbl fbTabler) (*fbsemantic.Row, error) {
	row, ok := tbl.(*fbsemantic.Row)
	if !ok {
		return nil, errors.New(codes.Internal, "MonoType is not a row")
	}
	return row, nil

}

// NumProperties returns the number of properties if this monotype is a row, and an error otherwise.
func (mt MonoType) NumProperties() (int, error) {
	row, err := getRow(mt.tbl)
	if err != nil {
		return 0, err
	}
	return row.PropsLength(), nil
}

// RowProperty returns a property given its ordinal position if this monotype is a row, and an error otherwise.
func (mt MonoType) RowProperty(i int) (*RowProperty, error) {
	row, err := getRow(mt.tbl)
	if err != nil {
		return nil, err
	}
	if i < 0 || i >= row.PropsLength() {
		return nil, errors.Newf(codes.Internal, "request for out-of-bounds property: %v of %v", i, row.PropsLength())
	}
	p := new(fbsemantic.Prop)
	if !row.Props(p, i) {
		return nil, errors.New(codes.Internal, "missing property")
	}
	return &RowProperty{fb: p}, nil
}

// Extends returns the extending type variable if this monotype is a row, and an error otherwise.
// If the type is a row but does not extend anything a false is returned.
func (mt MonoType) Extends() (MonoType, bool, error) {
	row, err := getRow(mt.tbl)
	if err != nil {
		return MonoType{}, false, err
	}
	v := row.Extends(nil)
	if v == nil {
		return MonoType{}, false, nil
	}
	return monoTypeFromVar(v), true, nil
}

// Argument represents a function argument.
type Argument struct {
	*fbsemantic.Argument
}

func newArgument(fb *fbsemantic.Argument) (*Argument, error) {
	if fb == nil {
		return nil, errors.Newf(codes.Internal, "nil argument")
	}
	return &Argument{Argument: fb}, nil
}

// TypeOf returns the type of the function argument.
func (a *Argument) TypeOf() (MonoType, error) {
	var tbl flatbuffers.Table
	if !a.T(&tbl) {
		return MonoType{}, errors.New(codes.Internal, "missing argument type")
	}
	argTy, err := NewMonoType(tbl, a.TType())
	if err != nil {
		return MonoType{}, err
	}
	return argTy, nil
}

// Property represents a property of a row.
type RowProperty struct {
	fb *fbsemantic.Prop
}

// Name returns the name of the property.
func (p *RowProperty) Name() string {
	return string(p.fb.K())
}

// TypeOf returns the type of the property.
func (p *RowProperty) TypeOf() (MonoType, error) {
	var tbl flatbuffers.Table
	if !p.fb.V(&tbl) {
		return MonoType{}, errors.Newf(codes.Internal, "missing property type")
	}
	return NewMonoType(tbl, p.fb.VType())
}

// String returns a string representation of this monotype.
func (mt MonoType) String() string {
	switch tk := mt.Kind(); tk {
	case Unknown:
		return "<" + fbsemantic.EnumNamesMonoType[fbsemantic.MonoType(tk)] + ">"
	case Basic:
		b, err := mt.Basic()
		if err != nil {
			return "<" + err.Error() + ">"
		}
		return strings.ToLower(fbsemantic.EnumNamesType[byte(b)])
	case Var:
		i, err := mt.VarNum()
		if err != nil {
			return "<" + err.Error() + ">"
		}
		return fmt.Sprintf("t%d", i)
	case Arr:
		et, err := mt.ElemType()
		if err != nil {
			return "<" + err.Error() + ">"
		}
		return "[" + et.String() + "]"
	case Row:
		var sb strings.Builder
		sb.WriteString("{")
		nprops, err := mt.NumProperties()
		if err != nil {
			return "<" + err.Error() + ">"
		}
		needBar := false
		for i := 0; i < nprops; i++ {
			if needBar {
				sb.WriteString(" | ")
			} else {
				needBar = true
			}
			prop, err := mt.RowProperty(i)
			if err != nil {
				return "<" + err.Error() + ">"
			}
			sb.WriteString(prop.Name() + ": ")
			ty, err := prop.TypeOf()
			if err != nil {
				return "<" + err.Error() + ">"
			}
			sb.WriteString(ty.String())
		}
		extends, ok, err := mt.Extends()
		if err != nil {
			return "<" + err.Error() + ">"
		}
		if ok {
			if needBar {
				sb.WriteString(" | ")
			}
			sb.WriteString(extends.String())
		}
		sb.WriteString("}")
		return sb.String()
	case Fun:
		var sb strings.Builder
		sb.WriteString("(")
		needComma := false
		nargs, err := mt.NumArguments()
		if err != nil {
			return "<" + err.Error() + ">"
		}
		for i := 0; i < nargs; i++ {
			arg, err := mt.Argument(i)
			if err != nil {
				return "<" + err.Error() + ">"
			}
			if needComma {
				sb.WriteString(", ")
			} else {
				needComma = true
			}
			if arg.Optional() {
				sb.WriteString("?")
			} else if arg.Pipe() {
				sb.WriteString("<-")
			}
			sb.WriteString(string(arg.Name()) + ": ")
			argTyp, err := arg.TypeOf()
			if err != nil {
				return "<" + err.Error() + ">"
			}
			sb.WriteString(argTyp.String())
		}
		sb.WriteString(") -> ")
		rt, err := mt.ReturnType()
		if err != nil {
			return "<" + err.Error() + ">"
		}
		sb.WriteString(rt.String())
		return sb.String()
	default:
		return "<" + fmt.Sprintf("unknown monotype (%v)", tk) + ">"
	}
}

func (l MonoType) Equal(r MonoType) bool {
	// TODO (algow): Remove this method, we will not support comparing types for equality.
	return false
}

func NewArrayType(elemType MonoType) MonoType {
	//builder := flatbuffers.NewBuilder(1024)
	//fbsemantic.ArrStart(builder)
	//fbsemantic.AddTType(builder, elemType.Type())
	//// TODO (algow): how do we inject the elemType?
	//fbsemantic.AddT(builder, flatbuffer.UOffsetT)

	return MonoType{}
}
func NewFunctionType() PolyType {
	// TODO (algow): needs both a list of vars constraints and the monotype
	return PolyType{}
}
func NewObjectType() MonoType {
	// TODO (algow): needs both a list of vars constraints and the monotype
	return MonoType{}
}
