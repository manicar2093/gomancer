package bootstrap

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/coditory/go-errors"
	"go/format"
	"io/fs"
	"os"
	"path"
)

func copyCoreToProject(projectDirName string, input InitProjectInput) error {
	return copyDirToProject(coreFS, "core", projectDirName, input)
}

func copyCmdToProject(projectDirName string, input InitProjectInput) error {
	return copyDirToProject(cmdFS, "cmd", projectDirName, input)
}

func copyDirToProject(atFs embed.FS, dir, projectDirName string, input InitProjectInput) error {
	return fs.WalkDir(atFs, dir, func(currentPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		projectDir := path.Join(projectDirName, currentPath)
		log.Debug("Copying file:", "currentPath", currentPath, "into", projectDir)

		if d.IsDir() {
			if err := os.MkdirAll(projectDir, 0755); err != nil {
				return err
			}

			return nil
		} else {

			fileExt := path.Ext(currentPath)

			content, err := fs.ReadFile(atFs, currentPath)
			if err != nil {
				fmt.Println("Error reading file:", currentPath, err)
				return err
			}

			content = bytes.Replace(content, []byte("\"github.com/manicar2093/gomancer/bootstrap"), []byte(fmt.Sprintf("\"%s", input.ModuleName)), -1)

			if fileExt == ".go" {
				content, err = format.Source(content)
				if err != nil {
					return errors.Wrap(err, "Error formatting .go file")
				}
			}

			if err := os.WriteFile(path.Join(projectDir), content, 0755); err != nil {
				return err
			}
		}

		return nil
	})
}
