package deps

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/manicar2093/gomancer/domain"
)

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
					Path:  fmt.Sprintf("%s/pkg", input.ModuleInfo.Name),
					Alias: "pkg",
				},
				Config: Dependency{
					Path:  fmt.Sprintf("%s/pkg/config", input.ModuleInfo.Name),
					Alias: "config",
				},
			},
			Core: Core{
				Dependency: Dependency{
					Path:  fmt.Sprintf("%s/core", input.ModuleInfo.Name),
					Alias: "core",
				},
				Converters: Dependency{
					Path:  fmt.Sprintf("%s/core/converters", input.ModuleInfo.Name),
					Alias: "converters",
				},
				Validator: Dependency{
					Path:  fmt.Sprintf("%s/core/validator", input.ModuleInfo.Name),
					Alias: "",
				},
				Logger: Dependency{
					Path:  fmt.Sprintf("%s/core/logger", input.ModuleInfo.Name),
					Alias: "",
				},
				Connections: Dependency{
					Path:  fmt.Sprintf("%s/core/connections", input.ModuleInfo.Name),
					Alias: "connections",
				},
				CommonReq: Dependency{
					Path:  fmt.Sprintf("%s/core/commonreq", input.ModuleInfo.Name),
					Alias: "commonreq",
				},
				AppErrors: Dependency{
					Path:  fmt.Sprintf("%s/core/apperrors", input.ModuleInfo.Name),
					Alias: "apperrors",
				},
			},
			Internal: Internal{
				Dependency: Dependency{
					Path:  fmt.Sprintf("%s/internal", input.ModuleInfo.Name),
					Alias: "internal",
				},
				Domain: Domain{
					Dependency: Dependency{
						Path:  fmt.Sprintf("%s/internal/domain", input.ModuleInfo.Name),
						Alias: "domain",
					},
					Models: Dependency{
						Path:  fmt.Sprintf("%s/internal/domain/models", input.ModuleInfo.Name),
						Alias: "models",
					},
				},
			},
			Cmd: Cmd{
				Dependency: Dependency{
					Path:  fmt.Sprintf("%s/cmd", input.ModuleInfo.Name),
					Alias: "",
				},
				Api: Api{
					Dependency: Dependency{
						Path:  fmt.Sprintf("%s/cmd/api", input.ModuleInfo.Name),
						Alias: "api",
					},
					Controllers: Dependency{
						Path:  fmt.Sprintf("%s/cmd/api/controllers", input.ModuleInfo.Name),
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
