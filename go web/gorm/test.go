package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
	u  User
)

type User struct {
	id    int
	name  string
	phone string
}

func initDB() (err error) {
	db, err = sql.Open("mysql", "root:85808353j@tcp(127.0.0.1:3306)/chapter4")
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

//查询单行数据
func queryRow() {
	err := db.QueryRow("select uid,name,phone from user where uid=?", 1).Scan(&u.id, &u.name, &u.phone)
	if err != nil {
		fmt.Printf("scan filed,err:%v\n", err)
		return
	}
	fmt.Printf("uid:%d name:%s phone:%s\n", u.id, u.name, u.phone)
}

//查询多行数据
func queryMultiRow() {
	row, err := db.Query("select uid,name,phone from user where uid > ?", 0)
	if err != nil {
		fmt.Printf("query filed,err:%v\n", err)
		return
	}
	defer row.Close()
	for row.Next() {
		err := row.Scan(&u.id, &u.name, &u.phone)
		if err != nil {
			fmt.Printf("scan filed,err:%v\n", err)
			return
		}
		fmt.Printf("uid:%d name:%s phone:%s\n", u.id, u.name, u.phone)
	}

}

//插入数据
func insertRow() {
	ret, err := db.Exec("insert into user (name,phone)values (?,?)", "白板", "139988557766")
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	uid, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed,err:%v\n", err)
		return
	}
	fmt.Printf("insert succedd, the id is %d \n", uid)
}

//更新数据
func updateRow() {
	ret, err := db.Exec("update user set name=? where uid= ?", "张三封", 5)
	if err != nil {
		fmt.Printf("update failed,err:%v\n", err)
		return
	}
	uid, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowAffected failed,err:%v\n", err)
		return
	}
	fmt.Printf("update succedd, affected rows: %d \n", uid)
}

//删除数据
func deleteRow() {
	ret, err := db.Exec("delete from user where uid=?", 8)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}
	uid, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowAffected failed,err:%v\n", err)
		return
	}
	fmt.Printf("delete succedd, affected rows: %d \n", uid)
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
	} else {
		fmt.Println("成功")
		deleteRow()
	}
}
