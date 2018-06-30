package app

import (
	"log"
	"net/http"
	"github.com/go-xorm/xorm"
	"github.com/surrexi/learning-golang/api.golang/src/system/router"
)

type Server struct {
	port string
	Db 	 *xorm.Engine
}

func NewServer() Server  {
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

	http.ListenAndServe(s.port, r.Router)
}