package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:fengjutian@(127.0.0.1:3306)/test_db")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err := db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}

	sql := "insert init teacher value ('dd', '1')"
	result,_ := db.Exec(sql)
	n,_ := result.RowsAffected()
	fmt.Println("受影响的记录数是："， n)


}