package gen

const TempHandler = `
package driver

import (
	"encoding/json"
	"fmt"
	"TempImportPkg/adapter/driver"
	"TempImportPkg/adapter/driver/dto"
	"TempImportPkg/domain/service"
	portDriver "TempImportPkg/port/driver"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type TempSvcNameCamelLowerHttpHandler struct {
	TempSvcNameCamelLowerService portDriver.TempSvcNameCaseCamelService
}

var (
	httpTempSvcNameCaseCamelOnce sync.Once
	httpTempSvcNameCaseCamelHand driver.HttpRouterInterface
)

func NewHttpTempSvcNameCaseCamelHandler() driver.HttpRouterInterface {
	httpTempSvcNameCaseCamelOnce.Do(func() {
		httpTempSvcNameCaseCamelHand = &TempSvcNameCamelLowerHttpHandler{
			TempSvcNameCamelLowerService: service.NewTempSvcNameCaseCamelService(),
		}
	})
	return httpTempSvcNameCaseCamelHand
}

// RegisterRouterPublic 注册外部API
func (h *TempSvcNameCamelLowerHttpHandler) RegisterRouterPublic(router *gin.RouterGroup) {
	router.GET("/TempSvcNameCaseSnake/:id", h.findTempSvcNameCaseCamelById) // 查询TempSvcNameCaseCamelById
	router.GET("/TempSvcNameCaseSnake", h.findTempSvcNameCaseCamelList)       // 查询TempSvcNameCaseCamel列表
	router.POST("/TempSvcNameCaseSnake", h.createTempSvcNameCaseCamel)        // 创建TempSvcNameCaseCamel
	router.PUT("/TempSvcNameCaseSnake/:id", h.updateTempSvcNameCaseCamel)     // 修改TempSvcNameCaseCamel
	router.DELETE("/TempSvcNameCaseSnake/:id", h.delTempSvcNameCaseCamel)     // 删除TempSvcNameCaseCamel
}

// RegisterRouterPrivate 注册内部API
func (h *TempSvcNameCamelLowerHttpHandler) RegisterRouterPrivate(router *gin.RouterGroup) {
}

func (h *TempSvcNameCamelLowerHttpHandler) findTempSvcNameCaseCamelById(c *gin.Context) {

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

func (h *TempSvcNameCamelLowerHttpHandler) findTempSvcNameCaseCamelList(c *gin.Context) {

	var reqForm dto.GetTempSvcNameCaseCamelListReq

	if err := c.ShouldBindQuery(&reqForm); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
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

func (h *TempSvcNameCamelLowerHttpHandler) createTempSvcNameCaseCamel(c *gin.Context) {

	var req dto.CreateTempSvcNameCaseCamelReq

	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id, err := h.TempSvcNameCamelLowerService.CreateTempSvcNameCaseCamel(c, req)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Header("Location", fmt.Sprintf("%s/%d", c.FullPath(), id))
	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (h *TempSvcNameCamelLowerHttpHandler) updateTempSvcNameCaseCamel(c *gin.Context) {

	var req dto.UpdateTempSvcNameCaseCamelReq

	if err := c.ShouldBindUri(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := h.TempSvcNameCamelLowerService.UpdateTempSvcNameCaseCamel(c, req.Id, req)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *TempSvcNameCamelLowerHttpHandler) delTempSvcNameCaseCamel(c *gin.Context) {

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
