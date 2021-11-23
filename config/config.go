package config

import (
	"blog_service/utils"
	"fmt"
	"os"
)

var (
	MysqlUri       string
	Host           string
)

func init() {
	data := utils.ParsJsonFile()

	MysqlUri = data["mysqlUri"].(string)
	Host = data["host"].(string)

	_, _ = fmt.Fprintf(os.Stdout, "%s", data["describe"])
}
