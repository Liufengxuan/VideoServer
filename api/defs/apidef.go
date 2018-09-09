package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

type Video struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}
