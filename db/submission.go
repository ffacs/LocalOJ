package db

import (
	"fmt"
)

//Submission describes a submission
type Submission struct {
	RunID    int64
	Subtime  string
	Runmem   uint64
	Runtime  uint64
	Status   string
	Lang     string
	Pid      string
	Username string
}

//InsertSubmission for inserting user
func InsertSubmission(sub Submission) (newID int64, err error) {
	newID, err = 0, nil
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Printf("Insert submission tx failed:%v\n", err)
		return 0, err
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO LocalOJ.submission (`sub_time`,`run_mem`,`run_time`,`status`,`lang`,`pid`,`user_name`) VALUES(now(),?,?,?,?,?,?)")
	if err != nil {
		fmt.Printf("Insert submission prepare failed:%v\n", err)
		return 0, err
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(0, 0, sub.Status, sub.Lang, sub.Pid, sub.Username)
	if err != nil {
		fmt.Printf("Insert submission failed:%v\n", err)
		return 0, err
	}
	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	newID, err = res.LastInsertId()
	fmt.Println("Insert Submission successfully")
	return
}

//UpdateSubmission update submissons status;
func UpdateSubmission(sub Submission) error {
	tx, err := DB.Begin()
	if err != nil {
		fmt.Printf("Update submission tx failed:%v\n", err)
		return err
	}
	//准备sql语句
	stmt, err := tx.Prepare("UPDATE LocalOJ.submission SET run_mem=?,run_time=?,status=? where run_id=?")
	if err != nil {
		fmt.Printf("Update submission prepare failed:%v\n", err)
		return err
	}
	//将参数传递到sql语句中并且执行
	_, err = stmt.Exec(sub.Runmem, sub.Runtime, sub.Status, sub.RunID)
	if err != nil {
		fmt.Printf("Update submission failed:%v\n", err)
		return err
	}
	//将事务提交
	tx.Commit()
	fmt.Println("Update Submission successfully")
	return nil
}

//QuerySubmission excutes query for all submissions
func QuerySubmission() []Submission {
	rows, err := DB.Query("SELECT * FROM LocalOJ.submission ")
	if err != nil {
		fmt.Printf("Query failed :%v\n", err)
		return nil
	}
	defer rows.Close()
	var res []Submission
	for rows.Next() {
		var sub Submission
		rows.Scan(&sub.RunID, &sub.Subtime, &sub.Runmem, &sub.Runtime, &sub.Status, &sub.Lang, &sub.Pid, &sub.Username)
		res = append(res, sub)
	}
	fmt.Println("Query Submission successfully")
	return res
}

//QuerySubmissionByPid returns Submission of problem_pid
func QuerySubmissionByPid(pid string) []Submission {
	rows, err := DB.Query("SELECT * FROM LocalOJ.submission where pid=?", pid)
	if err != nil {
		fmt.Printf("Query failed :%v\n", err)
		return nil
	}
	defer rows.Close()
	var res []Submission
	for rows.Next() {
		var sub Submission
		rows.Scan(&sub.RunID, &sub.Subtime, &sub.Runmem, &sub.Runtime, &sub.Status, &sub.Lang, &sub.Pid, &sub.Username)
		res = append(res, sub)
	}
	fmt.Println("Query Submission successfully")
	return res
}

//QuerySubmissionByUser excutes user's submission
func QuerySubmissionByUser(user *User) []Submission {
	rows, err := DB.Query("SELECT * FROM LocalOJ.submission where user_name=?", user.Name)
	if err != nil {
		fmt.Printf("Query failed :%v\n", err)
		return nil
	}
	defer rows.Close()
	var res []Submission
	for rows.Next() {
		var sub Submission
		rows.Scan(&sub.RunID, &sub.Subtime, &sub.Runmem, &sub.Runtime, &sub.Status, &sub.Lang, &sub.Pid, &sub.Username)
		res = append(res, sub)
	}
	fmt.Println("Query Submission successfully")
	return res
}
