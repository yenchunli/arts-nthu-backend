package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ExhibitionSvc struct {
	store Store
}

func NewRouter(store Store) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	exhibitionSvc := NewExhibitionSvc(store)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/exhibitions", exhibitionSvc.List)
		apiv1.GET("/exhibitions/:id", exhibitionSvc.Get)
		apiv1.POST("/exhibitions", exhibitionSvc.Create)
		apiv1.PUT("/exhibitions/:id", exhibitionSvc.Edit)
		apiv1.DELETE("/exhibitions/:id", exhibitionSvc.Delete)
	}

	return r
}

func NewExhibitionSvc(store Store) *ExhibitionSvc {
	return &ExhibitionSvc{store: store}
}

func (svc *ExhibitionSvc) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List Exhibitions",
	})
}

func (svc *ExhibitionSvc) Get(c *gin.Context) {
	exhibition, _ := svc.store.GetExhibition(1)
	fmt.Println(exhibition)
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Exhibition",
	})
}

func (svc *ExhibitionSvc) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create Exhibition",
	})
}

func (svc *ExhibitionSvc) Edit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Edit Exhibition",
	})
}

func (svc *ExhibitionSvc) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Exhibition",
	})
}
