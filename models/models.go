package models

import (
	"fmt"
	"github.com/charmbracelet/log"
	. "github.com/dave/jennifer/jen"
	"github.com/manicar2093/gomancer/deps"
	"github.com/manicar2093/gomancer/domain"
	"github.com/manicar2093/gomancer/parser"
	"github.com/manicar2093/gomancer/types"
	"github.com/rjNemo/underscore"

	"path"
)

func GenerateModel(input parser.GenerateModelInput, goDeps deps.Container, inCreation deps.Dependency) error {
	if err := doGenerateModel(input, goDeps, inCreation); err != nil {
		return err
	}
	if err := doGenerateTestGeneratorFunc(input); err != nil {
		return err
	}
	return nil
}

func doGenerateModel(input parser.GenerateModelInput, goDeps deps.Container, inCreation deps.Dependency) error {
	log.Info("Generating model...")
	attribs := append(
		[]Code{
			Id("Id").Add(types.QualifiersByType(input.IdAttribute.Type, goDeps)).Tag(
				domain.Tags(
					input.IdAttribute,
					domain.Validations{},
					domain.GormUuidTag, domain.JsonTag, domain.ParamTag, domain.MapstructureTag,
				),
			),
		},
		underscore.Map(input.Attributes, func(item parser.Attribute) Code {
			builder := Null().Id(item.PascalCase)
			itemType := types.QualifiersByType(item.Type, goDeps)
			if item.IsOptional {
				builder.Add(types.OptionalQualifier(goDeps)).Index(itemType)
			} else {
				builder.Add(itemType)
			}

			return builder.Tag(domain.Tags(
				item,
				domain.Validations{
					Required: !item.IsOptional,
				},
				domain.JsonTag, domain.MapstructureTag, domain.ParamTag,
			))
		})...,
	)

	destinyFile := NewFile(string(domain.ModelsPkg))
	destinyFile.Null().Type().Id(input.PascalCase).Struct(
		attribs...,
	)

	return destinyFile.Save(path.Join(string(domain.InternalDomainModelsPackagePath), input.SnakeCase+".go"))
}

func doGenerateTestGeneratorFunc(input parser.GenerateModelInput) error {
	log.Info("Generating model testing generator...")
	valuesDict := make(Dict)

	valuesDict[Id("Id")] = FakerCallByType(input.IdAttribute.Type)

	structValues := underscore.Reduce(input.Attributes, func(item parser.Attribute, acc Dict) Dict {
		fakeTypeQualifier := FakerCallByType(item.Type)
		if item.IsOptional {
			fakeTypeQualifier = Qual(domain.GoptionPkgPath, "Of").Call(fakeTypeQualifier)
		}
		acc[Id(item.PascalCase)] = fakeTypeQualifier
		return acc
	}, valuesDict)

	constructor := Null().Func().Id(fmt.Sprintf("Generate%s", input.PascalCase)).Params(
		Id("t").Id("testingI"),
		Id("args").Map(String()).Any(),
	).Add(domain.GetPackageQualifier(input.ModuleInfo.Name, domain.InternalDomainModelsPackagePath, input.PascalCase)).Block(
		Id("fak").Op(":=").Add(
			domain.GetPackageQualifier(input.ModuleInfo.Name, domain.InternalDomainModelsPackagePath, input.PascalCase),
		).Values(
			structValues,
		),
		Line(),
		Id("decode").Call(Id("t"), Id("args"), Op("&").Id("fak")),
		Line(),
		Return(Id("fak")),
	)

	destinyFile := NewFile(string(domain.GeneratorsPkg))
	destinyFile.ImportAlias(fakerPkgPath, "gofakeit")
	destinyFile.Add(constructor)
	return destinyFile.Save(path.Join(string(domain.PkgGeneratorsPackagePath), input.SnakeCase+".go"))
}
