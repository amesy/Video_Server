package defs

// request
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

// response
type SignedUp struct {
	Success   bool   `json: "success"`
	SessionId string `json: "session_id"`
}

// Data model
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
	// create_time  string // 该属性意义不大，只在DB model中有用。
}

type CommentInfo struct {
	Id      string
	VideoId string
	Author  string
	Content string
}

type SimpleSession struct {
	Username string // login name
	TTL      int64
}
