package api

import (
	"fmt"

	"github.com/PyMarcus/go_sqlc/token"
	"github.com/PyMarcus/go_sqlc/util"

	db "github.com/PyMarcus/go_sqlc/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
	config     util.Config
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.SymmetricKey)
	router := gin.Default()

	if err != nil {
		return nil, fmt.Errorf("cannot create token: %w", err)
	}
	server := &Server{store, router, tokenMaker, config}

	server.setupRouter(router)

	return server, nil
}

func (server Server) setupRouter(router *gin.Engine) {
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.GET("/users/:user", server.getUser)
}

func (s Server) Start(addr string) error {
	return s.router.Run(addr)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
