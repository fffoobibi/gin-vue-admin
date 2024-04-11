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

type TblKolResourceApi struct {
}

var tblKolResourceService = service.ServiceGroupApp.PkgCrawlReportsServiceGroup.TblKolResourceService

// CreateTblKolResource 创建tblKolResource表
// @Tags TblKolResource
// @Summary 创建tblKolResource表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgCrawlReports.TblKolResource true "创建tblKolResource表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tblKolResource/createTblKolResource [post]
func (tblKolResourceApi *TblKolResourceApi) CreateTblKolResource(c *gin.Context) {
	var tblKolResource pkgCrawlReports.TblKolResource
	err := c.ShouldBindJSON(&tblKolResource)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := tblKolResourceService.CreateTblKolResource(&tblKolResource); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTblKolResource 删除tblKolResource表
// @Tags TblKolResource
// @Summary 删除tblKolResource表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgCrawlReports.TblKolResource true "删除tblKolResource表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tblKolResource/deleteTblKolResource [delete]
func (tblKolResourceApi *TblKolResourceApi) DeleteTblKolResource(c *gin.Context) {
	id := c.Query("id")
	if err := tblKolResourceService.DeleteTblKolResource(id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTblKolResourceByIds 批量删除tblKolResource表
// @Tags TblKolResource
// @Summary 批量删除tblKolResource表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tblKolResource/deleteTblKolResourceByIds [delete]
func (tblKolResourceApi *TblKolResourceApi) DeleteTblKolResourceByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := tblKolResourceService.DeleteTblKolResourceByIds(ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTblKolResource 更新tblKolResource表
// @Tags TblKolResource
// @Summary 更新tblKolResource表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgCrawlReports.TblKolResource true "更新tblKolResource表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tblKolResource/updateTblKolResource [put]
func (tblKolResourceApi *TblKolResourceApi) UpdateTblKolResource(c *gin.Context) {
	var tblKolResource pkgCrawlReports.TblKolResource
	err := c.ShouldBindJSON(&tblKolResource)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := tblKolResourceService.UpdateTblKolResource(tblKolResource); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTblKolResource 用id查询tblKolResource表
// @Tags TblKolResource
// @Summary 用id查询tblKolResource表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgCrawlReports.TblKolResource true "用id查询tblKolResource表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tblKolResource/findTblKolResource [get]
func (tblKolResourceApi *TblKolResourceApi) FindTblKolResource(c *gin.Context) {
	id := c.Query("id")
	if retblKolResource, err := tblKolResourceService.GetTblKolResource(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retblKolResource": retblKolResource}, c)
	}
}

// GetTblKolResourceList 分页获取tblKolResource表列表
// @Tags TblKolResource
// @Summary 分页获取tblKolResource表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgCrawlReportsReq.TblKolResourceSearch true "分页获取tblKolResource表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tblKolResource/getTblKolResourceList [get]
func (tblKolResourceApi *TblKolResourceApi) GetTblKolResourceList(c *gin.Context) {
	var pageInfo pkgCrawlReportsReq.TblKolResourceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := tblKolResourceService.GetTblKolResourceInfoList(pageInfo); err != nil {
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

// GetKolResourceCrawlInfo 获取邮箱资源抓取统计数量
// @Router /tblCrawlStats/getKolResourceCrawlInfo [get]
func (tblKolResourceApi *TblKolResourceApi) GetKolResourceCrawlInfo(c *gin.Context) {
	var query pkgCrawlReportsReq.TblKolResourceSearchQuery
	err := c.ShouldBindQuery(&query)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, err := tblKolResourceService.GetKolResourceCrawlInfo(query); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}
