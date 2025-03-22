package domain

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type (
	Path string
	Pkg  string
)

var (
	ModelsPackagePath     Path = "internal/domain/models"
	GeneratorsPackagePath Path = "pkg/generators"
)

var (
	ModelsPkg     Pkg = "models"
	GeneratorsPkg Pkg = "generators"
)

func GetPackageQualifier(moduleName string, path Path, model string) jen.Code {
	return jen.Qual(fmt.Sprintf("%s/%s", moduleName, path), model)
}
