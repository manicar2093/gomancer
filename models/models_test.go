package models_test

import (
	"fmt"
	"github.com/manicar2093/gomancer/domain"
	"github.com/manicar2093/gomancer/models"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
	"path"
)

var _ = Describe("Models", func() {
	var (
		testPath        = "github.com/user/project_name"
		readFixtureFile = func(t GinkgoTestingT, filename string) string {
			d, err := os.ReadFile(fmt.Sprintf("./fixtures/%s", filename))
			if err != nil {
				t.Fail()
			}
			return string(d)
		}
		readFromModelsDir = func(t GinkgoTestingT, filename string) string {
			d, err := os.ReadFile(fmt.Sprintf("./internal/domain/models/%s.go", filename))
			if err != nil {
				t.Fail()
			}
			return string(d)
		}
		readFromGeneratorsDir = func(t GinkgoTestingT, filename string) string {
			d, err := os.ReadFile(fmt.Sprintf("./pkg/generators/%s.go", filename))
			if err != nil {
				t.Fail()
			}
			return string(d)
		}
	)

	BeforeEach(func() {
		os.MkdirAll(string(domain.ModelsPackagePath), os.ModePerm)
		os.MkdirAll(string(domain.GeneratorsPackagePath), os.ModePerm)
	})
	AfterEach(func() {
		os.RemoveAll("internal")
		os.RemoveAll("pkg")
	})

	Describe("GenerateModel", func() {
		It("creates text for model and function to generate models on test", func() {
			var input = models.GenerateModelInput{
				ModuleInfo: domain.ModuleInfo{
					Name: testPath,
				},
				TransformedText: domain.TransformedText{
					SnakeCase:  "post_test",
					PascalCase: "PostTest",
					CamelCase:  "postTest",
				},
				IdAttribute: models.IdAttribute{
					TransformedText: domain.TransformedText{
						SnakeCase:  "id",
						PascalCase: "Id",
						CamelCase:  "id",
					},
					Type: "uuid",
				},
				Attributes: []models.Attribute{
					{
						TransformedText: domain.TransformedText{
							SnakeCase:  "name",
							PascalCase: "Name",
							CamelCase:  "name",
						},
						Type:       "string",
						IsOptional: true,
					},
					{
						TransformedText: domain.TransformedText{
							SnakeCase:  "total_price",
							PascalCase: "TotalPrice",
							CamelCase:  "totalPrice",
						},
						Type: "decimal",
					},
					{
						TransformedText: domain.TransformedText{
							SnakeCase:  "price",
							PascalCase: "Price",
							CamelCase:  "price",
						},
						Type:       "decimal",
						IsOptional: true,
					},
					{
						TransformedText: domain.TransformedText{
							SnakeCase:  "age",
							PascalCase: "Age",
							CamelCase:  "age",
						},
						Type: "int",
					},
					{
						TransformedText: domain.TransformedText{
							SnakeCase:  "case_of_use",
							PascalCase: "CaseOfUse",
							CamelCase:  "caseOfUse",
						},
						Type:       "int32",
						IsOptional: true,
					},
				},
			}

			err := models.GenerateModel(input)

			Expect(err).ToNot(HaveOccurred())
			Expect(path.Join(string(domain.ModelsPackagePath), "post_test.go")).Should(BeAnExistingFile())
			Expect(readFromModelsDir(GinkgoT(), "post_test")).To(Equal(readFixtureFile(GinkgoT(), "model_1.txt")))
			Expect(path.Join(string(domain.GeneratorsPackagePath), "post_test.go")).Should(BeAnExistingFile())
			Expect(readFromGeneratorsDir(GinkgoT(), "post_test")).To(Equal(readFixtureFile(GinkgoT(), "model_2.txt")))
		})
	})
})
