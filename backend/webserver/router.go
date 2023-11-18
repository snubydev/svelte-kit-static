package webserver

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type ZooService interface {
	Search(text string) []string
}

var r *chi.Mux

func NewWebServer(zoo ZooService) *chi.Mux {

	r = chi.NewRouter()

	r.Get("/api/animals/search/", func(w http.ResponseWriter, r *http.Request) {
		searchString := r.URL.Query().Get("searchInput")
		if searchString == "" {
			_ = json.NewEncoder(w).Encode([]string{})
			return
		}
		result := zoo.Search(searchString)
		if len(result) == 0 {
			_ = json.NewEncoder(w).Encode([]string{})
			return
		}
		_ = json.NewEncoder(w).Encode(result)
	})

	patchRoutes(r, []string{"animals", "zoo"})

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	FileServer(r, "/", filesDir)
	return r
}

func Run(port string) {
	http.ListenAndServe(":"+port, r)
}

// PatchRoutes redirect /foo to /foo.html
func patchRoutes(r *chi.Mux, routes []string) {
	for _, route := range routes {
		path := "/" + route
		redirectURL := route + ".html"
		h := func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, redirectURL, http.StatusPermanentRedirect)
		}
		log.Printf("route: '%s'", route)
		r.Get(path, h)
	}
}
