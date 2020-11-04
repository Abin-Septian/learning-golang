package router

import (
	"learning-golang/api.example.com/pkg/types/routes"
	StatusHandler "learning-golang/api.example.com/src/controllers/v1/status"
	"log"
	"net/http"

	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

// Middleware function
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-App-Token")
		if len(token) < 1 {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		log.Println("V1 Middleware reached!")

		next.ServeHTTP(w, r)
	})
}

// GetRoutes function
func GetRoutes(DB *xorm.Engine) (SubRoute map[string]routes.SubRoutePackage) {

	db = DB

	StatusHandler.Init(db)

	/* ROUTES */
	SubRoute = map[string]routes.SubRoutePackage{
		"/v1": {
			Routes: routes.Routes{
				routes.Route{"Home", "GET", "/status", StatusHandler.Index},
			},
			Middleware: Middleware,
		},
	}

	return
}
