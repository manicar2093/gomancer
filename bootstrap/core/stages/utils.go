package stages

import (
	"fmt"
	"os"
	"slices"
)

// GetCurrentStage get ENVIRONMENT variable and checks if it's into supported stages. If it is not return an error
func GetCurrentStage() (string, error) {
	stage := os.Getenv("ENVIRONMENT")
	switch {
	case stage == "":
		return Dev, nil
	case slices.Contains(allStages, stage):
		return stage, nil

	default:
		return stage, fmt.Errorf("stage '%s' is not supported", stage)
	}
}
