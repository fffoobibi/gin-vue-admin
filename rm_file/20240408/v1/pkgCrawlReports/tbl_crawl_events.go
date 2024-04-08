package pkgCrawlReports

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/pkgCrawlReports"
    pkgCrawlReportsReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgCrawlReports/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type TblCrawlEventsApi struct {
}

var eService = service.ServiceGroupApp.PkgCrawlReportsServiceGroup.TblCrawlEventsService


// CreateTblCrawlEvents 创建时间表
// @Tags TblCrawlEvents
// @Summary 创建时间表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgCrawlReports.TblCrawlEvents true "创建时间表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /e/createTblCrawlEvents [post]
func (eApi *TblCrawlEventsApi) CreateTblCrawlEvents(c *gin.Context) {
	var e pkgCrawlReports.TblCrawlEvents
	err := c.ShouldBindJSON(&e)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := eService.CreateTblCrawlEvents(&e); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTblCrawlEvents 删除时间表
// @Tags TblCrawlEvents
// @Summary 删除时间表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgCrawlReports.TblCrawlEvents true "删除时间表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /e/deleteTblCrawlEvents [delete]
func (eApi *TblCrawlEventsApi) DeleteTblCrawlEvents(c *gin.Context) {
	ID := c.Query("ID")
	if err := eService.DeleteTblCrawlEvents(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTblCrawlEventsByIds 批量删除时间表
// @Tags TblCrawlEvents
// @Summary 批量删除时间表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /e/deleteTblCrawlEventsByIds [delete]
func (eApi *TblCrawlEventsApi) DeleteTblCrawlEventsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := eService.DeleteTblCrawlEventsByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTblCrawlEvents 更新时间表
// @Tags TblCrawlEvents
// @Summary 更新时间表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgCrawlReports.TblCrawlEvents true "更新时间表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /e/updateTblCrawlEvents [put]
func (eApi *TblCrawlEventsApi) UpdateTblCrawlEvents(c *gin.Context) {
	var e pkgCrawlReports.TblCrawlEvents
	err := c.ShouldBindJSON(&e)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := eService.UpdateTblCrawlEvents(e); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTblCrawlEvents 用id查询时间表
// @Tags TblCrawlEvents
// @Summary 用id查询时间表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgCrawlReports.TblCrawlEvents true "用id查询时间表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /e/findTblCrawlEvents [get]
func (eApi *TblCrawlEventsApi) FindTblCrawlEvents(c *gin.Context) {
	ID := c.Query("ID")
	if ree, err := eService.GetTblCrawlEvents(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"ree": ree}, c)
	}
}

// GetTblCrawlEventsList 分页获取时间表列表
// @Tags TblCrawlEvents
// @Summary 分页获取时间表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgCrawlReportsReq.TblCrawlEventsSearch true "分页获取时间表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /e/getTblCrawlEventsList [get]
func (eApi *TblCrawlEventsApi) GetTblCrawlEventsList(c *gin.Context) {
	var pageInfo pkgCrawlReportsReq.TblCrawlEventsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := eService.GetTblCrawlEventsInfoList(pageInfo); err != nil {
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
