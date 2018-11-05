package db

import (
	"go-box/common"

	"github.com/globalsign/mgo"
)

var (
	MgoDB *mgo.Database
)

func InitDB() {
	// mongod
	initMGO()

	// table
	initTable()
}

func initMGO() {
	config := common.GlobalConfig
	session, err := mgo.Dial(config.MongoURI)
	if err != nil {
		panic(err)
	}

	db := config.MongoName
	MgoDB = session.DB(db)
}

func initTable() {
	UserTable = MgoDB.C("User")
}
