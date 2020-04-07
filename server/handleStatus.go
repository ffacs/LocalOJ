package server

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
)

type Log struct{
	SubTime,Status,Time,Mem,Col,DirName string
}

type Status struct {
	Logs []Log
	ProID string
}

var color=map[string]string{
	"WA":"green",
	"AC":"red",
	"CE":"#6633FF",
	"TLE":"red",
	"MLE":"red",
}

func getList(ProID string) (logs []Log){
	dir:="./submitted/"+ProID+"/"
	paths,err:=ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("getList err:%v",err)
		return
	}
	for _,i:=range paths{
		name:=dir+i.Name()+"/sub.log"
		obj,_:=os.Open(name)
		line,_,_:=bufio.NewReader(obj).ReadLine()
		obj.Close()
		data:=strings.Split(string(line)," ")
		logs=append(logs,Log{data[0]+"  "+data[1],data[2],data[3],data[4],color[data[2]],i.Name()})
	}
	sort.Slice(logs, func(i, j int) bool {
		return logs[i].SubTime>logs[j].SubTime
	})
	return
}

func HandleStatus(writer http.ResponseWriter, request *http.Request){
	query:=request.URL.Query()
	ProID:=query.Get("ProID")
	temp,err:=template.ParseFiles("./static/status.temp")
	if err != nil {
		fmt.Println(err)
		writer.Write([]byte("502")) //Waiting for a 502 page
		return
	}
	var sta=Status{ProID:ProID,Logs:nil}
	sta.Logs=getList(ProID)
	temp.Execute(writer, sta)
}