package server

import (
	"database/sql"
	//"fmt"
	"github.com/gin-gonic/gin"
	store "github.com/yenchunli/arts-nthu-backend/store"
	"github.com/yenchunli/arts-nthu-backend/pkg/token"
	"github.com/yenchunli/arts-nthu-backend/util"
	"net/http"
	"strconv"
)
type ExhibitionSvc struct {
	store      store.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewExhibitionSvc(store store.Store, tokenMaker token.Maker, config util.Config) *ExhibitionSvc {
	return &ExhibitionSvc{store: store, tokenMaker: tokenMaker, config: config}
}

func (svc *ExhibitionSvc) List(ctx *gin.Context) {
	type request struct {
		Start int  `form:"start" binding:"required,min=1`
		Size  int  `form:"size" binding:"required,min=6, max=12`
		Type  string `form:"type"`
	}
	type response struct {
		Data 	[]store.Exhibition `json:"data"`
		MaxSize int    			   `json:"max_size"`
	}

	var req request
	var res response

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	switch req.Type {
	case "":
	case "visual_art":
	case "public_art":
	case "film_art":
	case "performing_art":
	case "ai_music":
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	arg := store.ListExhibitionsParams{
		Limit:  req.Size,
		Offset: (req.Start - 1) * req.Size,
		Type:   req.Type,
	}

	exhibitions, err := svc.store.ListExhibitions(arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	maxSize, err := svc.store.GetExhibitionsMaxSize()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	res.Data = exhibitions
	res.MaxSize = maxSize
	ctx.JSON(http.StatusOK, res)
	return
}

func (svc *ExhibitionSvc) Get(ctx *gin.Context) {
	type request struct {
		ID int `uri:"id" binding:"required,min=1"`
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
	return
}

func (svc *ExhibitionSvc) Create(ctx *gin.Context) {
	type request struct {
		Title          string          `json:"title" binding:"required"`
		TitleEn        string          `json:"title_en"`
		Subtitle       string          `json:"subtitle"`
		SubtitleEn     string          `json:"subtitle_en"`
		Type           string          `json:"type" binding:"required"`
		Cover          string          `json:"cover" binding:"required"`
		StartDate      string          `json:"start_date" binding:"required"`
		EndDate        string          `json:"end_date"`
		Draft          bool            `json:"draft"`
		Host           string          `json:"host"`
		HostEn         string          `json:"host_en"`
		Performer      string		   `json:"performer"`
		Location       string          `json:"location"`
		LocationEn     string          `json:"location_en"`
		DailyStartTime string          `json:"daily_start_time"`
		DailyEndTime   string          `json:"daily_end_time"`
		Category       string          `json:"category"`
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
		Type:           req.Type,
		Cover:          req.Cover,
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
	return
}

func (svc *ExhibitionSvc) Edit(ctx *gin.Context) {
	type request struct {
		Title          string          `json:"title"`
		TitleEn        string          `json:"title_en"`
		Subtitle       string          `json:"subtitle"`
		SubtitleEn     string          `json:"subtitle_en"`
		Type           string          `json:"type"`
		Cover          string          `json:"cover"`
		StartDate      string          `json:"start_date"`
		EndDate        string          `json:"end_date"`
		Draft          bool            `json:"draft"`
		Host           string          `json:"host"`
		HostEn         string          `json:"host_en"`
		Performer      string		   `json:"performer"`
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
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := store.EditExhibitionParams{
		ID:             int(id),
		Title:          req.Title,
		TitleEn:        req.TitleEn,
		Subtitle:       req.Subtitle,
		SubtitleEn:     req.SubtitleEn,
		Type:           req.Type,
		Cover:          req.Cover,
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
	return
}

func (svc *ExhibitionSvc) Delete(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	if svc.store.DeleteExhibition(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}