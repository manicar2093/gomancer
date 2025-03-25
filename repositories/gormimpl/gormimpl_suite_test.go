package gormimpl_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGormimpl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gormimpl Suite")
}
