package main

import (
	"fmt"
	"github.com/dipperin/dipperin-core/third-party/log"
	"github.com/urfave/cli"
	"github.com/yangchuang123456/cgo-pro/common"
	"github.com/yangchuang123456/cgo-pro/server/service"
	"os"
	"os/signal"
	"runtime"
	"sort"
	"syscall"
)


func appAction(c *cli.Context) error {
	fmt.Println("Hello cgoProServer!")

	//init rpc client
	endPointsConf := &service.EndPoints{
		HTTPHost : c.String(common.HttpHostFlagName),
		WSHost:c.String(common.WsHostFlagName),
		HttpEndpoint: c.Int(common.HttpPortFlagName),
		WsEndpoint:c.Int(common.WsPortFlagName),
		AllowHosts:[]string{"*"},
	}
	service.InitRpc(endPointsConf)

	return nil
}

func newApp() (nApp *cli.App) {
	nApp = cli.NewApp()
	nApp.Name = "cgoProServer"
	nApp.Version = "0.0.0"
	nApp.Usage = "cgoProServer for " + runtime.GOOS + "/" + runtime.GOARCH
	nApp.Description = ``

	nApp.Action = appAction
	nApp.Flags = common.Flags

	sort.Sort(cli.FlagsByName(nApp.Flags))
	sort.Sort(cli.CommandsByName(nApp.Commands))
	return nApp
}

func main(){
	log.Info("~~~~~~~~~cgoProServer start~~~~~~~~~~~~")
	app := newApp()

	if err := app.Run(os.Args); err != nil {
		panic("run cgoProServer failed: " + err.Error())
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	for {
		select {
			case s := <-c:
			log.Info("got system signal", "signal", s)
			os.Exit(0)
		}
	}
	log.Info("~~~~~~~~~cgoProServer end~~~~~~~~~~~~")
}

