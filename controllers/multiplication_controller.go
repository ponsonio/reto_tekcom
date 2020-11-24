package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ponsonio/reto_tekcom/services"
	"log"
	"net/http"
)

//validations can be here
type MultiplyRequest struct {
	A int64 `json:"a" binding:"required"`
	B int64 `json:"b" binding:"required"`
}

type MultiplyResponce struct {
	Result string `json:"result"`
}

var service services.Multiplicator

//this is to facilitate mock injection, but in this case _real_ obj is used per time constraints
func init() {
	service, _ = services.NewMultiplicator()
}

func Multiply(c *gin.Context) {
	req := MultiplyRequest{}
	err := c.ShouldBindJSON(&req)

	if err != nil {
		log.Println(fmt.Printf("Invalid request : %v", err))
		c.JSON(http.StatusBadRequest, "invalid request")
	}
	res, multErr := service.Multiply(req.A, req.B)

	if multErr != nil {
		log.Println(fmt.Printf("Error multiplying : %v", multErr))
		c.JSON(http.StatusInternalServerError, "An unknown error occurred")
	}
	log.Printf("mult res. %v", res)
	resp := MultiplyResponce{
		Result: res.String(),
	}
	c.JSON(http.StatusOK, resp)
}
