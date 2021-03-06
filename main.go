package main

import (
	"flag"
	"path"
	"runtime"

	"github.com/zhs007/tradingview_udf_go/base"
)

func main() {
	runtime.GOMAXPROCS(1)

	var curdir string
	flag.StringVar(&curdir, "path", "./", "run path")
	flag.Parse()

	base.LoadConfig(path.Join(curdir, "./config.yaml"), curdir)
	base.InitLogger()
	base.Info("restart...")
	cfg := base.GetConfig()

	StartServ(cfg.WebServAddr)
}
