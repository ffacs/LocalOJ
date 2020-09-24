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

//HandleDetails excutes deatail page
func HandleDetails(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	runid := query.Get("runid")

	src := make([]string, 0)
	det := make([]string, 0)

	paths, _ := ioutil.ReadDir("./submitted/" + runid)
	for _, path := range paths {
		file, _ := os.Open("./submitted/" + runid + "/" + path.Name())
		buffer := bufio.NewReader(file)
		if path.Name() == "sub.log" {
			for {
				line, _, err := buffer.ReadLine()
				if err == io.EOF {
					break
				}
				det = append(det, string(line))
			}
		} else {
			for {
				line, _, err := buffer.ReadLine()
				if err == io.EOF {
					break
				}
				src = append(src, string(line))
			}
		}
		_ = file.Close()
	}

	temp, err := template.ParseFiles("./static/details.temp")
	if err != nil {
		fmt.Println(err)
		_, _ = writer.Write([]byte("502")) //Waiting for a 502 page
		return
	}

	var Log = map[string]interface{}{
		"Source":  src,
		"Details": det,
	}

	_ = temp.Execute(writer, Log)
}
