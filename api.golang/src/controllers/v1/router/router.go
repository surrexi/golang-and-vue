package router

import (
    "net/http"
    StatusHandler "github.com/surrexi/learning-golang/api.golang/src/controllers/v1/status"
    "github.com/surrexi/learning-golang/api.golang/pkg/types/routes"
    "log"
)

func Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("X-App-Token")
        if len(token) < 1 {
            http.Error(w, "Not authorized", http.StatusUnauthorized)
            return
        }

        log.Println("Inside V1 Middleware")

        next.ServeHTTP(w, r)
    })
}

func GetRoutes() (SubRoute map[string]routes.SubRoutePackage) {
    SubRoute = map[string]routes.SubRoutePackage{
        "/v1": {
            Routes: routes.Routes{
                routes.Route{Name: "Status", Method: "GET", Pattern: "/status", HandlerFunc: StatusHandler.Index},
            },
            Middleware: Middleware,
        },
    }

    return
}
