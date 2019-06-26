package common

import "github.com/urfave/cli"

var (
	ServerIP         = "server_ip"
	HttpHostFlagName = "http_host"
	WsHostFlagName   = "ws_host"
	WsPortFlagName   = "ws_port"
	HttpPortFlagName = "http_port"
)
var Flags = []cli.Flag{
	cli.StringFlag{
		Name:  ServerIP,
		Usage: "set server ip",
		Value: "127.0.0.1",
	},
	cli.StringFlag{
		Name:  HttpHostFlagName,
		Usage: "set http host",
		Value: "127.0.0.1",
	},
	cli.StringFlag{
		Name:  WsHostFlagName,
		Usage: "set web socket host",
		Value: "127.0.0.1",
	},
	cli.IntFlag{
		Name:  HttpPortFlagName,
		Usage: "set http port",
		Value: 7001,
	},
	cli.IntFlag{
		Name:  WsPortFlagName,
		Usage: "set web socket port",
		Value: 7002,
	},
}
