package task

type service interface {
	GetTask()
	CreateTask()
	UpdateTask()
	DeleteTask()
}

type logger interface {
	Log(string)
}
