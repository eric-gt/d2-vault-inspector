package server

import (
	"encoding/json"
	"log"
	"net/http"

	"d2vaultinspector/cmd/web"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", templ.Handler(web.IndexPage()).ServeHTTP)

	r.Get("/health", s.healthHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/assets/*", fileServer)
	r.Get("/web", templ.Handler(web.HelloForm()).ServeHTTP)
	r.Post("/hello", web.HelloWebHandler)
	r.Post("/bungie-sign-in", s.BungieSigninHandler)

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

func (s *Server) BungieSigninHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Signed into Bungie"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
