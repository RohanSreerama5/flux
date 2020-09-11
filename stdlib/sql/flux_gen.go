// DO NOT EDIT: This file is autogenerated via the builtin command.

package sql

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
					Column: 11,
					Line:   4,
				},
				File:   "sql.flux",
				Source: "package sql\n\nbuiltin from : (driverName: string, dataSourceName: string, query: string) => [A]\nbuiltin to",
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
						Column: 13,
						Line:   3,
					},
					File:   "sql.flux",
					Source: "builtin from",
					Start: ast.Position{
						Column: 1,
						Line:   3,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 13,
							Line:   3,
						},
						File:   "sql.flux",
						Source: "from",
						Start: ast.Position{
							Column: 9,
							Line:   3,
						},
					},
				},
				Name: "from",
			},
			Ty: ast.TypeExpression{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 82,
							Line:   3,
						},
						File:   "sql.flux",
						Source: "(driverName: string, dataSourceName: string, query: string) => [A]",
						Start: ast.Position{
							Column: 16,
							Line:   3,
						},
					},
				},
				Constraints: []*ast.TypeConstraint{},
				Ty: &ast.FunctionType{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 82,
								Line:   3,
							},
							File:   "sql.flux",
							Source: "(driverName: string, dataSourceName: string, query: string) => [A]",
							Start: ast.Position{
								Column: 16,
								Line:   3,
							},
						},
					},
					Parameters: []*ast.ParameterType{&ast.ParameterType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 35,
									Line:   3,
								},
								File:   "sql.flux",
								Source: "driverName: string",
								Start: ast.Position{
									Column: 17,
									Line:   3,
								},
							},
						},
						Kind: "Required",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 27,
										Line:   3,
									},
									File:   "sql.flux",
									Source: "driverName",
									Start: ast.Position{
										Column: 17,
										Line:   3,
									},
								},
							},
							Name: "driverName",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 35,
										Line:   3,
									},
									File:   "sql.flux",
									Source: "string",
									Start: ast.Position{
										Column: 29,
										Line:   3,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 35,
											Line:   3,
										},
										File:   "sql.flux",
										Source: "string",
										Start: ast.Position{
											Column: 29,
											Line:   3,
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
									Column: 59,
									Line:   3,
								},
								File:   "sql.flux",
								Source: "dataSourceName: string",
								Start: ast.Position{
									Column: 37,
									Line:   3,
								},
							},
						},
						Kind: "Required",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 51,
										Line:   3,
									},
									File:   "sql.flux",
									Source: "dataSourceName",
									Start: ast.Position{
										Column: 37,
										Line:   3,
									},
								},
							},
							Name: "dataSourceName",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 59,
										Line:   3,
									},
									File:   "sql.flux",
									Source: "string",
									Start: ast.Position{
										Column: 53,
										Line:   3,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 59,
											Line:   3,
										},
										File:   "sql.flux",
										Source: "string",
										Start: ast.Position{
											Column: 53,
											Line:   3,
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
									Column: 74,
									Line:   3,
								},
								File:   "sql.flux",
								Source: "query: string",
								Start: ast.Position{
									Column: 61,
									Line:   3,
								},
							},
						},
						Kind: "Required",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 66,
										Line:   3,
									},
									File:   "sql.flux",
									Source: "query",
									Start: ast.Position{
										Column: 61,
										Line:   3,
									},
								},
							},
							Name: "query",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 74,
										Line:   3,
									},
									File:   "sql.flux",
									Source: "string",
									Start: ast.Position{
										Column: 68,
										Line:   3,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 74,
											Line:   3,
										},
										File:   "sql.flux",
										Source: "string",
										Start: ast.Position{
											Column: 68,
											Line:   3,
										},
									},
								},
								Name: "string",
							},
						},
					}},
					Return: &ast.ArrayType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 82,
									Line:   3,
								},
								File:   "sql.flux",
								Source: "[A]",
								Start: ast.Position{
									Column: 79,
									Line:   3,
								},
							},
						},
						ElementType: &ast.TvarType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 81,
										Line:   3,
									},
									File:   "sql.flux",
									Source: "A",
									Start: ast.Position{
										Column: 80,
										Line:   3,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 81,
											Line:   3,
										},
										File:   "sql.flux",
										Source: "A",
										Start: ast.Position{
											Column: 80,
											Line:   3,
										},
									},
								},
								Name: "A",
							},
						},
					},
				},
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 11,
						Line:   4,
					},
					File:   "sql.flux",
					Source: "builtin to",
					Start: ast.Position{
						Column: 1,
						Line:   4,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 11,
							Line:   4,
						},
						File:   "sql.flux",
						Source: "to",
						Start: ast.Position{
							Column: 9,
							Line:   4,
						},
					},
				},
				Name: "to",
			},
			Ty: ast.TypeExpression{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 112,
							Line:   4,
						},
						File:   "sql.flux",
						Source: "(<-tables: [A], driverName: string, dataSourceName: string, table: string, ?batchSize: int) => [A]",
						Start: ast.Position{
							Column: 14,
							Line:   4,
						},
					},
				},
				Constraints: []*ast.TypeConstraint{},
				Ty: &ast.FunctionType{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 112,
								Line:   4,
							},
							File:   "sql.flux",
							Source: "(<-tables: [A], driverName: string, dataSourceName: string, table: string, ?batchSize: int) => [A]",
							Start: ast.Position{
								Column: 14,
								Line:   4,
							},
						},
					},
					Parameters: []*ast.ParameterType{&ast.ParameterType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 28,
									Line:   4,
								},
								File:   "sql.flux",
								Source: "<-tables: [A]",
								Start: ast.Position{
									Column: 15,
									Line:   4,
								},
							},
						},
						Kind: "Pipe",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 23,
										Line:   4,
									},
									File:   "sql.flux",
									Source: "tables",
									Start: ast.Position{
										Column: 17,
										Line:   4,
									},
								},
							},
							Name: "tables",
						},
						Ty: &ast.ArrayType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 28,
										Line:   4,
									},
									File:   "sql.flux",
									Source: "[A]",
									Start: ast.Position{
										Column: 25,
										Line:   4,
									},
								},
							},
							ElementType: &ast.TvarType{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 27,
											Line:   4,
										},
										File:   "sql.flux",
										Source: "A",
										Start: ast.Position{
											Column: 26,
											Line:   4,
										},
									},
								},
								ID: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 27,
												Line:   4,
											},
											File:   "sql.flux",
											Source: "A",
											Start: ast.Position{
												Column: 26,
												Line:   4,
											},
										},
									},
									Name: "A",
								},
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 48,
									Line:   4,
								},
								File:   "sql.flux",
								Source: "driverName: string",
								Start: ast.Position{
									Column: 30,
									Line:   4,
								},
							},
						},
						Kind: "Required",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 40,
										Line:   4,
									},
									File:   "sql.flux",
									Source: "driverName",
									Start: ast.Position{
										Column: 30,
										Line:   4,
									},
								},
							},
							Name: "driverName",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 48,
										Line:   4,
									},
									File:   "sql.flux",
									Source: "string",
									Start: ast.Position{
										Column: 42,
										Line:   4,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 48,
											Line:   4,
										},
										File:   "sql.flux",
										Source: "string",
										Start: ast.Position{
											Column: 42,
											Line:   4,
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
									Column: 72,
									Line:   4,
								},
								File:   "sql.flux",
								Source: "dataSourceName: string",
								Start: ast.Position{
									Column: 50,
									Line:   4,
								},
							},
						},
						Kind: "Required",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 64,
										Line:   4,
									},
									File:   "sql.flux",
									Source: "dataSourceName",
									Start: ast.Position{
										Column: 50,
										Line:   4,
									},
								},
							},
							Name: "dataSourceName",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 72,
										Line:   4,
									},
									File:   "sql.flux",
									Source: "string",
									Start: ast.Position{
										Column: 66,
										Line:   4,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 72,
											Line:   4,
										},
										File:   "sql.flux",
										Source: "string",
										Start: ast.Position{
											Column: 66,
											Line:   4,
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
									Column: 87,
									Line:   4,
								},
								File:   "sql.flux",
								Source: "table: string",
								Start: ast.Position{
									Column: 74,
									Line:   4,
								},
							},
						},
						Kind: "Required",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 79,
										Line:   4,
									},
									File:   "sql.flux",
									Source: "table",
									Start: ast.Position{
										Column: 74,
										Line:   4,
									},
								},
							},
							Name: "table",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 87,
										Line:   4,
									},
									File:   "sql.flux",
									Source: "string",
									Start: ast.Position{
										Column: 81,
										Line:   4,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 87,
											Line:   4,
										},
										File:   "sql.flux",
										Source: "string",
										Start: ast.Position{
											Column: 81,
											Line:   4,
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
									Column: 104,
									Line:   4,
								},
								File:   "sql.flux",
								Source: "?batchSize: int",
								Start: ast.Position{
									Column: 89,
									Line:   4,
								},
							},
						},
						Kind: "Optional",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 99,
										Line:   4,
									},
									File:   "sql.flux",
									Source: "batchSize",
									Start: ast.Position{
										Column: 90,
										Line:   4,
									},
								},
							},
							Name: "batchSize",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 104,
										Line:   4,
									},
									File:   "sql.flux",
									Source: "int",
									Start: ast.Position{
										Column: 101,
										Line:   4,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 104,
											Line:   4,
										},
										File:   "sql.flux",
										Source: "int",
										Start: ast.Position{
											Column: 101,
											Line:   4,
										},
									},
								},
								Name: "int",
							},
						},
					}},
					Return: &ast.ArrayType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 112,
									Line:   4,
								},
								File:   "sql.flux",
								Source: "[A]",
								Start: ast.Position{
									Column: 109,
									Line:   4,
								},
							},
						},
						ElementType: &ast.TvarType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 111,
										Line:   4,
									},
									File:   "sql.flux",
									Source: "A",
									Start: ast.Position{
										Column: 110,
										Line:   4,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 111,
											Line:   4,
										},
										File:   "sql.flux",
										Source: "A",
										Start: ast.Position{
											Column: 110,
											Line:   4,
										},
									},
								},
								Name: "A",
							},
						},
					},
				},
			},
		}},
		Imports:  nil,
		Metadata: "parser-type=rust",
		Name:     "sql.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 12,
						Line:   1,
					},
					File:   "sql.flux",
					Source: "package sql",
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
							Column: 12,
							Line:   1,
						},
						File:   "sql.flux",
						Source: "sql",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "sql",
			},
		},
	}},
	Package: "sql",
	Path:    "sql",
}
