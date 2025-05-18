package entity

type (
	// define user data
	User struct {
		Username string
	}

	// define master data role
	Role struct {
		ID   string
		Name string
	}

	// define n to m mapping between User and Role
	UserRole struct {
		UserData User
		Role     Role
	}
)
