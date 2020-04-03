// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-04-03 16:10:26.2031854 +0800 CST m=+388.629443101

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/admin/DeleteAdmins": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员"
                ],
                "summary": "批量删除管理员",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/entity.DeleteIds"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"SUCCESS\"} {\"status\": false, \"message\": \"用户名或密码不能为空\"}  {\"status\": false, \"message\": \"ERROR\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/Login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员"
                ],
                "summary": "管理员登录",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/entity.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"SUCCESS\"} {\"status\": false, \"message\": \"用户名或密码不能为空\"}  {\"status\": false, \"message\": \"ERROR\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/addAdmin": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员"
                ],
                "summary": "添加管理员",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/entity.Register"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"SUCCESS\"} {\"status\": false, \"message\": \"用户名或密码不能为空\"}  {\"status\": false, \"message\": \"ERROR\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/deleteAdmin": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员"
                ],
                "summary": "删除管理员",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"SUCCESS\"} {\"status\": false, \"message\": \"用户名或密码不能为空\"}  {\"status\": false, \"message\": \"ERROR\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/queryAdmins": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员"
                ],
                "summary": "分页查询用户(默认前100条) 并返回总记录数",
                "parameters": [
                    {
                        "type": "string",
                        "description": "页大小",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "页数",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"SUCCESS\", \"data\":{}} {\"status\": false, \"message\": \"ERROR\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/updateAdminPass": {}
    },
    "definitions": {
        "entity.DeleteIds": {
            "type": "object",
            "properties": {
                "ids": {
                    "description": "ids",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "entity.Login": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "entity.Register": {
            "type": "object",
            "properties": {
                "administrator": {
                    "description": "超级管理员 Y | N",
                    "type": "string"
                },
                "online_username": {
                    "description": "当前登录用户名",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "entity.UpdatePass": {
            "type": "object",
            "properties": {
                "newPassword": {
                    "description": "密码",
                    "type": "string"
                },
                "oldPassword": {
                    "description": "当前登录用户名",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
