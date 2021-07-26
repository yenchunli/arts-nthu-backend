package server

import (
	"database/sql"
	//"fmt"
	"github.com/gin-gonic/gin"
	store "github.com/yenchunli/arts-nthu-backend/store"
	"github.com/yenchunli/arts-nthu-backend/util"
	"net/http"
)

func (server *Server) CreateUser(ctx *gin.Context) {
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

	user, err := server.store.CreateUser(arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
	return
}

func (server *Server) Login(ctx *gin.Context) {
	type request struct {
		Email    string `json:"email" binding:"required"`
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

	user, err := server.store.GetUserByEmail(req.Email)
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

	accessToken, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
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

func (server *Server) Info(ctx *gin.Context) {

	type response struct {
		Role         string `json:"role"`
		Introduction string `json:"introduction"`
		Avatar       string `json:"avatar"`
		Name         string `json:"name"`
	}

	res := response{
		Role:         "admin",
		Introduction: "",
		Avatar:       "",
		Name:         "",
	}

	ctx.JSON(http.StatusOK, res)
	return

}
