package api

import (
	db "github.com/PyMarcus/go_sqlc/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}
