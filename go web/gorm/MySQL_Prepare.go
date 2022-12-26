package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User1 struct {
	uid   int
	name  string
	phone string
}

var (
	db1   *sql.DB
	user1 User1
)

func InitDatabase() (err error) {
	db1, err = sql.Open("mysql", "root:85808353j@tcp(127.0.0.1:3306)/chapter4")
	if err != nil {
		return err
	}
	err = db1.Ping()
	if err != nil {
		return err
	}
	return nil
}

//预处理查询
func prepareQuery() {
	stmt, err := db1.Prepare("select uid,name,phone from user where uid > ? ")
	if err != nil {
		fmt.Printf("prepare failed,err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed,err:%v\n", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&user1.uid, &user1.name, &user1.phone)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
			return
		}
		fmt.Printf("uid:%d name:%s phpone:%s\n", user1.uid, user1.name, user1.phone)
	}
}

//预处理插入
func prepareInsert() {
	stmt, err := db1.Prepare("insert into user (name,phone)values (?,?)")
	if err != nil {
		fmt.Printf("prepare failed,err:%v\n", err)
		return
	}
	defer stmt.Close()
	ret, err := stmt.Exec("barry", 18799887766)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	fmt.Println(ret.LastInsertId())
	ret, err = stmt.Exec("jim", 18899888888)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	fmt.Println("insert success")
	fmt.Println(ret.LastInsertId())
}
func main() {
	err := InitDatabase()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	prepareQuery()
	prepareInsert()
}
