package handlers

import (
	"net/http"
	"log"
	"os"
	"bufio"
	"fmt"
)

func Comment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	
	comment := r.Form.Get("comment")
	fmt.Println(comment)
	
	if comment == ""{
		http.Redirect(w,r, "/watcher", http.StatusFound)
		return
	}

	commentsFile, err := os.OpenFile("./comments.list", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil{
		fmt.Println(err)
	}
	defer commentsFile.Close()
	
	_, err = commentsFile.WriteString(fmt.Sprintf("%s\n",comment))
	if err != nil{
		fmt.Println(err)
	}
	http.Redirect(w,r, "/watcher", http.StatusFound)
}

func GetComments() []string {
	var comments []string 

	commentsFile, err := os.Open("./comments.list")
	if err != nil {
		log.Fatal(err)
	}
	defer commentsFile.Close()

	scanner := bufio.NewScanner(commentsFile)
	for scanner.Scan(){
		comments = append(comments, scanner.Text())
	}

	return comments
}