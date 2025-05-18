package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdifr17/ScheduleManagement/repository"
	"github.com/mahdifr17/ScheduleManagement/usecase"
)

type EmployeeController struct {
	router          *gin.Engine
	employeeShiftUC usecase.EmployeeShiftUsecase
}

func SetupEmployeeController(
	router *gin.Engine,
	shiftRP repository.ShiftRepository,
	shiftRequestRP repository.ShiftRequestRepository,
) {
	c := new(EmployeeController)
	c.router = router
	c.employeeShiftUC = usecase.NewEmployeeShiftUsecaseV1(shiftRP, shiftRequestRP)
	c.setupController()
}

func (c *EmployeeController) setupController() {
	employeeAccess := c.router.Group("/employee", gin.BasicAuth(gin.Accounts{
		"eee": "eee", // user:eee password:eee
	}))

	// get assigned shift
	employeeAccess.GET("/shift/assigned", c.getAssignedShift())
	// get available shift
	employeeAccess.GET("/shift/available", c.getAvailableShift())
	// create shift request
	employeeAccess.POST("/shift-request", c.createShiftRequest())
	// get shift request -> get all
	employeeAccess.GET("/shift-request", c.getShiftRequest())
}

func (c *EmployeeController) getAssignedShift() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// call usecase
		c.employeeShiftUC.GetAssignedShift()

		// parse result to http response
	}
}

func (c *EmployeeController) getAvailableShift() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// call usecase
		c.employeeShiftUC.GetAvailableShift()

		// parse result to http response
	}
}

func (c *EmployeeController) createShiftRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// call usecase
		c.employeeShiftUC.CreateShiftRequest()

		// parse result to http response
	}
}

func (c *EmployeeController) getShiftRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// call usecase
		c.employeeShiftUC.GetShiftRequest()

		// parse result to http response
	}
}
