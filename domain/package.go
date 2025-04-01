package domain

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"path"
)

type (
	Path string
	Pkg  string
)

var (
	InternalPackagePath             Path = "internal"
	PkgPackagePath                  Path = "pkg"
	PrismaPackagePath               Path = "prisma"
	CmdPackagePath                  Path = "cmd"
	GithubPackagePath               Path = ".github"
	PkgGeneratorsPackagePath             = Path(path.Join(string(PkgPackagePath), "generators"))
	PkgConfigPackagePath                 = Path(path.Join(string(PkgPackagePath), "config"))
	PkgVersioningPackagePath             = Path(path.Join(string(PkgPackagePath), "versioning"))
	CmdApiPackagePath                    = Path(path.Join(string(CmdPackagePath), "api"))
	CmdApiControllersPackagePath         = Path(path.Join(string(CmdPackagePath), "api", "controllers"))
	PrismaSchemaPackagePath              = Path(path.Join(string(PrismaPackagePath), "schema"))
	InternalDomainModelsPackagePath      = Path(path.Join(string(InternalPackagePath), "domain", "models"))
	GithubWorkflowsPackagePath           = Path(path.Join(string(GithubPackagePath), "workflows"))
)

var (
	ModelsPkg     Pkg = "models"
	GeneratorsPkg Pkg = "generators"
)

func GetPackageQualifier(moduleName string, path Path, model string) *jen.Statement {
	return jen.Qual(fmt.Sprintf("%s/%s", moduleName, path), model)
}
