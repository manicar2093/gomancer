package domain_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"
	"strconv"

	"github.com/manicar2093/gomancer/domain"
)

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
				}
			)

			got, _, hasErr := domain.ParseArgs(args, moduleName, isPkUuid)

			Expect(hasErr).To(BeFalse())
			Expect(got).To(gstruct.MatchAllFields(gstruct.Fields{
				"TransformedText": gstruct.MatchAllFields(gstruct.Fields{
					"SnakeCase":  Equal("hello_world"),
					"PascalCase": Equal("HelloWorld"),
					"CamelCase":  Equal("helloWorld"),
				}),
				"ModuleInfo": gstruct.MatchAllFields(gstruct.Fields{
					"Name": Equal(moduleName),
				}),
				"IdAttribute": gstruct.MatchAllFields(gstruct.Fields{
					"TransformedText": gstruct.MatchAllFields(gstruct.Fields{
						"SnakeCase":  Equal("id"),
						"PascalCase": Equal("Id"),
						"CamelCase":  Equal("id"),
					}),
					"Type":       Equal("int"),
					"IsOptional": BeFalse(),
				}),
				"Attributes": gstruct.MatchAllElementsWithIndex(
					func(index int, _ interface{}) string {
						return strconv.Itoa(index)
					},
					gstruct.Elements{
						"0": gstruct.MatchAllFields(gstruct.Fields{
							"TransformedText": gstruct.MatchAllFields(gstruct.Fields{
								"SnakeCase":  Equal("arg_1"),
								"PascalCase": Equal("Arg1"),
								"CamelCase":  Equal("arg1"),
							}),
							"Type":       Equal("string"),
							"IsOptional": BeTrue(),
						}),
						"1": gstruct.MatchAllFields(gstruct.Fields{
							"TransformedText": gstruct.MatchAllFields(gstruct.Fields{
								"SnakeCase":  Equal("arg_2"),
								"PascalCase": Equal("Arg2"),
								"CamelCase":  Equal("arg2"),
							}),
							"Type":       Equal("int"),
							"IsOptional": BeFalse(),
						}),
						"2": gstruct.MatchAllFields(gstruct.Fields{
							"TransformedText": gstruct.MatchAllFields(gstruct.Fields{
								"SnakeCase":  Equal("arg_3"),
								"PascalCase": Equal("Arg3"),
								"CamelCase":  Equal("arg3"),
							}),
							"Type":       Equal("decimal"),
							"IsOptional": BeTrue(),
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
						"arg3:decimal:optional",
					}
				)

				_, errDetails, hasErr := domain.ParseArgs(args, moduleName, isPkUuid)

				Expect(hasErr).To(BeTrue())
				Expect(errDetails[0]).To(gstruct.MatchAllFields(gstruct.Fields{
					"Input": Equal("arg1:string:optionl"),
					"Err":   Equal("expected optional declaration, got 'optionl'"),
				}))

			})
		})
	})
})
