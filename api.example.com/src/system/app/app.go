package app

import (
	"log"
	"net/http"

	"learning-golang/api.example.com/src/system/router"

	"github.com/go-xorm/xorm"
)

// Server Struct
type Server struct {
	port string
	DB   *xorm.Engine
}

// NewServer init
func NewServer() Server {
	return Server{
		port: "",
	}
}

// Init all vals
func (s *Server) Init(port string, db *xorm.Engine) {
	log.Println("initializing server ...")
	s.port = ":" + port
	s.DB = db
}

// Start the server
func (s *Server) Start() {
	log.Println("starting server on port " + s.port)

	r := router.NewRouter()

	r.Init()

	http.ListenAndServe(s.port, r.Router)
}
