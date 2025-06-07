package deps_test

import (
	"github.com/dave/jennifer/jen"
	"github.com/manicar2093/gomancer/deps"
	"github.com/manicar2093/gomancer/testfixtures"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Deps", func() {
	Expect(true).To(BeTrue())

	var container deps.Container

	BeforeEach(func() {
		container = deps.Init(testfixtures.ModelBinaryIdSuccess.ModuleInfo.Name)
	})

	Describe("InCreation", func() {
		It("creates model package dependency", func() {
			file := jen.NewFile("in_creation_test")

			inCreation := deps.InCreation(testfixtures.ModelBinaryIdSuccess.ModuleInfo.Name, testfixtures.ModelBinaryIdSuccess.PackageEntityName)
			Expect(inCreation.Alias).To(Equal(testfixtures.ModelBinaryIdSuccess.PackageEntityName))
			Expect(inCreation.Path).To(Equal(testfixtures.TestPath + "/internal/" + testfixtures.ModelBinaryIdSuccess.PackageEntityName))
			inCreation.ImportAlias(file)
		})
	})

	Describe("Echo", func() {
		It("contains all echo usable callers", func() {
			file := jen.NewFile("echo_test")

			Expect(container.Echo.Alias).To(Equal("echo"))
			Expect(container.Echo.Path).To(Equal("github.com/labstack/echo/v4"))
			container.Echo.ImportAlias(file)

			Expect(container.Echo.Middlewares.Alias).To(Equal("middleware"))
			Expect(container.Echo.Middlewares.Path).To(Equal("github.com/labstack/echo/v4/middleware"))
			container.Echo.Middlewares.ImportAlias(file)
		})
	})

	Describe("EchoRoutesView", func() {
		It("contains all echoroutesview usable callers", func() {
			file := jen.NewFile("echoroutesview_test")

			Expect(container.EchoRoutesView.Alias).To(Equal("echoroutesview"))
			Expect(container.EchoRoutesView.Path).To(Equal("github.com/manicar2093/echoroutesview"))
			container.EchoRoutesView.ImportAlias(file)
		})
	})

	Describe("Std", func() {
		Describe("Net", func() {
			It("contains all net usable callers", func() {
				file := jen.NewFile("net_test")

				Expect(container.Std.Net.Alias).To(Equal("net"))
				Expect(container.Std.Net.Path).To(Equal("net"))
				container.Std.Net.ImportAlias(file)
			})

			Describe("Http", func() {
				It("contains all http usable callers", func() {
					file := jen.NewFile("http_test")

					Expect(container.Std.Net.Http.Alias).To(Equal("http"))
					Expect(container.Std.Net.Http.Path).To(Equal("net/http"))
					container.Std.Net.Http.ImportAlias(file)
				})
			})
		})

		Describe("Fmt", func() {
			It("contains all fmt usable callers", func() {
				file := jen.NewFile("fmt_test")

				Expect(container.Std.Fmt.Alias).To(Equal("fmt"))
				Expect(container.Std.Fmt.Path).To(Equal("fmt"))
				container.Std.Fmt.ImportAlias(file)
			})
		})

		Describe("Time", func() {
			It("contains all time usable callers", func() {
				file := jen.NewFile("time_test")

				Expect(container.Std.Time.Alias).To(Equal("time"))
				Expect(container.Std.Time.Path).To(Equal("time"))
				container.Std.Time.ImportAlias(file)
			})
		})

		Describe("Maps", func() {
			It("contains all maps usable callers", func() {
				file := jen.NewFile("maps_test")

				Expect(container.Std.Maps.Alias).To(Equal("maps"))
				Expect(container.Std.Maps.Path).To(Equal("maps"))
				container.Std.Maps.ImportAlias(file)
			})
		})

		Describe("Slices", func() {
			It("contains all slices usable callers", func() {
				file := jen.NewFile("slices_test")

				Expect(container.Std.Slices.Alias).To(Equal("slices"))
				Expect(container.Std.Slices.Path).To(Equal("slices"))
				container.Std.Slices.ImportAlias(file)
			})
		})
	})

	Describe("Gorm", func() {
		It("contains all gorm usable callers", func() {
			file := jen.NewFile("gorm_test")

			Expect(container.Gorm.Alias).To(Equal("gorm"))
			Expect(container.Gorm.Path).To(Equal("gorm.io/gorm"))
			container.Gorm.ImportAlias(file)

			Expect(container.Gorm.Clause.Alias).To(Equal("clause"))
			Expect(container.Gorm.Clause.Path).To(Equal("gorm.io/gorm/clause"))
			container.Gorm.Clause.ImportAlias(file)
		})
	})

	Describe("MapStructure", func() {
		It("contains all mapstructure usable callers", func() {
			file := jen.NewFile("mapstructure_test")

			Expect(container.MapStructure.Alias).To(Equal("mapstructure"))
			Expect(container.MapStructure.Path).To(Equal("github.com/go-viper/mapstructure/v2"))
			container.MapStructure.ImportAlias(file)
		})
	})

	Describe("GoFakeIt", func() {
		It("contains all gofakeit usable callers", func() {
			file := jen.NewFile("gofakeit_test")

			Expect(container.GoFakeIt.Alias).To(Equal("gofakeit"))
			Expect(container.GoFakeIt.Path).To(Equal("github.com/brianvoe/gofakeit/v7"))
			container.GoFakeIt.ImportAlias(file)
		})
	})

	Describe("Uuid", func() {
		It("contains all uuid usable callers", func() {
			file := jen.NewFile("uuid_test")

			Expect(container.Uuid.Alias).To(Equal("uuid"))
			Expect(container.Uuid.Path).To(Equal("github.com/google/uuid"))
			container.Uuid.ImportAlias(file)
		})
	})

	Describe("Goption", func() {
		It("contains all goption usable callers", func() {
			file := jen.NewFile("goption_test")

			Expect(container.Goption.Alias).To(Equal("goption"))
			Expect(container.Goption.Path).To(Equal("github.com/manicar2093/goption"))
			container.Goption.ImportAlias(file)
		})
	})

	Describe("GormPager", func() {
		It("contains all gormpager usable callers", func() {
			file := jen.NewFile("gormpager_test")

			Expect(container.GormPager.Alias).To(Equal("gormpager"))
			Expect(container.GormPager.Path).To(Equal("github.com/manicar2093/gormpager"))
			container.GormPager.ImportAlias(file)
		})
	})

	Describe("UDecimal", func() {
		It("contains all udecimal usable callers", func() {
			file := jen.NewFile("udecimal_test")

			Expect(container.UDecimal.Alias).To(Equal("udecimal"))
			Expect(container.UDecimal.Path).To(Equal("github.com/quagmt/udecimal"))
			container.UDecimal.ImportAlias(file)
		})
	})

	Describe("Project", func() {
		Context("Pkg", func() {
			It("contains all pkg usable callers", func() {
				file := jen.NewFile("pkg_test")

				Expect(container.Project.Pkg.Alias).To(Equal("pkg"))
				Expect(container.Project.Pkg.Path).To(Equal(testfixtures.TestPath + "/pkg"))
				container.Project.Pkg.ImportAlias(file)

				Expect(container.Project.Pkg.Config.Alias).To(Equal("config"))
				Expect(container.Project.Pkg.Config.Path).To(Equal(testfixtures.TestPath + "/pkg/config"))
				container.Project.Pkg.Config.ImportAlias(file)
			})
		})

		Context("when testing Core", func() {
			It("contains all core usable callers", func() {
				file := jen.NewFile("core_test")

				Expect(container.Project.Core.Alias).To(Equal("core"))
				Expect(container.Project.Core.Path).To(Equal(testfixtures.TestPath + "/core"))
				container.Project.Core.ImportAlias(file)

				Expect(container.Project.Core.Converters.Alias).To(Equal("converters"))
				Expect(container.Project.Core.Converters.Path).To(Equal(testfixtures.TestPath + "/core/converters"))
				container.Project.Core.Converters.ImportAlias(file)

				Expect(container.Project.Core.Validator.Alias).To(Equal("validator"))
				Expect(container.Project.Core.Validator.Path).To(Equal(testfixtures.TestPath + "/core/validator"))
				container.Project.Core.Validator.ImportAlias(file)

				Expect(container.Project.Core.Logger.Alias).To(Equal("logger"))
				Expect(container.Project.Core.Logger.Path).To(Equal(testfixtures.TestPath + "/core/logger"))
				container.Project.Core.Logger.ImportAlias(file)

				Expect(container.Project.Core.Connections.Alias).To(Equal("connections"))
				Expect(container.Project.Core.Connections.Path).To(Equal(testfixtures.TestPath + "/core/connections"))
				container.Project.Core.Connections.ImportAlias(file)

				Expect(container.Project.Core.CommonReq.Alias).To(Equal("commonreq"))
				Expect(container.Project.Core.CommonReq.Path).To(Equal(testfixtures.TestPath + "/core/commonreq"))
				container.Project.Core.CommonReq.ImportAlias(file)

				Expect(container.Project.Core.AppErrors.Alias).To(Equal("apperrors"))
				Expect(container.Project.Core.AppErrors.Path).To(Equal(testfixtures.TestPath + "/core/apperrors"))
				container.Project.Core.AppErrors.ImportAlias(file)
			})
		})

		Context("when testing Internal", func() {
			It("contains all internal usable callers", func() {
				file := jen.NewFile("internal_test")

				Expect(container.Project.Internal.Alias).To(Equal("internal"))
				Expect(container.Project.Internal.Path).To(Equal(testfixtures.TestPath + "/internal"))
				container.Project.Internal.ImportAlias(file)

				Expect(container.Project.Internal.Domain.Alias).To(Equal("domain"))
				Expect(container.Project.Internal.Domain.Path).To(Equal(testfixtures.TestPath + "/internal/domain"))
				container.Project.Internal.Domain.ImportAlias(file)

				Expect(container.Project.Internal.Domain.Models.Alias).To(Equal("models"))
				Expect(container.Project.Internal.Domain.Models.Path).To(Equal(testfixtures.TestPath + "/internal/domain/models"))
				container.Project.Internal.Domain.Models.ImportAlias(file)
			})
		})

		Context("when testing Cmd", func() {
			It("contains all cmd usable callers", func() {
				file := jen.NewFile("cmd_test")

				Expect(container.Project.Cmd.Alias).To(Equal("cmd"))
				Expect(container.Project.Cmd.Path).To(Equal(testfixtures.TestPath + "/cmd"))
				container.Project.Cmd.ImportAlias(file)

				Expect(container.Project.Cmd.Service.Alias).To(Equal("api"))
				Expect(container.Project.Cmd.Service.Path).To(Equal(testfixtures.TestPath + "/cmd/api"))
				container.Project.Cmd.Service.ImportAlias(file)

				Expect(container.Project.Cmd.Service.Controllers.Alias).To(Equal("controllers"))
				Expect(container.Project.Cmd.Service.Controllers.Path).To(Equal(testfixtures.TestPath + "/cmd/api/controllers"))
				container.Project.Cmd.Service.Controllers.ImportAlias(file)
			})
		})
	})

	Describe("Dependency", func() {
		It("should correctly generate import string", func() {
			dep := deps.Dependency{
				Path:  "github.com/example/package",
				Alias: "pkg",
			}

			Expect(dep.GenerateImportString()).To(Equal("pkg\t\"github.com/example/package\""))
		})

		It("should correctly import alias to file", func() {
			dep := deps.Dependency{
				Path:  "github.com/example/package",
				Alias: "pkg",
			}
			file := jen.NewFile("test")

			dep.ImportAlias(file)

			// Add a statement that uses the imported package to force it to be included in the output
			file.Func().Id("TestFunc").Params().Block(
				jen.Qual(dep.Path, "SomeFunction").Call(),
			)

			// Get the string representation of the file and check if it contains the expected import
			fileStr := file.GoString()
			Expect(fileStr).To(ContainSubstring(`import "github.com/example/package"`))
			Expect(fileStr).To(ContainSubstring(`pkg.SomeFunction()`))
		})
	})
})
