package env

import (
	"fmt"
	"os"
)

// GetEnvOrPanic look for the env variable. If not present will panic
func GetEnvOrPanic(name string) string {
	value := os.Getenv(name)
	if value == "" {
		panic(fmt.Sprintf("'%s' not found", name))
	}
	return value
}
