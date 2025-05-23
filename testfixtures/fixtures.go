package testfixtures

import "github.com/manicar2093/gomancer/domain"

const TestPath = "github.com/user/project_name"

var (
	ModelSuccess = domain.GenerateModelInput{
		PackageEntityName: "posttests",
		ModuleInfo: domain.ModuleInfo{
			Name: TestPath,
		},
		TransformedText: domain.TransformedText{
			SnakeCase:        "post_test",
			PascalCase:       "PostTest",
			CamelCase:        "postTest",
			LowerNoSpaceCase: "posttest",
		},
		IdAttribute: domain.Attribute{
			TransformedText: domain.TransformedText{
				SnakeCase:        "id",
				PascalCase:       "Id",
				CamelCase:        "id",
				LowerNoSpaceCase: "id",
			},
			Type:       "uuid",
			IsOptional: false,
		},
		Attributes: []domain.Attribute{
			{
				TransformedText: domain.TransformedText{
					SnakeCase:        "name",
					PascalCase:       "Name",
					CamelCase:        "name",
					LowerNoSpaceCase: "name",
				},
				Type:       "string",
				IsOptional: true,
			},
			{
				TransformedText: domain.TransformedText{
					SnakeCase:        "total_price",
					PascalCase:       "TotalPrice",
					CamelCase:        "totalPrice",
					LowerNoSpaceCase: "totalprice",
				},
				Type: "decimal",
			},
			{
				TransformedText: domain.TransformedText{
					SnakeCase:        "price",
					PascalCase:       "Price",
					CamelCase:        "price",
					LowerNoSpaceCase: "price",
				},
				Type:       "decimal",
				IsOptional: true,
			},
			{
				TransformedText: domain.TransformedText{
					SnakeCase:        "age",
					PascalCase:       "Age",
					CamelCase:        "age",
					LowerNoSpaceCase: "age",
				},
				Type: "int",
			},
			{
				TransformedText: domain.TransformedText{
					SnakeCase:        "case_of_use",
					PascalCase:       "CaseOfUse",
					CamelCase:        "caseOfUse",
					LowerNoSpaceCase: "caseffuse",
				},
				Type:       "int32",
				IsOptional: true,
			},
			{
				TransformedText: domain.TransformedText{
					SnakeCase:        "created_at",
					PascalCase:       "CreatedAt",
					CamelCase:        "createdAt",
					LowerNoSpaceCase: "createdat",
				},
				Type: "time",
			},
		},
	}
)
