package gen

const TempService = `
package service

import (
	"context"
	"TempImportPkg/adapter/driven"
	"TempImportPkg/adapter/driver/dto"
	"TempImportPkg/infra/po"
	"TempImportPkg/infra/utils/query"
	"TempImportPkg/infra/utils/struct"
	portDriven "TempImportPkg/port/driven"
	portDriver "TempImportPkg/port/driver"
	"sync"
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

func (svc *TempSvcNameCamelLowerService) FindTempSvcNameCaseCamelById(ctx context.Context, id int64) (res dto.GetTempSvcNameCaseCamelByIdRsp, err error) {
	TempSvcNameCamelLowerPo, err := svc.TempSvcNameCamelLowerRepo.FindById(ctx, id)
	if err != nil {
		return
	}

	// PO_to_DO
	err = _struct.CopyStruct(&res, TempSvcNameCamelLowerPo)

	return
}

func (svc *TempSvcNameCamelLowerService) FindTempSvcNameCaseCamelByQuery(ctx context.Context, queries []*query.Query) (res dto.GetTempSvcNameCaseCamelByQueryRsp, err error) {
	TempSvcNameCamelLowerPo, err := svc.TempSvcNameCamelLowerRepo.FindByQuery(ctx, queries)

	// PO_to_DO
	err = _struct.CopyStruct(&res, TempSvcNameCamelLowerPo)

	return
}

func (svc *TempSvcNameCamelLowerService) FindTempSvcNameCaseCamelList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []dto.GetTempSvcNameCaseCamelListRsp, err error) {
	total, TempSvcNameCamelLowerList, err := svc.TempSvcNameCamelLowerRepo.FindList(ctx, filter, args...)

	// POs_to_DOs
	res = make([]dto.GetTempSvcNameCaseCamelListRsp, 0)

	for _, TempSvcNameCamelLower := range TempSvcNameCamelLowerList {
		do := dto.GetTempSvcNameCaseCamelListRsp{}

		err = _struct.CopyStruct(&do, TempSvcNameCamelLower)
		if err != nil {
			return
		}

		res = append(res, do)
	}

	return
}

func (svc *TempSvcNameCamelLowerService) CreateTempSvcNameCaseCamel(ctx context.Context, req dto.CreateTempSvcNameCaseCamelReq) (id int64, err error) {
	var (
		TempSvcNameCamelLowerPo po.TempSvcNameCaseCamel
	)

	// DO_to_PO
	err = _struct.CopyStruct(&TempSvcNameCamelLowerPo, req)
	if err != nil {
		return
	}

	id, err = svc.TempSvcNameCamelLowerRepo.Insert(ctx, TempSvcNameCamelLowerPo)

	return
}

func (svc *TempSvcNameCamelLowerService) UpdateTempSvcNameCaseCamel(ctx context.Context, id int64, req dto.UpdateTempSvcNameCaseCamelReq) (err error) {
	var (
		TempSvcNameCamelLowerPo po.TempSvcNameCaseCamel
	)

	// DO_to_PO
	err = _struct.CopyStruct(&TempSvcNameCamelLowerPo, req)
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
