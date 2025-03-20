package bootstrap_test

import (
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
		readWithPath = func(filename string) (string, error) {
			data, err := os.ReadFile(dirWithPath(filename))
			return string(data), err
		}
	)

	AfterEach(func() {
		//if err := os.RemoveAll("test"); err != nil {
		//	GinkgoT().Log(err)
		//}
	})
	Describe("InitProject", func() {
		It("creates all needed directories and files to start a new project", func() {
			var input = bootstrap.InitProjectInput{
				ModuleName: "test",
			}

			err := bootstrap.InitProject(input)

			Expect(err).ToNot(HaveOccurred())
			Expect(dirWithPath(".env.dev")).Should(BeAnExistingFile())
			Expect(readWithPath(".env.dev")).To(ContainSubstring(`DATABASE_URL="postgresql://development:development@localhost:5432/test_dev?sslmode=disable"
ENVIRONMENT=dev
PORT=3000
`))
			Expect(dirWithPath("package.json")).Should(BeAnExistingFile())
			Expect(readWithPath("package.json")).Should(ContainSubstring(`{
    "devDependencies": {
        "prisma": "^6.3.0"
    },
    "dependencies": {
        "@prisma/client": "^6.3.0"
    }
}
`))
			Expect(dirWithPath("prisma/schema/schema.prisma")).Should(BeAnExistingFile())
			Expect(readWithPath("prisma/schema/schema.prisma")).Should(ContainSubstring(`// This is your Prisma schema file,
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
			Expect(dirWithPath("internal/domain/models/init.go")).Should(BeAnExistingFile())
			Expect(readWithPath("internal/domain/models/init.go")).Should(ContainSubstring(`// Package models contains all your models that represents your tables to go structs.
// Models is a package shared by all your app so that is way this package
// exists; to avoid circular dependencies
package models
`))
			Expect(dirWithPath("cmd/api/controllers/init.go")).Should(BeAnExistingFile())
			Expect(readWithPath("cmd/api/controllers/init.go")).Should(ContainSubstring(`package controllers

import "github.com/labstack/echo/v4"

type InitController struct {}

func NewInitController() *InitController {
    return &InitController{}
}

func (c *InitController) SetUpRoutes(group *echo.Group) {
    group.GET("/initial", c.GetHandler)
}

func (c *InitController) GetHandler(ctx echo.Context) error {
    return ctx.String("Hello from your new API :D")
}
`))
			Expect(dirWithPath("cmd/api/main.go")).Should(BeAnExistingFile())
			Expect(readWithPath("cmd/api/main.go")).Should(ContainSubstring(`package main

import (
    "github.com/labstack/echo/v4"
    "github.com/manicar2093/winter"
    "github.com/manicar2093/winter/connections"
    "test/pkg/config"
    "test/cmd/api/controllers"
    "github.com/manicar2093/winter/apperrors"
    "github.com/manicar2093/winter/validator"
    "github.com/manicar2093/echoroutesview"
)

func main() {
    var (
        echoInstance    = echo.New()
        baseEndpoint    = "/api/v1"
        conf            = converters.Must(winter.ParseConfig[config.Config]())
        dbConn          = connections.GetGormConnection(conf.DatabaseConnectionConfig)
    )
    echoInstance.Use(middleware.Logger())
    winter.RegisterController(baseEndpoint, controllers.NewInitController())
    echoroutesview.RegisterRoutesViewer(e)

    e.HTTPErrorHandler = apperrors.HandlerWEcho
    e.Validator = validator.NewGooKitValidator()

    e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
`))
			Expect(dirWithPath("pkg/generators/generators.go")).Should(BeAnExistingFile())
			Expect(readWithPath("pkg/generators/generators.go")).Should(ContainSubstring(`// Package generators contains all your models generators; functions used to generate mocked data and help you testing
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
			Expect(dirWithPath("pkg/config/config.go")).Should(BeAnExistingFile())
			Expect(readWithPath("pkg/config/config.go")).Should(ContainSubstring(`// Package config contains a struct with all your API configs
package config

import (
    "github.com/caarlos0/env/v10"
    "github.com/manicar2093/winter"
    "github.com/manicar2093/winter/connections"
)

type Config struct {
    winter.Config
    connections.DatabaseConnectionConfig
}`))
		})
	})
})
