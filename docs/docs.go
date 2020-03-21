// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support"
        },
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user": {
            "get": {
                "description": "Allows you to search for users using different filters and options.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Search for users.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Field to sort by. Allowed fields: username",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Direction to sort by. Allowed values: asc, desc. Default: asc",
                        "name": "sort_dir",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Starts results from this offset. Default: 0",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limits the amount of results to return. Default: 50",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resource.UserResource"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apperror.HttpError"
                        }
                    }
                }
            },
            "post": {
                "description": "Allows you to create a new user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a new user.",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/resource.UserCreateResource"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/resource.UserResource"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apperror.HttpError"
                        }
                    }
                }
            }
        },
        "/user/{username}": {
            "put": {
                "description": "Allows you to update an existing user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update a user.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/resource.UserUpdateResource"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resource.UserResource"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.HttpError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apperror.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apperror.HttpError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Allows you to delete an existing user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete a user.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/resource.UserDeleteResource"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resource.UserResource"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.HttpError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apperror.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apperror.HttpError"
                        }
                    }
                }
            }
        },
        "/user_type": {
            "get": {
                "description": "Allows you to search for user types using different filters and options.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user types"
                ],
                "summary": "Search for user types.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Type Name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Field to sort by. Allowed fields: name",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Direction to sort by. Allowed values: asc, desc. Default: asc",
                        "name": "sort_dir",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Starts results from this offset. Default: 0",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limits the amount of results to return. Default: 50",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resource.UserTypeResource"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apperror.HttpError"
                        }
                    }
                }
            },
            "post": {
                "description": "Allows you to create a new user type.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user types"
                ],
                "summary": "Create a new user type.",
                "parameters": [
                    {
                        "description": "User Type data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/resource.UserTypeCreateResource"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/resource.UserTypeResource"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apperror.HttpError"
                        }
                    }
                }
            }
        },
        "/user_type/{name}": {
            "put": {
                "description": "Allows you to update an existing user type.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user types"
                ],
                "summary": "Update a user type.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User Type data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/resource.UserTypeUpdateResource"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resource.UserTypeResource"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.HttpError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apperror.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apperror.HttpError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Allows you to delete an existing user type.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user types"
                ],
                "summary": "Delete a user type.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resource.UserTypeResource"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.HttpError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apperror.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apperror.HttpError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apperror.HttpError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "data": {
                    "type": "object",
                    "additionalProperties": true
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "resource.UserCreateResource": {
            "type": "object",
            "required": [
                "user_type_name",
                "username"
            ],
            "properties": {
                "disabled": {
                    "type": "boolean"
                },
                "user_type_name": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "resource.UserDeleteResource": {
            "type": "object"
        },
        "resource.UserResource": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "disabled": {
                    "type": "boolean"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_type": {
                    "type": "object",
                    "$ref": "#/definitions/resource.UserTypeResource"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "resource.UserTypeCreateResource": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "disabled": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "resource.UserTypeResource": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "disabled": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "resource.UserTypeUpdateResource": {
            "type": "object",
            "properties": {
                "disabled": {
                    "type": "boolean"
                }
            }
        },
        "resource.UserUpdateResource": {
            "type": "object",
            "required": [
                "user_type_name"
            ],
            "properties": {
                "disabled": {
                    "type": "boolean"
                },
                "user_type_name": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Users REST API",
	Description: "This is an autogenerated Users REST API.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
