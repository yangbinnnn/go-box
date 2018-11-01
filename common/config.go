package common

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	LogPath       string `json:"logpath"`
	HTTPAddr      string `json:"httpaddr"`
	WebIndex      string `json:"webindex"`
	WebStatic     string `json:"webstatic"`
	APIPrefix     string `json:"apiprefix"`
	MongoURI      string `json:"mongouri"`
	MongoName     string `json:"mongoname"`
	FileStoreType string `json:"filestoretype"`
	LocalFileHome string `json:"localfilehome"`
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
