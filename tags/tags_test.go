package tags_test

import (
	"github.com/manicar2093/gomancer/tags"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tags", func() {
	Expect(true).To(BeTrue())

	Describe("NewJson", func() {
		DescribeTable(
			"creates into map a json tag with given options",
			func(options tags.JsonOptions, expectedContent string) {
				var (
					generable = tags.NewJson(options)
				)

				tag, content := generable.Generate()

				Expect(tag).To(Equal(options.Name))
				Expect(content).To(Equal(expectedContent))
			},
			Entry(
				"alone",
				tags.JsonOptions{Name: "age"},
				"age",
			),
			Entry(
				"with omitempty",
				tags.JsonOptions{Name: "age", IsOmitEmpty: true},
				"age,omitempty",
			),
			Entry(
				"with omitzero",
				tags.JsonOptions{Name: "age", IsOmitZero: true},
				"age,omitzero",
			),
			Entry(
				"with omitempty and omitzero",
				tags.JsonOptions{Name: "age", IsOmitEmpty: true, IsOmitZero: true},
				"age,omitempty,omitzero",
			),
			Entry(
				"with ,inline",
				tags.JsonOptions{Name: "age", IsInline: true},
				",inline",
			),
			Entry(
				"with ,inline refusing omitempty and omitzero",
				tags.JsonOptions{Name: "age", IsOmitEmpty: true, IsOmitZero: true, IsInline: true},
				",inline",
			),
		)
	})
})
