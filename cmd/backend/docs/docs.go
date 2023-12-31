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
        "/admin/protocols": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This will list all available protocols in system including their corresponding tasks and their dependencies.\nMethod should not be used by regular users, because we do not want to allow them to see full list of tasks",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "protocol",
                    "admin"
                ],
                "summary": "Shows list of available protocols",
                "responses": {
                    "200": {
                        "description": "Full list of protocols",
                        "schema": {
                            "$ref": "#/definitions/main.ListResponse-main_ProtocolResponseAdmin"
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
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
                    "protocol",
                    "admin"
                ],
                "summary": "Creates or updates protocol in database",
                "parameters": [
                    {
                        "description": "Protocol body",
                        "name": "protocol",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/protocol.Protocol"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/protocols/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This will list all available protocols in system including their corresponding tasks and their dependencies.\nMethod should not be used by regular users, because we do not want to allow them to see full list of tasks",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "protocol",
                    "admin"
                ],
                "summary": "Shows protocol details with tasks",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Protocols ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Protocol details",
                        "schema": {
                            "$ref": "#/definitions/main.ProtocolResponseAdmin"
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/protocols": {
            "get": {
                "description": "This will list all available protocols in system including their corresponding tasks and their dependencies.\nMethod should not be used by regular users, because we do not want to allow them to see full list of tasks",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "protocol"
                ],
                "summary": "Shows list of available protocols",
                "responses": {
                    "200": {
                        "description": "Full list of protocols",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.ProtocolResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/protocols/progress": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "protocol"
                ],
                "summary": "Get protocol for startup by its id or vanity",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Startup ID",
                        "name": "startup_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Startup Vanity",
                        "name": "startup_vanity",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "1|0 show tasks without content",
                        "name": "skip_content",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.ProtocolProgress"
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/protocols/{id}": {
            "get": {
                "description": "This will list all available protocols in system including their corresponding tasks and their dependencies.\nMethod should not be used by regular users, because we do not want to allow them to see full list of tasks",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "protocol"
                ],
                "summary": "Shows protocols details",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Protocols ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Protocol details",
                        "schema": {
                            "$ref": "#/definitions/main.ProtocolResponse"
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/protocols/{id}/roles": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "protocol"
                ],
                "summary": "Shows list of available roles for specified protocol",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Protocol ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "All roles available in protocol",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/startup/{startup_id}/assign-protocol": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "protocol"
                ],
                "summary": "Assign protocol to startup",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Startup ID",
                        "name": "startup_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.assignProtocolRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/startup/{startup_id}/protocol-task/{task_id}": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "protocol"
                ],
                "summary": "Mark task as complete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Startup ID",
                        "name": "startup_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.completeTaskRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/startup/{startup_id}/role/{role_id}/protocol-roles": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "protocol"
                ],
                "summary": "Assign protocol-roles to startup-role",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Startup ID",
                        "name": "startup_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Role ID",
                        "name": "role_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.assignProtocolRolesRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/startups": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "startup"
                ],
                "summary": "Startup list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.ListResponse-main_StartupResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/startups/details": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "startup"
                ],
                "summary": "Startup details",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Startup ID",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Startup Vanity",
                        "name": "vanity",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.StartupResponse"
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.ListResponse-main_ProtocolResponseAdmin": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.ProtocolResponseAdmin"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "main.ListResponse-main_StartupResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.StartupResponse"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "main.ProtocolProgress": {
            "type": "object",
            "properties": {
                "complete": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.TaskResponse"
                    }
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "FinTech"
                },
                "pending": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.TaskResponse"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "main.ProtocolResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "FinTech"
                },
                "sample_tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/protocol.Task"
                    }
                },
                "total": {
                    "type": "integer",
                    "example": 30
                }
            }
        },
        "main.ProtocolResponseAdmin": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "FinTech"
                },
                "tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/protocol.Task"
                    }
                }
            }
        },
        "main.StartupResponse": {
            "type": "object",
            "properties": {
                "applicants": {
                    "type": "integer",
                    "example": 333
                },
                "avatar": {
                    "$ref": "#/definitions/startup.Image"
                },
                "company_type": {
                    "type": "string",
                    "example": "fintech"
                },
                "cover": {
                    "$ref": "#/definitions/startup.Image"
                },
                "deadline": {
                    "type": "integer",
                    "example": 10
                },
                "description": {
                    "type": "string",
                    "example": "Greatest startup engine"
                },
                "equity_min": {
                    "type": "number",
                    "example": 1
                },
                "equity_total": {
                    "type": "number",
                    "example": 25
                },
                "followers": {
                    "type": "integer",
                    "example": 100
                },
                "icon": {
                    "$ref": "#/definitions/startup.Image"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "logo": {
                    "$ref": "#/definitions/startup.Image"
                },
                "name": {
                    "type": "string",
                    "example": "Gusher"
                },
                "pitch": {
                    "type": "string",
                    "example": "Once upon a time in a galaxy far-far away..."
                },
                "title": {
                    "type": "string",
                    "example": "Gusher"
                },
                "topics": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/startup.Topic"
                    }
                },
                "user_description": {
                    "type": "string",
                    "example": "Hyperspace Pilot"
                },
                "user_equity": {
                    "type": "number",
                    "example": 10
                },
                "user_title": {
                    "type": "string",
                    "example": "Founder"
                },
                "vanity": {
                    "type": "string",
                    "example": "gusher"
                },
                "video": {
                    "$ref": "#/definitions/startup.Video"
                },
                "views": {
                    "type": "integer",
                    "example": 76633
                }
            }
        },
        "main.TaskResponse": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string",
                    "example": "html content..."
                },
                "id": {
                    "type": "integer",
                    "example": 123
                },
                "startupRole": {
                    "type": "string",
                    "example": "developer"
                },
                "title": {
                    "type": "string",
                    "example": "Develop company website"
                }
            }
        },
        "main.assignProtocolRequest": {
            "type": "object",
            "properties": {
                "protocol_id": {
                    "type": "integer"
                }
            }
        },
        "main.assignProtocolRolesRequest": {
            "type": "object",
            "properties": {
                "protocol_roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "main.completeTaskRequest": {
            "type": "object",
            "properties": {
                "complete": {
                    "type": "boolean"
                }
            }
        },
        "protocol.Protocol": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "SAAS Company"
                },
                "tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/protocol.Task"
                    }
                }
            }
        },
        "protocol.Task": {
            "type": "object",
            "properties": {
                "complete_fn": {
                    "type": "string",
                    "example": "websiteDeveloped"
                },
                "content": {
                    "type": "string",
                    "example": "html content..."
                },
                "dependsOn": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    },
                    "example": [
                        123,
                        444,
                        653,
                        4
                    ]
                },
                "id": {
                    "type": "integer",
                    "example": 123
                },
                "startupRole": {
                    "type": "string",
                    "example": "developer"
                },
                "title": {
                    "type": "string",
                    "example": "Develop company website"
                }
            }
        },
        "startup.Image": {
            "type": "object",
            "properties": {
                "resized": {
                    "type": "string"
                },
                "src": {
                    "type": "string"
                }
            }
        },
        "startup.Topic": {
            "type": "object",
            "properties": {
                "featured": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "industry": {
                    "type": "integer"
                },
                "interest": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "skill": {
                    "type": "integer"
                }
            }
        },
        "startup.Video": {
            "type": "object",
            "properties": {
                "job": {
                    "type": "string"
                },
                "src": {
                    "type": "string"
                },
                "src720": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Description for what is this security definition being used",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "3.0",
	Host:             "",
	BasePath:         "/v3",
	Schemes:          []string{},
	Title:            "Gusher backend API",
	Description:      "Gusher backend api",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
