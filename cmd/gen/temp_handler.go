package gen

const TempHandler = `
package handler

import (
	"encoding/json"
	"net/http"
	"sync"
	"TempImportPkg/adapter/driver/dto/request"
	"TempImportPkg/adapter/driver/router"
	portDriver "TempImportPkg/port/driver"
	"TempImportPkg/domain/service"

	"github.com/gin-gonic/gin"
)

type TempSvcNameCamelLowerHttpHandler struct {
	TempSvcNameCamelLowerService portDriver.TempSvcNameCaseCamelService
}

var (
	httpTempSvcNameCaseCamelOnce sync.Once
	httpTempSvcNameCaseCamelHand router.CommonRouter
)

func NewHttpTempSvcNameCaseCamelHandler() router.CommonRouter {
	httpTempSvcNameCaseCamelOnce.Do(func() {
		httpTempSvcNameCaseCamelHand = &TempSvcNameCamelLowerHttpHandler{
			TempSvcNameCamelLowerService: service.NewTempSvcNameCaseCamelService(),
		}
	})
	return httpTempSvcNameCaseCamelHand
}

func (h *TempSvcNameCamelLowerHttpHandler) InitRouter(router *gin.RouterGroup) {
	router.GET("/TempSvcNameCaseSnake/:id", h.findTempSvcNameCaseCamelById)
	router.GET("/TempSvcNameCaseSnake", h.findTempSvcNameCaseCamelList)
	router.POST("/TempSvcNameCaseSnake", h.createTempSvcNameCaseCamel)
	router.PUT("/TempSvcNameCaseSnake/:id", h.updateTempSvcNameCaseCamel)
	router.DELETE("/TempSvcNameCaseSnake/:id", h.delTempSvcNameCaseCamel)
}

func (h *TempSvcNameCamelLowerHttpHandler) findTempSvcNameCaseCamelById(c *gin.Context) {
	var req request.FindTempSvcNameCaseCamelByIdReq

	if err := c.ShouldBindUri(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.TempSvcNameCamelLowerService.FindTempSvcNameCaseCamelById(c, req.Id)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *TempSvcNameCamelLowerHttpHandler) findTempSvcNameCaseCamelList(c *gin.Context) {
	var reqForm request.GetTempSvcNameCaseCamelListReq

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
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total":   total,
		"entries": res,
	})
}

func (h *TempSvcNameCamelLowerHttpHandler) createTempSvcNameCaseCamel(c *gin.Context) {
	var req request.CreateTempSvcNameCaseCamelReq

	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id, err := h.TempSvcNameCamelLowerService.CreateTempSvcNameCaseCamel(c, req)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (h *TempSvcNameCamelLowerHttpHandler) updateTempSvcNameCaseCamel(c *gin.Context) {
	var req request.UpdateTempSvcNameCaseCamelReq

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
	var req request.DelTempSvcNameCaseCamelReq

	if err := c.ShouldBindUri(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := h.TempSvcNameCamelLowerService.DelTempSvcNameCaseCamel(c, req.Id)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
`
