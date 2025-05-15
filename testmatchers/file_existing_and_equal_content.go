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
	isFileExists          bool
	contentFile           []byte
}

func BeAnExistingFileAndEqualString(expectedStringContent string) types.GomegaMatcher {
	return &BeAnExistingFileAndEqualStringMatcher{
		ExpectedStringContent: expectedStringContent,
	}
}

func (c *BeAnExistingFileAndEqualStringMatcher) Match(actual interface{}) (success bool, err error) {
	sourceFilePath, ok := actual.(string)
	if !ok {
		return false, fmt.Errorf("BeAnExistingFileAndEqualStringMatcher matcher expects a file path")
	}
	existingFileMatcher := matchers.BeAnExistingFileMatcher{}
	c.isFileExists, err = existingFileMatcher.Match(sourceFilePath)
	if err != nil || !c.isFileExists {
		return c.isFileExists, err
	}

	sourceFilePathContent, err := os.ReadFile(sourceFilePath)
	if err != nil {
		return false, err
	}

	equalMatcher := matchers.EqualMatcher{
		Expected: string(sourceFilePathContent),
	}
	c.contentFile = sourceFilePathContent
	return equalMatcher.Match(c.ExpectedStringContent)

}

func (c *BeAnExistingFileAndEqualStringMatcher) FailureMessage(actual interface{}) (message string) {
	if !c.isFileExists {
		return format.Message(actual, "to exist")
	}
	return format.Message(string(c.contentFile), "to equal string content", c.ExpectedStringContent)
}

func (c *BeAnExistingFileAndEqualStringMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to exist and not match string")
}
