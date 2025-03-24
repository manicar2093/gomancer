package prisma_test

import (
	"github.com/manicar2093/gomancer/domain"
	"github.com/manicar2093/gomancer/migrations/prisma"
	"github.com/manicar2093/gomancer/testfixtures"
	"github.com/manicar2093/gomancer/testmatchers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
	"path"
)

var _ = Describe("Prisma", func() {

	BeforeEach(func() {
		os.MkdirAll(string(domain.PrismaSchemaPackagePath), os.ModePerm)
	})
	AfterEach(func() {
		os.RemoveAll(string(domain.PrismaPackagePath))
	})

	Describe("GenerateMigration", func() {
		It("should generate migration with prisma", func() {
			err := prisma.GenerateMigration(testfixtures.ModelSuccess)

			Expect(err).ToNot(HaveOccurred())

			Expect(path.Join(string(domain.PrismaSchemaPackagePath), "post_test.prisma")).Should(testmatchers.BeAnExistingFileWithEqualContent(path.Join("fixtures", "migration.txt")))
		})
	})

})
