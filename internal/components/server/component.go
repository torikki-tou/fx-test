package server

import (
	"strconv"
	"time"
)

type Server struct {
	port   int
	logger logger
}

func New(config config, logger logger) *Server {
	return &Server{
		port:   config.GetPort(),
		logger: logger,
	}
}

func (s *Server) Serve() {
	for {
		s.logger.Log("serving " + strconv.Itoa(s.port))
		time.Sleep(time.Second)
	}
}
