package commands

import (
	"errors"
	"github.com/dipperin/dipperin-core/third-party/log"
	"github.com/dipperin/dipperin-core/third-party/rpc"
	"github.com/urfave/cli"
	"reflect"
	"strconv"
	"strings"
)

var (
	Client *rpc.Client
	CliLog log.Logger
)

type rpcCaller struct{}

var callerRv = reflect.ValueOf(&rpcCaller{})

func RpcCall(c *cli.Context) {
	if Client == nil {
		panic("rpc Client not initialized")
	}

	CliLog.Info("the args is:", "args", c.Args())
	method := c.String("m")
	if method == "" {
		CliLog.Error("Please specify -m")
		return
	}

	rvf := callerRv.MethodByName(method)
	if rvf.Kind() != reflect.Func {
		CliLog.Error("not found method", "method_name", method)
		return
	}

	// call method
	rvf.Call([]reflect.Value{reflect.ValueOf(c)})
}

// get rpc method from name
func getServerRpcMethodByName(mName string) string {
	lm := strings.ToLower(string(mName[0])) + mName[1:]
	return "cgoProServer_" + lm
}

// get rpc parameters from map, return string list, delete space at end
func getRpcParamFromString(cParam string) []string {
	if cParam == "" {
		return []string{}
	}

	lm := strings.Split(cParam, ",")

	log.Info("the lm is:", "lm", lm)
	return lm
}

func getRpcMethodAndParam(c *cli.Context) (mName string, cParams []string, err error) {
	mName = c.String("m")
	if mName == "" {
		return "", []string{}, errors.New("the method name is nil")
	}
	params := c.String("p")

	log.Info("the params is:", "params", params)

	cParams = getRpcParamFromString(params)
	return mName, cParams, nil
}

func callNoParamFunc(c *cli.Context) {
	mName, params, err := getRpcMethodAndParam(c)
	if err != nil {
		log.Error("getRpcMethodAndParam error")
		return
	}

	if len(params) !=0 {
		log.Error("there should be no parameters")
		return
	}

	var resp interface{}
	if err := Client.Call(&resp, getServerRpcMethodByName(mName)); err != nil {
		log.Error("call current balance error", "err", err)
		return
	}
}

func (caller *rpcCaller) CgoTestF5(c *cli.Context) {
	callNoParamFunc(c)
}

func (caller *rpcCaller) CgoTestCallOtherFunctions(c *cli.Context) {
	callNoParamFunc(c)
}

func (caller *rpcCaller) CgoTestAddFunc(c *cli.Context) {
	mName, params, err := getRpcMethodAndParam(c)
	if err != nil {
		log.Error("getRpcMethodAndParam error")
		return
	}

	if len(params) != 2{
		log.Error("the AddFunc parameter number error")
		return
	}
	a ,err:= strconv.Atoi(params[0])
	if err !=nil{
		log.Error("the AddFunc parameter a error")
		return
	}

	b ,err:= strconv.Atoi(params[1])
	if err !=nil{
		log.Error("the AddFunc parameter b error")
		return
	}

	var result int
	if err := Client.Call(&result, getServerRpcMethodByName(mName),a,b); err != nil {
		log.Error("call current balance error", "err", err)
		return
	}

	log.Info("the call AddFunc result is:","result",result)
}


func (caller *rpcCaller) CgoTestCxxSum(c *cli.Context) {
	mName, params, err := getRpcMethodAndParam(c)
	if err != nil {
		log.Error("getRpcMethodAndParam error")
		return
	}

	//first params is dataLen
	dataLen ,err:= strconv.Atoi(params[0])
	if err !=nil{
		log.Error("the CgoTestCxxSum parameter dataLen error")
		return
	}

	if len(params) != (dataLen+1){
		log.Error("the CgoTestCxxSum parameter data error")
	}

	slice := make([]int32,0)
	for i:=0;i<dataLen;i++{
		data ,err:= strconv.Atoi(params[i+1])
		if err !=nil{
			log.Error("the CgoTestCxxSum parameter b error")
			return
		}
		slice = append(slice,int32(data))
	}

	log.Info("the slice is:","slice",slice)

	var result int
	if err := Client.Call(&result, getServerRpcMethodByName(mName),slice,dataLen); err != nil {
		log.Error("call CgoTestCxxSum error", "err", err)
		return
	}

	log.Info("the call CgoTestCxxSum result is:","result",result)
}
