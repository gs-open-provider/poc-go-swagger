// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a sample server for a News Feed.",
    "title": "Sample NewsFeed API",
    "contact": {
      "name": "NEWS API Support"
    },
    "version": "1.0.0"
  },
  "basePath": "/news-api/v1",
  "paths": {
    "/test": {
      "$ref": "./paths/test.yml"
    }
  },
  "definitions": {
    "text-response": {
      "$ref": "./definitions/text-response.yml"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a sample server for a News Feed.",
    "title": "Sample NewsFeed API",
    "contact": {
      "name": "NEWS API Support"
    },
    "version": "1.0.0"
  },
  "basePath": "/news-api/v1",
  "paths": {
    "/test": {
      "get": {
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/textResponse"
            }
          }
        }
      },
      "post": {
        "parameters": [
          {
            "name": "test-obj",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/textResponse"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/textResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "text-response": {
      "$ref": "#/definitions/textResponse"
    },
    "textResponse": {
      "type": "object",
      "properties": {
        "name": {
          "message": "string"
        }
      }
    }
  }
}`))
}
