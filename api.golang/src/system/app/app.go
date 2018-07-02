package app

import (
    "log"
    "net/http"
    "github.com/go-xorm/xorm"
    "github.com/surrexi/learning-golang/api.golang/src/system/router"
    "github.com/gorilla/handlers"
    "os"
    "time"
)

type Server struct {
    port string
    Db   *xorm.Engine
}

func NewServer() Server {
    return Server{}
}

// Init all values
func (s *Server) Init(port string, db *xorm.Engine) {
    log.Println("Initializing server ...")
    s.port = ":" + port
    s.Db = db
}

// Start the server
func (s *Server) Start() {
    log.Println("Starting server on port" + s.port)

    r := router.NewRouter()

    r.Init()

    handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),
        handlers.AllowedMethods([]string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-type", "Origin", "Cache-Control", "X-App-Token"}),
        handlers.ExposedHeaders([]string{""}),
        handlers.MaxAge(1000),
        handlers.AllowCredentials(),
    )(r.Router))

    handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

    NewServer := &http.Server{
        Handler:      handler,
        Addr:         "0.0.0.0" + s.port,
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Fatal(NewServer.ListenAndServe())
}
