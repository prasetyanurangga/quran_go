package config

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type JSONSuccessResult struct {
	Success bool `json:"success" example:"true"`
	Data interface{} `json:"data"`
	Count int `json:"count" example:"12"`
}

type JSONErrorResult struct {
	Success bool `json:"success" example:"false"`
	Count int `json:"count" example:"0"`
}

func SuccessResponse(c *gin.Context, data interface{}, count int) {
	c.JSON(http.StatusOK, JSONSuccessResult{
		Success: true,
		Data: data,
		Count: count,
	})
}

func ErrorResponse(c *gin.Context, count int) {
	c.JSON(http.StatusInternalServerError, JSONErrorResult{
		Success: false,
		Count: count,
	})
}