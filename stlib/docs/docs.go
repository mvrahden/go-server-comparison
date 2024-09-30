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
        "/helloworld": {
            "get": {
                "description": "getHelloWorld returns hello world",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "helloworld"
                ],
                "summary": "getHelloWorld returns hello world",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Message"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "getHelloWorld returns hello world with name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "helloworld"
                ],
                "summary": "getHelloWorld returns hello world with name",
                "parameters": [
                    {
                        "maxLength": 100,
                        "minLength": 2,
                        "type": "string",
                        "example": "John",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Message"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "getHelloWorld returns hello world with name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "helloworld"
                ],
                "summary": "getHelloWorld returns hello world with name",
                "parameters": [
                    {
                        "description": "what name",
                        "name": "query",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/NameModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Message"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ErrorResponse": {
            "type": "object",
            "required": [
                "error"
            ],
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "Message": {
            "description": "Response that contains a message",
            "type": "object",
            "required": [
                "message"
            ],
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "NameModel": {
            "description": "Model for put request - concatenates hello world with name",
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 2,
                    "example": "John"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "",
	Description:      "Response that contains a message",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
