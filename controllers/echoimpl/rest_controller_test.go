package echoimpl_test

import (
	"github.com/manicar2093/gomancer/domain"
	"github.com/manicar2093/gomancer/testfixtures"
	"github.com/manicar2093/gomancer/testmatchers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
	"path"

	"github.com/manicar2093/gomancer/controllers/echoimpl"
)

var _ = Describe("Controllers", func() {

	BeforeEach(func() {
		os.MkdirAll(string(domain.CmdServiceControllersPackagePath), os.ModePerm)
	})

	AfterEach(func() {
		os.RemoveAll(string(domain.CmdPackagePath))
	})

	Describe("GenerateRestController", func() {
		It("creates a new rest controller with given input", func() {
			var (
				input = testfixtures.ModelBinaryIdSuccess
			)

			err := echoimpl.GenerateRestController(input, testfixtures.ModelSuccessDepsContainer, testfixtures.ModelSuccessDepInCreation)

			Expect(err).ToNot(HaveOccurred())
			Expect(path.Join(string(domain.CmdServiceControllersPackagePath), "post_tests_rest.go")).Should(
				testmatchers.BeAnExistingFileWithEqualContent(path.Join("fixtures", "rest_controller.txt")),
			)
		})
	})

	Describe("GenerateWebController", func() {
		It("creates a new web controller with given input", func() {
			var (
				input = testfixtures.ModelBinaryIdSuccess
			)

			err := echoimpl.GenerateWebController(input, testfixtures.ModelSuccessDepsContainer, testfixtures.ModelSuccessDepInCreation)

			Expect(err).ToNot(HaveOccurred())
			//Expect(path.Join(string(domain.CmdServiceControllersPackagePath), "post_web_tests.go")).Should(
			//	testmatchers.BeAnExistingFileWithEqualContent(path.Join("fixtures", "controller.txt")),
			//)
		})
	})
})
