package db

import (
	"fmt"
	"time"
)

//Cookie contains client`s cookie information
type Cookie struct {
	SessionID string
	ID        string
}

//QueryUserBySession returns user structrue
func QueryUserBySession(cookie *Cookie) (*User, error) {
	user, err := QueryUserByName(cookie.ID)
	if err != nil {
		fmt.Println("QueryUserBySession failed")
		return nil, err
	}
	return user, nil
}

//QueryCookieBySessionID accepet sessionID returns cookie
func QueryCookieBySessionID(sessionID string) (*Cookie, error) {
	rows, err := DB.Query("SELECT sessionID,id FROM LocalOJ.cookie where sessionID=?", sessionID)
	if err != nil {
		fmt.Printf("QueryCookieBySessionID failed :%v\n", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		res := new(Cookie)
		rows.Scan(&res.SessionID, &res.ID)
		return res, nil
	}
	return nil, nil
}

//DeleteOuttimeSession delete cookie when user logout
func DeleteOuttimeSession() error {
	fmt.Println("GC engine Started successfully")
	for true {
		time.Sleep(time.Duration(300) * time.Second)
		//开启事务
		tx, err := DB.Begin()
		if err != nil {
			fmt.Printf("Delete cookie tx failed:%v\n", err)
			return err
		}
		//准备sql语句
		stmt, err := tx.Prepare("DELETE FROM LocalOJ.cookie WHERE time_to_sec(timediff(now(),LAtime))>=600")
		if err != nil {
			fmt.Printf("Delete cookie prepare failed:%v\n", err)
			return err
		}
		//将参数传递到sql语句中并且执行
		_, err = stmt.Exec()
		if err != nil {
			fmt.Printf("Delete cookie failed:%v\n", err)
			return err
		}
		//将事务提交
		tx.Commit()
	}
	return nil
}

//InsertSession insert new cookie
func InsertSession(cookie *Cookie) error {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Printf("Insert cookie tx failed:%v\n", err)
		return err
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO LocalOJ.cookie (`sessionID`,`id`,`LAtime`) VALUES(?,?,now())")
	if err != nil {
		fmt.Printf("Insert cookie prepare failed:%v\n", err)
		return err
	}
	//将参数传递到sql语句中并且执行
	_, err = stmt.Exec(cookie.SessionID, cookie.ID)
	if err != nil {
		fmt.Printf("Insert cookie failed:%v\n", err)
		return err
	}
	//将事务提交
	tx.Commit()
	fmt.Println("Insert cookie successfully")
	return nil
}

//UpdateLAtime updates cookie's Last Active time
func UpdateLAtime(cookie *Cookie) error {
	tx, err := DB.Begin()
	if err != nil {
		fmt.Printf("Update LAtime tx failed:%v\n", err)
		return err
	}
	//准备sql语句
	stmt, err := tx.Prepare("UPDATE LocalOJ.cookie SET LAtime=now() where sessionID=?")
	if err != nil {
		fmt.Printf("Update LAtime prepare failed:%v\n", err)
		return err
	}
	//将参数传递到sql语句中并且执行
	_, err = stmt.Exec(cookie.SessionID)
	if err != nil {
		fmt.Printf("Update LAtime failed:%v\n", err)
		return err
	}
	//将事务提交
	tx.Commit()
	fmt.Println("Update LAtime successfully")
	return nil
}
