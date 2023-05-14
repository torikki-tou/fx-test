package logger

type Logger struct {
	dir string
}

func New(config config) *Logger {
	return &Logger{
		dir: config.GetDir(),
	}
}

func (l *Logger) Log(message string) {
	println("log:", l.dir, message)
}
