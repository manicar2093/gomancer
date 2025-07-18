package parser_test

import (
	"github.com/manicar2093/gomancer/parser"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"
	"strconv"
)

func idSliceGetter(index int, _ interface{}) string {
	return strconv.Itoa(index)
}

var _ = Describe("Parser", func() {
	Describe("ParseArgs", func() {
		It("should parse arguments", func() {
			var (
				isPkUuid   = false
				moduleName = "test"
				args       = []string{
					"HelloWorld",
					"arg1:string:optional",
					"arg2:int",
					"arg3:decimal:optional",
					"arg4:enum/terror/lol:optional",
				}
			)

			got, errorss, hasErr := parser.ParseArgs(args, moduleName, isPkUuid)

			Expect(errorss).To(BeEmpty())
			Expect(hasErr).To(BeFalse())
			Expect(got).To(gstruct.MatchAllFields(gstruct.Fields{
				"PackageEntityName": Equal("helloworlds"),
				"TransformedText": gstruct.MatchAllFields(gstruct.Fields{
					"SnakeCase":        Equal("hello_world"),
					"PascalCase":       Equal("HelloWorld"),
					"CamelCase":        Equal("helloWorld"),
					"LowerNoSpaceCase": Equal("helloworld"),
				}),
				"ModuleInfo": gstruct.MatchAllFields(gstruct.Fields{
					"Name": Equal(moduleName),
				}),
				"IdAttribute": gstruct.MatchAllFields(gstruct.Fields{
					"TransformedText": gstruct.MatchAllFields(gstruct.Fields{
						"SnakeCase":        Equal("id"),
						"PascalCase":       Equal("Id"),
						"CamelCase":        Equal("id"),
						"LowerNoSpaceCase": Equal("id"),
					}),
					"Type":        Equal("int"),
					"IsOptional":  BeFalse(),
					"EnumStrings": BeEmpty(),
				}),
				"Attributes": gstruct.MatchAllElementsWithIndex(
					idSliceGetter,
					gstruct.Elements{
						"0": gstruct.MatchAllFields(gstruct.Fields{
							"TransformedText": gstruct.MatchAllFields(gstruct.Fields{
								"SnakeCase":        Equal("arg1"),
								"PascalCase":       Equal("Arg1"),
								"CamelCase":        Equal("arg1"),
								"LowerNoSpaceCase": Equal("arg1"),
							}),
							"Type":        Equal("string"),
							"IsOptional":  BeTrue(),
							"EnumStrings": BeEmpty(),
						}),
						"1": gstruct.MatchAllFields(gstruct.Fields{
							"TransformedText": gstruct.MatchAllFields(gstruct.Fields{
								"SnakeCase":        Equal("arg2"),
								"PascalCase":       Equal("Arg2"),
								"CamelCase":        Equal("arg2"),
								"LowerNoSpaceCase": Equal("arg2"),
							}),
							"Type":        Equal("int"),
							"IsOptional":  BeFalse(),
							"EnumStrings": BeEmpty(),
						}),
						"2": gstruct.MatchAllFields(gstruct.Fields{
							"TransformedText": gstruct.MatchAllFields(gstruct.Fields{
								"SnakeCase":        Equal("arg3"),
								"PascalCase":       Equal("Arg3"),
								"CamelCase":        Equal("arg3"),
								"LowerNoSpaceCase": Equal("arg3"),
							}),
							"Type":        Equal("decimal"),
							"IsOptional":  BeTrue(),
							"EnumStrings": BeEmpty(),
						}),
						"3": gstruct.MatchAllFields(gstruct.Fields{
							"TransformedText": gstruct.MatchAllFields(gstruct.Fields{
								"SnakeCase":        Equal("arg4"),
								"PascalCase":       Equal("Arg4"),
								"CamelCase":        Equal("arg4"),
								"LowerNoSpaceCase": Equal("arg4"),
							}),
							"Type":       Equal("enum"),
							"IsOptional": BeTrue(),
							"EnumStrings": gstruct.MatchAllElementsWithIndex(idSliceGetter, gstruct.Elements{
								"0": gstruct.MatchAllFields(gstruct.Fields{
									"SnakeCase":        Equal("terror"),
									"PascalCase":       Equal("Terror"),
									"CamelCase":        Equal("terror"),
									"LowerNoSpaceCase": Equal("terror"),
								}),
								"1": gstruct.MatchAllFields(gstruct.Fields{
									"SnakeCase":        Equal("lol"),
									"PascalCase":       Equal("Lol"),
									"CamelCase":        Equal("lol"),
									"LowerNoSpaceCase": Equal("lol"),
								}),
							}),
						}),
					},
				),
			}))
		})

		When("optional option is misspell", func() {
			It("should return error", func() {
				var (
					isPkUuid   = false
					moduleName = "test"
					args       = []string{
						"HelloWorld",
						"arg1:string:optionl",
						"arg2:int",
						"arg3:decimal:",
					}
				)

				_, errDetails, hasErr := parser.ParseArgs(args, moduleName, isPkUuid)

				Expect(hasErr).To(BeTrue())
				Expect(errDetails[0]).To(gstruct.MatchAllFields(gstruct.Fields{
					"Input": Equal("arg1:string:optionl"),
					"Err":   Equal("expected optional declaration, got 'optionl'"),
					"Index": Not(BeZero()),
				}))
				Expect(errDetails).To(gstruct.MatchAllElementsWithIndex(
					idSliceGetter, gstruct.Elements{
						"0": gstruct.MatchAllFields(gstruct.Fields{
							"Input": Equal("arg1:string:optionl"),
							"Err":   Equal("expected optional declaration, got 'optionl'"),
							"Index": Not(BeZero()),
						}),
						"1": gstruct.MatchAllFields(gstruct.Fields{
							"Input": Equal("arg3:decimal:"),
							"Err":   Equal("expected optional declaration, got '<empty>'"),
							"Index": Not(BeZero()),
						}),
					}),
				)
			})
		})

		When("type is not supported", func() {
			It("should return error details", func() {
				var (
					isPkUuid   = false
					moduleName = "test"
					args       = []string{
						"HelloWorld",
						"arg1:asdf:apsdf",
						"arg2:sdfg",
						"arg3:dfgh",
					}
				)

				_, errDetails, hasErr := parser.ParseArgs(args, moduleName, isPkUuid)

				Expect(hasErr).To(BeTrue())
				Expect(errDetails).To(gstruct.MatchAllElementsWithIndex(
					idSliceGetter, gstruct.Elements{
						"0": gstruct.MatchAllFields(gstruct.Fields{
							"Input": Equal("arg1:asdf:apsdf"),
							"Err":   Equal("type 'asdf' is not supported"),
							"Index": Not(BeZero()),
						}),
						"1": gstruct.MatchAllFields(gstruct.Fields{
							"Input": Equal("arg2:sdfg"),
							"Err":   Equal("type 'sdfg' is not supported"),
							"Index": Not(BeZero()),
						}),
						"2": gstruct.MatchAllFields(gstruct.Fields{
							"Input": Equal("arg3:dfgh"),
							"Err":   Equal("type 'dfgh' is not supported"),
							"Index": Not(BeZero()),
						}),
					}),
				)
			})
		})

		When("any input has no enough data", func() {
			It("should return error details", func() {
				var (
					isPkUuid   = false
					moduleName = "test"
					args       = []string{
						"HelloWorld",
						"arg1",
						"arg2:",
						"arg3:string",
						"arg4:string:optional",
					}
				)

				_, errDetails, hasErr := parser.ParseArgs(args, moduleName, isPkUuid)

				Expect(hasErr).To(BeTrue())
				Expect(errDetails).To(gstruct.MatchAllElementsWithIndex(
					idSliceGetter, gstruct.Elements{
						"0": gstruct.MatchAllFields(gstruct.Fields{
							"Input": Equal("arg1"),
							"Err":   Equal("not enough data to continue. Remember syntax: attribute:type:optional"),
							"Index": Not(BeZero()),
						}),
						"1": gstruct.MatchAllFields(gstruct.Fields{
							"Input": Equal("arg2:"),
							"Err":   Equal("type '' is not supported"),
							"Index": Not(BeZero()),
						}),
					}),
				)

			})
		})
	})
})
