package domain

import (
	"github.com/charmbracelet/log"
	"golang.org/x/mod/modfile"
	"os"
)

func GetGoModName() (string, error) {
	log.Info("Getting Go module name...")
	goModBytes, err := os.ReadFile("go.mod")
	if err != nil {
		return "", err
	}

	modName := modfile.ModulePath(goModBytes)
	log.Infof("Go module detected: %s", modName)

	return modName, nil
}
