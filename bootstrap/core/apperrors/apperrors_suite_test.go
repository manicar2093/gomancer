package apperrors_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestApperrors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Apperrors Suite")
}
