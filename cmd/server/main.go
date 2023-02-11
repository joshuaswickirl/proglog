package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joshuaswickirl/proglog/internal/server"
)

func main() {
	srv := newServerWithHandlers(":8080")
	log.Fatal(srv.ListenAndServe())
}

func newServerWithHandlers(addr string) *http.Server {
	httpsrv := server.NewHTTPServer()
	r := mux.NewRouter()
	r.HandleFunc("/", httpsrv.HandleProduce).Methods("POST")
	r.HandleFunc("/", httpsrv.HandleConsume).Methods("GET")
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
