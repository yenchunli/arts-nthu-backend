package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yenchunli/arts-nthu-backend/store"
	"github.com/yenchunli/arts-nthu-backend/pkg/token"
	"github.com/yenchunli/arts-nthu-backend/util"
	"net/http"
	"time"
)

type Server struct {
	config     util.Config
	store      store.Store // Database Interface
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(config util.Config, store store.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.JWTTokenKey)
	router := NewRouter(config, store, tokenMaker)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	return &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
		router:     router,
	}, nil

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
