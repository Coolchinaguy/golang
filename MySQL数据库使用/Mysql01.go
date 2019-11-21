package main

import (
	"database/sql"
	"fmt"
	_ "mysql-master"
)

func main() {
	 //1.打开连接
	 db,err := sql.Open("mysql","root:@tcp(localhost:3306)/go-first")
	 if err != nil {
	 	fmt.Println("sql err = ",err)
		 return
	 }

	 defer  db.Close()

	 db.Ping()
	 //2.预处理sql
	 //?表示占位符
	stm,err:=db.Prepare("insert  into people values (default ,?,?)")
	if err != nil {
		fmt.Println("prepare err = ",err)
		return
	}

	defer  stm.Close()



	//参数和占位符对应
	result,err :=stm.Exec("张三","北京")
	if err != nil {
		fmt.Println("exec err = ",err)
		return
	}

	//3.获取结果

	counter,err:=result.RowsAffected()
	if err != nil {
		fmt.Println("rows err = ",err)
		return
	}

	if counter > 0 {
		fmt.Println("输出成功")
	}else {
		fmt.Println("输出失败")
	}

	//可能需要获取到新增时主键的值
	id,err :=result.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId err = ",err)
		return
	}

	fmt.Println(id)
}
