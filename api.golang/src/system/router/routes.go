package router

import (
    "net/http"
    "github.com/surrexi/learning-golang/api.golang/pkg/types/routes"
    HomeHandler "github.com/surrexi/learning-golang/api.golang/src/controllers/home"
    "github.com/go-xorm/xorm"
)

func Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        next.ServeHTTP(w, r)
    })
}

func GetRoutes(db *xorm.Engine) routes.Routes {

    HomeHandler.Init(db)

    return routes.Routes{
        routes.Route{Name: "Home", Method: "GET", Pattern: "/", HandlerFunc: HomeHandler.Index},
    }
}
