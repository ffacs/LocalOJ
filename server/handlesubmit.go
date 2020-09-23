package server

import (
	"fmt"
	"html/template"
	"net/http"
)

//HandleSubmit handles submission
func HandleSubmit(w http.ResponseWriter, r *http.Request) {
	user := checklogin(w, r)
	if user == nil {
		return
	}
	query := r.URL.Query()
	ProID := query.Get("ProID")
	temp, err := template.ParseFiles("./static/submit.temp")
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("502")) //Waiting for a 502 page
		return
	}
	temp.Execute(w, ProID)
}
