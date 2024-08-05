package handlers

import (
	"net/http"
	"os"
)

//Serve static files and avoid listing them to the web page
func ServeStatic(w http.ResponseWriter, r *http.Request) {
    file := "." + r.URL.Path

    
    info, err := os.Stat(file)
    if err != nil {
        // file is not in the directory
        renderErrorPage(w, "File not found", http.StatusNotFound)
        return
    }

    if info.IsDir() {
        // if it is a directory
        renderErrorPage(w, "File not found", http.StatusNotFound)
        return
    }

    http.ServeFile(w, r, file)
}