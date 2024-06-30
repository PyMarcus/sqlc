package api

import (
	"database/sql"
	"net/http"

	db "github.com/PyMarcus/go_sqlc/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD BRL"`
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type listAccountRequest struct {
	PageId   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (s Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := db.CreateAccountParams{Owner: req.Owner, Currency: req.Currency, Balance: "0"}
	acc, err := s.store.CreateAccount(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, acc)
}

func (s Server) getAccount(ctx *gin.Context) {
	var uri getAccountRequest
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	account, err := s.store.GetAccount(ctx, uri.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNoContent, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (s Server) listAccount(ctx *gin.Context) {
	var query listAccountRequest

	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	param := db.ListAccountsParams{
		Limit:  query.PageSize,
		Offset: (query.PageId - 1) * query.PageSize,
	}

	accounts, err := s.store.ListAccounts(ctx, param)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNoContent, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}
