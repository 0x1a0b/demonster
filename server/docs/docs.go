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
        "termsOfService": "http://localhost:8080/tos.html",
        "contact": {
            "name": "Demonster API Support",
            "url": "https://github.com/0x1a0b/demonster/issues",
            "email": "support@localhost.local"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/0x1a0b/demonster/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/actions/request": {
            "post": {
                "description": "Test Endpoint for Smuggling Verification",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Echoes body and headers",
                "operationId": "get-pinged",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.ActionsResponse"
                        }
                    }
                }
            }
        },
        "/alive": {
            "get": {
                "description": "Test Backend Aliveness",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Ping Pong the Server",
                "operationId": "get-pinged",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Pong"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.ActionsResponse": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "body_len": {
                    "type": "integer"
                },
                "headers": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "main.Pong": {
            "type": "object",
            "properties": {
                "response": {
                    "type": "string"
                },
                "time": {
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
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "Demonster API",
	Description: "This is an example api to demonstrate things",
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
