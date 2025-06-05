package tags_test

import (
	"github.com/manicar2093/gomancer/tags"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validate", func() {
	Describe("NewValidate", func() {
		It("creates a required tag", func() {
			var (
				requiredTag = tags.NewValidateRequired()
				generable   = tags.NewValidate(requiredTag)
			)

			tag, content := generable.Generate()

			Expect(tag).To(Equal("validate"))
			Expect(content).To(Equal("required"))
		})

		It("creates a required_uuid tag", func() {
			var (
				requiredUuidTag = tags.NewValidateRequiredUuid()
				generable       = tags.NewValidate(requiredUuidTag)
			)

			tag, content := generable.Generate()

			Expect(tag).To(Equal("validate"))
			Expect(content).To(Equal("required_uuid"))
		})
	})
})
