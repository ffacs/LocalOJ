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
	uname := user.Name

	var sub = db.Submission{
		RunID:    0,
		Subtime:  "",
		Username: uname,
		Runmem:   0,
		Runtime:  0,
		Status:   "pending",
		Lang:     lang,
		Pid:      ProID,
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

	w.Write([]byte("<script language=\"javascript\" type=\"text/javascript\">window.location.href=\"/status?ProID=-1\";</script>"))

}
