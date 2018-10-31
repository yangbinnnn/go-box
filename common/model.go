package common

import (
	"time"
)

type User struct {
	Email        string `json:"email" bson:"_id"`
	Name         string `json:"name" bson:"name"`
	RegisterTime time.Time
}
