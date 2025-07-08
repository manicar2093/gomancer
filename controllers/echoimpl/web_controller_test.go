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

var _ = Describe("WebController", func() {
	BeforeEach(func() {
		os.MkdirAll(string(domain.CmdServiceControllersPackagePath), os.ModePerm)
	})

	AfterEach(func() {
		os.RemoveAll(string(domain.CmdPackagePath))
	})

	Describe("GenerateWebController", func() {
		It("creates a new web controller with given input", func() {
			var (
				input = testfixtures.ModelBinaryIdSuccess
			)

			err := echoimpl.GenerateWebController(input, testfixtures.ModelSuccessDepsContainer, testfixtures.ModelSuccessDepInCreation)

			Expect(err).ToNot(HaveOccurred())
			Expect(path.Join(string(domain.CmdServiceControllersPackagePath), "post_tests_web.go")).Should(
				testmatchers.BeAnExistingFileWithEqualContent(path.Join("fixtures", "web_controller.txt")),
			)
			Expect(path.Join(string(domain.CmdServiceControllersPackagePath), "posttestpages", "post_test_form.templ")).Should(
				testmatchers.BeAnExistingFileWithEqualContent(path.Join("fixtures", "web_form.txt")),
			)
			Expect(path.Join(string(domain.CmdServiceControllersPackagePath), "posttestpages", "post_test_table.templ")).Should(
				testmatchers.BeAnExistingFileWithEqualContent(path.Join("fixtures", "web_table.txt")),
			)
			Expect(path.Join(string(domain.CmdServiceControllersPackagePath), "posttestpages", "post_test_show.templ")).Should(
				testmatchers.BeAnExistingFileWithEqualContent(path.Join("fixtures", "web_show.txt")),
			)
			Expect(path.Join(string(domain.CmdServiceControllersPackagePath), "posttestpages", "post_test_register.templ")).Should(
				testmatchers.BeAnExistingFileWithEqualContent(path.Join("fixtures", "web_register.txt")),
			)
			Expect(path.Join(string(domain.CmdServiceControllersPackagePath), "posttestpages", "post_test_index.templ")).Should(
				testmatchers.BeAnExistingFileWithEqualContent(path.Join("fixtures", "web_index.txt")),
			)
		})
	})
})
