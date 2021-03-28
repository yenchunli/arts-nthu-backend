package main

import (
	"database/sql"
	//"fmt"
	"github.com/gin-gonic/gin"
	store "github.com/yenchunli/go-nthu-artscenter-server/store"
	"net/http"
	"strconv"
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

func (svc *ExhibitionSvc) List(ctx *gin.Context) {
	type request struct {
		Start int32 `form:"start" binding:"required,min=1`
		Size  int32 `form:"size" binding:"required,min=6, max=12`
	}
	var req request
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := store.ListExhibitionsParams{
		Limit:  req.Size,
		Offset: (req.Start - 1) * req.Size,
	}

	exhibitions, err := svc.store.ListExhibitions(arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, exhibitions)
}

func (svc *ExhibitionSvc) Get(ctx *gin.Context) {
	type request struct {
		ID int32 `uri:"id" binding:"required,min=1"`
	}
	var req request
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	exhibition, err := svc.store.GetExhibition(req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, exhibition)
}

func (svc *ExhibitionSvc) Create(ctx *gin.Context) {
	type request struct {
		Title          string          `json:"title" binding:"required"`
		TitleEn        string          `json:"title_en"`
		Subtitle       string          `json:"subtitle" binding:"required"`
		SubtitleEn     string          `json:"subtitle_en"`
		StartDate      string          `json:"start_date" binding:"required"`
		EndDate        string          `json:"end_date"`
		Draft          bool            `json:"draft"`
		Host           string          `json:"host"`
		HostEn         string          `json:"host_en"`
		Performer      store.Performer `json:"performer"`
		Location       string          `json:"location"`
		LocationEn     string          `json:"location_en"`
		DailyStartTime string          `json:"daily_start_time"`
		DailyEndTime   string          `json:"daily_end_time"`
		Category       string          `json:"category" binding:"required"`
		Description    string          `json:"description" binding:"required"`
		DescriptionEn  string          `json:"description_en"`
		Content        string          `json:"content" binding:"required"`
		ContentEn      string          `json:"content_en"`
	}
	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := store.CreateExhibitionParams{
		Title:          req.Title,
		TitleEn:        req.TitleEn,
		Subtitle:       req.Subtitle,
		SubtitleEn:     req.SubtitleEn,
		StartDate:      req.StartDate,
		EndDate:        req.EndDate,
		Draft:          req.Draft,
		Host:           req.Host,
		HostEn:         req.HostEn,
		Performer:      req.Performer,
		Location:       req.Location,
		LocationEn:     req.LocationEn,
		DailyStartTime: req.DailyStartTime,
		DailyEndTime:   req.DailyEndTime,
		Category:       req.Category,
		Description:    req.Description,
		DescriptionEn:  req.DescriptionEn,
		Content:        req.Content,
		ContentEn:      req.ContentEn,
	}

	exhibition, err := svc.store.CreateExhibition(arg)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, exhibition)
}

func (svc *ExhibitionSvc) Edit(ctx *gin.Context) {
	type request struct {
		Title          string          `json:"title"`
		TitleEn        string          `json:"title_en"`
		Subtitle       string          `json:"subtitle"`
		SubtitleEn     string          `json:"subtitle_en"`
		StartDate      string          `json:"start_date"`
		EndDate        string          `json:"end_date"`
		Draft          bool            `json:"draft"`
		Host           string          `json:"host"`
		HostEn         string          `json:"host_en"`
		Performer      store.Performer `json:"performer"`
		Location       string          `json:"location"`
		LocationEn     string          `json:"location_en"`
		DailyStartTime string          `json:"daily_start_time"`
		DailyEndTime   string          `json:"daily_end_time"`
		Category       string          `json:"category"`
		Description    string          `json:"description"`
		DescriptionEn  string          `json:"description_en"`
		Content        string          `json:"content"`
		ContentEn      string          `json:"content_en"`
	}

	var req request
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := store.EditExhibitionParams{
		ID:             int32(id),
		Title:          req.Title,
		TitleEn:        req.TitleEn,
		Subtitle:       req.Subtitle,
		SubtitleEn:     req.SubtitleEn,
		StartDate:      req.StartDate,
		EndDate:        req.EndDate,
		Draft:          req.Draft,
		Host:           req.Host,
		HostEn:         req.HostEn,
		Performer:      req.Performer,
		Location:       req.Location,
		LocationEn:     req.LocationEn,
		DailyStartTime: req.DailyStartTime,
		DailyEndTime:   req.DailyEndTime,
		Category:       req.Category,
		Description:    req.Description,
		DescriptionEn:  req.DescriptionEn,
		Content:        req.Content,
		ContentEn:      req.ContentEn,
	}
	exhibition, err := svc.store.EditExhibitions(arg)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, exhibition)
}

func (svc *ExhibitionSvc) Delete(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	if svc.store.DeleteExhibition(int32(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}
