package models

import (
	"fmt"
	"github.com/charmbracelet/log"
	. "github.com/dave/jennifer/jen"
	"github.com/jinzhu/inflection"
	"github.com/manicar2093/gomancer/deps"
	"github.com/manicar2093/gomancer/domain"
	"github.com/manicar2093/gomancer/parser"
	"github.com/manicar2093/gomancer/tags"
	"github.com/manicar2093/gomancer/types"
	"github.com/rjNemo/underscore"

	"path"
)

func GenerateModel(input parser.GenerateModelInput, goDeps deps.Container) error {
	if err := doGenerateModel(input, goDeps); err != nil {
		return err
	}
	if err := doGenerateTestGeneratorFunc(input, goDeps); err != nil {
		return err
	}
	return nil
}

func doGenerateModel(input parser.GenerateModelInput, goDeps deps.Container) error {
	log.Info("Generating model...")
	// Id does not need required validation due this is handled in core package via GetById or GetByUuid structs
	// which has needed validations
	attribs := append(
		[]Code{
			Id("Id").Add(types.QualifiersByType(input.IdAttribute.Type, goDeps, "", false)).Tag(
				tags.Tags(
					tags.NewJson(tags.JsonOptions{
						Name: input.IdAttribute.SnakeCase,
					}),
					tags.NewEcho(tags.EchoOptions{
						Name: input.IdAttribute.SnakeCase,
						Tag:  tags.EchoParam,
					}),
					tags.NewMapstructure(tags.MapstructureOptions{
						Name: input.IdAttribute.SnakeCase,
					}),
					tags.NewGormPK(tags.GormPKOptions{
						IsUuid: input.IdAttribute.Type == string(types.TypeUuid),
					}),
				),
			),
		},
		underscore.Map(input.Attributes, func(item parser.Attribute) Code {
			builder := Null().Id(item.PascalCase)
			itemType := types.QualifiersByType(item.Type, goDeps, item.PascalCase, false)
			if item.IsOptional {
				builder.Add(types.OptionalQualifier(goDeps)).Index(itemType)
			} else {
				builder.Add(itemType)
			}

			var validationsTags []tags.ValidateGenerable

			if !item.IsOptional {
				if item.Type == string(types.TypeUuid) {
					validationsTags = append(validationsTags, tags.ValidateRequiredUuid{})
				} else {
					validationsTags = append(validationsTags, tags.ValidateRequired{})
				}
			}

			itemTags := []tags.Generable{
				tags.NewJson(tags.JsonOptions{
					Name: item.SnakeCase,
				}),
				tags.NewMapstructure(tags.MapstructureOptions{
					Name: item.SnakeCase,
				}),
				tags.NewEcho(tags.EchoOptions{
					Name: item.SnakeCase,
					Tag:  tags.EchoParam,
				}),
			}
			if len(validationsTags) > 0 {
				itemTags = append(itemTags, tags.NewValidate(validationsTags...))
			}
			return builder.Tag(tags.Tags(
				itemTags...,
			))
		})...,
	)

	destinyFile := NewFile(string(domain.ModelsPkg))
	destinyFile.Null().Type().Id(input.PascalCase).Struct(
		attribs...,
	).Line()

	enums := underscore.Filter(input.Attributes, func(item parser.Attribute) bool {
		return item.Type == string(types.TypeEnum)
	})
	underscore.Each(enums, func(item parser.Attribute) {
		destinyFile.Null().Type().Id(item.PascalCase).String()

		destinyFile.Const().DefsFunc(func(group *Group) {
			for _, enumItem := range item.EnumStrings {
				group.Id(enumItem.PascalCase).Id(item.PascalCase).Op("=").Lit(enumItem.SnakeCase)
			}
		}).Line()
	})

	return destinyFile.Save(path.Join(string(domain.InternalDomainModelsPackagePath), inflection.Plural(input.SnakeCase)+".go"))
}

func doGenerateTestGeneratorFunc(input parser.GenerateModelInput, goDeps deps.Container) error {
	log.Info("Generating model testing generator...")
	valuesDict := make(Dict)

	valuesDict[Id("Id")] = FakerCallByType(input.IdAttribute.Type, goDeps, input.IdAttribute)

	structValues := underscore.Reduce(input.Attributes, func(item parser.Attribute, acc Dict) Dict {
		fakeTypeQualifier := FakerCallByType(item.Type, goDeps, item)
		if item.IsOptional {
			fakeTypeQualifier = Qual(domain.GoptionPkgPath, "Of").Call(fakeTypeQualifier)
		}
		acc[Id(item.PascalCase)] = fakeTypeQualifier
		return acc
	}, valuesDict)

	constructor := Null().Func().Id(fmt.Sprintf("Generate%s", input.PascalCase)).Params(
		Id("t").Id("testingI"),
		Id("args").Map(String()).Any(),
	).Qual(goDeps.Project.Internal.Domain.Models.Path, input.PascalCase).Block(
		Id("fak").Op(":=").Qual(
			goDeps.Project.Internal.Domain.Models.Path, input.PascalCase,
		).Values(
			structValues,
		),
		Line(),
		Id("decode").Call(Id("t"), Id("args"), Op("&").Id("fak")),
		Line(),
		Return(Id("fak")),
	)

	destinyFile := NewFile(string(domain.GeneratorsPkg))
	destinyFile.ImportAlias(goDeps.GoFakeIt.Path, goDeps.GoFakeIt.Alias)
	destinyFile.Add(constructor)
	return destinyFile.Save(path.Join(string(domain.PkgGeneratorsPackagePath), inflection.Plural(input.SnakeCase)+".go"))
}
