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
