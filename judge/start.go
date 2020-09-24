package judge

import (
	"ffacs/LocalOJ/db"
	"fmt"
	"strings"
)

//Judgement is a struct which contains the information of submission
type Judgement struct {
	Dir  string
	Name string
	Sub  db.Submission
}

//JudgeQueue is judge queue
var JudgeQueue chan Judgement

//StartJudge starts judge
func StartJudge(ind int) {
	for true {

		onejudge := <-JudgeQueue // Waitting for Judgement

		dir, name, sub := onejudge.Dir, onejudge.Name, onejudge.Sub

		fmt.Printf("JudgeMachine %v ready to judge runid : %v\n", ind, sub.RunID)

		sub.Status = "running" //Ready for judging
		db.UpdateSubmission(sub)

		status, logs, rtime, rmem := Parse(sub.Pid, sub.Lang, name)

		logs = strings.Replace(logs, dir+"/", "", -1) //删除输出中的文件信息

		sub.Status = status
		sub.Runtime = rtime
		sub.Runmem = rmem
		db.UpdateSubmission(sub)

		saveInfo(dir, logs, sub)

	}
}

//Init the JudgeQueue and make JudgeRunning
func init() {
	JudgeQueue = make(chan Judgement, 50)
	for i := 0; i < 10; i++ {
		go StartJudge(i)
	}
	fmt.Println("JudgeMachine is ready")
}
