package v1

import (
	"fmt"
	"net/http"
	"user-service/internal/service"
	"user-service/internal/util"

	"github.com/akanshgupta98/go-logger/v2"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.RouterGroup) {
	g.GET("/", HealthCheck)
	g.POST("/user-profile", CreateProfile)
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Health": "Ok",
	})

}

func CreateProfile(c *gin.Context) {
	var request CreateProfileRequest
	err := util.ReadJSON(c, &request)
	if err != nil {
		logger.Errorf("invalid request payload: %s", err.Error())
		util.ErrorJSON(c, err, http.StatusBadRequest)
		return
	}

	data := service.CreateProfileInput{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		UserID:    int64(request.UserID),
		Country:   request.Country,
	}
	err = service.CreateProfile(data)
	if err != nil {
		fmt.Printf("unable to create user profile for id: %d. Err: %v", request.UserID, err)
		logger.Errorf("unable to create user profile for id: %d. Err: %v", request.UserID, err)
		util.ErrorJSON(c, err)
		return
	}
	response := CreateProfileResponse{
		Message: "Profile created successfully",
	}
	util.WriteJSON(c, response, http.StatusCreated)

}
