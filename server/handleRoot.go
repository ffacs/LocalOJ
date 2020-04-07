package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func HandleRoot(w http.ResponseWriter, r *http.Request){
		fileObj ,err:=os.Open("./index.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer fileObj.Close()
		data ,err:=ioutil.ReadAll(fileObj)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Write(data)
}