package utils

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func ParsePostData(data io.ReadCloser) (rData map[string]interface{}, err error) {

	res, err := ioutil.ReadAll(data)
	if err != nil {
		return nil, err
	}

	var r interface{}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}

	rData, ok := r.(map[string]interface{})
	if !ok {
		return nil, errors.New("parse data err 断言失败")
	}

	return rData, nil
}

type Z map[string]interface{}

func JSON(res Z) (data []byte) {
	data, _ = json.Marshal(res)
	return
}

func ParsJsonFile() map[string]interface{} {
	conf, err := os.Open("./config.json")
	if err != nil {
		log.Fatalf("打开配置文件错误:%s", err)
	}
	data, err := ioutil.ReadAll(conf)
	if err != nil {
		log.Fatalf("读取配置文件错误:%s", err)
	}
	var a interface{}
	err = json.Unmarshal(data, &a)
	if err != nil {
		log.Fatalf("读取配置文件错误:%s", err)
	}

	return a.(map[string]interface{})
}
