package defs

type UserCredential struct {
	Username string `json:"username"`
	Pwd		string 	`json:"pwd"`
}

type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}


type SimpleSession struct {
	Username string
	TTL int64
}
