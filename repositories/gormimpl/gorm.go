package gormimpl

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
	"github.com/manicar2093/gomancer/domain"
	"github.com/rjNemo/underscore"
	"path"
)

type (
	generatorData struct {
		repositoryStructName string
		dbKey                string
		receiverStatement    Code
		receiverVar          string
	}
	generatorType func(domain.GenerateModelInput, generatorData) Code
)

func GenerateRepository(input domain.GenerateModelInput) error {
	repositoryStructName := fmt.Sprintf("%sRepository", input.PascalCase)
	receiverVar := "c"
	data := generatorData{
		repositoryStructName: repositoryStructName,
		dbKey:                "db",
		receiverStatement:    Id(receiverVar).Op("*").Add(Id(repositoryStructName)),
		receiverVar:          receiverVar,
	}
	f := NewFile(input.SnakeCase)
	generators := []generatorType{
		generateRepoStruct,
		generateRepoConstructor,
		generateSaveMethod,
		generateGetByIdMethod,
		generatedGetAllPaginatedMethod,
		generatePartialUpdateFunction,
		generateDeleteByIdFunction,
	}

	underscore.Each(generators, func(generator generatorType) {
		f.Add(generator(input, data))
	})

	return f.Save(
		path.Join(
			string(domain.InternalPackagePath),
			input.SnakeCase,
			"repository_gomancer.go",
		),
	)
}

func generateRepoStruct(input domain.GenerateModelInput, generatorData generatorData) Code {
	return Type().Id(generatorData.repositoryStructName).Struct(
		Id(generatorData.dbKey).Op("*").Qual(domain.WinterConnPkgPath, "ConnWrapper"),
	).Line().Line()
}

func generateRepoConstructor(input domain.GenerateModelInput, generatorData generatorData) Code {
	return Func().Id(fmt.Sprintf("New%sRepository", input.PascalCase)).Params(
		Id(generatorData.dbKey).Op("*").Qual(domain.WinterConnPkgPath, "ConnWrapper"),
	).Op("*").Id(generatorData.repositoryStructName).Block(
		Return(
			Op("&").Id(generatorData.repositoryStructName).Values(
				Dict{
					Id(generatorData.dbKey): Id(generatorData.dbKey),
				},
			),
		),
	).Line().Line()
}

func generateSaveMethod(input domain.GenerateModelInput, generatorData generatorData) Code {
	return Func().Params(generatorData.receiverStatement).Id("Save").Params(
		Id("input").Op("*").Add(
			domain.GetPackageQualifier(input.ModuleInfo.Name, domain.InternalDomainModelsPackagePath, input.PascalCase)),
	).Error().Block(
		If(
			Id("res").Op(":=").Id(generatorData.receiverVar).Dot(generatorData.dbKey).Dot("Save").Call(
				Id("input"),
			),
			Id("res").Dot("Error").Op("!=").Nil(),
		).Block(
			Return(Id("res").Dot("Error")),
		),
		Return(Nil()),
	).Line().Line()
}

func generateGetByIdMethod(input domain.GenerateModelInput, generatorData generatorData) Code {
	return Comment(" get method")
}

func generatedGetAllPaginatedMethod(input domain.GenerateModelInput, generatorData generatorData) Code {
	return Comment(" get paginated method")
}

func generatePartialUpdateFunction(input domain.GenerateModelInput, generatorData generatorData) Code {
	return Comment(" partial update function")
}

func generateDeleteByIdFunction(input domain.GenerateModelInput, generatorData generatorData) Code {
	return Comment("delete by id")
}
