package pkgCrawlReports

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgCrawlReports"
	pkgCrawlReportsReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgCrawlReports/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TblKolResourceCleanApi struct {
}

var tblKolResourceCleanService = service.ServiceGroupApp.PkgCrawlReportsServiceGroup.TblKolResourceCleanService

// CreateTblKolResourceClean 创建tblKolResourceClean清洗表
// @Tags TblKolResourceClean
// @Summary 创建tblKolResourceClean清洗表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgCrawlReports.TblKolResourceClean true "创建tblKolResourceClean清洗表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tblKolResourceClean/createTblKolResourceClean [post]
func (tblKolResourceCleanApi *TblKolResourceCleanApi) CreateTblKolResourceClean(c *gin.Context) {
	var tblKolResourceClean pkgCrawlReports.TblKolResourceClean
	err := c.ShouldBindJSON(&tblKolResourceClean)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := tblKolResourceCleanService.CreateTblKolResourceClean(&tblKolResourceClean); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTblKolResourceClean 删除tblKolResourceClean清洗表
// @Tags TblKolResourceClean
// @Summary 删除tblKolResourceClean清洗表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgCrawlReports.TblKolResourceClean true "删除tblKolResourceClean清洗表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tblKolResourceClean/deleteTblKolResourceClean [delete]
func (tblKolResourceCleanApi *TblKolResourceCleanApi) DeleteTblKolResourceClean(c *gin.Context) {
	id := c.Query("id")
	if err := tblKolResourceCleanService.DeleteTblKolResourceClean(id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTblKolResourceCleanByIds 批量删除tblKolResourceClean清洗表
// @Tags TblKolResourceClean
// @Summary 批量删除tblKolResourceClean清洗表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tblKolResourceClean/deleteTblKolResourceCleanByIds [delete]
func (tblKolResourceCleanApi *TblKolResourceCleanApi) DeleteTblKolResourceCleanByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := tblKolResourceCleanService.DeleteTblKolResourceCleanByIds(ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTblKolResourceClean 更新tblKolResourceClean清洗表
// @Tags TblKolResourceClean
// @Summary 更新tblKolResourceClean清洗表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgCrawlReports.TblKolResourceClean true "更新tblKolResourceClean清洗表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tblKolResourceClean/updateTblKolResourceClean [put]
func (tblKolResourceCleanApi *TblKolResourceCleanApi) UpdateTblKolResourceClean(c *gin.Context) {
	var tblKolResourceClean pkgCrawlReports.TblKolResourceClean
	err := c.ShouldBindJSON(&tblKolResourceClean)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := tblKolResourceCleanService.UpdateTblKolResourceClean(tblKolResourceClean); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTblKolResourceClean 用id查询tblKolResourceClean清洗表
// @Tags TblKolResourceClean
// @Summary 用id查询tblKolResourceClean清洗表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgCrawlReports.TblKolResourceClean true "用id查询tblKolResourceClean清洗表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tblKolResourceClean/findTblKolResourceClean [get]
func (tblKolResourceCleanApi *TblKolResourceCleanApi) FindTblKolResourceClean(c *gin.Context) {
	id := c.Query("id")
	if retblKolResourceClean, err := tblKolResourceCleanService.GetTblKolResourceClean(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retblKolResourceClean": retblKolResourceClean}, c)
	}
}

// GetTblKolResourceCleanList 分页获取tblKolResourceClean清洗表列表
// @Tags TblKolResourceClean
// @Summary 分页获取tblKolResourceClean清洗表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgCrawlReportsReq.TblKolResourceCleanSearch true "分页获取tblKolResourceClean清洗表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tblKolResourceClean/getTblKolResourceCleanList [get]
func (tblKolResourceCleanApi *TblKolResourceCleanApi) GetTblKolResourceCleanList(c *gin.Context) {
	var pageInfo pkgCrawlReportsReq.TblKolResourceCleanSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := tblKolResourceCleanService.GetTblKolResourceCleanInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
