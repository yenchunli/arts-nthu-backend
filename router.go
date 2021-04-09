package main

import (
	"database/sql"
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yenchunli/go-nthu-artscenter-server/middleware"
	store "github.com/yenchunli/go-nthu-artscenter-server/store"
	"github.com/yenchunli/go-nthu-artscenter-server/token"
	"github.com/yenchunli/go-nthu-artscenter-server/util"
	"net/http"
	"strconv"
)

type ExhibitionSvc struct {
	store      store.Store
	tokenMaker token.Maker
	config     Config
}

type UserSvc struct {
	store      store.Store
	tokenMaker token.Maker
	config     Config
}

type NewsSvc struct {
	store      store.Store
	tokenMaker token.Maker
	config     Config
}

func NewRouter(config Config, store store.Store, tokenMaker token.Maker) *gin.Engine {
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

func NewUserSvc(store store.Store, tokenMaker token.Maker, config Config) *UserSvc {
	return &UserSvc{store: store, tokenMaker: tokenMaker, config: config}
}

func NewExhibitionSvc(store store.Store, tokenMaker token.Maker, config Config) *ExhibitionSvc {
	return &ExhibitionSvc{store: store, tokenMaker: tokenMaker, config: config}
}

func NewNewsSvc(store store.Store, tokenMaker token.Maker, config Config) *NewsSvc {
	return &NewsSvc{store: store, tokenMaker: tokenMaker, config: config}
}

func (svc *UserSvc) Create(ctx *gin.Context) {
	type request struct {
		Username string `json:"username" binding:"required,alphanum"`
		Password string `json:"password" binding:"required,min=6"`
		FullName string `json:"full_name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
	}
	type response struct {
		Username         string `json:"username"`
		FullName         string `json:"full_name"`
		Email            string `json:"email"`
		PasswordChangeAt int64  `json:"password_change_at"`
		CreatedAt        int64  `json:"created_at"`
	}
	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := store.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}
	user, err := svc.store.CreateUser(arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
	return
}

func (svc *UserSvc) Login(ctx *gin.Context) {
	type request struct {
		Username string `json:"username" binding:"required,alphanum"`
		Password string `json:"password" binding:"required,min=6"`
	}
	type response struct {
		AccessToken string `json:"access_token"`
		Username    string `json:"username"`
	}

	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := svc.store.GetUser(req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := svc.tokenMaker.CreateToken(
		user.Username,
		svc.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := response{
		AccessToken: accessToken,
		Username:    user.Username,
	}
	ctx.JSON(http.StatusOK, res)
	return
}



func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
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
