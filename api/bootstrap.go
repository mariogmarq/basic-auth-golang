package api

import (
	"GoLandPruebas/api/middleware"
	"GoLandPruebas/api/routes/login"
	"GoLandPruebas/api/routes/usuarios"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func (s *Server) setupRoutes() {
	s.engine.GET("/users", middleware.LoginMiddleware, usuarios.GetAllUsuarios)
	s.engine.GET("/users/:id", middleware.LoginMiddleware, usuarios.GetUsuarios)
	s.engine.POST("/users", login.ValidateLoginPost, usuarios.CreateUser)
	s.engine.POST("/login", login.ValidateLoginPost ,login.Login)
}

func NewServer() Server {
	eng := gin.New()
	eng.Use(gin.Logger(), gin.Recovery())
	srv := Server{eng}
	srv.setupRoutes()
	return srv
}

func (s Server) Run() error {
	return s.engine.Run()
}