package controller

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"log"
)

const (
	FirebaseProject = `test1-1c797`
	FirebaseConfig = `firebase.json`
)

type Server struct {
	Router    *gin.Engine
	App       *firebase.App
	Firestore *firestore.Client
}

func InitServer() *Server {
	ctx := context.Background()
	opt := option.WithCredentialsFile(FirebaseConfig)
	cfg := &firebase.Config{ProjectID: FirebaseProject}
	app, err := firebase.NewApp(ctx,cfg,opt)
	if err != nil {
		log.Fatalf(`failed to connect to firebase: `+err.Error())
	}
	fire, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf(`failed to connect to firestore: `+err.Error())
	}
	s := &Server{}
	s.App = app
	s.Firestore = fire
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
