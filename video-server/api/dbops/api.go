package dbops

import (
	"database/sql"
	"log"
	"time"
	"video-server/api/defs"
	"video-server/api/utils"
)

func AddUserCrendential(loginName string,pwd string) error {
	stmtIns,err := dbCoon.Prepare("INSERT INTO users (login_name,pwd) VALUES (?,?)")
	if err != nil {
		panic(err)
	}
	_, _ = stmtIns.Exec(loginName, pwd)
	_ = stmtIns.Close()
	return nil
}

func GetUserCrendential(loginName string)(string,error){
	stmtOut,err := dbCoon.Prepare("select pwd From users where login_name = ?")
	if err != nil {
		log.Printf("%s",err)
		return "",err
	}
	var pwd string
	_ = stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()

	return pwd,nil

}

func DeleteUser(loginName string,pwd string) error  {
	stmtDel,err := dbCoon.Prepare("delete from users where lgoin_name=? and pwd=?")
	if err != nil {
		log.Printf("%s",err)
		return err
	}

	stmtDel.Exec(loginName,pwd)
	stmtDel.Close()
	return nil
}

func AddNewVideo(aid int,name string)(*defs.VideoInfo,error){
	// create uuid
	vid,err := utils.NewUUID()
	if err != nil {
		return nil,err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006,15:04:05")
	stmtIns,err :=dbCoon.Prepare("insert into video_info (id,author_id,name,display_ctime) values(?,?,?,?)")
	 _,err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil{
		return nil,err
	}

	res := &defs.VideoInfo{
		Id:vid,
		AuthorId:aid,
		Name:name,
		DisplayCtime:ctime,
	}

	defer stmtIns.Close()

	return res,nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut,err := dbCoon.Prepare("select author_id,name,display_ctime from video_info where id=?")
	if err != nil {
		return nil,err
	}

	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows{
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}

	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbCoon.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	  }

	defer stmtDel.Close()
	return nil
}

