// Code generated by swaggo/swag. DO NOT EDIT.

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
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/brands": {
            "get": {
                "description": "Get a list of Brand.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Brand"
                ],
                "summary": "Get all Brand.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Brand"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Creating a new Brand.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Brand"
                ],
                "summary": "Create New Brand.",
                "parameters": [
                    {
                        "description": "the body to create a new Brand",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.BrandInput"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Brand"
                        }
                    }
                }
            }
        },
        "/brands/{id}": {
            "get": {
                "description": "Get an Brand by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Brand"
                ],
                "summary": "Get Brand.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Brand id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Brand"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Delete a Brand by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Brand"
                ],
                "summary": "Delete one Brand.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Brand id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Update Brand by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Brand"
                ],
                "summary": "Update Brand.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Brand id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to update age rating category",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.BrandInput"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Brand"
                        }
                    }
                }
            }
        },
        "/brands/{id}/models": {
            "get": {
                "description": "Get all Models by BrandId.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Brand"
                ],
                "summary": "Get Models.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Brand id",
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
                                "$ref": "#/definitions/models.Model"
                            }
                        }
                    }
                }
            }
        },
        "/brands/{id}/phones": {
            "get": {
                "description": "Get all Phones by BrandId.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Brand"
                ],
                "summary": "Get Phones.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Brand id",
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
                                "$ref": "#/definitions/models.Phone"
                            }
                        }
                    }
                }
            }
        },
        "/colors": {
            "get": {
                "description": "Get a list of Color.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Color"
                ],
                "summary": "Get all Color.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Color"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Creating a new Color.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Color"
                ],
                "summary": "Create New Color.",
                "parameters": [
                    {
                        "description": "the body to create a new Color",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ColorInput"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Color"
                        }
                    }
                }
            }
        },
        "/colors/{id}": {
            "get": {
                "description": "Get an Color by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Color"
                ],
                "summary": "Get Color.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Color id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Color"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Delete a Color by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Color"
                ],
                "summary": "Delete one Color.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Color id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Update Color by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Color"
                ],
                "summary": "Update Color.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Color id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "the body to update age rating category",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ColorInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Color"
                        }
                    }
                }
            }
        },
        "/colors/{id}/phones": {
            "get": {
                "description": "Get all Phones by ColorId.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Color"
                ],
                "summary": "Get Phones.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Color id",
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
                                "$ref": "#/definitions/models.Phone"
                            }
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Logging in to get jwt token to access admin or user api by roles.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login as as user.",
                "parameters": [
                    {
                        "description": "the body to login a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/models": {
            "get": {
                "description": "Get a list of Model.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Model"
                ],
                "summary": "Get all Model.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Model"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Creating a new Model.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Model"
                ],
                "summary": "Create New Model.",
                "parameters": [
                    {
                        "description": "the body to create a new Model",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ModelInput"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Model"
                        }
                    }
                }
            }
        },
        "/models/{id}": {
            "get": {
                "description": "Get an Model by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Model"
                ],
                "summary": "Get Model.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Model id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Model"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Delete a Model by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Model"
                ],
                "summary": "Delete one Model.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Model id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Update Model by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Model"
                ],
                "summary": "Update Model.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Model id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "the body to update age rating category",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ModelInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Model"
                        }
                    }
                }
            }
        },
        "/models/{id}/phones": {
            "get": {
                "description": "Get all Phones by ModelId.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Model"
                ],
                "summary": "Get Phones.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Model id",
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
                                "$ref": "#/definitions/models.Phone"
                            }
                        }
                    }
                }
            }
        },
        "/phones": {
            "get": {
                "description": "Get a list of Phone.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Phone"
                ],
                "summary": "Get all Phone.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Phone"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Creating a new Phone.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Phone"
                ],
                "summary": "Create New Phone.",
                "parameters": [
                    {
                        "description": "the body to create a new Phone",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.PhoneInput"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Phone"
                        }
                    }
                }
            }
        },
        "/phones/{id}": {
            "get": {
                "description": "Get an Phone by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Phone"
                ],
                "summary": "Get Phone.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Phone id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Phone"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Delete a Phone by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Phone"
                ],
                "summary": "Delete one Phone.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Phone id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Update Phone by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Phone"
                ],
                "summary": "Update Phone.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Phone id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "the body to update age rating category",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.PhoneInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Phone"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "registering a user from public access.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register a user.",
                "parameters": [
                    {
                        "description": "the body to register a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.RegisterInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.BrandInput": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "controller.ColorInput": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "controller.LoginInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controller.ModelInput": {
            "type": "object",
            "properties": {
                "brand_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "controller.PhoneInput": {
            "type": "object",
            "properties": {
                "brand_id": {
                    "type": "integer"
                },
                "color_id": {
                    "type": "integer"
                },
                "model_id": {
                    "type": "integer"
                },
                "price": {
                    "type": "string"
                },
                "storage": {
                    "type": "integer"
                }
            }
        },
        "controller.RegisterInput": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Brand": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Color": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Model": {
            "type": "object",
            "properties": {
                "brand_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Phone": {
            "type": "object",
            "properties": {
                "brand_id": {
                    "type": "integer"
                },
                "color_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "model_id": {
                    "type": "integer"
                },
                "price": {
                    "type": "string"
                },
                "storage": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
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
