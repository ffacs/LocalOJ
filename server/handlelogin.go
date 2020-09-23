package server

import (
	"ffacs/LocalOJ/db"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var letter = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func getSessionID() string {
	rand.Seed(time.Now().UnixNano())
	sessionID := make([]byte, 16)
	tot := int64(len(letter))
	for i := range sessionID {
		sessionID[i] = letter[rand.Int63()%tot]
	}
	return string(sessionID)
}

// Handlelogin set cookie and change page
func Handlelogin(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	psw := r.PostFormValue("psw")
	temp, err := template.ParseFiles("./static/announcement.temp")
	if err != nil {
		fmt.Println("Login failed when parse template")
		w.Write([]byte("502 Error")) //Waiting for a 502 page
		return
	}
	var msg []string
	user, err := db.QueryUserByName(id)
	if err != nil {
		msg = append(msg, "502 Error!")
		fmt.Printf("Login failed when query uer by name : %v", err)
		temp.Execute(w, msg)
		return
	}
	if user == nil || psw != user.Psw {
		msg = append(msg, "No such user or uncorrect passwoed")
		temp.Execute(w, msg)
		return
	}
	sessionID := getSessionID()
	cookie := http.Cookie{
		Name:  "sessionID",
		Value: sessionID,
	}
	http.SetCookie(w, &cookie)
	if err = db.InsertSession(&db.Cookie{SessionID: sessionID, ID: user.Name}); err != nil {
		msg = append(msg, "502 Error!")
		fmt.Println("Login failed when insert session")
		temp.Execute(w, msg)
		return
	}
	msg = append(msg, "Login successfully,welcome "+id)
	temp.Execute(w, msg)
}
