package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yenchunli/arts-nthu-backend/internal/middleware"
	"github.com/yenchunli/arts-nthu-backend/pkg/token"
	store "github.com/yenchunli/arts-nthu-backend/store"
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
	//router := NewRouter(config, store, tokenMaker)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.router = server.NewRouter()

	return server, nil

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

func (server *Server) NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/exhibitions", server.ListExhibitions)
		apiv1.GET("/exhibitions/:id", server.GetExhibition)

		apiv1.GET("/news", server.ListNews)
		apiv1.GET("/news/:id", server.GetNews)

		apiv1.POST("/users/login", server.Login)
	}

	apiv1_auth := r.Group("/api/v1").Use(middleware.JWT(server.tokenMaker))
	{
		apiv1_auth.POST("/exhibitions", server.CreateExhibition)
		apiv1_auth.PUT("/exhibitions/:id", server.EditExhibition)
		apiv1_auth.DELETE("/exhibitions/:id", server.DeleteExhibition)

		apiv1_auth.POST("/news", server.CreateNews)
		apiv1_auth.PUT("/news/:id", server.EditNews)
		apiv1_auth.DELETE("/news/:id", server.DeleteNews)

		apiv1_auth.POST("/upload", server.UploadImage)

		apiv1_auth.POST("/users", server.CreateUser)
		apiv1_auth.POST("/users/info", server.CreateUser)
	}

	return r
}
