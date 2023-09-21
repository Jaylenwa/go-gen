package gen

const TempHandler = `
package driver

import (
	"encoding/json"
	"fmt"
	"TempImportPkg/adapter/driver/dto"
	"TempImportPkg/domain/service"
	"TempImportPkg/infra/utils/validate"
	portDriver "TempImportPkg/port/driver"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type TempSvcNameCamelLowerHandler struct {
	TempSvcNameCamelLowerService portDriver.TempSvcNameCaseCamelService
}

var (
	TempSvcNameCamelLowerOnce sync.Once
	TempSvcNameCamelLowerHand RouterInterface
)

func NewTempSvcNameCaseCamelHandler() RouterInterface {
	TempSvcNameCamelLowerOnce.Do(func() {
		TempSvcNameCamelLowerHand = &TempSvcNameCamelLowerHandler{
			TempSvcNameCamelLowerService: service.NewTempSvcNameCaseCamelService(),
		}
	})
	return TempSvcNameCamelLowerHand
}

// RegisterRouterPublic 注册外部API
func (h *TempSvcNameCamelLowerHandler) RegisterRouterPublic(router *gin.RouterGroup) {
	router.GET("/TempSvcNameCaseSnake/:id", h.findTempSvcNameCaseCamelById) // 查询TempSvcNameCaseCamelById
	router.GET("/TempSvcNameCaseSnake", h.findTempSvcNameCaseCamelList)       // 查询TempSvcNameCaseCamel列表
	router.POST("/TempSvcNameCaseSnake", h.createTempSvcNameCaseCamel)        // 创建TempSvcNameCaseCamel
	router.PUT("/TempSvcNameCaseSnake/:id", h.updateTempSvcNameCaseCamel)     // 修改TempSvcNameCaseCamel
	router.DELETE("/TempSvcNameCaseSnake/:id", h.delTempSvcNameCaseCamel)     // 删除TempSvcNameCaseCamel
}

// RegisterRouterPrivate 注册内部API
func (h *TempSvcNameCamelLowerHandler) RegisterRouterPrivate(router *gin.RouterGroup) {
}

func (h *TempSvcNameCamelLowerHandler) findTempSvcNameCaseCamelById(c *gin.Context) {

	var req dto.FindTempSvcNameCaseCamelByIdReq

	if err := c.ShouldBindUri(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.TempSvcNameCamelLowerService.FindTempSvcNameCaseCamelById(c, req.Id)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h *TempSvcNameCamelLowerHandler) findTempSvcNameCaseCamelList(c *gin.Context) {

	var reqForm dto.GetTempSvcNameCaseCamelListReq

	if err := c.ShouldBindQuery(&reqForm); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := validate.Validate(reqForm)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var filter map[string]interface{}
	reqBytes, err := json.Marshal(&reqForm)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(reqBytes, &filter)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	total, res, err := h.TempSvcNameCamelLowerService.FindTempSvcNameCaseCamelList(c, filter)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total":   total,
		"entries": res,
	})

}

func (h *TempSvcNameCamelLowerHandler) createTempSvcNameCaseCamel(c *gin.Context) {

	var req dto.CreateTempSvcNameCaseCamelReq

	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := validate.Validate(req)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	id, err := h.TempSvcNameCamelLowerService.CreateTempSvcNameCaseCamel(c, &req)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Header("Location", fmt.Sprintf("%s/%d", c.FullPath(), id))
	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (h *TempSvcNameCamelLowerHandler) updateTempSvcNameCaseCamel(c *gin.Context) {

	var req dto.UpdateTempSvcNameCaseCamelReq

	if err := c.ShouldBindUri(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := validate.Validate(req)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = h.TempSvcNameCamelLowerService.UpdateTempSvcNameCaseCamel(c, req.Id, &req)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *TempSvcNameCamelLowerHandler) delTempSvcNameCaseCamel(c *gin.Context) {

	var req dto.DelTempSvcNameCaseCamelReq

	if err := c.ShouldBindUri(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := h.TempSvcNameCamelLowerService.DelTempSvcNameCaseCamel(c, req.Id)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Status(http.StatusNoContent)
}

`
