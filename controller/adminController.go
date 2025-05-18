package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdifr17/ScheduleManagement/repository"
	"github.com/mahdifr17/ScheduleManagement/usecase"
)

type AdminController struct {
	router       *gin.Engine
	adminShiftUC usecase.AdminShiftUsecase
}

func SetupAdminController(
	router *gin.Engine,
	shiftRP repository.ShiftRepository,
	shiftRequestRP repository.ShiftRequestRepository,
) {
	c := new(AdminController)
	c.router = router
	c.adminShiftUC = usecase.NewAdminShiftUsecaseV1(shiftRP, shiftRequestRP)
	c.setupController()
}

func (c *AdminController) setupController() {
	adminAccess := c.router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"aaa": "aaa", // user:foo password:bar
	}))

	// create shift
	adminAccess.POST("/shift", c.createShift())
	// edit shift
	adminAccess.PUT("/shift", c.editShift())
	// get all shift request
	adminAccess.GET("/shift", c.getAllShiftRequest())
	// get all finalized shift
	adminAccess.GET("/shift/final", c.getAllFinalizedShift())
	// approve shift request
	adminAccess.POST("/shift-request/approve", c.approveShiftRequest())
	// reject shift request
	adminAccess.POST("/shift-request/reject", c.rejectShiftRequest())
}

func (c *AdminController) createShift() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// call usecase
		c.adminShiftUC.CreateShift()

		// parse result to http response
	}
}

func (c *AdminController) editShift() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// call usecase
		c.adminShiftUC.EditShift()

		// parse result to http response
	}
}

func (c *AdminController) getAllShiftRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// call usecase
		c.adminShiftUC.GetAllShiftRequest()

		// parse result to http response
	}
}

func (c *AdminController) getAllFinalizedShift() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// call usecase
		c.adminShiftUC.GetAllFinalizedShift()

		// parse result to http response
	}
}

func (c *AdminController) approveShiftRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// call usecase
		c.adminShiftUC.ApproveShiftRequest()

		// parse result to http response
	}
}

func (c *AdminController) rejectShiftRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// call usecase
		c.adminShiftUC.RejectShiftRequest()

		// parse result to http response
	}
}
