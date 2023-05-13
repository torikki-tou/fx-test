package task

type repository interface {
	GetTask()
	CreateTask()
	UpdateTask()
	DeleteTask()
}

type logger interface {
	Log(string)
}
