package dataTypes

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	MsgChannel chan string `json:"-"`
}

func (u User) String() string {
	uJSON, err := json.MarshalIndent(u, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error converting User to JSON: %v", err)
	}
	return string(uJSON)
}
