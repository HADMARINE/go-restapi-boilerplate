package models

import (
	"github.com/zebresel-com/mongodm"
)

// User is model of User
type User struct {
	mongodm.DocumentBase `json:",inline" bson:",inline"`

	Userid    string `json:"userid" bson:"userid"`
	Password  string `json:"password" bson:"password"`
	Authority string `json:"authority" bson:"authority"`
}
