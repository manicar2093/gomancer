package testmatchers

import (
	"fmt"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"
	"os"
)

type BeAnExistingFileAndEqualStringMatcher struct {
	ExpectedStringContent string
}

func BeAnExistingFileAndEqualString(expectedStringContent string) types.GomegaMatcher {
	return &BeAnExistingFileAndEqualStringMatcher{
		expectedStringContent,
	}
}

func (c BeAnExistingFileAndEqualStringMatcher) Match(actual interface{}) (success bool, err error) {
	sourceFilePath, ok := actual.(string)
	if !ok {
		return false, fmt.Errorf("BeAnExistingFileAndEqualStringMatcher matcher expects a file path")
	}
	existingFileMatcher := matchers.BeAnExistingFileMatcher{}
	isExisting, err := existingFileMatcher.Match(sourceFilePath)
	if err != nil || !isExisting {
		return isExisting, err
	}

	sourceFilePathContent, err := os.ReadFile(sourceFilePath)
	if err != nil {
		return false, err
	}

	equalMatcher := matchers.EqualMatcher{
		Expected: string(sourceFilePathContent),
	}
	return equalMatcher.Match(c.ExpectedStringContent)

}

func (c BeAnExistingFileAndEqualStringMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to exist and equal string")
}

func (c BeAnExistingFileAndEqualStringMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to exist and not match string")
}
