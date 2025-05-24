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

var _ = Describe("Controller", func() {

	BeforeEach(func() {
		os.MkdirAll(string(domain.CmdApiControllersPackagePath), os.ModePerm)
	})

	AfterEach(func() {
		os.RemoveAll(string(domain.CmdPackagePath))
	})

	Describe("GenerateController", func() {
		It("creates a new controller with given input", func() {
			var (
				input = testfixtures.ModelSuccess
			)

			err := echoimpl.GenerateController(input, testfixtures.ModelSuccessDepsContainer, testfixtures.ModelSuccessDepInCreation)

			Expect(err).ToNot(HaveOccurred())
			Expect(path.Join(string(domain.CmdApiControllersPackagePath), "post_test.go")).Should(
				testmatchers.BeAnExistingFileWithEqualContent(path.Join("fixtures", "controller.txt")),
			)
		})
	})
})
