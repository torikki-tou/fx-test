package task

type db interface {
	Select()
	Insert()
	Update()
	Delete()
}

type logger interface {
	Log(string)
}
