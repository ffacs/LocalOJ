package Judge

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Makefile(data []byte,ProID,lang string) (dir,name string,err error){
	rand.Seed(time.Now().UnixNano())
	path,_:=os.Getwd()
	dir=path+"/submitted/"+ProID+"/"+strconv.Itoa(rand.Int())
	name=dir+"/main"+suffix[lang]
	err=os.Mkdir(dir,0666)
	if err!=nil{
		return
	}
	err=ioutil.WriteFile(name,data,0666)
	if err!=nil{
		return
	}
	return
}

