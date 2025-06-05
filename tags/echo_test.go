package tags_test

import (
	"github.com/manicar2093/gomancer/tags"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tags", func() {
	Describe("NewQuery", func() {
		DescribeTable(
			"creates tag and value for echo query tag with given options",
			func(options tags.EchoOptions, expectedTag string, expectedContent string) {
				var (
					generable = tags.NewEcho(options)
				)

				tag, content := generable.Generate()

				Expect(tag).To(Equal(expectedTag))
				Expect(content).To(Equal(expectedContent))
			},
			Entry(
				"query tag",
				tags.EchoOptions{Name: "age", Tag: tags.EchoQuery},
				tags.EchoQuery,
				"age",
			),
			Entry(
				"param tag",
				tags.EchoOptions{Name: "age", Tag: tags.EchoParam},
				tags.EchoParam,
				"age",
			),
			Entry(
				"header tag",
				tags.EchoOptions{Name: "age", Tag: tags.EchoHeader},
				tags.EchoHeader,
				"age",
			),
			Entry(
				"form tag",
				tags.EchoOptions{Name: "age", Tag: tags.EchoForm},
				tags.EchoForm,
				"age",
			),
		)
	})
})
