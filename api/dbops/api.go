package dbops

import (
	"Video_server/api/defs"
	"Video_server/api/utils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

// // 已在同级目录下conn.go文件中初始化完成,起到复用的效果。
// func openConn() *sql.DB {
// 	dbConn, err := sql.Open("mysql", "mysqlcli:12345678@tcp(10.68.7.24:3306)/video_server?charset=utf8")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	return dbConn
// }

// 增加用户凭证.
func AddUserCredential(loginName string, pwd string) error {
	// 使用+拼接容易被撞库或攻击.
	// 使用Prepare预编译.
	stmtIns, err := dbConn.Prepare("Insert INTO users (login_name, pwd) VALUES (?, ?)")

	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)

	if err != nil {
		return err
	}

	defer stmtIns.Close()

	return nil

}

// 获取指定用户对应的密码。
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd from users WHERE login_name=?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	defer stmtOut.Close()

	return pwd, nil
}

// 删除，用两个参数保证校验合理无误。
func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? and pwd=?")
	if err != nil {
		log.Printf("Delete User info %s", err)
		return err
	}

	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	defer stmtDel.Close()

	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	// create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	// video创建时间和写入DB的时间，是串行的。不会出现错乱的情况。
	// DB中的时间方便用来排序，video创建时间方便在页面展示。
	t := time.Now()

	// 既定写法，不要更改。
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info  
		(id, author_id, name, display_ctime) VALUES(?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}

	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}

	defer stmtIns.Close()
	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE video_info=?")
	if err != nil {
		return nil, err
	}

	var (
		aid   int
		name  string
		ctime string
	)

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &ctime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	defer stmtOut.Close()

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}

	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE video_id=?")
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
