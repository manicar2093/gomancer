package env_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGetenv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Getenv Suite")
}
