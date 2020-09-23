package server

import (
	"ffacs/LocalOJ/db"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

//HandleRoot handles / dir
func HandleRoot(w http.ResponseWriter, r *http.Request) {

	sessionID, err := r.Cookie("sessionID")
	if err != nil {
		fmt.Printf("Get cookie failed :%v\n", err)
		fileObj, err := os.Open("./index.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer fileObj.Close()
		data, err := ioutil.ReadAll(fileObj)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Write(data)
		return
	}

	cookie, err := db.QueryCookieBySessionID(sessionID.Value)
	if err != nil {
		fmt.Println("HanleRoot failed")
		return
	}
	if cookie == nil {
		fileObj, err := os.Open("./index.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer fileObj.Close()
		data, err := ioutil.ReadAll(fileObj)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Write(data)
		return
	}
	// handle logined page
	temp, err := template.ParseFiles("./static/index.temp")
	if err != nil {
		fmt.Printf("root failed when parse template : %v\n", err)
		w.Write([]byte("502 Error!")) //Waiting for a 502 page
		return
	}
	temp.Execute(w, cookie.ID)
}
