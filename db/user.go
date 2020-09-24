package db

import (
	"fmt"
)

//User describes a user
type User struct {
	Psw   string
	Name  string
	Email string
}

//QueryUserByName returns ID
func QueryUserByName(name string) (*User, error) {
	rows, err := DB.Query("SELECT * FROM LocalOJ.user where name=?", name)
	if err != nil {
		fmt.Printf("Query failed :%v\n", err)
		return nil, err
	}
	fmt.Println("Query user by name successfully")
	defer rows.Close()
	for rows.Next() {
		var res = new(User)
		rows.Scan(&(res.Psw), &(res.Name), &(res.Email))
		return res, nil
	}
	return nil, nil
}

//InsertUser for inserting user
func InsertUser(user *User) (newID int64, err error) {
	newID, err = 0, nil
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO LocalOJ.user (`psw`,`name`,`email`) VALUES (?,?,?)")
	if err != nil {
		fmt.Printf("Prepare fail:%v\n", err)
		return
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(user.Psw, user.Name, user.Email)
	if err != nil {
		fmt.Println("Exec fail")
		return
	}
	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	newID, err = res.LastInsertId()
	return
}
