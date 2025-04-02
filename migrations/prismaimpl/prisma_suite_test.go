package prismaimpl_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPrisma(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Prisma Suite")
}
