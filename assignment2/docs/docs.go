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
        "/orders": {
            "get": {
                "description": "Get all orders",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "List orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetAllOrdersResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.MessageErrData"
                        }
                    }
                }
            },
            "post": {
                "description": "Create an order by json",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Create order",
                "parameters": [
                    {
                        "description": "Create order request body",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.NewOrderResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.MessageErrData"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errs.MessageErrData"
                        }
                    }
                }
            }
        },
        "/orders/{id}": {
            "get": {
                "description": "Get an order by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Find order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetOrderByIDResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.MessageErrData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errs.MessageErrData"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an order by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Delete order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.DeleteOrderByIDResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.MessageErrData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errs.MessageErrData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.MessageErrData"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update an order by json",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Update order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update order request body",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetOrderByIDResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.MessageErrData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errs.MessageErrData"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errs.MessageErrData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.DeleteOrderByIDResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Order with id 69 has been successfully deleted"
                },
                "statusCode": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "dto.GetAllOrdersResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.OrderData"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "statusCode": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "dto.GetOrderByIDResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/dto.OrderData"
                },
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "statusCode": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "dto.ItemData": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string",
                    "example": "2023-03-19T18:55:00+07:00"
                },
                "description": {
                    "type": "string",
                    "example": "Ini adalah sebuah barang yang dipesan"
                },
                "id": {
                    "type": "integer",
                    "example": 2
                },
                "itemCode": {
                    "type": "string",
                    "example": "BRNG-001"
                },
                "orderId": {
                    "type": "integer",
                    "example": 1
                },
                "quantity": {
                    "type": "integer",
                    "example": 1
                },
                "updatedAt": {
                    "type": "string",
                    "example": "2023-03-19T18:55:00+07:00"
                }
            }
        },
        "dto.NewItemRequest": {
            "type": "object",
            "required": [
                "description",
                "itemCode",
                "quantity"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Ini adalah sebuah barang yang dipesan"
                },
                "itemCode": {
                    "type": "string",
                    "example": "BRNG-001"
                },
                "quantity": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "dto.NewOrderRequest": {
            "type": "object",
            "required": [
                "customerName",
                "items"
            ],
            "properties": {
                "customerName": {
                    "type": "string",
                    "example": "Taufik Hidayat"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.NewItemRequest"
                    }
                },
                "orderedAt": {
                    "type": "string",
                    "example": "2023-03-19T18:55:00+07:00"
                }
            }
        },
        "dto.NewOrderResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/dto.NewOrderRequest"
                },
                "message": {
                    "type": "string",
                    "example": "Order with id 69 has been successfully created"
                },
                "statusCode": {
                    "type": "integer",
                    "example": 201
                }
            }
        },
        "dto.OrderData": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string",
                    "example": "2023-03-19T18:55:00+07:00"
                },
                "customerName": {
                    "type": "string",
                    "example": "Taufik Hidayat"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.ItemData"
                    }
                },
                "updatedAt": {
                    "type": "string",
                    "example": "2023-03-19T18:55:00+07:00"
                }
            }
        },
        "errs.MessageErrData": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "BAD_REQUEST"
                },
                "message": {
                    "type": "string",
                    "example": "This is an error message"
                },
                "statusCode": {
                    "type": "integer",
                    "example": 400
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
