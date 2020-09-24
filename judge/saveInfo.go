package judge

import (
	"ffacs/LocalOJ/db"
	"fmt"
	"os"
	"strconv"
)

func saveInfo(dir, logs string, sub db.Submission) {
	file, err := os.OpenFile(dir+"/sub.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(sub.Status + " " + strconv.Itoa(int(sub.Runtime)) + "ms" + " " + ParseMemory(sub.Runmem) + "\n" + logs)
	if err != nil {
		fmt.Println(err)
		return
	}
}
