package testmatchers

import (
	"fmt"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"
	"os"
)

type BeAnExistingFileWithEqualFileContentMatcher struct {
	EquallySourceFilePath string
}

func BeAnExistingFileWithEqualContent(equallySourceFilePath string) types.GomegaMatcher {
	return &BeAnExistingFileWithEqualFileContentMatcher{
		equallySourceFilePath,
	}
}

func (c BeAnExistingFileWithEqualFileContentMatcher) Match(actual interface{}) (success bool, err error) {
	mustExistingFilePath, ok := actual.(string)
	if !ok {
		return false, fmt.Errorf("BeAnExistingFileWithEqualFileContentMatcher matcher expects a file path")
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

func (c BeAnExistingFileWithEqualFileContentMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to exist and match content")
}

func (c BeAnExistingFileWithEqualFileContentMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to exist and not match content (this has no sense any way ._.)")
}
