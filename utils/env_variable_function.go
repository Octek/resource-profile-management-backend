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
