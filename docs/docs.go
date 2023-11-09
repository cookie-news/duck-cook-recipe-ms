// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/page/{page}": {
            "get": {
                "description": "Retorna as receita paginadas",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe"
                ],
                "summary": "Retorna as receita paginadas",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Número da page",
                        "name": "page",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Recipe name",
                        "name": "nameRecipe",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Número da page",
                        "name": "nameIngredin",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Pagination"
                        }
                    }
                }
            }
        },
        "/recipe": {
            "put": {
                "description": "Atualizar a receita",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe"
                ],
                "summary": "Atualizar a receita",
                "parameters": [
                    {
                        "description": "Dados da receita",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Recipe"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Recipe"
                        }
                    }
                }
            },
            "post": {
                "description": "Adicionar uma nova receita",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe"
                ],
                "summary": "Adicionar nova receita",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID da Receita",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID do Usuário",
                        "name": "idUser",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Título da Receita",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Descrição da Receita",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Tempo de Preparação da Receita",
                        "name": "preparationTime",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Metodo de Preparação da Receita",
                        "name": "preparationMethod",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Imagem da Receita",
                        "name": "images",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ingredientes da Receita",
                        "name": "ingredients",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Recipe"
                        }
                    }
                }
            }
        },
        "/recipe/{id}": {
            "get": {
                "description": "Retorna a receita",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe"
                ],
                "summary": "Retorna a receita",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Recipe"
                        }
                    }
                }
            },
            "delete": {
                "description": "Adicionar uma nova receita",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe"
                ],
                "summary": "Adicionar nova receita",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/recipe/{id}/comment": {
            "get": {
                "description": "Comentarios da receita",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment-recipe"
                ],
                "summary": "Comentarios da receita",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe ID",
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
                                "$ref": "#/definitions/entity.CommentRecipe"
                            }
                        }
                    }
                }
            }
        },
        "/recipe/{id}/like": {
            "get": {
                "description": "likes da receita",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "like-recipe"
                ],
                "summary": "Likes da receita",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "int"
                        }
                    }
                }
            }
        },
        "/user/{id}/recipe": {
            "get": {
                "description": "Retonar as receitas do usuário",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe"
                ],
                "summary": "Retonar as receitas do usuário",
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
                            "$ref": "#/definitions/entity.Recipe"
                        }
                    }
                }
            }
        },
        "/user/{id}/recipe/{idRecipe}/comment": {
            "post": {
                "description": "Comenta na receita baseado no usuário",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment-recipe"
                ],
                "summary": "Comenta na receita",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "idRecipe",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Comentário",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.CommentRecipe"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Recipe"
                        }
                    }
                }
            }
        },
        "/user/{id}/recipe/{idRecipe}/comment/{idComment}": {
            "delete": {
                "description": "Delete o comentário da receita",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment-recipe"
                ],
                "summary": "Deleta o comentário da receita",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Recipe ID",
                        "name": "idRecipe",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Comment ID",
                        "name": "idComment",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/{id}/recipe/{idRecipe}/like": {
            "post": {
                "description": "Like na receita baseado no usuário",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "like-recipe"
                ],
                "summary": "Like na receita",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "idRecipe",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Recipe"
                        }
                    }
                }
            }
        },
        "/user/{id}/recipe/{idRecipe}/like/{idLike}": {
            "delete": {
                "description": "Delete o like da receita",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "like-recipe"
                ],
                "summary": "Deleta o like da receita",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Recipe ID",
                        "name": "idRecipe",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Like ID",
                        "name": "idLike",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.CommentRecipe": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "entity.Ingredients": {
            "type": "object",
            "properties": {
                "measure": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "qty": {
                    "type": "number"
                }
            }
        },
        "entity.Pagination": {
            "type": "object",
            "properties": {
                "currentPage": {
                    "type": "integer"
                },
                "items": {},
                "next": {
                    "type": "integer"
                },
                "previous": {
                    "type": "integer"
                },
                "recordPerPage": {
                    "type": "integer"
                },
                "totalPage": {
                    "type": "integer"
                }
            }
        },
        "entity.Recipe": {
            "type": "object",
            "required": [
                "description",
                "idUser",
                "preparationMethod",
                "preparationTime",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "idUser": {
                    "type": "string"
                },
                "ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Ingredients"
                    }
                },
                "preparationMethod": {
                    "type": "string"
                },
                "preparationTime": {
                    "type": "integer",
                    "format": "int64",
                    "example": 600
                },
                "title": {
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
