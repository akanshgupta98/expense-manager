package handlers

import (
	"auth-service/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"server health": "OK",
	})
}

func LogMiddleWare() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		log.Printf("API: %s, Method: %s, Host: %s", ctx.Request.URL, ctx.Request.Method, ctx.Request.Host)
		ctx.Next()

	}

}

func Registration(c *gin.Context) {

	var payload RegistrationPayload
	var response RegistrationResponse
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.Error(err)
		return
	}

	serviceData := service.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}
	err = service.RegisterUser(serviceData)
	if err != nil {

		c.Error(err)
		return
	}

	response = RegistrationResponse{
		APIResponse: APIResponse{
			Error:   false,
			Message: "user registered successfully.",
			Data:    payload,
		},
	}
	c.JSON(http.StatusCreated, response)

}
