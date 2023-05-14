package server

type config interface {
	GetPort() int
}

type logger interface {
	Log(string)
}
