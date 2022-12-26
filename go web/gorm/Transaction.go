package main

import (
	"database/sql"
	"fmt"
)

var (
	db2 *sql.DB
)

func InitD() (err error) {
	db2, err = sql.Open("mysql", "root:85808353j@tcp(127.0.0.1:3306)/chapter4")
	if err != nil {
		return err
	}
	err = db2.Ping()
	if err != nil {
		return err
	}
	return nil
}
func transaction() {
	tx, err := db2.Begin() //开启事务
	if err != nil {
		if tx != nil { //可以理解成tx是事务日志，若日志为空则没必要回滚，只有日志不为空且报错时才需要回滚
			tx.Rollback() //回滚
		}
		fmt.Printf("begin trans failed,err:%v\n", err)
		return
	}
	_, err = tx.Exec("update user set name='james' where uid=?", 1)
	if err != nil {
		tx.Rollback() //回滚
		fmt.Printf("exec sql1 failed,err:%v\n", err)
		return
	}
	_, err = tx.Exec("update user set name='james' where uid=?", 3)
	if err != nil {
		tx.Rollback() //回滚
		fmt.Printf("exec sql2 failed,err:%v\n", err)
		return
	}
	tx.Commit() //提交事务
	fmt.Println("exec transaction success!")
}
func main() {
	err := InitD()
	if err != nil {
		fmt.Printf("init failed,err:%v\n", err)
	}
	transaction()
}
