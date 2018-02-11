package main

import (
	"github.com/henrylee2cn/pholcus/config"
	"github.com/henrylee2cn/pholcus/exec"
	_ "github.com/laidingqing/feichong/spider/pholcus_lib"
)

func main() {
	config.MGO_CONN_STR = "127.0.0.1:27017"
	config.DB_NAME = "localDb"
	config.MGO_CONN_CAP = 1024
	exec.DefaultRun("web")
}
