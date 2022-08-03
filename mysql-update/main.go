package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	dsn := "root:@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
func updateData() {
	sqlStr := "update user_tbl set username=?, password=? where id=?"
	ret, err := db.Exec(sqlStr, "kite2", "kite123", "2")
	if err != nil {
		fmt.Printf("更新失败, err: %v\n", err)
		return
	}

	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("更新失败, err: %v\n", err)
		return
	}

	fmt.Printf("更新成功, 更新的行数： %d.\n", rows)

}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("初始化失败！,err:%v\n", err)
		return
	} else {
		fmt.Printf("初始化成功\n")
	}

	updateData()
}
