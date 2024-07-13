package api

import (
	"fmt"
	"log"

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
	log.Println(config.SymmetricKey)

	tokenMaker, err := token.NewPasetoMaker(config.SymmetricKey)
	router := gin.Default()

	if err != nil {
		return nil, fmt.Errorf("cannot create token: %w", err)
	}
	server := &Server{store, router, tokenMaker, config}

	server.setupRouter(router, tokenMaker)

	return server, nil
}

func (server Server) setupRouter(router *gin.Engine, tokenMaker token.Maker) {
	// auth
	authRoutes := router.Group("/").Use(authMiddleware(tokenMaker))
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)
	// unprotected
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
