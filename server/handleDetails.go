package server

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func HandleDetails(writer http.ResponseWriter, request *http.Request){
	query:=request.URL.Query()
	ProID:=query.Get("ProID")
	DirName:=query.Get("Dir")


	src:=make([]string,0)
	det:=make([]string,0)

	paths,_ :=ioutil.ReadDir("./submitted/"+ProID+"/"+DirName)
	for _,path:=range paths{
		file,_:=os.Open("./submitted/"+ProID+"/"+DirName+"/"+path.Name())
		buffer:=bufio.NewReader(file)
		if path.Name()=="sub.log" {
			for  {
				line,_,err:=buffer.ReadLine()
				if err ==io.EOF {
					break
				}
				det=append(det,string(line))
			}
		}else{
			for  {
				line,_,err:=buffer.ReadLine()
				if err ==io.EOF {
					break
				}
				src=append(src,string(line))
			}
		}
		_ = file.Close()
	}

	temp,err:=template.ParseFiles("./static/details.temp")
	if err != nil {
		fmt.Println(err)
		_, _ = writer.Write([]byte("502")) //Waiting for a 502 page
		return
	}

	var Log = map[string]interface{}{
		"Source":src,
		"Details":det,
	}

	_ = temp.Execute(writer, Log)
}