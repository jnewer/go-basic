package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id       int
	username string
	password string
}

// 定义一个全局对象db
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

func queryRowDemo() {
	sqlStr := "select id, username, password from user_tbl where id=?"
	var u user
	// 确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.username, &u.password)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}

	fmt.Printf("id:%d name:%s age:%s\n", u.id, u.username, u.password) //id:1 name:tom age:123
}

func queryMultiRow() {
	sqlStr := "select id, username, password from user_tbl where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}

		fmt.Printf("id:%d name:%s age:%s\n", u.id, u.username, u.password)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("初始化失败！,err: %v\n", err)
		return
	}

	fmt.Println("初始化成功")

	// queryRowDemo()
	queryMultiRow()

}
