package domain

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/manicar2093/gomancer/parser"
	"path"
)

type (
	Path        string
	Pkg         string
	CorePackage string
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

	CoreConnectionsPkg CorePackage = "connections"
	CoreCommonReqPkg   CorePackage = "commonreq"
)

var (
	ModelsPkg     Pkg = "models"
	GeneratorsPkg Pkg = "generators"
)

func GetPackageQualifier(moduleName string, path Path, model string) *jen.Statement {
	return jen.Qual(fmt.Sprintf("%s/%s", moduleName, path), model)
}

func GenerateCorePackage(info parser.ModuleInfo) string {
	return fmt.Sprintf("%s/core", info.Name)
}

func GetCorePackage(info parser.ModuleInfo, pkg CorePackage) string {
	return fmt.Sprintf("%s/%s", GenerateCorePackage(info), pkg)
}
