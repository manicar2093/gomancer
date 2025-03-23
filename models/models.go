package models

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
	"github.com/manicar2093/gomancer/domain"
	"github.com/rjNemo/underscore"
	"log"
	"path"
)

type (
	Attribute struct {
		domain.TransformedText
		Type       string
		IsOptional bool
	}
	GenerateModelInput struct {
		domain.TransformedText
		ModuleInfo  domain.ModuleInfo
		IdAttribute Attribute
		Attributes  []Attribute
	}
)

func GenerateModel(input GenerateModelInput) error {
	if err := doGenerateModel(input); err != nil {
		return err
	}
	if err := doGenerateTestGeneratorFunc(input); err != nil {
		return err
	}
	return nil
}

func doGenerateModel(input GenerateModelInput) error {
	log.Println("Generating model...")
	attribs := append(
		[]Code{
			Id("Id").Add(domain.QualifiersByType(input.IdAttribute.Type)).Tag(
				domain.Tags(
					input.IdAttribute.SnakeCase,
					domain.Validations{},
					domain.GormUuidTag, domain.JsonTag, domain.ParamTag, domain.MapstructureTag,
				),
			),
		},
		underscore.Map(input.Attributes, func(item Attribute) Code {
			builder := Null().Id(item.PascalCase)
			itemType := domain.QualifiersByType(item.Type)
			if item.IsOptional {
				builder.Add(domain.OptionalQualifier()).Index(itemType)
			} else {
				builder.Add(itemType)
			}

			return builder.Tag(domain.Tags(
				item.SnakeCase,
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

func doGenerateTestGeneratorFunc(input GenerateModelInput) error {
	log.Println("Generating model testing generator...")
	valuesDict := make(Dict)

	valuesDict[Id("Id")] = FakerCallByType(input.IdAttribute.Type)

	structValues := underscore.Reduce(input.Attributes, func(item Attribute, acc Dict) Dict {
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
