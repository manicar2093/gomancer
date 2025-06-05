package tags_test

import (
	"github.com/manicar2093/gomancer/tags"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tags", func() {
	Describe("NewMapstructure", func() {
		DescribeTable(
			"creates tag and value for mapstructure tag with given options",
			func(options tags.MapstructureOptions, expectedContent string) {
				var (
					generable = tags.NewMapstructure(options)
				)

				tag, content := generable.Generate()

				Expect(tag).To(Equal("mapstructure"))
				Expect(content).To(Equal(expectedContent))
			},
			Entry(
				"alone",
				tags.MapstructureOptions{Name: "age"},
				"age",
			),
			Entry(
				"with ,squash refusing tag name",
				tags.MapstructureOptions{Name: "age", IsSquash: true},
				",squash",
			),
			Entry(
				"with ,remain refusing tag name",
				tags.MapstructureOptions{Name: "age", IsRemain: true},
				",remain",
			),
			Entry(
				"with ,omitempty",
				tags.MapstructureOptions{Name: "age", IsOmitEmpty: true},
				"age,omitempty",
			),
		)
	})
})
