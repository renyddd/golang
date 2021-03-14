package main

import (
	"crypto/tls"
	"time"
)

type Server struct {
	Add string
	Port int
	Protocol string
	Timeout time.Duration
	MaxCounts int
	TLS *tls.Config
}

type Option func(server *Server)

func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func Timeout(t time.Duration) Option {
	return func(s *Server) {
		s.Timeout = t
	}
}

func MaxCounts(m int) Option {
	return func(s *Server) {
		s.MaxCounts = m
	}
}

func NewServer(addr string, port int, options ...Option) (*Server, error) {
	// 对必填参数进行赋值，可选参数赋予默认值
	srv := &Server{
		Add: addr,
		Port: port,
		Protocol: "tcp",
		Timeout: 30 * time.Second,
		MaxCounts: 1000,
		TLS: nil,
	}

	for _, optionFunc := range options {
		optionFunc(srv)
	}

	return srv, nil
}