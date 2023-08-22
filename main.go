package main

import (
	"StarRocks_License/app"
	"StarRocks_License/util"
)

func main() {
	util.Parm()
	util.Loggers()
	app.LicenseCrontab()
}
