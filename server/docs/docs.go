// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-04-08 10:29:49.1316794 +0800 CST m=+442.666881001

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
                        "description": "desc",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/entity.ResponseData"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/addGoodsType": {
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
                "summary": "添加商品种类",
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
                            "$ref": "#/definitions/entity.AddGoodsType"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/entity.ResponseData"
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
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/entity.ResponseData"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/deleteAdmins": {
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
                        "description": "desc",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/entity.ResponseData"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/deleteGoodsType": {
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
                "summary": "下架商品",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/entity.ResponseData"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/deleteGoodsTypes": {
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
                "summary": "批量下架商品",
                "parameters": [
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
                        "description": "desc",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/entity.ResponseData"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/login": {
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
                        "description": "desc",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/entity.ResponseData"
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
                        "description": "desc",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/entity.ResponseData"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/queryByGoodsTypeID": {
            "get": {
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
                "summary": "查看商品种类详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商品种类ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/entity.ResponseData"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/queryGoodsTypesByLimitOffset": {
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
                "summary": "分页查询商品种类(默认前100条) 并返回总记录数",
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
                        "description": "desc",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/entity.ResponseData"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/updateAdminPass": {},
        "/api/v1/admin/updateGoodsType": {
            "put": {
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
                "summary": "修改商品种类信息",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/entity.UpdateGoodsType"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/entity.ResponseData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.AddGoodsType": {
            "type": "object",
            "properties": {
                "goods_batch_number": {
                    "description": "生产批号",
                    "type": "string"
                },
                "goods_date": {
                    "description": "生产日期",
                    "type": "string"
                },
                "goods_image": {
                    "description": "商品图片",
                    "type": "string"
                },
                "goods_name": {
                    "description": "商品名称",
                    "type": "string"
                },
                "goods_prince": {
                    "description": "商品销售价",
                    "type": "number"
                },
                "goods_specs": {
                    "description": "商品规格 1.盒 2.瓶 3.支",
                    "type": "string",
                    "default": "1",
                    "enum": [
                        "1",
                        "2",
                        "3"
                    ]
                },
                "goods_state": {
                    "description": "商品状态 1.下架  2.在售",
                    "type": "string",
                    "default": "2",
                    "enum": [
                        "1",
                        "2",
                        "3"
                    ]
                },
                "goods_unitprince": {
                    "description": "商品成本价",
                    "type": "number"
                }
            }
        },
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
        "entity.ResponseData": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "entity.UpdateGoodsType": {
            "type": "object",
            "properties": {
                "goods_batch_number": {
                    "description": "生产批号",
                    "type": "string"
                },
                "goods_date": {
                    "description": "生产日期",
                    "type": "string"
                },
                "goods_id": {
                    "description": "商品ID",
                    "type": "integer"
                },
                "goods_image": {
                    "description": "商品图片",
                    "type": "string"
                },
                "goods_name": {
                    "description": "商品名称",
                    "type": "string"
                },
                "goods_prince": {
                    "description": "商品销售价",
                    "type": "number"
                },
                "goods_specs": {
                    "description": "商品规格 1.盒 2.瓶 3.支",
                    "type": "string",
                    "default": "1",
                    "enum": [
                        "1",
                        "2",
                        "3"
                    ]
                },
                "goods_state": {
                    "description": "商品状态 1.下架  2.在售",
                    "type": "string",
                    "default": "2",
                    "enum": [
                        "1",
                        "2",
                        "3"
                    ]
                },
                "goods_unitprince": {
                    "description": "商品成本价",
                    "type": "number"
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
                    "description": "旧密码",
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
