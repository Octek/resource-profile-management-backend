package utils

import "os"

const (
	EnvironmentVariableNotSet    = " environment variable not set"
	DB_SERVICE_CONNECTION_STRING = "DB_SERVICE_CONNECTION_STRING"
	SWAGGER_HOST_URL             = "SWAGGER_HOST_URL"
)

func GetConnectionString() string {
	connectionString, ok := os.LookupEnv(DB_SERVICE_CONNECTION_STRING)
	if !ok {
		panic(DB_SERVICE_CONNECTION_STRING + EnvironmentVariableNotSet)
	}
	return connectionString
}

func GetSwaggerHostUrl() string {
	swaggerHostUrl, ok := os.LookupEnv(SWAGGER_HOST_URL)
	if !ok {
		panic(SWAGGER_HOST_URL + EnvironmentVariableNotSet)
	}
	return swaggerHostUrl
}

const (
	RequestSchemaInvalid                           = "The request schema is invalid: %v"
	SomethingWentWrongWhileCreatingSkillCategories = "Something went wrong  while creating the skill categories: %v"
	SomethingWentWrongWhileDeletingSkillCategories = "Something went wrong while deleting the skill category: %v"
	SomethingWentWrongWhileUpdatingSkillCategory   = "Something went wrong while updating the skill category: %v"
	SomethingWentWrongWhileGettingSkillCategory    = "Something went wrong while getting the skill category: %v"
	InvalidJsonBody                                = "The JSON body is invalid: %v"
	SuccessfullyCreatedSkillsCategories            = "The Skills categories has been created successfully"
	InvalidIntegerValueLimitMessage                = "Invalid integer value for the limit : %v"
	InvalidIntegerValueOffsetMessage               = "Invalid integer value for the offset : %v"
	DefaultLimit                                   = "20"
	DefaultOffset                                  = "0"
	DefaultOrderBy                                 = "created_at desc"
	SomethingWentWrong                             = "Something went wrong: %v"
	SuccessfullyDeletedSkillsCategory              = "Skill category has been successfully deleted"
	SuccessfullyUpdatedSkillsCategory              = "Skill category has been successfully Updated"
	Success                                        = "Success"
	RequiredSkillCategoryNames                     = "Skill category Name(s) missing. Provide at least one ID."
)
