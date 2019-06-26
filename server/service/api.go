package service

import "github.com/vntchain/go-vnt/log"

type CgoProApi struct {
	service *CgoService
}

func MakeCgoProApi(service *CgoService)*CgoProApi{
	return &CgoProApi{
		service:service,
	}
}

func (api *CgoProApi) CgoTestF5(){
	api.service.CgoTestF5()
}

func (api *CgoProApi) CgoTestCallOtherFunctions(){
	api.service.CgoTestCallOtherFunctions()
}

func (api *CgoProApi) CgoTestAddFunc(a,b int) int{
	return api.service.CgoTestAddFunc(a,b)
}

func (api *CgoProApi) CgoTestCxxSum(data []int32,dataLen int) int{
	log.Info("the data is:","data",data)
	return api.service.CgoTestCxxSum(data,dataLen)
}