// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/LeandroAlcantara-1997",
            "email": "leandro1997silva97@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://choosealicense.com/licenses/mit/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/appointment": {
            "get": {
                "description": "Get all appointments",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointment"
                ],
                "summary": "Get all appointments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.AppResponse"
                            }
                        }
                    },
                    "404": {
                        "description": "Appointment not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "An error happened in database",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/appointment/available": {
            "get": {
                "description": "get all available appointments",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointment"
                ],
                "summary": "Get available appointments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.AppResponse"
                            }
                        }
                    },
                    "404": {
                        "description": "Appointment not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "An error happened in database",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/appointment/salon/{id}": {
            "get": {
                "description": "get by salon ID and return an appointment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointment"
                ],
                "summary": "Get appointments by salon id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Salon ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.AppResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Cannot read path",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Appointment not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "An error happened in database",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/appointment/user/{id}": {
            "get": {
                "description": "Get by user ID and return an appointment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointment"
                ],
                "summary": "Get appointments by user id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.AppResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Cannot read path",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Appointment not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "An error happened in database",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/appointment/{id}": {
            "put": {
                "description": "Get Appointment by ID and body for update",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointment"
                ],
                "summary": "Update an appointment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Appointment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Appointment",
                        "name": "appointment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AppResponse"
                        }
                    },
                    "400": {
                        "description": "Cannot read path",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Appointment not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "An error happened in database",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "get string by ID and delete an appointment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointment"
                ],
                "summary": "Delete appointments by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Appointment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Cannot read path",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Appointment not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "An error happened in database",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/appointment/{id}/{user}": {
            "put": {
                "description": "cancel appointment by ID and user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointment"
                ],
                "summary": "Cancel an appointment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Appointment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Cannot read path",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Appointment not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "An error happened in database",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AppResponse": {
            "type": "object",
            "properties": {
                "appointment_date": {
                    "type": "string",
                    "example": "2022-06-23T21:12:02.000000001Z"
                },
                "id": {
                    "type": "string",
                    "example": "62b65300e1d7eab1ea9a681d"
                },
                "salon_id": {
                    "type": "integer",
                    "example": 1
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1/appointment",
	Schemes:          []string{},
	Title:            "Appointment API",
	Description:      "This is a service for make appointments .",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
