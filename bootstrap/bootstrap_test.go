package bootstrap_test

import (
	"github.com/manicar2093/gomancer/testmatchers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
	"path"

	"github.com/manicar2093/gomancer/bootstrap"
)

var _ = Describe("Main", func() {

	var (
		testPath    = "test"
		dirWithPath = func(input string) string {
			return path.Join(testPath, input)
		}
	)

	AfterEach(func() {
		if err := os.RemoveAll("test"); err != nil {
			GinkgoT().Log(err)
		}
	})
	Describe("InitProject", func() {
		It("creates all needed directories and files to start a new project", func() {
			var input = bootstrap.InitProjectInput{
				ModuleName: "test",
			}

			err := bootstrap.InitProject(input)

			Expect(err).ToNot(HaveOccurred())

			Expect(dirWithPath(".env")).Should(testmatchers.BeAnExistingFileAndEqualString(`DATABASE_URL="postgresql://development:development@localhost:5432/test_dev?sslmode=disable"
ENVIRONMENT=dev
PORT=3000
`))
			Expect(dirWithPath("package.json")).Should(testmatchers.BeAnExistingFileAndEqualString(`{
    "devDependencies": {
        "prisma": "^6.3.0"
    },
    "dependencies": {
        "@prisma/client": "^6.3.0"
    }
}
`))

			Expect(dirWithPath("prisma/schema/schema.prisma")).Should(testmatchers.BeAnExistingFileAndEqualString(`// This is your Prisma schema file,
// learn more about it in the docs: https://pris.ly/d/prisma-schema

generator client {
    provider        = "prisma-client-js"
    previewFeatures = ["prismaSchemaFolder"]
}

datasource db {
    provider = "postgresql"
    url      = env("DATABASE_URL")
}
`))

			Expect(dirWithPath("internal/domain/models/init.go")).Should(testmatchers.BeAnExistingFileAndEqualString(`// Package models contains all your models that represents your tables to go structs.
// Models is a package shared by all your app so that is way this package
// exists; to avoid circular dependencies
package models
`))
			Expect(dirWithPath("cmd/api/controllers/init.go")).Should(testmatchers.BeAnExistingFileAndEqualString(`package controllers

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
`))
			Expect(dirWithPath("cmd/api/main.go")).Should(testmatchers.BeAnExistingFileAndEqualString(`package main

import (
    "fmt"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/manicar2093/echoroutesview"
    "github.com/manicar2093/winter"
    "github.com/manicar2093/winter/apperrors"
    "github.com/manicar2093/winter/converters"
    "github.com/manicar2093/winter/validator"
    "test/pkg/config"
    "test/cmd/api/controllers"
)

func main() {
    var (
        echoInstance = echo.New()
        baseEndpoint = "/api/v1"
        baseGroup    = echoInstance.Group(baseEndpoint)
        conf         = converters.Must(winter.ParseConfig[config.Config]())
    )
    echoInstance.Use(middleware.Logger())
    winter.RegisterController(baseGroup, controllers.NewInitController())
    echoroutesview.RegisterRoutesViewer(echoInstance)

    echoInstance.HTTPErrorHandler = apperrors.HandlerWEcho
    echoInstance.Validator = validator.NewGooKitValidator()

    echoInstance.Logger.Fatal(echoInstance.Start(fmt.Sprintf(":%d", conf.Port)))
}
`))
			Expect(dirWithPath("pkg/generators/generators.go")).Should(testmatchers.BeAnExistingFileAndEqualString(`// Package generators contains all your models generators; functions used to generate mocked data and help you testing
package generators

import "github.com/go-viper/mapstructure/v2"

type testingI interface {
    mock.TestingT
    Fatal(args ...any)
    Cleanup(func())
}

func decode(t testingI, args map[string]any, holder any) {
    if err := mapstructure.Decode(args, &holder); err != nil {
        t.Fatal(err)
    }
}
`))
			Expect(dirWithPath("pkg/config/config.go")).Should(testmatchers.BeAnExistingFileAndEqualString(`// Package config contains a struct with all your API configs
package config

import (
    "github.com/manicar2093/winter"
    "github.com/manicar2093/winter/connections"
)

type Config struct {
    winter.Config
    connections.DatabaseConnectionConfig
}
`))
			Expect(dirWithPath("pkg/versioning/version.go")).Should(testmatchers.BeAnExistingFileAndEqualString(`// Package versioning contains a constant with the current version of your API code
package versioning

const Version = "0.0.0"
`))
			Expect(dirWithPath(".cz.toml")).Should(testmatchers.BeAnExistingFileAndEqualString(`[tool]
[tool.commitizen]
name = "cz_conventional_commits"
version = "0.0.0"
tag_format = "v$version"
version_files = [
    ".cz.toml:version",
    "pkg/versioning/version.go:Version",
]
bump_message = "bump: Semantic Release Bot: Realease Version: $new_version 🤖🚀 [skip ci]"
`))
			Expect(dirWithPath(".github/workflows/bump_version.yml")).Should(testmatchers.BeAnExistingFileAndEqualString(`name: Bump Version

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
                    go-version: 1.23.x
            - name: Set GOBIN
                run: go env -w GOBIN=/usr/local/bin
            - name: Create bump and changelog
                uses: commitizen-tools/commitizen-action@master
                with:
                    github_token: ${{ secrets.PERSONAL_ACCESS_TOKEN }}
                    branch: main
`))
			Expect(dirWithPath("go.mod")).Should(testmatchers.BeAnExistingFileAndEqualString(`module test

go 1.23
`))
			Expect(dirWithPath("Taskfile.yml")).Should(testmatchers.BeAnExistingFileAndEqualString(`# https://taskfile.dev

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
`))
			Expect(dirWithPath(".air.toml")).Should(testmatchers.BeAnExistingFileAndEqualString(`root = "."
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
`))

			Expect(dirWithPath("README.md")).Should(testmatchers.BeAnExistingFileAndEqualString(`# test

This is the initial code for the greatest project ever made.

## Tools

Yep, you need to install some tools to make this work:

- [Go](https://go.dev/): That's why you are here, right? :P
- [Taskfile](https://taskfile.dev/): A task runner used instead Makefile
- [Air](https://github.com/air-verse/air): Live reload tool for Go apps
- [Node](https://nodejs.org/es): Needed to run npx command

## Thank you!

Be yourself. Find joy by what you do. Happy coding :)
`))
		})
	})
})
