package main

import (
	"net/http"
	"path/filepath"
	"two-x-dot-org/components"
	"two-x-dot-org/utils"

	"github.com/a-h/templ"
)

var log *utils.Logger

func main() {
	utils.InitLoggers()
	log = utils.MainLogger

	http.Handle("/", templ.Handler(components.GenHomePage()))
	http.Handle("/code", templ.Handler(components.GenCodePage()))
	http.Handle("/music", templ.Handler(components.GenMusicPage()))
	http.Handle("/gallery", templ.Handler(components.GenGalleryPage()))
	http.Handle("/blog", templ.Handler(components.GenBlogPage()))
	http.Handle("/about", templ.Handler(components.GenAboutPage()))

	staticFiles := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFiles))

	// allows adjacent .js, .css, ..etc files next to their components
	http.Handle("/static/components/", http.StripPrefix("/static/components/",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ext := filepath.Ext(r.URL.Path)
			if ext == ".go" || ext == ".templ" {
				http.NotFound(w, r)
				return
			}
			http.FileServer(http.Dir("./components")).ServeHTTP(w, r)
		}),
	))

	base := http.DefaultServeMux
	handler := getOnly(base)
	handler = hostAllowlist("two-x.org", "www.two-x.org")(handler)

	log.Println("Listening on 3000...")
	if err := http.ListenAndServe(":3000", handler); err != nil {
		log.Fatalln("ERROR: could not start server", err)
	}
}

func getOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow normal page loads + CORS preflight
		if r.Method == http.MethodGet || r.Method == http.MethodOptions {
			next.ServeHTTP(w, r)
			return
		}

		// Block everything else
		w.WriteHeader(http.StatusMethodNotAllowed)
	})
}

func hostAllowlist(allowed ...string) func(http.Handler) http.Handler {
	set := map[string]struct{}{}
	for _, h := range allowed {
		set[h] = struct{}{}
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if _, ok := set[r.Host]; !ok {
				http.NotFound(w, r)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
