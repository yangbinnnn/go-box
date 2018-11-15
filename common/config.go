package common

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Debug      bool   `json:"debug"`
	LogPath    string `json:"logpath"`
	HTTPAddr   string `json:"httpaddr"`
	GrpcAddr   string `json:"grpcaddr"`
	GrpcEnable bool   `json:"grpcenable"`
	WebIndex   string `json:"webindex"`
	WebStatic  string `json:"webstatic"`
	DocStatic  string `json:"docstatic"`
	APIPrefix  string `json:"apiprefix"`
	MongoURL   string `json:"mongourl"`
	RedisURL   string `json:"redisurl"`
	File       struct {
		Type string `json:"type"`
		Home string `json:"home"`
	} `json:"file"`
}

var GlobalConfig *Config

// Init config
func InitConfig(p string) {
	data, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err.Error())
	}
	config := &Config{}
	err = json.Unmarshal(data, config)
	if err != nil {
		panic(err.Error())
	}
	GlobalConfig = config
}
