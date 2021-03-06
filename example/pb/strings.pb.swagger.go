// Code generated by protoc-gen-go-http-server.
// source: pb/strings.proto

package strings

import (
	"net/http"

	"github.com/doroginin/protobuf/protoc-gen-go-http-server/swagger"
)

var SwaggerJSONHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	w.WriteHeader(http.StatusOK)
	w.Write(_swaggerJSON)
})

var SwaggerUIHandler = swaggerui.NewHTTPHandler()

var _swaggerJSON = []byte(`{
  "swagger": "2.0",
  "info": {
    "title": "pb/strings.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/strings/to_lower": {
      "post": {
        "operationId": "ToLower",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/StringResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/StringRequest"
            }
          }
        ],
        "tags": [
          "Strings"
        ]
      }
    },
    "/strings/to_upper/{str}": {
      "get": {
        "operationId": "ToUpper",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/StringResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "str",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Strings"
        ]
      }
    },
    "/v1/strings/to_upper/{str}": {
      "get": {
        "operationId": "ToUpper2",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/StringResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "str",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Strings"
        ]
      }
    }
  },
  "definitions": {
    "StringRequest": {
      "type": "object",
      "properties": {
        "str": {
          "type": "string"
        }
      }
    },
    "StringResponse": {
      "type": "object",
      "properties": {
        "str": {
          "type": "string"
        }
      }
    }
  }
}
`)
