package server

import (
	"ffacs/LocalOJ/db"
	"fmt"
	"html/template"
	"net/http"
)

func checklogin(w http.ResponseWriter, r *http.Request) *db.User {
	announcement, err := template.ParseFiles("./static/announcement.temp")
	var msg []string
	if err != nil {
		fmt.Println("Test failed when parse template")
		w.Write([]byte("502 Error")) //Waiting for a 502 page
		return nil
	}
	sessionID, err := r.Cookie("sessionID")
	if err != nil {
		msg = append(msg, "Please login in first")
		announcement.Execute(w, msg)
		return nil
	}
	cookie, err := db.QueryCookieBySessionID(sessionID.Value)
	if err != nil {
		msg = append(msg, "502 Error!")
		fmt.Println("HanleRoot failed")
		announcement.Execute(w, msg)
		return nil
	}

	if cookie == nil {
		msg = append(msg, "Please login in first")
		announcement.Execute(w, msg)
		return nil
	}

	user, err := db.QueryUserBySession(cookie)
	if err != nil {
		msg = append(msg, "502 Error!")
		fmt.Println("HandleRoot failed")
		announcement.Execute(w, msg)
		return nil
	}
	return user
}
