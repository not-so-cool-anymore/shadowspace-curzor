package handlers

import (
	"net/http"
	"fmt"
	"io/ioutil"
)
func GetImage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	imageURL := r.Form.Get("imageURL")
	

	if imageURL == ""{
		return
	}

	fmt.Println(fmt.Sprintf(
		">>> Requested image on: %s", imageURL,
	))

	response, err := http.Get(imageURL)
	if err != nil{
		fmt.Fprint(w, err)
		return
	}
	fmt.Println(fmt.Sprintf(
		">>> Status code: %d", response.StatusCode,
	))

	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	
	w.Write(data)
	return
}