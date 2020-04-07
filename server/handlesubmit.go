package server

import (
	"fmt"
	"html/template"
	"net/http"
)

func HandleSubmit(writer http.ResponseWriter, request *http.Request){
	query:=request.URL.Query()
	ProID:=query.Get("ProID")
	temp,err:=template.ParseFiles("./static/submit.temp")
	if err != nil {
		fmt.Println(err)
		 writer.Write([]byte("502")) //Waiting for a 502 page
		return
	}
	temp.Execute(writer, ProID)
}