package server

import (
	"ffacs/LocalOJ/db"
	"fmt"
	"html/template"
	"net/http"
)

type status struct {
	Logs []db.Submission
}

var color = map[string]string{
	"WA":  "green",
	"AC":  "red",
	"CE":  "#6633FF",
	"TLE": "red",
	"MLE": "red",
}

var lang = map[string]string{
	"0": "c",
	"1": "c++11",
	"2": "python3",
}

//HandleStatus shows status
func HandleStatus(writer http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("./static/status.temp")
	if err != nil {
		fmt.Println(err)
		writer.Write([]byte("502")) //Waiting for a 502 page
		return
	}
	var Sta status
	Sta.Logs = db.QuerySubmission()

	for i, j := 0, len(Sta.Logs)-1; i < j; i, j = i+1, j-1 {
		Sta.Logs[i], Sta.Logs[j] = Sta.Logs[j], Sta.Logs[i]
	}
	for i := range Sta.Logs {
		Sta.Logs[i].Lang = lang[Sta.Logs[i].Lang]
	}
	temp.Execute(writer, Sta)
}
