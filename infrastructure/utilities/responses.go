package utils

import (
	"context"
	"loan_tracker/config"
	"log"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Response represents the structure of an HTTP response.
type Response struct {
	Message string      `json:"message" bson:"message"`   // Description or error message
	Data    interface{} `json:"data,omitempty" bson:"data,omitempty"` // Payload data, if any
}

// Status codes and corresponding messages
const (
	StatusOK                  = 200
	StatusCreated             = 201
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusInternalServerError = 500

	MsgOK                  = "Operation successful"
	MsgCreated             = "Resource created successfully"
	MsgBadRequest          = "Invalid request"
	MsgUnauthorized        = "Unauthorized access"
	MsgForbidden           = "Access forbidden"
	MsgNotFound            = "Resource not found"
	MsgInternalServerError = "Internal server error"
)

// Result sends a standardized JSON response.
func Result(httpStatusCode int, data interface{}, message string, c *gin.Context) {
	//get the request id from the context
	reqID := requestid.Get(c)
	logEntry := bson.M{
		"request_id": reqID,
		"status":     httpStatusCode,
		"message":    message,
		"data":       data,
		"timestamp":  time.Now(),
	}

	_, err := config.LogCollection.InsertOne(context.TODO(), logEntry)
	if err != nil {
		log.Printf("Failed to log to MongoDB: %v", err)
	}
	c.JSON(httpStatusCode, Response{
		Message: message,
		Data:    data,
	})
}


//custom response sends customized responses
func CustomResponse(httpStatusCode int, message string, c *gin.Context) {
	Result(httpStatusCode, nil, message, c)
}

// Success response utilities
func Success(c *gin.Context) {
	Result(StatusOK, nil, MsgOK, c)
}

func SuccessWithMessage(message string, c *gin.Context) {
	Result(StatusOK, nil, message, c)
}

func SuccessWithData(data interface{}, c *gin.Context) {
	Result(StatusOK, data, MsgOK, c)
}

func SuccessWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(StatusOK, data, message, c)
}

// Error response utilities
func Error(c *gin.Context) {
	Result(StatusInternalServerError, nil, MsgInternalServerError, c)
}

func ErrorWithMessage(message string, c *gin.Context) {
	Result(StatusInternalServerError, nil, message, c)
}

func ErrorWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(StatusInternalServerError, data, message, c)
}

func BadRequest(c *gin.Context) {
	Result(StatusBadRequest, nil, MsgBadRequest, c)
}

func Unauthorized(c *gin.Context) {
	Result(StatusUnauthorized, nil, MsgUnauthorized, c)
}

func Forbidden(c *gin.Context) {
	Result(StatusForbidden, nil, MsgForbidden, c)
}

func NotFound(c *gin.Context) {
	Result(StatusNotFound, nil, MsgNotFound, c)
}
