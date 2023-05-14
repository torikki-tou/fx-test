package user

type repository interface {
	GetUser()
	CreateUser()
	UpdateUser()
	DeleteUser()
}

type logger interface {
	Log(string)
}
