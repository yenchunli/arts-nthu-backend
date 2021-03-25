package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	config Config
	store  Store	// Database Interface
	router *gin.Engine
	//email EmailSender
}

func NewServer(config Config, store Store, router *gin.Engine) *Server {
	return &Server{
		config: config,
		store:  store,
		router: router,
	}
}

func (server *Server) Run() {
	s := &http.Server{
		Addr:           server.config.ServerAddress,
		Handler:        server.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
