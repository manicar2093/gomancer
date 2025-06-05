package tags_test

import (
	"github.com/manicar2093/gomancer/tags"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gorm", func() {
	Expect(true).To(BeTrue())

	Describe("NewPrimaryKey", func() {
		DescribeTable(
			"creates tag and value for gorm tag for primary key",
			func(options tags.GormPKOptions, expectedContent string) {
				var (
					generable = tags.NewGormPK(options)
				)

				tag, content := generable.Generate()

				Expect(tag).To(Equal("gorm"))
				Expect(content).To(Equal(expectedContent))
			},
			Entry("int id", tags.GormPKOptions{}, "primaryKey"),
			Entry("uuid id", tags.GormPKOptions{IsUuid: true}, "default:gen_random_uuid();primaryKey"),
		)
	})
})
