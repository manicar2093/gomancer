package gormimpl_test

import (
	"os"
	"path"

	"github.com/manicar2093/gomancer/domain"
	"github.com/manicar2093/gomancer/testfixtures"
	"github.com/manicar2093/gomancer/testmatchers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/manicar2093/gomancer/repositories/gormimpl"
)

var _ = Describe("Gorm", func() {

	AfterEach(func() {
		os.RemoveAll(string(domain.InternalPackagePath))
	})

	Describe("GenerateRepository", func() {
		It("creates a repository with given data", func() {

			err := gormimpl.GenerateRepository(
				testfixtures.ModelBinaryIdSuccess,
				testfixtures.ModelSuccessDepsContainer,
				testfixtures.ModelSuccessDepInCreation,
			)

			Expect(err).ToNot(HaveOccurred())
			Expect(path.Join(string(domain.InternalPackagePath), "posttests", "repository_gomancer.go")).Should(
				testmatchers.BeAnExistingFileWithEqualContent(path.Join("fixtures", "repository.txt")),
			)
			Expect(path.Join(string(domain.InternalPackagePath), "posttests", "posttests_suite_test.go")).Should(
				testmatchers.BeAnExistingFileWithEqualContent(path.Join("fixtures", "entity_suite_test.txt")),
			)
			Expect(path.Join(string(domain.InternalPackagePath), "posttests", "repository_gomancer_test.go")).Should(
				testmatchers.BeAnExistingFileWithEqualContent(path.Join("fixtures", "repository_test.txt")),
			)
		})
	})
})
