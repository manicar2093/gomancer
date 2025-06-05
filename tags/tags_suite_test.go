package tags_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTags(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tags Suite")
}
