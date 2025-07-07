package deps

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"path"
)

func Init(moduleName string) Container {
	return Container{
		Echo: Echo{
			Dependency: Dependency{
				Path:  "github.com/labstack/echo/v4",
				Alias: "echo",
			},
			Middlewares: Dependency{
				Path:  "github.com/labstack/echo/v4/middleware",
				Alias: "middleware",
			},
		},
		EchoRoutesView: EchoRoutesView{
			Dependency{
				Path:  "github.com/manicar2093/echoroutesview",
				Alias: "echoroutesview",
			},
		},
		Std: Std{
			Net: Net{
				Dependency: Dependency{
					Path:  "net",
					Alias: "net",
				},
				Http: Dependency{
					Path:  "net/http",
					Alias: "http",
				},
			},
			Fmt: Fmt{
				Dependency{
					Path:  "fmt",
					Alias: "fmt",
				},
			},
			Time: Dependency{
				Path:  "time",
				Alias: "time",
			},
			Maps: Dependency{
				Path:  "maps",
				Alias: "maps",
			},
			Slices: Dependency{
				Path:  "slices",
				Alias: "slices",
			},
		},
		Gorm: Gorm{
			Dependency: Dependency{
				Path:  "gorm.io/gorm",
				Alias: "gorm",
			},
			Clause: Dependency{
				Path:  "gorm.io/gorm/clause",
				Alias: "clause",
			},
		},
		I18n: I18n{
			Dependency: Dependency{
				Path:  "github.com/invopop/ctxi18n/i18n",
				Alias: "i18n",
			},
		},
		GookitValidate: GookitValidate{
			Dependency: Dependency{
				Path:  "github.com/gookit/validate",
				Alias: "validate",
			},
		},
		MapStructure: Dependency{
			Path:  "github.com/go-viper/mapstructure/v2",
			Alias: "mapstructure",
		},
		GoFakeIt: Dependency{
			Path:  "github.com/brianvoe/gofakeit/v7",
			Alias: "gofakeit",
		},
		Uuid: Dependency{
			Path:  "github.com/google/uuid",
			Alias: "uuid",
		},
		Goption: Dependency{
			Path:  "github.com/manicar2093/goption",
			Alias: "goption",
		},
		GormPager: Dependency{
			Path:  "github.com/manicar2093/gormpager",
			Alias: "gormpager",
		},
		UDecimal: Dependency{
			Path:  "github.com/quagmt/udecimal",
			Alias: "udecimal",
		},
		Project: Project{
			Pkg: Pkg{
				Dependency: Dependency{
					Path:  buildPath(moduleName, "pkg"),
					Alias: "pkg",
				},
				Config: Dependency{
					Path:  buildPath(moduleName, "pkg", "config"),
					Alias: "config",
				},
			},
			Core: Core{
				Dependency: Dependency{
					Path:  buildPath(moduleName, "core"),
					Alias: "core",
				},
				Converters: Dependency{
					Path:  buildPath(moduleName, "core", "converters"),
					Alias: "converters",
				},
				Validator: Dependency{
					Path:  buildPath(moduleName, "core", "validator"),
					Alias: "validator",
				},
				Logger: Dependency{
					Path:  buildPath(moduleName, "core", "logger"),
					Alias: "logger",
				},
				Connections: Dependency{
					Path:  buildPath(moduleName, "core", "connections"),
					Alias: "connections",
				},
				CommonReq: Dependency{
					Path:  buildPath(moduleName, "core", "commonreq"),
					Alias: "commonreq",
				},
				AppErrors: Dependency{
					Path:  buildPath(moduleName, "core", "apperrors"),
					Alias: "apperrors",
				},
				CoreTpls: CoreTpls{
					Dependency: Dependency{
						Path:  buildPath(moduleName, "core", "coretpls"),
						Alias: "coretpls",
					},
					Toast: Dependency{
						Path:  buildPath(moduleName, "core", "coretpls", "toast"),
						Alias: "toast",
					},
					DateTime: Dependency{
						Path:  buildPath(moduleName, "core", "coretpls", "datetime"),
						Alias: "datetime",
					},
					FormTag: Dependency{
						Path:  buildPath(moduleName, "core", "coretpls", "formtag"),
						Alias: "formtag",
					},
					Link: Dependency{
						Path:  buildPath(moduleName, "core", "coretpls", "link"),
						Alias: "link",
					},
				},
			},
			Internal: Internal{
				Dependency: Dependency{
					Path:  buildPath(moduleName, "internal"),
					Alias: "internal",
				},
				Domain: Domain{
					Dependency: Dependency{
						Path:  buildPath(moduleName, "internal", "domain"),
						Alias: "domain",
					},
					Models: Dependency{
						Path:  buildPath(moduleName, "internal", "domain", "models"),
						Alias: "models",
					},
				},
			},
			Cmd: Cmd{
				Dependency: Dependency{
					Path:  buildPath(moduleName, "cmd"),
					Alias: "cmd",
				},
				Service: Service{
					Dependency: Dependency{
						Path:  buildPath(moduleName, "cmd", "service"),
						Alias: "service",
					},
					Controllers: Controllers{
						Dependency: Dependency{
							Path:  buildPath(moduleName, "cmd", "service", "controllers"),
							Alias: "controllers",
						},
						InitPages: Dependency{
							Path:  buildPath(moduleName, "cmd", "service", "controllers", "initpages"),
							Alias: "initpages",
						},
					},
					Ui: Ui{
						Components: Components{
							Button: Dependency{
								Path:  buildPath(moduleName, "cmd", "service", "ui", "components", "button"),
								Alias: "button",
							},
							Drawer: Dependency{
								Path:  buildPath(moduleName, "cmd", "service", "ui", "components", "drawer"),
								Alias: "drawer",
							},
							Form: Dependency{
								Path:  buildPath(moduleName, "cmd", "service", "ui", "components", "form"),
								Alias: "form",
							},
							Icon: Dependency{
								Path:  buildPath(moduleName, "cmd", "service", "ui", "components", "icon"),
								Alias: "icon",
							},
							Input: Dependency{
								Path:  buildPath(moduleName, "cmd", "service", "ui", "components", "input"),
								Alias: "input",
							},
							Label: Dependency{
								Path:  buildPath(moduleName, "cmd", "service", "ui", "components", "label"),
								Alias: "label",
							},
							Pagination: Dependency{
								Path:  buildPath(moduleName, "cmd", "service", "ui", "components", "pagination"),
								Alias: "pagination",
							},
							Popover: Dependency{
								Path:  buildPath(moduleName, "cmd", "service", "ui", "components", "popover"),
								Alias: "popover",
							},
							SelectBox: Dependency{
								Path:  buildPath(moduleName, "cmd", "service", "ui", "components", "selectbox"),
								Alias: "selectbox",
							},
							Table: Dependency{
								Path:  buildPath(moduleName, "cmd", "service", "ui", "components", "table"),
								Alias: "table",
							},
							Toggle: Dependency{
								Path:  buildPath(moduleName, "cmd", "service", "ui", "components", "toggle"),
								Alias: "toggle",
							},
						},
						Layouts: Dependency{
							Path:  buildPath(moduleName, "cmd", "service", "ui", "layouts"),
							Alias: "layouts",
						},
					},
					Translations: Dependency{
						Path:  buildPath(moduleName, "cmd", "service", "translations"),
						Alias: "translations",
					},
				},
			},
		},
	}
}

func InCreation(moduleName, packageEntityName string) Dependency {
	return Dependency{
		Path:  buildPath(moduleName, "internal", packageEntityName),
		Alias: packageEntityName,
	}
}

func (c Dependency) ImportAlias(file *jen.File) {
	file.ImportName(c.Path, c.Alias)
}
func (c Dependency) GenerateImportString() string {
	return fmt.Sprintf("%s\t\"%s\"", c.Alias, c.Path)
}

// buildPath creates a path by joining the module name with the given path components
func buildPath(moduleName string, components ...string) string {
	return path.Join(append([]string{moduleName}, components...)...)
}
