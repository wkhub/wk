package user

var currentUser *User

// WkHome represent the WK_HOME directory
type User struct {
	Config *UserConfig
	Home   Home
}

func Current() *User {
	if currentUser == nil {
		currentUser = &User{
			Config: getUserConfig(),
			Home:   WkHome(),
		}
		// userConfig.Set
	}
	return currentUser
}
