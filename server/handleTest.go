package server

import (
	"ffacs/LocalOJ/db"
	"ffacs/LocalOJ/judge"
	"fmt"
	"net/http"
)

//HandleTest receive http request and ready to test
func HandleTest(w http.ResponseWriter, r *http.Request) {

	user := checklogin(w, r)
	if user == nil {
		return
	}

	source := r.PostFormValue("source")
	ProID := r.PostFormValue("ProID")
	lang := r.PostFormValue("language")
	uid := user.ID

	var sub = db.Submission{
		RunID:   0,
		Subtime: "",
		Userid:  uid,
		Runmem:  0,
		Runtime: 0,
		Status:  "pending",
		Lang:    lang,
		Pid:     ProID,
	}
	newID, err := db.InsertSubmission(sub)
	if err != nil {
		return
	}
	sub.RunID = newID

	sub.Status = "Submitted" //Ready for judging
	db.UpdateSubmission(sub)

	dir, name, err := judge.Makefile([]byte(source), sub)
	if err != nil {
		fmt.Println("handleTest makefile failed")
		sub.Status = "UKE"
		db.UpdateSubmission(sub)
		return
	}

	judge.JudgeQueue <- judge.Judgement{
		Dir:  dir,
		Name: name,
		Sub:  sub,
	}

	w.Write([]byte("<script language=\"javascript\" type=\"text/javascript\">window.location.href=\"/status\";</script>"))

	//需要改成302跳转到status界面
	// w.Write([]byte(logs)) //返回浏览器信息
	// w.Write([]byte("\n\n" + time.Now().Format("2006-01-02 15:04:05") + ": Done\n"))
	// w.Write([]byte("\n\nstatus:" + status))
}
