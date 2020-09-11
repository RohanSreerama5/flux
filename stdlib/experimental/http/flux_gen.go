// DO NOT EDIT: This file is autogenerated via the builtin command.

package http

import (
	ast "github.com/influxdata/flux/ast"
	runtime "github.com/influxdata/flux/runtime"
)

func init() {
	runtime.RegisterPackage(pkgAST)
}

var pkgAST = &ast.Package{
	BaseNode: ast.BaseNode{
		Errors: nil,
		Loc:    nil,
	},
	Files: []*ast.File{&ast.File{
		BaseNode: ast.BaseNode{
			Errors: nil,
			Loc: &ast.SourceLocation{
				End: ast.Position{
					Column: 12,
					Line:   5,
				},
				File:   "http.flux",
				Source: "package http\n\n// Get submits an HTTP get request to the specified URL with headers\n// Returns HTTP status code and body as a byte array\nbuiltin get",
				Start: ast.Position{
					Column: 1,
					Line:   1,
				},
			},
		},
		Body: []ast.Statement{&ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 12,
						Line:   5,
					},
					File:   "http.flux",
					Source: "builtin get",
					Start: ast.Position{
						Column: 1,
						Line:   5,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 12,
							Line:   5,
						},
						File:   "http.flux",
						Source: "get",
						Start: ast.Position{
							Column: 9,
							Line:   5,
						},
					},
				},
				Name: "get",
			},
			Ty: ast.TypeExpression{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 136,
							Line:   5,
						},
						File:   "http.flux",
						Source: "(url: string, ?headers: A, ?timeout: duration) => {statusCode: int , body: bytes , headers: B} where A: Record, B: Record",
						Start: ast.Position{
							Column: 15,
							Line:   5,
						},
					},
				},
				Constraints: []*ast.TypeConstraint{&ast.TypeConstraint{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 125,
								Line:   5,
							},
							File:   "http.flux",
							Source: "A: Record",
							Start: ast.Position{
								Column: 116,
								Line:   5,
							},
						},
					},
					Kinds: []*ast.Identifier{&ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 125,
									Line:   5,
								},
								File:   "http.flux",
								Source: "Record",
								Start: ast.Position{
									Column: 119,
									Line:   5,
								},
							},
						},
						Name: "Record",
					}},
					Tvar: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 117,
									Line:   5,
								},
								File:   "http.flux",
								Source: "A",
								Start: ast.Position{
									Column: 116,
									Line:   5,
								},
							},
						},
						Name: "A",
					},
				}, &ast.TypeConstraint{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 136,
								Line:   5,
							},
							File:   "http.flux",
							Source: "B: Record",
							Start: ast.Position{
								Column: 127,
								Line:   5,
							},
						},
					},
					Kinds: []*ast.Identifier{&ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 136,
									Line:   5,
								},
								File:   "http.flux",
								Source: "Record",
								Start: ast.Position{
									Column: 130,
									Line:   5,
								},
							},
						},
						Name: "Record",
					}},
					Tvar: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 128,
									Line:   5,
								},
								File:   "http.flux",
								Source: "B",
								Start: ast.Position{
									Column: 127,
									Line:   5,
								},
							},
						},
						Name: "B",
					},
				}},
				Ty: &ast.FunctionType{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 109,
								Line:   5,
							},
							File:   "http.flux",
							Source: "(url: string, ?headers: A, ?timeout: duration) => {statusCode: int , body: bytes , headers: B}",
							Start: ast.Position{
								Column: 15,
								Line:   5,
							},
						},
					},
					Parameters: []*ast.ParameterType{&ast.ParameterType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 27,
									Line:   5,
								},
								File:   "http.flux",
								Source: "url: string",
								Start: ast.Position{
									Column: 16,
									Line:   5,
								},
							},
						},
						Kind: "Required",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 19,
										Line:   5,
									},
									File:   "http.flux",
									Source: "url",
									Start: ast.Position{
										Column: 16,
										Line:   5,
									},
								},
							},
							Name: "url",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 27,
										Line:   5,
									},
									File:   "http.flux",
									Source: "string",
									Start: ast.Position{
										Column: 21,
										Line:   5,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 27,
											Line:   5,
										},
										File:   "http.flux",
										Source: "string",
										Start: ast.Position{
											Column: 21,
											Line:   5,
										},
									},
								},
								Name: "string",
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 40,
									Line:   5,
								},
								File:   "http.flux",
								Source: "?headers: A",
								Start: ast.Position{
									Column: 29,
									Line:   5,
								},
							},
						},
						Kind: "Optional",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 37,
										Line:   5,
									},
									File:   "http.flux",
									Source: "headers",
									Start: ast.Position{
										Column: 30,
										Line:   5,
									},
								},
							},
							Name: "headers",
						},
						Ty: &ast.TvarType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 40,
										Line:   5,
									},
									File:   "http.flux",
									Source: "A",
									Start: ast.Position{
										Column: 39,
										Line:   5,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 40,
											Line:   5,
										},
										File:   "http.flux",
										Source: "A",
										Start: ast.Position{
											Column: 39,
											Line:   5,
										},
									},
								},
								Name: "A",
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 60,
									Line:   5,
								},
								File:   "http.flux",
								Source: "?timeout: duration",
								Start: ast.Position{
									Column: 42,
									Line:   5,
								},
							},
						},
						Kind: "Optional",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 50,
										Line:   5,
									},
									File:   "http.flux",
									Source: "timeout",
									Start: ast.Position{
										Column: 43,
										Line:   5,
									},
								},
							},
							Name: "timeout",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 60,
										Line:   5,
									},
									File:   "http.flux",
									Source: "duration",
									Start: ast.Position{
										Column: 52,
										Line:   5,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 60,
											Line:   5,
										},
										File:   "http.flux",
										Source: "duration",
										Start: ast.Position{
											Column: 52,
											Line:   5,
										},
									},
								},
								Name: "duration",
							},
						},
					}},
					Return: &ast.RecordType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 109,
									Line:   5,
								},
								File:   "http.flux",
								Source: "{statusCode: int , body: bytes , headers: B}",
								Start: ast.Position{
									Column: 65,
									Line:   5,
								},
							},
						},
						Properties: []*ast.PropertyType{&ast.PropertyType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 81,
										Line:   5,
									},
									File:   "http.flux",
									Source: "statusCode: int",
									Start: ast.Position{
										Column: 66,
										Line:   5,
									},
								},
							},
							Name: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 76,
											Line:   5,
										},
										File:   "http.flux",
										Source: "statusCode",
										Start: ast.Position{
											Column: 66,
											Line:   5,
										},
									},
								},
								Name: "statusCode",
							},
							Ty: &ast.NamedType{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 81,
											Line:   5,
										},
										File:   "http.flux",
										Source: "int",
										Start: ast.Position{
											Column: 78,
											Line:   5,
										},
									},
								},
								ID: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 81,
												Line:   5,
											},
											File:   "http.flux",
											Source: "int",
											Start: ast.Position{
												Column: 78,
												Line:   5,
											},
										},
									},
									Name: "int",
								},
							},
						}, &ast.PropertyType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 95,
										Line:   5,
									},
									File:   "http.flux",
									Source: "body: bytes",
									Start: ast.Position{
										Column: 84,
										Line:   5,
									},
								},
							},
							Name: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 88,
											Line:   5,
										},
										File:   "http.flux",
										Source: "body",
										Start: ast.Position{
											Column: 84,
											Line:   5,
										},
									},
								},
								Name: "body",
							},
							Ty: &ast.NamedType{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 95,
											Line:   5,
										},
										File:   "http.flux",
										Source: "bytes",
										Start: ast.Position{
											Column: 90,
											Line:   5,
										},
									},
								},
								ID: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 95,
												Line:   5,
											},
											File:   "http.flux",
											Source: "bytes",
											Start: ast.Position{
												Column: 90,
												Line:   5,
											},
										},
									},
									Name: "bytes",
								},
							},
						}, &ast.PropertyType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 108,
										Line:   5,
									},
									File:   "http.flux",
									Source: "headers: B",
									Start: ast.Position{
										Column: 98,
										Line:   5,
									},
								},
							},
							Name: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 105,
											Line:   5,
										},
										File:   "http.flux",
										Source: "headers",
										Start: ast.Position{
											Column: 98,
											Line:   5,
										},
									},
								},
								Name: "headers",
							},
							Ty: &ast.TvarType{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 108,
											Line:   5,
										},
										File:   "http.flux",
										Source: "B",
										Start: ast.Position{
											Column: 107,
											Line:   5,
										},
									},
								},
								ID: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 108,
												Line:   5,
											},
											File:   "http.flux",
											Source: "B",
											Start: ast.Position{
												Column: 107,
												Line:   5,
											},
										},
									},
									Name: "B",
								},
							},
						}},
						Tvar: nil,
					},
				},
			},
		}},
		Imports:  nil,
		Metadata: "parser-type=rust",
		Name:     "http.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 13,
						Line:   1,
					},
					File:   "http.flux",
					Source: "package http",
					Start: ast.Position{
						Column: 1,
						Line:   1,
					},
				},
			},
			Name: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 13,
							Line:   1,
						},
						File:   "http.flux",
						Source: "http",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "http",
			},
		},
	}},
	Package: "http",
	Path:    "experimental/http",
}
