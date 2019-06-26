package service

import (
	"fmt"
	"github.com/dipperin/dipperin-core/third-party/log"
	"github.com/dipperin/dipperin-core/third-party/rpc"
	wrapper2 "github.com/yangchuang123456/cgo-pro/server/C++lib/wrapper"
	"github.com/yangchuang123456/cgo-pro/server/clib/wrapper"
)

type EndPoints struct {
	HTTPHost string `toml:",omitempty"`
	WSHost string `toml:",omitempty"`
	HttpEndpoint int
	WsEndpoint   int
	AllowHosts   []string
}

func (conf EndPoints) HttpEndPoint() string {
	if conf.HTTPHost == "" {
		return ""
	}
	return fmt.Sprintf("%s:%d", conf.HTTPHost, conf.HttpEndpoint)
}
func (conf EndPoints) WsEndPoint() string {
	if conf.WSHost == "" {
		return ""
	}
	return fmt.Sprintf("%s:%d", conf.WSHost, conf.WsEndpoint)
}


func makeRpcService(conf *EndPoints, apis []rpc.API) *RpcService {
	return &RpcService{
		httpEndpoint: conf.HttpEndPoint(),
		wsEndpoint:   conf.WsEndPoint(),
		apis:         apis,
		allowHosts:   conf.AllowHosts,
	}
}

func InitRpc(conf *EndPoints) {
	service := &CgoService{}
	rpcApi := MakeCgoProApi(service)

	rpcService := makeRpcService(conf,
		[]rpc.API{
			{
				Namespace: "cgoProServer",
				Version:   "0.0.0",
				Service:   rpcApi,
				Public:    true,
			},
		},
	)

	go rpcService.Start()
}


type CgoService struct {
}

func (service *CgoService) CgoTestF5() {
	log.Info("call wrapper CallF5WithF start")
	wrapper.CallF5WithF()
	log.Info("call wrapper CallF5WithF end")
	return
}

func (service *CgoService) CgoTestCallOtherFunctions(){
	log.Info("call wrapper CgoTestCallOtherFunctions start")
	wrapper.CallOtherFunctions()
	log.Info("call wrapper CgoTestCallOtherFunctions end")
}

func (service *CgoService) CgoTestAddFunc(a,b int) int{
	return wrapper.AddFunc(a,b)
}

func (service *CgoService) CgoTestCxxSum(data []int32,dataLen int) int{
	return wrapper2.GetVectorDataSum(data,dataLen)
}