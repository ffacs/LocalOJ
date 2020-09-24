package server

import (
	"ffacs/LocalOJ/db"
	"fmt"
	"html/template"
	"net/http"
)

type userinfo struct {
	user db.User
	Logs []db.Submission
}

//Handleuserpage excutes user's page
func Handleuserpage(w http.ResponseWriter, r *http.Request) {
	user := checklogin(w, r)
	if user == nil {
		return
	}
	var data userinfo

	data.user = *user
	data.Logs = db.QuerySubmissionByUser(user)
	for i, j := 0, len(data.Logs)-1; i < j; i, j = i+1, j-1 {
		data.Logs[i], data.Logs[j] = data.Logs[j], data.Logs[i]
	}
	for i := range data.Logs {
		data.Logs[i].Lang = lang[data.Logs[i].Lang]
	}
	temp, err := template.ParseFiles("./static/status.temp")
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("502")) //Waiting for a 502 page
		return
	}
	temp.Execute(w, data)
}
