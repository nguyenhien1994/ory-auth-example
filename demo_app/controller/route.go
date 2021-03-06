package controller

import (
	"simple-go-auth/controller/api"
)

func (s *Server) InitializeRoutes() {
	// Group of APIs that need to authen
	authGroup := s.Router.Group("/")
	authGroup.Use(AuthenHandler())
	{
		authGroup.PUT("/user/:userid", Authorize(), api.UpdateUser)
		authGroup.GET("/user/:userid", Authorize(), api.GetUser)
	}
}
