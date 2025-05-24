package testfixtures

import (
	"github.com/manicar2093/gomancer/deps"
	"github.com/manicar2093/gomancer/parser"
)

const TestPath = "github.com/user/project_name"

var (
	ModelSuccess = parser.GenerateModelInput{
		PackageEntityName: "posttests",
		ModuleInfo: parser.ModuleInfo{
			Name: TestPath,
		},
		TransformedText: parser.TransformedText{
			SnakeCase:        "post_test",
			PascalCase:       "PostTest",
			CamelCase:        "postTest",
			LowerNoSpaceCase: "posttest",
		},
		IdAttribute: parser.Attribute{
			TransformedText: parser.TransformedText{
				SnakeCase:        "id",
				PascalCase:       "Id",
				CamelCase:        "id",
				LowerNoSpaceCase: "id",
			},
			Type:       "uuid",
			IsOptional: false,
		},
		Attributes: []parser.Attribute{
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "name",
					PascalCase:       "Name",
					CamelCase:        "name",
					LowerNoSpaceCase: "name",
				},
				Type:       "string",
				IsOptional: true,
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "total_price",
					PascalCase:       "TotalPrice",
					CamelCase:        "totalPrice",
					LowerNoSpaceCase: "totalprice",
				},
				Type: "decimal",
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "price",
					PascalCase:       "Price",
					CamelCase:        "price",
					LowerNoSpaceCase: "price",
				},
				Type:       "decimal",
				IsOptional: true,
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "age",
					PascalCase:       "Age",
					CamelCase:        "age",
					LowerNoSpaceCase: "age",
				},
				Type: "int",
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "case_of_use",
					PascalCase:       "CaseOfUse",
					CamelCase:        "caseOfUse",
					LowerNoSpaceCase: "caseffuse",
				},
				Type:       "int32",
				IsOptional: true,
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "created_at",
					PascalCase:       "CreatedAt",
					CamelCase:        "createdAt",
					LowerNoSpaceCase: "createdat",
				},
				Type: "time",
			},
		},
	}

	ModelSuccessDepsContainer = deps.Init(ModelSuccess.ModuleInfo.Name)
	ModelSuccessDepInCreation = deps.InCreation(ModelSuccess.ModuleInfo.Name, ModelSuccess.PackageEntityName)
)
