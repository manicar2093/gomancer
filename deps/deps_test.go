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
		container = deps.Init(testfixtures.ModelBinaryIdSuccess.ModuleInfo.Name, testfixtures.ModelBinaryIdSuccess.TransformedText.LowerNoSpaceCase)
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

	Describe("I18n", func() {
		It("contains all i18n usable callers", func() {
			file := jen.NewFile("i18n_test")

			Expect(container.I18n.Alias).To(Equal("i18n"))
			Expect(container.I18n.Path).To(Equal("github.com/invopop/ctxi18n/i18n"))
			container.Echo.ImportAlias(file)
		})
	})

	Describe("GookitValidate", func() {
		It("contains all gookit usable callers", func() {
			file := jen.NewFile("gookit_test")

			Expect(container.GookitValidate.Alias).To(Equal("validate"))
			Expect(container.GookitValidate.Path).To(Equal("github.com/gookit/validate"))
			container.Echo.ImportAlias(file)
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

		Describe("StrConv", func() {
			It("contains all strconv usable callers", func() {
				file := jen.NewFile("fmt_test")

				Expect(container.Std.StrConv.Alias).To(Equal("strconv"))
				Expect(container.Std.StrConv.Path).To(Equal("strconv"))
				container.Std.Fmt.ImportAlias(file)
			})
		})

		Describe("Errors", func() {
			It("contains all errors usable callers", func() {
				file := jen.NewFile("fmt_test")

				Expect(container.Std.Errors.Alias).To(Equal("errors"))
				Expect(container.Std.Errors.Path).To(Equal("errors"))
				container.Std.Errors.ImportAlias(file)
			})
		})

		Describe("Testing", func() {
			It("contains all testing usable callers", func() {
				file := jen.NewFile("testing_test")

				Expect(container.Std.Testing.Alias).To(Equal("testing"))
				Expect(container.Std.Testing.Path).To(Equal("testing"))
				container.Std.Testing.ImportAlias(file)
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

				Expect(container.Project.Core.Web.Alias).To(Equal("web"))
				Expect(container.Project.Core.Web.Path).To(Equal(testfixtures.TestPath + "/core/web"))
				container.Project.Core.Web.ImportAlias(file)

				Expect(container.Project.Core.Web.Layouts.Alias).To(Equal("layouts"))
				Expect(container.Project.Core.Web.Layouts.Path).To(Equal(testfixtures.TestPath + "/core/web/layouts"))
				container.Project.Core.Web.Layouts.ImportAlias(file)

				Expect(container.Project.Core.Web.Components.Alias).To(Equal("components"))
				Expect(container.Project.Core.Web.Components.Path).To(Equal(testfixtures.TestPath + "/core/web/components"))
				container.Project.Core.Web.Components.ImportAlias(file)

				Expect(container.Project.Core.Web.Components.Toast.Alias).To(Equal("toast"))
				Expect(container.Project.Core.Web.Components.Toast.Path).To(Equal(testfixtures.TestPath + "/core/web/components/toast"))
				container.Project.Core.Web.Components.Toast.ImportAlias(file)

				Expect(container.Project.Core.Web.Components.DateTime.Alias).To(Equal("datetime"))
				Expect(container.Project.Core.Web.Components.DateTime.Path).To(Equal(testfixtures.TestPath + "/core/web/components/datetime"))
				container.Project.Core.Web.Components.DateTime.ImportAlias(file)

				Expect(container.Project.Core.Web.Components.FormTag.Alias).To(Equal("formtag"))
				Expect(container.Project.Core.Web.Components.FormTag.Path).To(Equal(testfixtures.TestPath + "/core/web/components/formtag"))
				container.Project.Core.Web.Components.FormTag.ImportAlias(file)

				Expect(container.Project.Core.Web.Components.Link.Alias).To(Equal("link"))
				Expect(container.Project.Core.Web.Components.Link.Path).To(Equal(testfixtures.TestPath + "/core/web/components/link"))
				container.Project.Core.Web.Components.Link.ImportAlias(file)

				Expect(container.Project.Core.Web.Components.FormErrors.Alias).To(Equal("formerrors"))
				Expect(container.Project.Core.Web.Components.FormErrors.Path).To(Equal(testfixtures.TestPath + "/core/web/components/formerrors"))
				container.Project.Core.Web.Components.FormErrors.ImportAlias(file)

				Expect(container.Project.Core.Web.Components.Button.Alias).To(Equal("button"))
				Expect(container.Project.Core.Web.Components.Button.Path).To(Equal(testfixtures.TestPath + "/core/web/components/button"))
				container.Project.Core.Web.Components.Button.ImportAlias(file)

				Expect(container.Project.Core.Web.Components.Drawer.Alias).To(Equal("drawer"))
				Expect(container.Project.Core.Web.Components.Drawer.Path).To(Equal(testfixtures.TestPath + "/core/web/components/drawer"))
				container.Project.Core.Web.Components.Drawer.ImportAlias(file)

				Expect(container.Project.Core.Web.Components.Form.Alias).To(Equal("form"))
				Expect(container.Project.Core.Web.Components.Form.Path).To(Equal(testfixtures.TestPath + "/core/web/components/form"))
				container.Project.Core.Web.Components.Form.ImportAlias(file)

				Expect(container.Project.Core.Web.Components.Icon.Alias).To(Equal("icon"))
				Expect(container.Project.Core.Web.Components.Icon.Path).To(Equal(testfixtures.TestPath + "/core/web/components/icon"))
				container.Project.Core.Web.Components.Icon.ImportAlias(file)

				Expect(container.Project.Core.Web.Components.Input.Alias).To(Equal("input"))
				Expect(container.Project.Core.Web.Components.Input.Path).To(Equal(testfixtures.TestPath + "/core/web/components/input"))
				container.Project.Core.Web.Components.Input.ImportAlias(file)

				Expect(container.Project.Core.Web.Components.Label.Alias).To(Equal("label"))
				Expect(container.Project.Core.Web.Components.Label.Path).To(Equal(testfixtures.TestPath + "/core/web/components/label"))
				container.Project.Core.Web.Components.Label.ImportAlias(file)

				Expect(container.Project.Core.Web.Components.Pagination.Alias).To(Equal("pagination"))
				Expect(container.Project.Core.Web.Components.Pagination.Path).To(Equal(testfixtures.TestPath + "/core/web/components/pagination"))
				container.Project.Core.Web.Components.Pagination.ImportAlias(file)

				Expect(container.Project.Core.Web.Components.Popover.Alias).To(Equal("popover"))
				Expect(container.Project.Core.Web.Components.Popover.Path).To(Equal(testfixtures.TestPath + "/core/web/components/popover"))
				container.Project.Core.Web.Components.Popover.ImportAlias(file)

				Expect(container.Project.Core.Web.Components.SelectBox.Alias).To(Equal("selectbox"))
				Expect(container.Project.Core.Web.Components.SelectBox.Path).To(Equal(testfixtures.TestPath + "/core/web/components/selectbox"))
				container.Project.Core.Web.Components.SelectBox.ImportAlias(file)

				Expect(container.Project.Core.Web.Components.Table.Alias).To(Equal("table"))
				Expect(container.Project.Core.Web.Components.Table.Path).To(Equal(testfixtures.TestPath + "/core/web/components/table"))
				container.Project.Core.Web.Components.Table.ImportAlias(file)

				Expect(container.Project.Core.Web.Components.Toggle.Alias).To(Equal("toggle"))
				Expect(container.Project.Core.Web.Components.Toggle.Path).To(Equal(testfixtures.TestPath + "/core/web/components/toggle"))
				container.Project.Core.Web.Components.Toggle.ImportAlias(file)

				Expect(container.Project.Pkg.TestFunc.Alias).To(Equal("testfunc"))
				Expect(container.Project.Pkg.TestFunc.Path).To(Equal(testfixtures.TestPath + "/pkg/testfunc"))
				container.Project.Pkg.TestFunc.ImportAlias(file)
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

				Expect(container.Project.Cmd.Service.Alias).To(Equal("service"))
				Expect(container.Project.Cmd.Service.Path).To(Equal(testfixtures.TestPath + "/cmd/service"))
				container.Project.Cmd.Service.ImportAlias(file)

				Expect(container.Project.Cmd.Service.Controllers.Alias).To(Equal("controllers"))
				Expect(container.Project.Cmd.Service.Controllers.Path).To(Equal(testfixtures.TestPath + "/cmd/service/controllers"))
				container.Project.Cmd.Service.Controllers.ImportAlias(file)

				Expect(container.Project.Cmd.Service.Controllers.ModelPages.Alias).To(Equal("posttestpages"))
				Expect(container.Project.Cmd.Service.Controllers.ModelPages.Path).To(Equal(testfixtures.TestPath + "/cmd/service/controllers/posttestpages"))
				container.Project.Cmd.Service.Controllers.ModelPages.ImportAlias(file)

				Expect(container.Project.Cmd.Service.Controllers.InitPages.Alias).To(Equal("initpages"))
				Expect(container.Project.Cmd.Service.Controllers.InitPages.Path).To(Equal(testfixtures.TestPath + "/cmd/service/controllers/initpages"))
				container.Project.Cmd.Service.Controllers.ImportAlias(file)
			})
		})
	})

	Describe("Ginkgo", func() {
		It("contains all ginkgo usable callers", func() {
			file := jen.NewFile("ginkgo_test")

			Expect(container.Ginkgo.Alias).To(Equal("."))
			Expect(container.Ginkgo.Path).To(Equal("github.com/onsi/ginkgo/v2"))
			container.Ginkgo.ImportAlias(file)
		})

	})

	Describe("Gomega", func() {
		It("contains all gomega usable callers", func() {
			file := jen.NewFile("gomega_test")

			Expect(container.Gomega.Alias).To(Equal("."))
			Expect(container.Gomega.Path).To(Equal("github.com/onsi/gomega"))
			container.Gomega.ImportAlias(file)

			Expect(container.Gomega.GStruct.Alias).To(Equal("gstruct"))
			Expect(container.Gomega.GStruct.Path).To(Equal("github.com/onsi/gomega/gstruct"))
			container.Gomega.GStruct.ImportAlias(file)
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
			Expect(fileStr).To(ContainSubstring(`import pkg "github.com/example/package"`))
			Expect(fileStr).To(ContainSubstring(`pkg.SomeFunction()`))
		})
	})
})
