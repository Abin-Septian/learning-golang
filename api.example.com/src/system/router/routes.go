package router

import (
	"log"
	"net/http"

	"learning-golang/api.example.com/pkg/types/routes"
	HomeHandler "learning-golang/api.example.com/src/controllers/home"
)

// Middleware function
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Global middleware reached!")
		next.ServeHTTP(w, r)
	})
}

// GetRoutes function
func GetRoutes() routes.Routes {

	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},
	}
}
