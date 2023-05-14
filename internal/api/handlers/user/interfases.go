package user

type service interface {
	GetUser()
	CreateUser()
	UpdateUser()
	DeleteUser()
}

type logger interface {
	Log(string)
}
