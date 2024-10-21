// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/skills/categories": {
            "get": {
                "security": [
                    {
                        "ApiAuthKey": []
                    }
                ],
                "description": "Get all skill Categories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Skills Categories"
                ],
                "summary": "Get all skill Categories",
                "operationId": "Get-skill-categories",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "example - 50",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "example - 0",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "example - created_at desc,updated_at desc",
                        "name": "orderBy",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiAuthKey": []
                    }
                ],
                "description": "Create skill Categories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Skills Categories"
                ],
                "summary": "Create skill Categories",
                "operationId": "Create-skill-categories",
                "parameters": [
                    {
                        "description": "Skills Categories",
                        "name": "CreateSkillCategoryRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/skills.CreateSkillCategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseMessage"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseMessage"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseMessage"
                        }
                    }
                }
            }
        },
        "/skills/categories/{id}": {
            "get": {
                "security": [
                    {
                        "ApiAuthKey": []
                    }
                ],
                "description": "Get skill category",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Skills Categories"
                ],
                "summary": "Get skill category",
                "operationId": "get-skill-category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Skill Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiAuthKey": []
                    }
                ],
                "description": "Delete skill category",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Skills Categories"
                ],
                "summary": "Delete skill category",
                "operationId": "delete-skill-category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Skill Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "ApiAuthKey": []
                    }
                ],
                "description": "Update skill category",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Skills Categories"
                ],
                "summary": "Update skill category",
                "operationId": "update-skill-category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Skill Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Skills Categories",
                        "name": "SkillCategoryUpdateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/skills.SkillCategoryUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/experience/add-user-experience": {
            "post": {
                "description": "Adds new experiences for a given user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "experience"
                ],
                "summary": "Add experiences for user",
                "operationId": "add-experience",
                "parameters": [
                    {
                        "description": "AddUserExperienceRequest",
                        "name": "AddUserExperienceRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/experience.AddUserExperienceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseMessage"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseMessage"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseMessage"
                        }
                    }
                }
            }
        },
        "/experience/delete-user-experience/{id}": {
            "delete": {
                "description": "delete user experience by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "experience"
                ],
                "summary": "Delete user experience by id",
                "operationId": "delete-user-experience-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "GetExperienceRequest",
                        "name": "GetExperienceRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/experience.GetExperienceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/experience/get-user-experience-details/{id}": {
            "get": {
                "description": "get user experience details by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "experience"
                ],
                "summary": "Get user experience details by id",
                "operationId": "get-user-experience-details-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "GetExperienceRequest",
                        "name": "GetExperienceRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/experience.GetExperienceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/experience/update-user-experience/{id}": {
            "patch": {
                "description": "Updates experience",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "experience"
                ],
                "summary": "Update experience",
                "operationId": "update-experience",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "UpdateExpRequest",
                        "name": "UpdateExpRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/experience.UpdateExpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/delete-user{id}": {
            "delete": {
                "description": "delete user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete user by id",
                "operationId": "delete-user-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/get-all-user-list": {
            "get": {
                "description": "get all user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get all user",
                "operationId": "get-all-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "example - 50",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "example - 0",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "example - created_at desc",
                        "name": "orderBy",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/get-user-details/{id}": {
            "get": {
                "description": "get user details by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user details by id",
                "operationId": "get-user-details-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/update-user/{id}": {
            "patch": {
                "description": "Updates user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update user",
                "operationId": "update-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "UpdateUser",
                        "name": "UpdateUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/create-user": {
            "post": {
                "description": "creates a new complete user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create user",
                "operationId": "create-user",
                "parameters": [
                    {
                        "description": "CreateUserRequest",
                        "name": "CreateUserRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseMessage"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "skills.CreateSkillCategoryRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "skills.SkillCategoryUpdateRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "experience.AddUserExperienceRequest": {
            "type": "object",
            "required": [
                "user_id"
            ],
            "properties": {
                "experiences": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/experience.ExpRequest"
                    }
                },
                "skill_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "experience.ExpRequest": {
            "type": "object",
            "required": [
                "company",
                "position",
                "start_date"
            ],
            "properties": {
                "company": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "is_currently_working": {
                    "type": "boolean"
                },
                "position": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                }
            }
        },
        "experience.GetExperienceRequest": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "experience.UpdateExpRequest": {
            "type": "object",
            "required": [
                "company",
                "position",
                "start_date"
            ],
            "properties": {
                "company": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "is_currently_working": {
                    "type": "boolean"
                },
                "position": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "user.CreateUserRequest": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "mobile_number": {
                    "type": "string"
                },
                "user_category_id": {
                    "type": "integer"
                }
            }
        },
        "user.UpdateUser": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string"
                },
                "certifications": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "mobile_number": {
                    "type": "string"
                },
                "user_category_id": {
                    "type": "integer"
                },
                "video_url": {
                    "type": "string"
                }
            }
        },
        "utils.ResponseMessage": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
