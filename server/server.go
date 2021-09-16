package server

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"ig.engagements/log"
	"ig.engagements/models"
	"ig.engagements/service"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type ServerRouter struct {
	srv     *chi.Mux
	logger  log.Logger
	service *service.IgEngagementService
}

func NewServerRouter(logger log.Logger, service *service.IgEngagementService) *ServerRouter {
	router := chi.NewRouter()
	// Middlewares
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Accept"},
		AllowCredentials: true,
		Debug:            false,
		AllowedMethods:   []string{"HEAD", "OPTIONS", "GET", "POST", "PUT", "DELETE"},
	}).Handler)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	return &ServerRouter{
		srv:     router,
		logger:  logger,
		service: service,
	}
}

func (s *ServerRouter) mountEndpoints() {
	s.srv.Get("/engagements-rate/{username}", s.getEngagements)
	s.fileServer("/", "./public")
}

func (s *ServerRouter) getEngagements(writer http.ResponseWriter, request *http.Request) {
	username := chi.URLParam(request, "username")
	engagementRatings, err := s.service.GetEngagementRating(username)
	if err != nil {
		writer.WriteHeader(400)
		_ = json.NewEncoder(writer).Encode(&models.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(&models.Response{
		Success: true,
		Message: "Successfully got ratings",
		Data:    engagementRatings,
	})
}

func (s *ServerRouter) Serve(address string) error {
	s.mountEndpoints()
	return http.ListenAndServe(address, s.srv)
}

func (s *ServerRouter) fileServer(public string, static string) {

	if strings.ContainsAny(public, "{}*") {
		panic("fileServer does not permit URL parameters.")
	}

	root, _ := filepath.Abs(static)
	if _, err := os.Stat(root); os.IsNotExist(err) {
		panic("Static Documents Directory Not Found")
	}

	fs := http.StripPrefix(public, http.FileServer(http.Dir(root)))

	if public != "/" && public[len(public)-1] != '/' {
		s.srv.Get(public, http.RedirectHandler(public+"/", 301).ServeHTTP)
		public += "/"
	}

	s.srv.Get(public+"*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file := strings.Replace(r.RequestURI, public, "/", 1)
		if _, err := os.Stat(root + file); os.IsNotExist(err) {
			http.ServeFile(w, r, path.Join(root, "index.html"))
			return
		}
		fs.ServeHTTP(w, r)
	}))
}
