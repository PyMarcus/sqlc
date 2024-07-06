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
