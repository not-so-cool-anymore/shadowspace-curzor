package handlers

import (
	"net/http"
)

func ServeIndexPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/html/index.html")
}

func ServeAwardsPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/html/files.html")
}

func ServeContactsPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/html/upload.html")
}

func ServeProjectsPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/html/admin.html")
}

func ServeWatcherPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/html/watcher.html")
}

func ServeImageViewerPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/html/image-viewer.html")
}
