package handlers
import (
	"fmt"
    "io/ioutil"
	"log"
	"io/fs"
	"html/template"
	"net/http"
)

func ListDir() []fs.FileInfo {
	fmt.Println("Listing")
	files, err := ioutil.ReadDir("uploads")
	if err != nil {
        log.Fatal(err)
    } 
	for _, f := range files {
		fmt.Println(f.Name())
	}

	return files
}

func TemplateList(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("static/html/files.html"))
    files := ListDir()

	t.Execute(w, files)  // merge.
}