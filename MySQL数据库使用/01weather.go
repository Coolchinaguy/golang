package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	_ "mysql-master"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	//fmt.Println(25+rand.Intn(8))
	//1.打开连接
	db,err := sql.Open("mysql","root:@tcp(localhost:3306)/weather")
	if err != nil {
		fmt.Println("sql err = ",err)
		return
	}

	defer  db.Close()

	db.Ping()
	//2.预处理sql
	//?表示占位符
	stm,err:=db.Prepare("insert  into weather values (default ,?,?)")
	if err != nil {
		fmt.Println("prepare err = ",err)
		return
	}

	defer  stm.Close()


	result,err := stm.Exec("5",25+rand.Intn(5))
	if err != nil {
		fmt.Println("Exec err = ",err)
		return
	}

	count,err :=result.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected err = ",err)
		return
	}

	if count > 0 {
		fmt.Println("输出成功")
	}else {
		fmt.Println("输出失败")
	}

	id,err :=result.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId err = ",err)
		return
	}
	fmt.Println(id)
}
