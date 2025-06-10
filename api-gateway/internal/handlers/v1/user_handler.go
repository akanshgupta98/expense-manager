package v1

import (
	"api-gateway/internal/util"
	"io"
	"net/http"

	"github.com/akanshgupta98/go-logger/v2"
	"github.com/gin-gonic/gin"
)

func FetchUsers(c *gin.Context) {

	req, err := http.NewRequest("GET", "http://user-service/api/user/v1/", c.Request.Body)
	if err != nil {
		logger.Errorf("unable to create request for user-service: %s", err.Error())
		util.ErrorJSON(c, err, nil)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Errorf("unable to send request to user-service: %s", err.Error())
		util.ErrorJSON(c, err, nil)
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	util.WriteJSON(c, data, resp.StatusCode)

}
