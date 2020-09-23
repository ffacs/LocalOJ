package judge

import (
	"ffacs/LocalOJ/db"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

//Makefile is for making dir and file to store code
func Makefile(data []byte, sub db.Submission) (dir, name string, err error) { //需要加入数据库
	path, _ := os.Getwd()
	// dir is the dirctory where code restored
	// name is the name of source code
	dir = path + "/submitted/" + strconv.Itoa(int(sub.RunID))
	name = dir + "/main" + suffix[sub.Lang]
	if err = os.Mkdir(dir, 0666); err != nil {
		fmt.Println(err)
		return
	}
	if err = ioutil.WriteFile(name, data, 0666); err != nil {
		fmt.Println(err)
		return
	}
	return
}
