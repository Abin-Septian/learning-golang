package router

import (
	"log"
	"net/http"

	"learning-golang/api.example.com/pkg/types/routes"
	AuthHandler "learning-golang/api.example.com/src/controllers/auth"
	HomeHandler "learning-golang/api.example.com/src/controllers/home"

	"github.com/go-xorm/xorm"
)

// Middleware function
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Global middleware reached!")
		next.ServeHTTP(w, r)
	})
}

// GetRoutes function
func GetRoutes(db *xorm.Engine) routes.Routes {

	HomeHandler.Init(db)
	AuthHandler.Init(db)

	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},
		routes.Route{"AuthStore", "POST", "/auth/login", AuthHandler.Login},
		routes.Route{"AuthCheck", "GET", "/auth/check", AuthHandler.Check},
	}
}
