package bootstrap

import (
	"bytes"
	"fmt"
	"github.com/charmbracelet/log"
	"go/format"
	"io/fs"
	"os"
	"path"
)

func copyCoreToProject(projectDirName string, input InitProjectInput) error {
	return fs.WalkDir(coreFS, "core", func(currentPath string, d fs.DirEntry, err error) error {
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

			//FIXME: this should not be, but now due logger package I put it here to avoid errors. When adding new logger impl this should
			// be erased
			if fileExt != ".go" {
				return nil
			}

			content, err := fs.ReadFile(coreFS, currentPath)
			if err != nil {
				fmt.Println("Error reading file:", currentPath, err)
				return err
			}

			content = bytes.Replace(content, []byte("\"github.com/manicar2093/gomancer/bootstrap"), []byte(fmt.Sprintf("\"%s", input.ModuleName)), -1)

			content, err = format.Source(content)
			if err != nil {
				return err
			}

			if err := os.WriteFile(path.Join(projectDir), content, 0755); err != nil {
				return err
			}
		}

		return nil
	})
}
