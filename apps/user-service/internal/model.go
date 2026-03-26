package internal

import (
	"fmt"
)
type User Struct {

	Id		uuid.UUID	`json:id`
	Username string `json:username`
	Password string `json:password`
	Github	string `json:github``
}