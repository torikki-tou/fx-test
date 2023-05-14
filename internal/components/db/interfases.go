package db

type config interface {
	GetConnectionString() string
}

type logger interface {
	Log(string)
}
