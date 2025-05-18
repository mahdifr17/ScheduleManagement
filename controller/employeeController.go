package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupEmployeeController(r *gin.Engine) {
	employeeAccess := r.Group("/employee", gin.BasicAuth(gin.Accounts{
		"eee": "eee", // user:eee password:eee
	}))

	// get assigned shift
	employeeAccess.GET("/shift/assigned", func(c *gin.Context) {
		c.String(http.StatusOK, http.StatusText(http.StatusOK))
	})

	// get available shift
	employeeAccess.GET("/shift/available", func(c *gin.Context) {
		c.String(http.StatusOK, http.StatusText(http.StatusOK))
	})

	// create shift request
	employeeAccess.POST("/shift-request", func(c *gin.Context) {
		c.String(http.StatusOK, http.StatusText(http.StatusOK))
	})

	// get shift request -> get all
	employeeAccess.GET("/shift-request", func(c *gin.Context) {
		c.String(http.StatusOK, http.StatusText(http.StatusOK))
	})
}
