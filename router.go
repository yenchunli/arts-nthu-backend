package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	store "github.com/yenchunli/go-nthu-artscenter-server/store"
	"net/http"
)

type ExhibitionSvc struct {
	store store.Store
}

func NewRouter(store store.Store) *gin.Engine {
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

func NewExhibitionSvc(store store.Store) *ExhibitionSvc {
	return &ExhibitionSvc{store: store}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (svc *ExhibitionSvc) List(c *gin.Context) {
	type request struct {
		Start int32 `form:"start" binding:"required,min=1`
		Size  int32 `form:"size" binding:"required,min=6, max=12`
	}
	var req request
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := store.ListExhibitionsParams{
		Limit:  req.Size,
		Offset: (req.Start - 1) * req.Size,
	}

	exhibitions, err := svc.store.ListExhibitions(arg)

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, exhibitions)
}

func (svc *ExhibitionSvc) Get(c *gin.Context) {
	type request struct {
		ID int8 `uri:"id" binding:"required,min=1"`
	}
	var req request
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	exhibition, err := svc.store.GetExhibition(req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, exhibition)
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
