package main

import (
	"ffacs/LocalOJ/server"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/submit",server.HandleSubmit)
	http.HandleFunc("/",server.HandleRoot)
	http.HandleFunc("/test",server.HandleTest)
	http.HandleFunc("/status",server.HandleStatus)
	http.HandleFunc("/details",server.HandleDetails)
	http.Handle("/JudgeOnline/", http.StripPrefix("/JudgeOnline/", http.FileServer(http.Dir("JudgeOnline"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/Semantic/", http.StripPrefix("/Semantic/", http.FileServer(http.Dir("Semantic"))))
	http.Handle("/submitted/", http.StripPrefix("/submitted/", http.FileServer(http.Dir("submitted"))))
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("image"))))
	err:=http.ListenAndServe(":80",nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}