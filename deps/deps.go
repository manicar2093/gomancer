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
					Controllers: Dependency{
						Path:  buildPath(moduleName, "cmd", "service", "controllers"),
						Alias: "controllers",
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
