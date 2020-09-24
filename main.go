package main

import (
	"ffacs/LocalOJ/db"
	"ffacs/LocalOJ/server"
	"fmt"
	"net/http"
)

func main() {

	go db.DeleteOuttimeSession() //Start GC

	http.HandleFunc("/submit", server.HandleSubmit)
	http.HandleFunc("/", server.HandleRoot)
	http.HandleFunc("/test", server.HandleTest)
	http.HandleFunc("/status", server.HandleStatus)
	http.HandleFunc("/details", server.HandleDetails)

	http.HandleFunc("/register", server.HandleRegister)
	http.HandleFunc("/login", server.Handlelogin)
	http.HandleFunc("/logout", server.HandleLogout)
	http.HandleFunc("/userpage", server.Handleuserpage)

	http.Handle("/JudgeOnline/", http.StripPrefix("/JudgeOnline/", http.FileServer(http.Dir("JudgeOnline"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/Semantic/", http.StripPrefix("/Semantic/", http.FileServer(http.Dir("Semantic"))))
	http.Handle("/submitted/", http.StripPrefix("/submitted/", http.FileServer(http.Dir("submitted"))))
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("image"))))

	if err := http.ListenAndServe(":80", nil); err != nil { //Start Listen and serve
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
