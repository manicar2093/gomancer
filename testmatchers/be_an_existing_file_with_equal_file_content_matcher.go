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

	existingFileMatcher *matchers.BeAnExistingFileMatcher
	equalMatcher        *matchers.EqualMatcher
	failExistingMatcher bool
	failEqualMatcher    bool
	expectedContent     string
}

func BeAnExistingFileWithEqualContent(equallySourceFilePath string) types.GomegaMatcher {
	return &BeAnExistingFileWithEqualFileContentMatcher{
		EquallySourceFilePath: equallySourceFilePath,
		existingFileMatcher:   &matchers.BeAnExistingFileMatcher{},
		equalMatcher:          &matchers.EqualMatcher{},
	}
}

func (c *BeAnExistingFileWithEqualFileContentMatcher) Match(actual interface{}) (success bool, err error) {
	mustExistingFilePath, ok := actual.(string)
	if !ok {
		return false, fmt.Errorf("BeAnExistingFileWithEqualFileContentMatcher matcher expects a file path")
	}

	success, err = c.existingFileMatcher.Match(mustExistingFilePath)
	if err != nil || !success {
		c.failExistingMatcher = true
		return success, err
	}

	mustExistingSourceContent, err := os.ReadFile(mustExistingFilePath)
	if err != nil {
		return false, err
	}

	equallySourceContent, err := os.ReadFile(c.EquallySourceFilePath)
	if err != nil {
		return false, err
	}

	c.equalMatcher.Expected = string(mustExistingSourceContent)
	success, err = c.equalMatcher.Match(equallySourceContent)
	if err != nil || !success {
		c.failEqualMatcher = true
		c.expectedContent = string(equallySourceContent)
		return success, err
	}
	return

}

func (c *BeAnExistingFileWithEqualFileContentMatcher) FailureMessage(actual interface{}) (message string) {
	if c.failExistingMatcher {
		return c.existingFileMatcher.FailureMessage(actual)
	}
	if c.failEqualMatcher {
		return c.equalMatcher.FailureMessage(c.expectedContent)
	}
	return format.Message(actual, "to exist and match content")
}

func (c *BeAnExistingFileWithEqualFileContentMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	if c.failExistingMatcher {
		return c.existingFileMatcher.NegatedFailureMessage(actual)
	}
	if c.failEqualMatcher {
		return c.equalMatcher.NegatedFailureMessage(actual)
	}
	return format.Message(actual, "not to exist and not match content (this has no sense any way ._.)")
}
