package v1

import (
	"auth-service/internal/service"
	"auth-service/internal/util"
	"net/http"

	"github.com/akanshgupta98/go-logger"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(gr *gin.RouterGroup) {
	gr.GET("/", HealthCheck)
	gr.POST("/register", Registration)
	gr.POST("/login", Login)
}

func HealthCheck(c *gin.Context) {
	logger.Debugf("Health Check invoked")
	c.JSON(http.StatusOK, gin.H{
		"server health": "OK",
	})
}

func Registration(c *gin.Context) {

	var payload RegistrationPayload
	var response RegistrationResponse
	// Read payload
	err := util.ReadJSON(c, &payload)
	if err != nil {
		util.ErrorJSON(c, err)
		return
	}

	// send to service layer for processing.
	serviceData := service.RegisterUserInput{
		Email:    payload.Email,
		Password: payload.Password,
	}
	data, err := service.RegisterUser(serviceData)
	if err != nil {
		util.ErrorJSON(c, err)
		return
	}

	response = RegistrationResponse{
		APIResponse: APIResponse{
			Message: "user registered successfully.",
			Data:    data,
		},
	}

	util.WriteJSON(c, response, http.StatusCreated)

}

func Login(c *gin.Context) {
	var payload LoginPayload
	err := util.ReadJSON(c, &payload)
	if err != nil {
		util.ErrorJSON(c, err, http.StatusBadRequest)
		return
	}
	data := service.Login{
		Email:    payload.Email,
		Password: payload.Password,
	}
	token, err := service.LoginUser(data)
	if err != nil {
		util.ErrorJSON(c, err)
		return
	}

	response := LoginResponse{
		APIResponse: APIResponse{
			Message: "login successfull",
			Data:    token,
		},
	}
	util.WriteJSON(c, response)
}
