package server

import (
	"fmt"
	"html/template"
	"net/http"
)

//HandleLogout sets sessionID empty
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	announcement, err := template.ParseFiles("./static/announcement.temp")
	var msg []string
	if err != nil {
		fmt.Println("Test failed when parse template")
		w.Write([]byte("502 Error")) //Waiting for a 502 page
		return
	}
	http.SetCookie(w, &http.Cookie{Name: "sessionID", Value: ""})
	msg = append(msg, "See you~")
	announcement.Execute(w, msg)
}
