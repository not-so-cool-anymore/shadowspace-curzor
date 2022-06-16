package main

import (
	"flag"
	"net/http"
	"time"

	"curzor/handlers"
	"curzor/middlewares"
	"github.com/gorilla/mux"
)

func readConfigFile() {

}

func configure() *http.Server {
	var staticFilesDir string
	// var uploadedFilesDir string

	flag.StringVar(&staticFilesDir, "staticFilesDir", "static/", "Directory to serve static files from. Defaults to the current dir.")
	flag.Parse()

	// flag.StringVar(&uploadedFilesDir, "uploadedFilesDir", "uploaded_files/", "Directory to serve the uploaded files to the service.")
	// flag.Parse()

	router := mux.NewRouter()
	router.Use(middlewares.HeadersMiddleware)

	// serve static files on :host:/static/<file>
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticFilesDir)))).Methods("GET", "HEAD")

	// serve static files on :host:/uploaded_files/<file>
	// router.PathPrefix("/uploaded_files/").Handler(http.StripPrefix("/uploaded_files/", http.FileServer(http.Dir(uploadedFilesDir)))).Methods("GET", "HEAD")

	// serve static content
	router.HandleFunc("/", handlers.ServeIndexPage).Methods("GET")
	router.HandleFunc("/files", handlers.TemplateList).Methods("GET")
	router.HandleFunc("/upload", handlers.ServeContactsPage).Methods("GET")
	router.HandleFunc("/upload", handlers.Upload).Methods("POST")
	router.HandleFunc("/admin", handlers.ServeProjectsPage).Methods("GET")
	router.HandleFunc("/authenticate", handlers.Authenticate).Methods("POST")
	router.HandleFunc("/image-viewer", handlers.ServeImageViewerPage).Methods("GET")
	router.HandleFunc("/users", handlers.ListUsers).Methods("GET")


	router.HandleFunc("/watcher", handlers.Watch).Methods("GET")
	router.HandleFunc("/ping", handlers.Ping).Methods("POST")
	router.HandleFunc("/comment", handlers.Comment).Methods("POST")
	router.HandleFunc("/get-image", handlers.GetImage).Methods("POST")

	server := &http.Server{
		Handler:      router,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return server
}
