package v1

import (
	"api-gateway/internal/util"
	"io"
	"net/http"

	"github.com/akanshgupta98/go-logger/v2"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	req, err := http.NewRequest("POST", "http://auth-service/auth/login", c.Request.Body)
	if err != nil {
		logger.Errorf("unable to create request: %s", err.Error())
		util.ErrorJSON(c, err, nil)
		return
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("unable to forward request to auth-service: %s", err.Error())
		util.ErrorJSON(c, err, nil)
		return
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	util.WriteJSON(c, data, resp.StatusCode)

}

func Register(c *gin.Context) {
	req, err := http.NewRequest("POST", "http://auth-service/auth/users", c.Request.Body)
	if err != nil {
		logger.Errorf("unable to create request: %s", err.Error())
		util.ErrorJSON(c, err, nil)
		return
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("unable to forward request to auth-service: %s", err.Error())
		util.ErrorJSON(c, err, nil)
		return
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	util.WriteJSON(c, data, resp.StatusCode)
}
