package handlers

import (
	"net/http"
	"text/template"
	"os/exec"
	"fmt"
	"strings"
)
func Watch(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("static/html/watcher.html"))
    comments := GetComments()

	t.Execute(w, comments)  // merge.
}

func Ping(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	if strings.Contains(r.Form.Get("ip"), ";"){
		fmt.Fprint(w, "Forbidden character detected.")
		return
	}
	command := fmt.Sprintf("ping -c 1 \"%s\"", r.Form.Get("ip"))
	fmt.Println(command)
	
	cmd := exec.Command("/bin/sh", "-c", command)
	std_out, err := cmd.Output()
	
	if err != nil {
		fmt.Fprint(w, "Failed to execute command!")
	}

	stdOutStr := string(std_out)

	fmt.Fprint(w, stdOutStr)
}