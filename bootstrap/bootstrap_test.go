package bootstrap_test

import (
	"github.com/manicar2093/gomancer/bootstrap"
	"github.com/manicar2093/gomancer/testmatchers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
	"path"
)

var _ = Describe("Main", Ordered, func() {

	var (
		testPath    = "test"
		dirWithPath = func(input string) string {
			return path.Join(testPath, input)
		}
	)

	Describe("InitProject", func() {

		BeforeAll(func() {
			var input = bootstrap.InitProjectInput{
				ModuleName: testPath,
			}

			got, err := bootstrap.InitProject(input)

			Expect(err).ToNot(HaveOccurred())
			Expect(got).ToNot(BeZero())
		})

		AfterAll(func() {
			if err := os.RemoveAll(testPath); err != nil {
				GinkgoT().Log(err)
			}
		})

		It("creates .env file", func() {
			content := `DATABASE_URL="postgresql://development:development@localhost:5432/test_dev?sslmode=disable"
ENVIRONMENT=dev
PORT=3000
`
			Expect(dirWithPath(".env")).Should(testmatchers.BeAnExistingFileAndEqualString(content))
		})

		It("creates package.json file", func() {
			content := `{
    "devDependencies": {
        "prisma": "^6.3.0"
    },
    "dependencies": {
        "@prisma/client": "^6.3.0"
    },
    "prisma": {
        "schema": "./prisma"
    }
}
`
			Expect(dirWithPath("package.json")).Should(testmatchers.BeAnExistingFileAndEqualString(content))

		})

		It("creates prisma/schema/schema.prisma file", func() {
			content := `// This is your Prisma schema file,
// learn more about it in the docs: https://pris.ly/d/prisma-schema

generator client {
    provider        = "prisma-client-js"
}

datasource db {
    provider = "postgresql"
    url      = env("DATABASE_URL")
}
`
			Expect(dirWithPath("prisma/schema/schema.prisma")).Should(testmatchers.BeAnExistingFileAndEqualString(content))
		})

		It("creates internal/domain/models/init.go file", func() {
			content := `// Package models contains all your models that represents your tables to go structs.
// Models is a package shared by all your app so that is way this package
// exists; to avoid circular dependencies
package models
`
			Expect(dirWithPath("internal/domain/models/init.go")).Should(testmatchers.BeAnExistingFileAndEqualString(content))
		})

		It("creates cmd/api/controllers/init.go file", func() {
			content := `package controllers

import (
    "github.com/labstack/echo/v4"
    "net/http"
)

type InitController struct {}

func NewInitController() *InitController {
    return &InitController{}
}

func (c *InitController) SetUpRoutes(group *echo.Group) {
    group.GET("/initial", c.GetHandler)
}

func (c *InitController) GetHandler(ctx echo.Context) error {
    return ctx.String(http.StatusOK, "Hello from your new API :D")
}
`
			Expect(dirWithPath("cmd/api/controllers/init.go")).Should(testmatchers.BeAnExistingFileAndEqualString(content))
		})

		It("creates cmd/api/main.go file", func() {
			content := `package main

import (
    fmt	"fmt"
    echo	"github.com/labstack/echo/v4"
    middleware	"github.com/labstack/echo/v4/middleware"
    echoroutesview	"github.com/manicar2093/echoroutesview"
    core	"test/core"
    apperrors	"test/core/apperrors"
    converters	"test/core/converters"
    validator	"test/core/validator"
    logger	"test/core/logger"
    config	"test/pkg/config"
    controllers	"test/cmd/api/controllers"
)

func main() {
    var (
        echoInstance = echo.New()
        baseEndpoint = "/api/v1"
        baseGroup    = echoInstance.Group(baseEndpoint)
        conf         = converters.Must(core.ParseConfig[config.Config]())
    )
    logger.Config()
    echoInstance.Use(middleware.Logger())
    core.RegisterController(baseGroup, controllers.NewInitController())
    echoroutesview.RegisterRoutesViewer(echoInstance)

    echoInstance.HTTPErrorHandler = apperrors.HandlerWEcho
    echoInstance.Validator = validator.NewGooKitValidator()

    echoInstance.Logger.Fatal(echoInstance.Start(fmt.Sprintf(":%d", conf.Port)))
}
`
			Expect(dirWithPath("cmd/api/main.go")).Should(testmatchers.BeAnExistingFileAndEqualString(content))
		})

		It("creates pkg/generators/generators.go file", func() {
			content := `// Package generators contains all your models generators; functions used to generate mocked data and help you testing
package generators

import mapstructure	"github.com/go-viper/mapstructure/v2"

type testingI interface {
    Fatal(args ...any)
}

func decode(t testingI, args map[string]any, holder any) {
    if err := mapstructure.Decode(args, &holder); err != nil {
        t.Fatal(err)
    }
}
`
			Expect(dirWithPath("pkg/generators/generators.go")).Should(testmatchers.BeAnExistingFileAndEqualString(content))
		})

		It("creates pkg/config/config.go file", func() {
			content := `// Package config contains a struct with all your API configs
package config

import (
    core	"test/core"
    connections	"test/core/connections"
)

type Config struct {
    core.Config
    connections.DatabaseConnectionConfig
}
`
			Expect(dirWithPath("pkg/config/config.go")).Should(testmatchers.BeAnExistingFileAndEqualString(content))
		})

		It("creates pkg/versioning/version.go file", func() {
			content := `// Package versioning contains a constant with the current version of your API code
package versioning

const Version = "0.0.0"
`
			Expect(dirWithPath("pkg/versioning/version.go")).Should(testmatchers.BeAnExistingFileAndEqualString(content))
		})

		It("creates .cz.toml file", func() {
			content := `[tool]
[tool.commitizen]
name = "cz_conventional_commits"
version = "0.0.0"
tag_format = "v$version"
version_files = [
    ".cz.toml:version",
    "pkg/versioning/version.go:Version",
]
bump_message = "bump: Semantic Release Bot: Realease Version: $new_version 🤖🚀 [skip ci]"
`
			Expect(dirWithPath(".cz.toml")).Should(testmatchers.BeAnExistingFileAndEqualString(content))
		})

		It("creates .github/workflows/bump_version.yml file", func() {
			content := `name: Bump Version

on:
    push:
        branches:
            - main

jobs:
    bump-version:
        if: ${{!contains(github.event.head_commit.message, '[skip ci]')}}
        runs-on: ubuntu-latest
        name: "Bump version and create changelog with commitizen"
        steps:
            - name: Check out
                uses: actions/checkout@v2
                with:
                    token: "${{ secrets.PERSONAL_ACCESS_TOKEN }}"
                    fetch-depth: 0
            - name: Set up Go
                uses: actions/setup-go@v3
                with:
                    go-version: 1.24.x
            - name: Set GOBIN
                run: go env -w GOBIN=/usr/local/bin
            - name: Create bump and changelog
                uses: commitizen-tools/commitizen-action@master
                with:
                    github_token: ${{ secrets.PERSONAL_ACCESS_TOKEN }}
                    branch: main
`
			Expect(dirWithPath(".github/workflows/bump_version.yml")).Should(testmatchers.BeAnExistingFileAndEqualString(content))
		})

		It("creates go.mod file", func() {
			content := `module test

go 1.24
`
			Expect(dirWithPath("go.mod")).Should(testmatchers.BeAnExistingFileAndEqualString(content))
		})

		It("creates Taskfile.yml file", func() {
			content := `# https://taskfile.dev

version: '3'

tasks:
    build:
        desc: Build your API to deploy anywhere
        cmds:
            - go build -o .bin/api/server cmd/api/*.go
    fmt:
        desc: Format all your Golang and Prisma code
        cmds:
            - go fmt ./...
            - npx prisma format
    version:
        desc: Shows your current API version
        cmds:
            - cz version -p
    dev:
        desc: Start project with air using your .env file
        dotenv: ['.env']
        cmds:
            - air
    run:
        desc: Start project from build
        dotenv: ['.env']
        deps: [build]
        cmds:
            - ./.bin/api/server
`
			Expect(dirWithPath("Taskfile.yml")).Should(testmatchers.BeAnExistingFileAndEqualString(content))
		})

		It("creates .air.toml file", func() {
			content := `root = "."
tmp_dir = "tmp"

[build]
# Just plain old shell command. You could use make as well.
cmd = "task build"
# Binary file yields from cmd.
bin = "./.bin/api/server"
# Customize binary, can setup environment variables when run your app.
full_bin = ""
# Watch these filename extensions.
include_ext = ["go", "tpl", "tmpl", "html"]
# Ignore these filename extensions or directories.
exclude_dir = ["assets", "tmp", "vendor", "node_modules", "bin", "temporal", "prisma", "scripts", "static"]
# Watch these directories if you specified.
include_dir = []
# Watch these files.
include_file = []
# Exclude files.
exclude_file = []
# Exclude specific regular expressions.
exclude_regex = ["_test\\.go"]
# Exclude unchanged files.
exclude_unchanged = true
# Follow symlink for directories
follow_symlink = true
# This log file places in your tmp_dir.
log = "air.log"
# Poll files for changes instead of using fsnotify.
poll = false
# Poll interval (defaults to the minimum interval of 500ms).
poll_interval = 500 # ms
# It's not necessary to trigger build each time file changes if it's too frequent.
delay = 0 # ms
# Stop running old binary when build errors occur.
stop_on_error = true
# Send Interrupt signal before killing process (windows does not support this feature)
send_interrupt = false
# Delay after sending Interrupt signal
kill_delay = 500 # ms
# Rerun binary or not
rerun = false
# Delay after each executions
rerun_delay = 500
# Add additional arguments when running binary (bin/full_bin). Will run './tmp/main hello world'.
args_bin = []

[log]
# Show log time
time = false
# Only show main log (silences watcher, build, runner)
main_only = false

[color]
# Customize each part's color. If no color found, use the raw app log.
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"

[misc]
# Delete tmp directory on exit
clean_on_exit = true

[screen]
clear_on_rebuild = true
keep_scroll = true
`
			Expect(dirWithPath(".air.toml")).Should(testmatchers.BeAnExistingFileAndEqualString(content))
		})

		It("creates README.md file", func() {
			content := `# test

This is the initial code for the greatest project ever made.

## Tools

Yep, you need to install some tools to make this work:

- [Go](https://go.dev/): That's why you are here, right? :P
- [Taskfile](https://taskfile.dev/): A task runner used instead Makefile
- [Air](https://github.com/air-verse/air): Live reload tool for Go apps
- [Node](https://nodejs.org/es): Needed to run npx command

## Thank you!

Be yourself. Find joy by what you do. Happy coding :)
`
			Expect(dirWithPath("README.md")).Should(testmatchers.BeAnExistingFileAndEqualString(content))
		})

		It("creates core directory", func() {
			Expect(dirWithPath("core")).Should(BeADirectory())
			Expect(dirWithPath("core/apperrors")).Should(BeADirectory())
			Expect(dirWithPath("core/commonreq")).Should(BeADirectory())
			Expect(dirWithPath("core/connections")).Should(BeADirectory())
			Expect(dirWithPath("core/converters")).Should(BeADirectory())
			Expect(dirWithPath("core/echoer")).Should(BeADirectory())
			Expect(dirWithPath("core/env")).Should(BeADirectory())
			Expect(dirWithPath("core/httphealthcheck")).Should(BeADirectory())
			Expect(dirWithPath("core/logger")).Should(BeADirectory())
			Expect(dirWithPath("core/stages")).Should(BeADirectory())
			Expect(dirWithPath("core/validator")).Should(BeADirectory())
		})
	})
})
