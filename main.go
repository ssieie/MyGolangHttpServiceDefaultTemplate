package main

import (
	"blog_service/model"
	"blog_service/web"
	"log"
)

func main() {
	err := model.InitDB()
	if err != nil {
		log.Fatalf("init DB err %s \n", err.Error())
	}

	err = web.InitHttp()
	if err != nil {
		log.Fatalf("http server err %s \n", err.Error())
	}

}
