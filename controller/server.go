package controller

import (
	"be01gofire/model/mBank"
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

const (
	FirebaseProject = `test1-1c797`
	FirebaseConfig = `firebase.json`
	MysqlDsn = `be01:pass123@/be01?parseTime=True&charset=utf8`
)

type Server struct {
	Router    *gin.Engine
	App       *firebase.App
	Firestore *firestore.Client
	Db        *gorm.DB
}

func InitServer() *Server {
	rand.Seed(time.Now().UnixNano())
	// connect ke firebase
	//ctx := context.Background()
	//opt := option.WithCredentialsFile(FirebaseConfig)
	//cfg := &firebase.Config{ProjectID: FirebaseProject}
	//app, err := firebase.NewApp(ctx,cfg,opt)
	//if err != nil {
	//	log.Fatalf(`failed to connect to firebase: `+err.Error())
	//}
	//fire, err := app.Firestore(ctx)
	//if err != nil {
	//	log.Fatalf(`failed to connect to firestore: `+err.Error())
	//}
	// connect ke mysql
	db, err := gorm.Open(mysql.Open(MysqlDsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(`failed to connect to mysql: `+err.Error())
	}
	db.AutoMigrate(&mBank.Account{},&mBank.Auth{},&mBank.Transaction{})
	// set struct
	s := &Server{}
	//s.App = app
	//s.Firestore = fire
	s.Router = gin.Default()
	// add CORS
	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true
	cfg.AllowMethods = []string{"GET","POST"}
	cfg.AllowHeaders = []string{"Authorization","Origin","Accept","X-Requested-With"," Content-Type", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	s.Router.Use(cors.New(cfg))
	s.Db = db
	// https://chenyitian.gitbooks.io/gin-web-framework/content/docs/26.html
	s.Router.LoadHTMLGlob("view/*")
	return s
}

func (s *Server) Listen(port string) {
	s.Router.Run(port)	
}

func (s *Server) AssignHandler(route string, handler Handler) {
	s.Router.GET(route, func(context *gin.Context) {
		handler(&Ctx{
			Server: s,
			Context: context,
		})
	})
	s.Router.POST(route, func(context *gin.Context) {
		handler(&Ctx{
			Server: s,
			Context: context,
		})
	})
}
