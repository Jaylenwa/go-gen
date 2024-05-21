package gen

const TempService = `
package service

import (
	"context"
	adapterDriven "TempImportPkg/adapter/driven"
	"TempImportPkg/adapter/driver/dto/request"
	"TempImportPkg/adapter/driver/dto/response"
	"TempImportPkg/infra/po"
	portDriven "TempImportPkg/port/driven"
	portDriver "TempImportPkg/port/driver"
	"sync"

	"github.com/jinzhu/copier"
)

var (
	TempSvcNameCamelLowerServiceOnce sync.Once
	TempSvcNameCamelLowerServiceImpl portDriver.TempSvcNameCaseCamelService
)

type TempSvcNameCamelLowerService struct {
	TempSvcNameCamelLowerRepo portDriven.TempSvcNameCaseCamelRepo
}

var _ portDriver.TempSvcNameCaseCamelService = &TempSvcNameCamelLowerService{}

func NewTempSvcNameCaseCamelService() portDriver.TempSvcNameCaseCamelService {
	TempSvcNameCamelLowerServiceOnce.Do(func() {
		TempSvcNameCamelLowerServiceImpl = &TempSvcNameCamelLowerService{
			TempSvcNameCamelLowerRepo: adapterDriven.NewTempSvcNameCaseCamelRepo(),
		}
	})
	return TempSvcNameCamelLowerServiceImpl

}

func (svc *TempSvcNameCamelLowerService) FindTempSvcNameCaseCamelById(ctx context.Context, id int64) (res response.GetTempSvcNameCaseCamelByIdRsp, err error) {
	TempSvcNameCamelLowerPo, err := svc.TempSvcNameCamelLowerRepo.FindById(ctx, id)
	if err != nil {
		return
	}

	// PO_to_DO
	err = copier.Copy(&res, TempSvcNameCamelLowerPo)

	return
}

func (svc *TempSvcNameCamelLowerService) FindTempSvcNameCaseCamelList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []response.GetTempSvcNameCaseCamelListRsp, err error) {
	total, TempSvcNameCamelLowerList, err := svc.TempSvcNameCamelLowerRepo.FindList(ctx, filter, args...)

	// POs_to_DOs
	res = make([]response.GetTempSvcNameCaseCamelListRsp, 0)

	for _, TempSvcNameCamelLower := range TempSvcNameCamelLowerList {
		do := response.GetTempSvcNameCaseCamelListRsp{}

		err = copier.Copy(&do, TempSvcNameCamelLower)
		if err != nil {
			return
		}

		res = append(res, do)
	}

	return
}

func (svc *TempSvcNameCamelLowerService) CreateTempSvcNameCaseCamel(ctx context.Context, req request.CreateTempSvcNameCaseCamelReq) (id int64, err error) {
	var (
		TempSvcNameCamelLowerPo po.TempSvcNameCaseCamel
	)

	// DO_to_PO
	err = copier.Copy(&TempSvcNameCamelLowerPo, req)
	if err != nil {
		return
	}

	id, err = svc.TempSvcNameCamelLowerRepo.Insert(ctx, TempSvcNameCamelLowerPo)

	return
}

func (svc *TempSvcNameCamelLowerService) UpdateTempSvcNameCaseCamel(ctx context.Context, id int64, req request.UpdateTempSvcNameCaseCamelReq) (err error) {
	var (
		TempSvcNameCamelLowerPo po.TempSvcNameCaseCamel
	)

	// DO_to_PO
	err = copier.Copy(&TempSvcNameCamelLowerPo, req)
	if err != nil {
		return
	}

	err = svc.TempSvcNameCamelLowerRepo.Update(ctx, id, TempSvcNameCamelLowerPo)

	return
}

func (svc *TempSvcNameCamelLowerService) DelTempSvcNameCaseCamel(ctx context.Context, id int64) (err error) {
	err = svc.TempSvcNameCamelLowerRepo.Delete(ctx, id)

	return
}
`
