package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/manicar2093/gomancer/bootstrap"
	"github.com/spf13/cobra"
)

var BootstrapCmd = &cobra.Command{
	Use:   "new",
	Short: "Starts a new project",
	Long: `Does Golang needs a fullstack framework?...I don't thing so. 

Go has many cool packages waiting there to be used and Gomancer gather them into a code generator to be blazingly fast at coding :D

It uses Prisma, Echo, Winter and many others packages to create a friendly way to develop API (by the moment)

Project structure:

<project_name>/
├── .air.toml
├── .cz.toml
├── .env
├── README.md
├── Taskfile.yml
├── go.mod
├── package.json
├── .github
│   └── workflows
│       └── bump_version.yml
├── cmd
│   └── api
│       ├── main.go
│       └── controllers
│           └── init.go
├── internal
│   └── domain
│       └── models
│           └── init.go
├── pkg
│   ├── config
│   │   └── config.go
│   ├── generators
│   │   └── generators.go
│   └── versioning
│       └── version.go
└── prisma
    └── schema
        └── schema.prisma
`,
	Example: "gomancer bootstrap github.com/great-dev/revolutionary",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectDirName, err := bootstrap.InitProject(bootstrap.InitProjectInput{
			ModuleName: args[0],
		})

		log.Infof(`
✅ Ready!
🏁 Next steps:

	➡️ Run 'cd %s'
	➡️ Run 'go mod download' to install deps
	➡️ Run 'npm i' to install prisma deps
	➡️ Run 'gomancer gen' command to create your first API resource
	➡️ Enjoy! 😎
`, projectDirName)

		return err
	},
}
