package handlers

import (
	"auth-service/internal/amqp"
	"auth-service/internal/service"
	"auth-service/internal/util"
	"net/http"

	"github.com/akanshgupta98/expense-manager/contracts/eventspb"
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

	eventData := &eventspb.UserCreatedEvent{
		UserId:    data.UserID,
		Email:     payload.Email,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Country:   payload.Country,
	}

	err = amqp.PublishUserCreated(eventData)
	if err != nil {
		logger.Warnf("unable to publish event on message broker: %s", err.Error())

	} else {
		logger.Infof("published event of user creation")
	}

	util.WriteJSON(c, response, http.StatusCreated)

}

// Just for testing. Will be movedd to user-service
func FetchAllUsers(c *gin.Context) {
	var result []User
	data, err := service.FetchAllUsers()
	if err != nil {
		util.ErrorJSON(c, err)
		return
	}
	for _, u := range data {
		user := User{
			// Name:     u.Name,
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
