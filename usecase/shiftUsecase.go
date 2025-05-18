package usecase

type (
	AdminShiftUsecase interface {
		CreateShift()
		EditShift()
		GetAllShiftRequest()
		ApproveShiftRequest()
		RejectShiftRequest()
		GetAllFinalizedShift()
	}

	EmployeeShiftUsecase interface {
		GetAssignedShift()
		GetAvailableShift() // unassigned shift
		CreateShiftRequest()
		GetShiftRequest()
	}
)
