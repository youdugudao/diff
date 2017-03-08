package config

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"errors"
)

var Password string
var Port string

func init() {
	if Port == "" {
		type basicData struct {
			Password string
			Port     string
		}
		f, err := ioutil.ReadFile("./config/basic.json")
		if err != nil {
			panic("读取basic配置文件失败" + err.Error())
		}
		data := basicData{}
		err = json.Unmarshal(f, &data)
		if err != nil {
			panic("解析basic配置文件失败" + err.Error())
		}
		Password = data.Password
		Port = data.Port
		log.Println("读取basic配置文件成功")
	}
}

func GetKeywords() (keywords map[string]int, err error) {
	keywords = make(map[string]int)
	var keys_file [][]string
	f, err := ioutil.ReadFile("./config/keywords.json")
	if err != nil {
		err = errors.New("读取keywords配置文件失败" + err.Error())
		return
	}
	err = json.Unmarshal(f, &keys_file)
	if err != nil {
		err = errors.New("解析keywords配置文件失败" + err.Error())
	}
	for k, v := range keys_file {
		for _, vv := range v {
			keywords[vv] = k
		}
	}
	return
}
