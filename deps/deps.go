package deps

import (
	"github.com/dave/jennifer/jen"
	"github.com/manicar2093/gomancer/domain"
	"path"
)

// buildPath creates a path by joining the module name with the given path components
func buildPath(moduleName string, components ...string) string {
	return path.Join(append([]string{moduleName}, components...)...)
}

func Init(input domain.GenerateModelInput) Container {
	return Container{
		Echo: Echo{
			Dependency: Dependency{
				Path:  "github.com/labstack/echo/v4",
				Alias: "echo",
			},
			Middleware: Dependency{
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
					Path:  buildPath(input.ModuleInfo.Name, "pkg"),
					Alias: "pkg",
				},
				Config: Dependency{
					Path:  buildPath(input.ModuleInfo.Name, "pkg", "config"),
					Alias: "config",
				},
			},
			Core: Core{
				Dependency: Dependency{
					Path:  buildPath(input.ModuleInfo.Name, "core"),
					Alias: "core",
				},
				Converters: Dependency{
					Path:  buildPath(input.ModuleInfo.Name, "core", "converters"),
					Alias: "converters",
				},
				Validator: Dependency{
					Path:  buildPath(input.ModuleInfo.Name, "core", "validator"),
					Alias: "",
				},
				Logger: Dependency{
					Path:  buildPath(input.ModuleInfo.Name, "core", "logger"),
					Alias: "",
				},
				Connections: Dependency{
					Path:  buildPath(input.ModuleInfo.Name, "core", "connections"),
					Alias: "connections",
				},
				CommonReq: Dependency{
					Path:  buildPath(input.ModuleInfo.Name, "core", "commonreq"),
					Alias: "commonreq",
				},
				AppErrors: Dependency{
					Path:  buildPath(input.ModuleInfo.Name, "core", "apperrors"),
					Alias: "apperrors",
				},
			},
			Internal: Internal{
				Dependency: Dependency{
					Path:  buildPath(input.ModuleInfo.Name, "internal"),
					Alias: "internal",
				},
				Domain: Domain{
					Dependency: Dependency{
						Path:  buildPath(input.ModuleInfo.Name, "internal", "domain"),
						Alias: "domain",
					},
					Models: Dependency{
						Path:  buildPath(input.ModuleInfo.Name, "internal", "domain", "models"),
						Alias: "models",
					},
				},
				InCreation: Dependency{
					Path:  buildPath(input.ModuleInfo.Name, "internal", input.PackageEntityName),
					Alias: input.PackageEntityName,
				},
			},
			Cmd: Cmd{
				Dependency: Dependency{
					Path:  buildPath(input.ModuleInfo.Name, "cmd"),
					Alias: "",
				},
				Api: Api{
					Dependency: Dependency{
						Path:  buildPath(input.ModuleInfo.Name, "cmd", "api"),
						Alias: "api",
					},
					Controllers: Dependency{
						Path:  buildPath(input.ModuleInfo.Name, "cmd", "api", "controllers"),
						Alias: "controllers",
					},
				},
			},
		},
	}
}

func (c Dependency) ImportAlias(file *jen.File) {
	file.ImportName(c.Path, c.Alias)
}
