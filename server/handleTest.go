package server

import (
	"ffacs/LocalOJ/Judge"
	"fmt"
	"net/http"
	"strings"
	"time"
)



func HandleTest(w http.ResponseWriter, r *http.Request){
	startTime:=time.Now().Format("2006-01-02 15:04:05")
	w.Write([]byte(startTime+": Start test\n\n\n"))
	args:=r.URL.Query()
	source:=args.Get("source")
	ProID:=args.Get("ProID")
	lang:=args.Get("language")
	dir,name,err:=Judge.Makefile([]byte(source),ProID,lang)
	if err != nil {
		w.Write([]byte("UKE"))
		fmt.Printf("Makefile failed: %v\n",err)
		return
	}

	status,logs,rtime,rmem:=Judge.Parse(ProID,lang,name)

	logs=strings.Replace(logs,dir+"/","",-1) //删除输出中的文件信息

	Judge.SaveInfo(startTime,dir,status,logs,rtime,rmem) //存储测评信息

	w.Write([]byte(logs)) //返回浏览器信息
	w.Write([]byte("\n\n"+time.Now().Format("2006-01-02 15:04:05")+": Done\n"))
	w.Write([]byte("\n\nstatus:"+status))
}
