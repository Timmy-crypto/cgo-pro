package main

import (
	"fmt"
	"github.com/dipperin/dipperin-core/third-party/log"
	"github.com/dipperin/dipperin-core/third-party/rpc"
	"github.com/urfave/cli"
	"github.com/yangchuang123456/cgo-pro/cmd/commands"
	"github.com/yangchuang123456/cgo-pro/common"
	"os"
	"runtime"
	"sort"
)

var commonFlags = []cli.Flag{
	cli.StringFlag{Name: "m", Usage: "method"},
	cli.StringFlag{Name: "p", Usage: "parameters"},
}

var CliCommands = []cli.Command{
	{
		Name:    "rpc",
		Aliases: []string{},
		Usage:   "call rpc function",
		Flags:   commonFlags,
		Action: func(c *cli.Context) error {
			commands.RpcCall(c)
			return nil
		},
		Hidden:   false,
		HideHelp: false,
	},
	{
		Name:    "quit",
		Aliases: []string{"exit"},
		Usage:   "quit",
		Action: func(c *cli.Context) error {
			os.Exit(0)
			return cli.NewExitError("", 0)
		},
		Hidden:   false,
		HideHelp: false,
	},
}

func InitLogger(){
	commands.CliLog = log.New()
	commands.CliLog.SetHandler(log.MultiHandler(log.CliOutHandler))
}

func InitRpcClient(serverIp string,port int) {
	commands.CliLog.Info("init rpc client", "port", port)
	var err error
	//if client, err = rpc.Dial(fmt.Sprintf("http://%v:%d", "127.0.0.1", port)); err != nil {
	//	panic("init rpc client failed: " + err.Error())
	//}
	wsURL := fmt.Sprintf("ws://%v:%d", serverIp, port)
	//l.Info("init rpc client", "wsURL", wsURL)
	if commands.Client, err = rpc.Dial(wsURL); err != nil {
		panic("init rpc client failed: " + err.Error())
	}
}

func appAction(c *cli.Context) error {
	fmt.Println("Hello cgoProCli!")

	//init cmd logger
	InitLogger()

	//init rpc client
	ip := c.String(common.ServerIP)
	port := c.Int(common.WsPortFlagName)
	InitRpcClient(ip,port)

	//new console and start
	console := commands.NewConsole(commands.Executor(c),commands.CgoProCliCompleter)
	console.Prompt.Run()

	return nil
}

func newApp() (nApp *cli.App) {
	nApp = cli.NewApp()
	nApp.Name = "cgoProCli"
	nApp.Version = "0.0.0"
	nApp.Usage = "cgoProCli commandline tool for " + runtime.GOOS + "/" + runtime.GOARCH
	nApp.Description = ``

	nApp.Action = appAction
	nApp.Flags = common.Flags
	nApp.Commands = CliCommands

	sort.Sort(cli.FlagsByName(nApp.Flags))
	sort.Sort(cli.CommandsByName(nApp.Commands))
	return nApp
}


func main() {
	app := newApp()
	if err := app.Run(os.Args); err != nil {
		panic("run cgoProCli failed: " + err.Error())
	}
}
