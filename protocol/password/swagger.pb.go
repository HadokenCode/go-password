package password

const (
	swagger = `{
  "swagger": "2.0",
  "info": {
    "title": "password.proto",
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
    "/v1/password": {
      "post": {
        "operationId": "Encode",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/passwordEncodedPasswordRes"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/passwordPasswordReq"
            }
          }
        ],
        "tags": [
          "Password"
        ]
      }
    },
    "/v1/validate": {
      "post": {
        "operationId": "Validate",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/passwordPasswordValidationRes"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/passwordPasswordReq"
            }
          }
        ],
        "tags": [
          "Password"
        ]
      }
    }
  },
  "definitions": {
    "passwordEncodedPasswordRes": {
      "type": "object",
      "properties": {
        "error": {
          "$ref": "#/definitions/passwordError"
        },
        "hash": {
          "type": "string"
        }
      }
    },
    "passwordError": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "passwordPasswordReq": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string"
        },
        "hash": {
          "type": "string"
        }
      }
    },
    "passwordPasswordValidationRes": {
      "type": "object",
      "properties": {
        "error": {
          "$ref": "#/definitions/passwordError"
        },
        "valid": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    }
  }
}
`
)
