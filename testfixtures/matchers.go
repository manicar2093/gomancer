package testfixtures

import (
	"fmt"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"
	"os"
)

type BeAnExistingFileWithEqualContentMatcher struct {
	EquallySourceFilePath string
}

func BeAnExistingFileWithEqualContent(equallySourceFilePath string) types.GomegaMatcher {
	return &BeAnExistingFileWithEqualContentMatcher{
		equallySourceFilePath,
	}
}

func (c BeAnExistingFileWithEqualContentMatcher) Match(actual interface{}) (success bool, err error) {
	//println("REALLY!!!!!!! :O -------------")
	mustExistingFilePath, ok := actual.(string)
	if !ok {
		return false, fmt.Errorf("BeAnExistingFileWithEqualContentMatcher matcher expects a file path")
	}
	existingFileMatcher := matchers.BeAnExistingFileMatcher{}
	isExisting, err := existingFileMatcher.Match(mustExistingFilePath)
	if err != nil || !isExisting {
		return isExisting, err
	}

	mustExistingSourceContent, err := os.ReadFile(mustExistingFilePath)
	if err != nil {
		return false, err
	}

	equallySourceContent, err := os.ReadFile(c.EquallySourceFilePath)
	if err != nil {
		return false, err
	}

	equalMatcher := matchers.EqualMatcher{
		Expected: mustExistingSourceContent,
	}
	return equalMatcher.Match(equallySourceContent)

}

func (c BeAnExistingFileWithEqualContentMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to exist and match content")
}

func (c BeAnExistingFileWithEqualContentMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to exist and not match content (this has no sense any way ._.)")
}
