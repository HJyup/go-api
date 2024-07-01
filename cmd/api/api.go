package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"go-api/service/user"
	"log"
	"net/http"
)

type Server struct {
	address string
	db      *sql.DB
}

func NewServer(address string, db *sql.DB) *Server {
	return &Server{
		address: address,
		db:      db,
	}
}

func (server *Server) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(server.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRouters(subrouter)

	log.Println("Server is running on", server.address)

	return http.ListenAndServe(server.address, router)
}
