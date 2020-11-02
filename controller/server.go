package controller

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func InitServer() *Server {
	s := &Server{}
	s.Router = gin.Default()
	// https://chenyitian.gitbooks.io/gin-web-framework/content/docs/26.html
	s.Router.LoadHTMLGlob("view/*")
	return s
}

func (s *Server) Listen(port string) {
	s.Router.Run(port)	
}

func (s *Server) AssignHandler(route string, handler Handler) {
	s.Router.Any(route, func(context *gin.Context) {
		handler(&Ctx{
			Server: s,
			Context: context,
		})
	})
}
