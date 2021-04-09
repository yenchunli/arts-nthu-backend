package server

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	store "github.com/yenchunli/arts-nthu-backend/store"
	"github.com/yenchunli/arts-nthu-backend/pkg/token"
	"github.com/yenchunli/arts-nthu-backend/util"
	"net/http"
	"strconv"
)
type NewsSvc struct {
	store      store.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewNewsSvc(store store.Store, tokenMaker token.Maker, config util.Config) *NewsSvc {
	return &NewsSvc{store: store, tokenMaker: tokenMaker, config: config}
}

func (svc *NewsSvc) List(ctx *gin.Context) {
	type request struct {
		Start int  `form:"start" binding:"required,min=1`
		Size  int  `form:"size" binding:"required,min=6, max=12`
		Type  string `form:"type"`
	}
	var req request
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	switch req.Type {
	case "":
	case "exhibition":
	case "information":
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	arg := store.ListNewsParams{
		Limit:  req.Size,
		Offset: (req.Start - 1) * req.Size,
		Type:   req.Type,
	}

	news, err := svc.store.ListNews(arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, news)
	return
}

func (svc *NewsSvc) Get(ctx *gin.Context) {
	type request struct {
		ID int `uri:"id" binding:"required,min=1"`
	}
	var req request
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	news, err := svc.store.GetNews(req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, news)
	return
}

func (svc *NewsSvc) Create(ctx *gin.Context) {
	type request struct {
		Author    string `json:"author"  binding:"required"`
		Title     string `json:"title"  binding:"required"`
		TitleEn   string `json:"title_en"`
		StartDate string `json:"start_date"  binding:"required"`
		Type 	  string `json:"type"  binding:"required"`
		Draft     bool 	 `json:"draft"`
		Content   string `json:"content"  binding:"required"`
		ContentEn string `json:"content_en"`
	}
	
	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload, _ := ctx.MustGet("authorization_payload").(*token.Payload)

	arg := store.CreateNewsParams{
		Username : payload.Username,
		Author   : req.Author   ,
		Title    : req.Title    ,
		TitleEn  : req.TitleEn  ,
		StartDate: req.StartDate,
		Type 	 : req.Type 	 ,
		Draft    : req.Draft    ,
		Content  : req.Content  ,
		ContentEn: req.ContentEn,
	}

	news, err := svc.store.CreateNews(arg)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, news)
	return
}

func (svc *NewsSvc) Edit(ctx *gin.Context) {
	type request struct {
		Author    string `json:"author"  binding:"required"`
		Title     string `json:"title"  binding:"required"`
		TitleEn   string `json:"title_en"`
		StartDate string `json:"start_date"  binding:"required"`
		Type 	  string `json:"type"  binding:"required"`
		Draft     bool 	 `json:"draft"`
		Content   string `json:"content"  binding:"required"`
		ContentEn string `json:"content_en"`
	}

	var req request
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}


	payload, _ := ctx.MustGet("authorization_payload").(*token.Payload)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := store.EditNewsParams{
		ID: id,
		Username : payload.Username,
		Author   : req.Author   ,
		Title    : req.Title    ,
		TitleEn  : req.TitleEn  ,
		StartDate: req.StartDate,
		Type 	 : req.Type 	 ,
		Draft    : req.Draft    ,
		Content  : req.Content  ,
		ContentEn: req.ContentEn,
	}

	news, err := svc.store.EditNews(arg)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, news)
	return
}

func (svc *NewsSvc) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if svc.store.DeleteNews(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}