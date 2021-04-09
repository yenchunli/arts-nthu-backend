package server

import (

	"github.com/gin-gonic/gin"
	"github.com/yenchunli/arts-nthu-backend/middleware"
	store "github.com/yenchunli/arts-nthu-backend/store"
	"github.com/yenchunli/arts-nthu-backend/pkg/token"
	"github.com/yenchunli/arts-nthu-backend/util"
)

func NewRouter(config util.Config, store store.Store, tokenMaker token.Maker) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	exhibitionSvc := NewExhibitionSvc(store, tokenMaker, config)
	userSvc := NewUserSvc(store, tokenMaker, config)
	newsSvc := NewNewsSvc(store, tokenMaker, config)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/exhibitions", exhibitionSvc.List)
		apiv1.GET("/exhibitions/:id", exhibitionSvc.Get)

		apiv1.GET("/news", newsSvc.List)
		apiv1.GET("/news/:id", newsSvc.Get)

		apiv1.POST("/users", userSvc.Create)
		apiv1.POST("/users/login", userSvc.Login)
	}

	apiv1_auth := r.Group("/api/v1").Use(middleware.JWT(tokenMaker))
	{
		apiv1_auth.POST("/exhibitions", exhibitionSvc.Create)
		apiv1_auth.PUT("/exhibitions/:id", exhibitionSvc.Edit)
		apiv1_auth.DELETE("/exhibitions/:id", exhibitionSvc.Delete)

		apiv1_auth.POST("/news", newsSvc.Create)
		apiv1_auth.PUT("/news/:id", newsSvc.Edit)
		apiv1_auth.DELETE("/news/:id", newsSvc.Delete)
	}

	return r
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

