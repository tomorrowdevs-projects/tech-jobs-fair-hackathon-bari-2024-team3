package dataTypes

type User struct {
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	MsgChannel chan string `json:"-"`
}

func (q User) String() string {
	return string(q.Name)
}
