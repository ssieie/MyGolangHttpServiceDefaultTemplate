package web

import (
	"blog_service/config"
	"blog_service/middleware"
	"net/http"
)

func InitHttp() (err error) {
	var service Service

	// service
	http.Handle("/login", middleware.Cross(http.HandlerFunc(service.Login)))
	err = http.ListenAndServe(config.Host, nil)

	if err != nil {
		return err
	}
	return nil
}
