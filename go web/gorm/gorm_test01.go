package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type GormUser struct {
	ID       uint   `json:"id"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {
	db, err := gorm.Open("mysql", "root:85808353j@tcp(127.0.0.1:3306)/chapter4")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormUser := new(GormUser)
	res := db.First(&gormUser, "phone=?", "18888888888")
	fmt.Println(gormUser)
	fmt.Println("找到的记录数：", res.RowsAffected)
}
