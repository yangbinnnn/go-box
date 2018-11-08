package db

import (
	"go-box/common"

	"github.com/globalsign/mgo"
	"github.com/go-redis/redis"
)

var (
	MgoDB   *mgo.Database
	RedisDB *redis.Client
)

func InitDB() {
	// mongod
	initMGO()

	// redis
	initRedis()

	// table
	initTable()
}

func initMGO() {
	config := common.GlobalConfig
	opts, err := mgo.ParseURL(config.MongoURL)
	if err != nil {
		panic(err)
	}
	session, err := mgo.DialWithInfo(opts)
	if err != nil {
		panic(err)
	}

	MgoDB = session.DB(opts.Database)
}

func initRedis() {
	config := common.GlobalConfig
	opts, err := redis.ParseURL(config.RedistURL)
	if err != nil {
		panic(err)
	}
	RedisDB = redis.NewClient(opts)
}

func initTable() {
	UserTable = MgoDB.C("User")
}
