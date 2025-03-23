package models_test

import (
	"github.com/manicar2093/gomancer/domain"
	"github.com/manicar2093/gomancer/models"
	"github.com/manicar2093/gomancer/testfixtures"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
	"path"
)

var _ = Describe("Models", func() {

	BeforeEach(func() {
		os.MkdirAll(string(domain.InternalDomainModelsPackagePath), os.ModePerm)
		os.MkdirAll(string(domain.PkgGeneratorsPackagePath), os.ModePerm)
	})
	AfterEach(func() {
		os.RemoveAll(string(domain.InternalPackagePath))
		os.RemoveAll(string(domain.PkgPackagePath))
	})

	Describe("GenerateModel", func() {
		It("creates text for model and function to generate models on test", func() {
			var input = testfixtures.ModelSuccess

			err := models.GenerateModel(input)

			Expect(err).ToNot(HaveOccurred())
			Expect(path.Join(string(domain.InternalDomainModelsPackagePath), "post_test.go")).Should(
				testfixtures.BeAnExistingFileWithEqualContent(path.Join("fixtures", "model_success.txt")),
			)
			Expect(path.Join(string(domain.PkgGeneratorsPackagePath), "post_test.go")).Should(
				testfixtures.BeAnExistingFileWithEqualContent(path.Join("fixtures", "generator_success.txt")),
			)
		})
	})
})
