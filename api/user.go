package api

import (
	"database/sql"
	"net/http"

	db "github.com/PyMarcus/go_sqlc/db/sqlc"
	"github.com/PyMarcus/go_sqlc/util"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Fullname string `json:"full_name" binding:"required"`
}

type createUserResponse struct {
	Username  string `json:"username" binding:"required,alphanum"`
	Email     string `json:"email" binding:"required,email"`
	Fullname  string `json:"full_name" binding:"required"`
	CreatedAt string `json:"created_at"`
}

type getUserRequest struct {
	Username string `uri:"user" binding:"required"`
}

func (s Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashed, err := util.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	args := db.CreateUserParams{Username: req.Username, Email: req.Email, HashedPassword: hashed, FullName: req.Fullname}
	user, err := s.store.CreateUser(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := createUserResponse{
		Username:  user.Username,
		Fullname:  user.FullName,
		CreatedAt: user.CreatedAt.String(),
		Email:     user.Email,
	}

	ctx.JSON(http.StatusCreated, response)
}

func (s Server) getUser(ctx *gin.Context) {
	var uri getUserRequest
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := s.store.GetUser(ctx, uri.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNoContent, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := createUserResponse{
		Username:  user.Username,
		Fullname:  user.FullName,
		CreatedAt: user.CreatedAt.String(),
		Email:     user.Email,
	}

	ctx.JSON(http.StatusOK, response)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string             `json:"access_token"`
	User        createUserResponse `json:"user"`
}

func (s Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := s.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil{
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	token, err := s.tokenMaker.CreateToken(req.Username, s.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	userResponse := createUserResponse{Username: user.Username, Email: user.Email, Fullname: user.FullName, CreatedAt: user.CreatedAt.String()}
	response := loginUserResponse{token, userResponse}
	ctx.JSON(http.StatusOK, response)
}
