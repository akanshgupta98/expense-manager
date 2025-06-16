package v1

import (
	"expense-service/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(mux *gin.RouterGroup) {

	mux.GET("/", health)
	mux.POST("/expense", addExpense)

}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Health": "OK",
	})
}

func addExpense(c *gin.Context) {
	var payload AddExpensePayload
	err := util.ReadJSON(c, payload)
	if err != nil {
		util.ErrorJSON(c, err, http.StatusBadRequest)
		return
	}

}
