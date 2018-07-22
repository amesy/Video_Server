package dbops

import (
	"Video_Server/api/defs"
	"Video_Server/api/utils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

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
		(video_id, author_id, name, display_ctime, create_time) VALUES(?, ?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}

	_, err = stmtIns.Exec(vid, aid, name, ctime, t)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{
		Id:           vid,
		AuthorId:     aid,
		Name:         name,
		DisplayCtime: ctime,
	}

	defer stmtIns.Close()
	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE video_id=?")
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

	if err == sql.ErrNoRows {
		return nil, nil
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

// 评论 - 增、改

func AddNewComments(aid int, vid string, content string) error {
	// create uuid
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	t := time.Now()

	stmtIns, err := dbConn.Prepare(`INSERT INTO comments
		(comments_id, video_id, author_id, content, time) VALUES(?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(id, vid, aid, content, t)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.CommentInfo, error) {
	// users表 join comments表
	// comments -> author id, video id -> comments
	// users -> author id -> login name
	stmtOut, err := dbConn.Prepare(`SELECT comments.comments_id, users.login_name, comments.content FROM 
		comments INNER JOIN users ON comments.author_id=users.users_id WHERE comments.video_id=? AND 
		comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)

	var res []*defs.CommentInfo

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err = rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}

		c := &defs.CommentInfo{
			Id:      id,
			VideoId: vid,
			Author:  name,
			Content: content,
		}

		res = append(res, c)
	}

	defer stmtOut.Close()

	return res, nil
}
