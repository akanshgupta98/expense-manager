package handlers

import (
	"auth-service/internal/service"
	"auth-service/internal/util"
	"net/http"

	"github.com/akanshgupta98/go-logger"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	logger.Debugf("Health Check invoked")
	c.JSON(http.StatusOK, gin.H{
		"server health": "OK",
	})
}

func LogMiddleWare() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		logger.Infof("API: %s, Method: %s, Host: %s", ctx.Request.URL, ctx.Request.Method, ctx.Request.Host)
		ctx.Next()

	}

}

func Registration(c *gin.Context) {

	var payload RegistrationPayload
	var response RegistrationResponse
	err := util.ReadJSON(c, &payload)
	if err != nil {
		util.ErrorJSON(c, err)
		return
	}

	serviceData := service.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}
	err = service.RegisterUser(serviceData)
	if err != nil {
		util.ErrorJSON(c, err)
		return
	}

	response = RegistrationResponse{
		APIResponse: APIResponse{
			Message: "user registered successfully.",
			Data:    payload,
		},
	}
	util.WriteJSON(c, response, http.StatusCreated)

}

func FetchAllUsers(c *gin.Context) {
	var result []User
	data, err := service.FetchAllUsers()
	if err != nil {
		util.ErrorJSON(c, err)
		return
	}
	for _, u := range data {
		user := User{
			Name:     u.Name,
			Password: u.Password,
			Email:    u.Email,
		}
		result = append(result, user)
	}
	response := FetchUsersResponse{
		Users: result,
	}
	util.WriteJSON(c, response)

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
