package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupAdminController(r *gin.Engine) {
	adminAccess := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"aaa": "aaa", // user:foo password:bar
	}))

	// create shift
	adminAccess.POST("/shift", func(c *gin.Context) {
		c.String(http.StatusOK, http.StatusText(http.StatusOK))
	})

	// edit shift
	adminAccess.PUT("/shift", func(c *gin.Context) {
		c.String(http.StatusOK, http.StatusText(http.StatusOK))
	})

	// get all shift request
	adminAccess.GET("/shift", func(c *gin.Context) {
		c.String(http.StatusOK, http.StatusText(http.StatusOK))
	})

	// get all finalized shift
	adminAccess.GET("/shift/final", func(c *gin.Context) {
		c.String(http.StatusOK, http.StatusText(http.StatusOK))
	})

	// approve shift request
	adminAccess.POST("/shift-request/approve", func(c *gin.Context) {
		c.String(http.StatusOK, http.StatusText(http.StatusOK))
	})

	// reject shift request
	adminAccess.POST("/shift-request/reject", func(c *gin.Context) {
		c.String(http.StatusOK, http.StatusText(http.StatusOK))
	})
}
