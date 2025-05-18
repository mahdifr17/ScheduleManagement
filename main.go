package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdifr17/ScheduleManagement/controller"
	"github.com/mahdifr17/ScheduleManagement/repository"
)

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// load repository
	shiftRP := repository.NewShiftRepositoryMemory()
	shiftRequestRP := repository.NewShiftRequestRepositoryMemory()

	// healthcheck
	r.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, http.StatusText(http.StatusOK))
	})

	// login
	r.POST("/login", func(c *gin.Context) {
		c.String(http.StatusOK, http.StatusText(http.StatusOK))
	})

	// route admin access
	controller.SetupAdminController(r, shiftRP, shiftRequestRP)

	// route employee access
	controller.SetupEmployeeController(r, shiftRP, shiftRequestRP)

	return r
}
