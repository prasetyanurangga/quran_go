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
            "name": "API Support",
            "email": "angganurprasetya4@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ayah/get_by_id_ayah/{id}": {
            "get": {
                "description": "Get Ayah By Id Ayah",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ayah"
                ],
                "summary": "Get Ayah By Id Ayah",
                "operationId": "get-ayah-by-id-ayah",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "IdAyah",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/config.JSONSuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "count": {
                                            "type": "integer"
                                        },
                                        "data": {
                                            "$ref": "#/definitions/model.Ayah"
                                        },
                                        "success": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/config.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "count": {
                                            "type": "integer"
                                        },
                                        "success": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/ayah/get_by_id_surah/{id}": {
            "get": {
                "description": "Get Ayah By Id Surah",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ayah"
                ],
                "summary": "Get Ayah By Id Surah",
                "operationId": "get-ayah-by-id-surah",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "IdSurah",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/config.JSONSuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "count": {
                                            "type": "integer"
                                        },
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Ayah"
                                            }
                                        },
                                        "success": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/config.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "count": {
                                            "type": "integer"
                                        },
                                        "success": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/ayah/get_by_juz/{juz}": {
            "get": {
                "description": "Get Ayah By Juz",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ayah"
                ],
                "summary": "Get Ayah By Juz",
                "operationId": "get-ayah-by-juz",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Juz",
                        "name": "juz",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/config.JSONSuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "count": {
                                            "type": "integer"
                                        },
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Ayah"
                                            }
                                        },
                                        "success": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/config.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "count": {
                                            "type": "integer"
                                        },
                                        "success": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/ayah/search/{keyword}": {
            "get": {
                "description": "Search Ayah",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ayah"
                ],
                "summary": "Search Ayah",
                "operationId": "search-ayah",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Keyword",
                        "name": "keyword",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/config.JSONSuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "count": {
                                            "type": "integer"
                                        },
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Ayah"
                                            }
                                        },
                                        "success": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/config.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "count": {
                                            "type": "integer"
                                        },
                                        "success": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/surah/get_all/": {
            "get": {
                "description": "Get All Surah",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Surah"
                ],
                "summary": "Get All Surah",
                "operationId": "get-all-surah",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/config.JSONSuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "count": {
                                            "type": "integer"
                                        },
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Surah"
                                            }
                                        },
                                        "success": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/config.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "count": {
                                            "type": "integer"
                                        },
                                        "success": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/surah/get_by_id/{id}": {
            "get": {
                "description": "Get Surah By Id Surah",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Surah"
                ],
                "summary": "Get Surah By Id Surah",
                "operationId": "get-surah-by-id-surah",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id Surah",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/config.JSONSuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "count": {
                                            "type": "integer"
                                        },
                                        "data": {
                                            "$ref": "#/definitions/model.Surah"
                                        },
                                        "success": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/config.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "count": {
                                            "type": "integer"
                                        },
                                        "success": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/surah/get_by_id_with_ayah/{id}": {
            "get": {
                "description": "Get Surah By Id Surah With Ayah",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Surah"
                ],
                "summary": "Get Surah By Id Surah With Ayah",
                "operationId": "get-surah-by-id-surah-with-ayah",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id Surah",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/config.JSONSuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "count": {
                                            "type": "integer"
                                        },
                                        "data": {
                                            "$ref": "#/definitions/model.Surah"
                                        },
                                        "success": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/config.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "count": {
                                            "type": "integer"
                                        },
                                        "success": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/surah/search/{keyword}": {
            "get": {
                "description": "Get Search Surah",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Surah"
                ],
                "summary": "Get Search Surah",
                "operationId": "search-surah",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Keyword",
                        "name": "keyword",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/config.JSONSuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "count": {
                                            "type": "integer"
                                        },
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Surah"
                                            }
                                        },
                                        "success": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/config.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "count": {
                                            "type": "integer"
                                        },
                                        "success": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "config.JSONErrorResult": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer",
                    "example": 0
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "config.JSONSuccessResult": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer",
                    "example": 12
                },
                "data": {
                    "type": "object"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "model.Ayah": {
            "type": "object",
            "properties": {
                "arabText": {
                    "type": "string"
                },
                "audio_1": {
                    "type": "string"
                },
                "audio_2": {
                    "type": "string"
                },
                "audio_3": {
                    "type": "string"
                },
                "enText": {
                    "type": "string"
                },
                "hizbQuarter": {
                    "type": "integer"
                },
                "idAyahInQuran": {
                    "type": "integer"
                },
                "idAyahInSurah": {
                    "type": "integer"
                },
                "idSurah": {
                    "type": "integer"
                },
                "indoText": {
                    "type": "string"
                },
                "juz": {
                    "type": "integer"
                },
                "manzil": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "ruku": {
                    "type": "integer"
                },
                "sajda_obligatory": {
                    "type": "boolean"
                },
                "sajda_recommended": {
                    "type": "boolean"
                },
                "tafsirLong": {
                    "type": "string"
                },
                "tafsirShort": {
                    "type": "string"
                },
                "transliterationEn": {
                    "type": "string"
                }
            }
        },
        "model.Surah": {
            "type": "object",
            "properties": {
                "ayah": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Ayah"
                    }
                },
                "idSurah": {
                    "type": "integer"
                },
                "nameLong": {
                    "type": "string"
                },
                "nameShort": {
                    "type": "string"
                },
                "nameTranslationEn": {
                    "type": "string"
                },
                "nameTranslationId": {
                    "type": "string"
                },
                "nameTransliterationEn": {
                    "type": "string"
                },
                "nameTransliterationId": {
                    "type": "string"
                },
                "numberOfAyah": {
                    "type": "integer"
                },
                "revelationArab": {
                    "type": "string"
                },
                "revelationEn": {
                    "type": "string"
                },
                "revelationId": {
                    "type": "string"
                },
                "tafsir": {
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
	Title:       "Quran Go",
	Description: "This is quran rest api",
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
