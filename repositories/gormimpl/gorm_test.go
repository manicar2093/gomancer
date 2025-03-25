package gormimpl_test

import (
	"github.com/manicar2093/gomancer/domain"
	"github.com/manicar2093/gomancer/testfixtures"
	"github.com/manicar2093/gomancer/testmatchers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
	"path"

	"github.com/manicar2093/gomancer/repositories/gormimpl"
)

var _ = Describe("Gorm", func() {

	BeforeEach(func() {
		os.MkdirAll(path.Join(string(domain.InternalPackagePath), testfixtures.ModelSuccess.SnakeCase), os.ModePerm)
	})

	AfterEach(func() {
		os.RemoveAll(string(domain.InternalPackagePath))
	})

	Describe("GenerateRepository", func() {
		It("creates a repository with given data", func() {
			var input = testfixtures.ModelSuccess

			err := gormimpl.GenerateRepository(input)

			Expect(err).ToNot(HaveOccurred())
			Expect(path.Join(string(domain.InternalPackagePath), "post_test", "repository_gomancer.go")).Should(
				testmatchers.BeAnExistingFileWithEqualContent(path.Join("fixtures", "repository.txt")),
			)
		})
	})
})
