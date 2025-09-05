package testfixtures

import (
	"github.com/manicar2093/gomancer/deps"
	"github.com/manicar2093/gomancer/parser"
)

const TestPath = "github.com/user/project_name"

var (
	ModelBinaryIdSuccess = parser.GenerateModelInput{
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
					SnakeCase:        "an_optional_int",
					PascalCase:       "AnOptionalInt",
					CamelCase:        "anOptionalInt",
					LowerNoSpaceCase: "anoptionalint",
				},
				Type:       "int",
				IsOptional: true,
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "an_int",
					PascalCase:       "AnInt",
					CamelCase:        "anInt",
					LowerNoSpaceCase: "anint",
				},
				Type: "int",
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "an_optional_int8",
					PascalCase:       "AnOptionalInt8",
					CamelCase:        "anOptionalInt8",
					LowerNoSpaceCase: "anoptionalint8",
				},
				Type:       "int8",
				IsOptional: true,
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "an_int8",
					PascalCase:       "AnInt8",
					CamelCase:        "anInt8",
					LowerNoSpaceCase: "anint8",
				},
				Type: "int8",
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "an_optional_int16",
					PascalCase:       "AnOptionalInt16",
					CamelCase:        "anOptionalInt16",
					LowerNoSpaceCase: "anoptionalint16",
				},
				Type:       "int16",
				IsOptional: true,
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "an_int16",
					PascalCase:       "AnInt16",
					CamelCase:        "anInt16",
					LowerNoSpaceCase: "anint16",
				},
				Type: "int16",
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "an_optional_int32",
					PascalCase:       "AnOptionalInt32",
					CamelCase:        "anOptionalInt32",
					LowerNoSpaceCase: "anoptionalint32",
				},
				Type:       "int32",
				IsOptional: true,
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "an_int32",
					PascalCase:       "AnInt32",
					CamelCase:        "anInt32",
					LowerNoSpaceCase: "anint32",
				},
				Type: "int32",
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "an_optional_int64",
					PascalCase:       "AnOptionalInt64",
					CamelCase:        "anOptionalInt64",
					LowerNoSpaceCase: "anoptionalint64",
				},
				Type:       "int64",
				IsOptional: true,
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "an_int64",
					PascalCase:       "AnInt64",
					CamelCase:        "anInt64",
					LowerNoSpaceCase: "anint64",
				},
				Type: "int64",
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "an_optional_float32",
					PascalCase:       "AnOptionalFloat32",
					CamelCase:        "anOptionalFloat32",
					LowerNoSpaceCase: "anoptionalfloat32",
				},
				Type:       "float32",
				IsOptional: true,
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "an_float32",
					PascalCase:       "AnFloat32",
					CamelCase:        "anFloat32",
					LowerNoSpaceCase: "anfloat32",
				},
				Type: "float32",
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "an_optional_float64",
					PascalCase:       "AnOptionalFloat64",
					CamelCase:        "anOptionalFloat64",
					LowerNoSpaceCase: "anoptionalfloat64",
				},
				Type:       "float64",
				IsOptional: true,
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "an_float64",
					PascalCase:       "AnFloat64",
					CamelCase:        "anFloat64",
					LowerNoSpaceCase: "anfloat64",
				},
				Type: "float64",
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "optional_string",
					PascalCase:       "OptionalString",
					CamelCase:        "optionalString",
					LowerNoSpaceCase: "optionalstring",
				},
				Type:       "string",
				IsOptional: true,
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "string",
					PascalCase:       "String",
					CamelCase:        "string",
					LowerNoSpaceCase: "string",
				},
				Type: "string",
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "optional_bool",
					PascalCase:       "OptionalBool",
					CamelCase:        "optionalBool",
					LowerNoSpaceCase: "optionalbool",
				},
				Type:       "bool",
				IsOptional: true,
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "bool",
					PascalCase:       "Bool",
					CamelCase:        "bool",
					LowerNoSpaceCase: "bool",
				},
				Type: "bool",
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "optional_time",
					PascalCase:       "OptionalTime",
					CamelCase:        "optionalTime",
					LowerNoSpaceCase: "optionaltime",
				},
				Type:       "time",
				IsOptional: true,
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "time",
					PascalCase:       "Time",
					CamelCase:        "time",
					LowerNoSpaceCase: "time",
				},
				Type: "time",
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "optional_decimal",
					PascalCase:       "OptionalDecimal",
					CamelCase:        "optionalDecimal",
					LowerNoSpaceCase: "optionaldecimal",
				},
				Type:       "decimal",
				IsOptional: true,
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "decimal",
					PascalCase:       "Decimal",
					CamelCase:        "decimal",
					LowerNoSpaceCase: "decimal",
				},
				Type: "decimal",
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "optional_uuid",
					PascalCase:       "OptionalUuid",
					CamelCase:        "optionalUuid",
					LowerNoSpaceCase: "optionaluuid",
				},
				Type:       "uuid",
				IsOptional: true,
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "uuid",
					PascalCase:       "Uuid",
					CamelCase:        "uuid",
					LowerNoSpaceCase: "uuid",
				},
				Type: "uuid",
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "optional_enum",
					PascalCase:       "OptionalEnum",
					CamelCase:        "optionalEnum",
					LowerNoSpaceCase: "optionalenum",
				},
				Type:       "enum",
				IsOptional: true,
				EnumStrings: []parser.TransformedText{
					{
						SnakeCase:        "optional_enum1",
						PascalCase:       "OptionalEnum1",
						CamelCase:        "optionalEnum1",
						LowerNoSpaceCase: "optionalenum1",
					},
					{
						SnakeCase:        "optional_enum2",
						PascalCase:       "OptionalEnum2",
						CamelCase:        "optionalEnum2",
						LowerNoSpaceCase: "optionalenum2",
					},
					{
						SnakeCase:        "optional_enum3",
						PascalCase:       "OptionalEnum3",
						CamelCase:        "optionalEnum3",
						LowerNoSpaceCase: "optionalenum3",
					},
				},
			},
			{
				TransformedText: parser.TransformedText{
					SnakeCase:        "enum",
					PascalCase:       "Enum",
					CamelCase:        "enum",
					LowerNoSpaceCase: "enum",
				},
				Type: "enum",
				EnumStrings: []parser.TransformedText{
					{
						SnakeCase:        "enum1",
						PascalCase:       "Enum1",
						CamelCase:        "enum1",
						LowerNoSpaceCase: "enum1",
					},
					{
						SnakeCase:        "enum2",
						PascalCase:       "Enum2",
						CamelCase:        "enum2",
						LowerNoSpaceCase: "enum2",
					},
					{
						SnakeCase:        "enum3",
						PascalCase:       "Enum3",
						CamelCase:        "enum3",
						LowerNoSpaceCase: "enum3",
					},
				},
			},
		},
	}

	ModelSuccessDepsContainer = deps.Init(ModelBinaryIdSuccess.ModuleInfo.Name, ModelBinaryIdSuccess.TransformedText.LowerNoSpaceCase)
	ModelSuccessDepInCreation = deps.InCreation(ModelBinaryIdSuccess.ModuleInfo.Name, ModelBinaryIdSuccess.PackageEntityName)
)
