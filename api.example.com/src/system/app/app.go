package app

import (
	"log"
	"net/http"
	"os"
	"time"

	"learning-golang/api.example.com/src/system/router"

	"github.com/go-xorm/xorm"
	"github.com/gorilla/handlers"
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

	r.Init(s.DB)

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "PATCH", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Cache-Control", "X-App-Token"}),
		handlers.ExposedHeaders([]string{""}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(r.Router))
	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	newServer := &http.Server{
		Handler:      handler,
		Addr:         "127.0.0.1" + s.port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(newServer.ListenAndServe())
}
