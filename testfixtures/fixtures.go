package testfixtures

import "github.com/manicar2093/gomancer/domain"

const TestPath = "github.com/user/project_name"

var (
	ModelSuccess = domain.GenerateModelInput{
		ModuleInfo: domain.ModuleInfo{
			Name: TestPath,
		},
		TransformedText: domain.TransformedText{
			SnakeCase:  "post_test",
			PascalCase: "PostTest",
			CamelCase:  "postTest",
		},
		IdAttribute: domain.Attribute{
			TransformedText: domain.TransformedText{
				SnakeCase:  "id",
				PascalCase: "Id",
				CamelCase:  "id",
			},
			Type:       "uuid",
			IsOptional: false,
		},
		Attributes: []domain.Attribute{
			{
				TransformedText: domain.TransformedText{
					SnakeCase:  "name",
					PascalCase: "Name",
					CamelCase:  "name",
				},
				Type:       "string",
				IsOptional: true,
			},
			{
				TransformedText: domain.TransformedText{
					SnakeCase:  "total_price",
					PascalCase: "TotalPrice",
					CamelCase:  "totalPrice",
				},
				Type: "decimal",
			},
			{
				TransformedText: domain.TransformedText{
					SnakeCase:  "price",
					PascalCase: "Price",
					CamelCase:  "price",
				},
				Type:       "decimal",
				IsOptional: true,
			},
			{
				TransformedText: domain.TransformedText{
					SnakeCase:  "age",
					PascalCase: "Age",
					CamelCase:  "age",
				},
				Type: "int",
			},
			{
				TransformedText: domain.TransformedText{
					SnakeCase:  "case_of_use",
					PascalCase: "CaseOfUse",
					CamelCase:  "caseOfUse",
				},
				Type:       "int32",
				IsOptional: true,
			},
			{
				TransformedText: domain.TransformedText{
					SnakeCase:  "created_at",
					PascalCase: "CreatedAt",
					CamelCase:  "createdAt",
				},
				Type: "time",
			},
		},
	}
)
