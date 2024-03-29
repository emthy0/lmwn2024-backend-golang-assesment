// Package openapi_spec Code generated by swaggo/swag. DO NOT EDIT
package openapi_spec

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Kuranasaki",
            "email": "kongphop.c@kuranasaki.work"
        },
        "license": {
            "name": "NIGGY License",
            "url": "https://github.com/Admin-OR-1-1/2110336-SE2-Crafty/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/covid/summary": {
            "get": {
                "description": "Generate Summary of Covid Patient in Thailand grouped by Province and Age",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Covid"
                ],
                "responses": {
                    "200": {
                        "description": "Fetch and Summarize Covid Data Success",
                        "schema": {
                            "$ref": "#/definitions/covidsummaryapi.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/covidsummaryapi.APIResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "covidsummaryapi.APIResponse": {
            "type": "object",
            "properties": {
                "Age": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "Province": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                }
            }
        },
        "covidsummaryapi.APIResponseError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "LMWN Assesment Covid Summary API",
	Description:      "This is a covid summary API server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
