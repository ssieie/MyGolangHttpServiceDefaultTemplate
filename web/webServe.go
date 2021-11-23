package web

import (
	X "blog_service/utils"
	"net/http"
)

type Service struct{}

func (z *Service) Login(w http.ResponseWriter, r *http.Request) {
	data, err := X.ParsePostData(r.Body)
	if err != nil {
		_, _ = w.Write(X.JSON(X.Z{
			"code":    1,
			"message": "请求参数错误！",
		}))
		return
	}
	_, _ = w.Write(X.JSON(X.Z{
		"code": 0,
		"data": data,
	}))
}
