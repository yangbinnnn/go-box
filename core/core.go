package core

import (
	"go-box/db"

	"github.com/globalsign/mgo"
)

var (
	UserTable *mgo.Collection
)

// InitCore put your business here
func InitCore() {
	go MSGLoop()

	// database tables
	initTable()
}

func initTable() {
	UserTable = db.MgoDB.C("User")
}
