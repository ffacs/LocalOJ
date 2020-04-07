package Judge

import (
	"fmt"
	"os"
	"strconv"
)

func SaveInfo (startTime ,dir,status,logs string,time int64,mem uint64){
	file,err:=os.OpenFile(dir+"/sub.log",os.O_CREATE|os.O_WRONLY,0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	_,err=file.WriteString(startTime+" "+status+" "+strconv.Itoa(int(time))+"ms"+" "+ParseMemory(mem)+"\n"+logs)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()
}