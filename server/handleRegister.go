package server

import (
	"ffacs/LocalOJ/db"
	"fmt"
	"html/template"
	"net/http"
)

//HandleRegister provide register page
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	psw := r.PostFormValue("psw")
	email := r.PostFormValue("email")

	//User registion information
	userRG := db.User{
		Psw:   psw,
		Name:  name,
		Email: email,
	}

	temp, err := template.ParseFiles("./static/announcement.temp")
	if err != nil {
		fmt.Printf("Register failed when parse template : %v\n", err)
		w.Write([]byte("502 Error!")) //Waiting for a 502 page
		return
	}
	var msg []string
	user, err := db.QueryUserByName(name)
	if err != nil {
		msg = append(msg, "502 Error!")
		fmt.Println("Register failed when query user by name")
		temp.Execute(w, msg)
		return
	}
	if user != nil {
		msg = append(msg, "This name has been used!")
		temp.Execute(w, msg)
		return
	}
	newID, err := db.InsertUser(&userRG)
	if err != nil {
		fmt.Println("Register failed when insert user")
		msg = append(msg, "502 Error!")
		temp.Execute(w, msg)
		return
	}
	msg = append(msg, "Register successfully,your name is "+name)
	msg = append(msg, "You can use name and password to login now")
	temp.Execute(w, msg)
	fmt.Printf("Register Done,new ID is %v,name is %v\n", newID, userRG.Name)
}
