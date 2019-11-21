package dbops

import "database/sql"

var (
	dbCoon *sql.DB
	err error
)

func init(){
	dbConn,err :=sql.Open("mysql",
		"root:514814@tcp(localhost:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err)
	}
}