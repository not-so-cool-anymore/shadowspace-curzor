package handlers
import (
	"net/http"
	"fmt"
)
func Authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	
	if r.Form.Get("username") == "admin" && r.Form.Get("password") == "12345"{
			fmt.Fprintf(w, "Successfully logged. Use this info for the report.")
	}

	fmt.Println(r.Form.Get("username"))
	fmt.Println(r.Form.Get("password"))

	fmt.Fprintf(w,"Fail! Try again!")
}