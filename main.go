package main

import (
	"net/http"
	// "path/filepath"
	// "two-x-dot-org/components"
	"two-x-dot-org/utils"
	// "github.com/a-h/templ"
)

var log *utils.Logger

func main() {
	utils.InitLoggers()
	log = utils.MainLogger

	log.Println("Server is under construction...")

	// http.Handle("/", templ.Handler(components.GenHomePage()))
	// http.Handle("/code", templ.Handler(components.GenCodePage()))
	// http.Handle("/music", templ.Handler(components.GenMusicPage()))
	// http.Handle("/gallery", templ.Handler(components.GenGalleryPage()))
	// http.Handle("/blog", templ.Handler(components.GenBlogPage()))
	// http.Handle("/about", templ.Handler(components.GenAboutPage()))
	//
	// staticFiles := http.FileServer(http.Dir("./static"))
	// http.Handle("/static/", http.StripPrefix("/static/", staticFiles))
	//
	// // allows adjacent .js, .css, ..etc files next to their components
	// http.Handle("/static/components/", http.StripPrefix("/static/components/",
	// 	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		ext := filepath.Ext(r.URL.Path)
	// 		if ext == ".go" || ext == ".templ" {
	// 			http.NotFound(w, r)
	// 			return
	// 		}
	// 		http.FileServer(http.Dir("./components")).ServeHTTP(w, r)
	// 	}),
	// ))

	log.Println("Listening on 3000...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalln("ERROR: could not start server", err)
	}
}
