package router

import (
    "github.com/gorilla/mux"
    "github.com/surrexi/learning-golang/api.golang/pkg/types/routes"
    V1SubRoutes "github.com/surrexi/learning-golang/api.golang/src/controllers/v1/router"
    "github.com/go-xorm/xorm"
)

type Router struct {
    Router *mux.Router
}

func (r *Router) Init(db *xorm.Engine) {
    r.Router.Use(Middleware)

    baseRoutes := GetRoutes(db)
    for _, route := range baseRoutes {
        r.Router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    v1SubRoutes := V1SubRoutes.GetRoutes(db)
    for name, pack := range v1SubRoutes {
        r.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
    }
}

func (r *Router) AttachSubRouterWithMiddleware(path string, subRoutes routes.Routes, middleware mux.MiddlewareFunc) (SubRouter *mux.Router) {
    SubRouter = r.Router.PathPrefix(path).Subrouter()
    SubRouter.Use(middleware)

    for _, route := range subRoutes {
        SubRouter.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return
}

func NewRouter() (r Router) {
    r.Router = mux.NewRouter().StrictSlash(true)

    return
}
