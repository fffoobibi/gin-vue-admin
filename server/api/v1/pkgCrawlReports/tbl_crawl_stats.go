package pkgCrawlReports

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgCrawlReports"
	pkgCrawlReportsReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgCrawlReports/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TblCrawlStatsApi struct {
}

var tblCrawlStatsService = service.ServiceGroupApp.PkgCrawlReportsServiceGroup.TblCrawlStatsService

// CreateTblCrawlStats 创建tblCrawlStats表
// @Tags TblCrawlStats
// @Summary 创建tblCrawlStats表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgCrawlReports.TblCrawlStats true "创建tblCrawlStats表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tblCrawlStats/createTblCrawlStats [post]
func (tblCrawlStatsApi *TblCrawlStatsApi) CreateTblCrawlStats(c *gin.Context) {
	var tblCrawlStats pkgCrawlReports.TblCrawlStats
	err := c.ShouldBindJSON(&tblCrawlStats)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := tblCrawlStatsService.CreateTblCrawlStats(&tblCrawlStats); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTblCrawlStats 删除tblCrawlStats表
// @Tags TblCrawlStats
// @Summary 删除tblCrawlStats表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgCrawlReports.TblCrawlStats true "删除tblCrawlStats表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tblCrawlStats/deleteTblCrawlStats [delete]
func (tblCrawlStatsApi *TblCrawlStatsApi) DeleteTblCrawlStats(c *gin.Context) {
	ID := c.Query("ID")
	if err := tblCrawlStatsService.DeleteTblCrawlStats(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTblCrawlStatsByIds 批量删除tblCrawlStats表
// @Tags TblCrawlStats
// @Summary 批量删除tblCrawlStats表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tblCrawlStats/deleteTblCrawlStatsByIds [delete]
func (tblCrawlStatsApi *TblCrawlStatsApi) DeleteTblCrawlStatsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := tblCrawlStatsService.DeleteTblCrawlStatsByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTblCrawlStats 更新tblCrawlStats表
// @Tags TblCrawlStats
// @Summary 更新tblCrawlStats表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgCrawlReports.TblCrawlStats true "更新tblCrawlStats表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tblCrawlStats/updateTblCrawlStats [put]
func (tblCrawlStatsApi *TblCrawlStatsApi) UpdateTblCrawlStats(c *gin.Context) {
	var tblCrawlStats pkgCrawlReports.TblCrawlStats
	err := c.ShouldBindJSON(&tblCrawlStats)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := tblCrawlStatsService.UpdateTblCrawlStats(tblCrawlStats); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTblCrawlStats 用id查询tblCrawlStats表
// @Tags TblCrawlStats
// @Summary 用id查询tblCrawlStats表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgCrawlReports.TblCrawlStats true "用id查询tblCrawlStats表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tblCrawlStats/findTblCrawlStats [get]
func (tblCrawlStatsApi *TblCrawlStatsApi) FindTblCrawlStats(c *gin.Context) {
	ID := c.Query("ID")
	if retblCrawlStats, err := tblCrawlStatsService.GetTblCrawlStats(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retblCrawlStats": retblCrawlStats}, c)
	}
}

// GetTblCrawlStatsList 分页获取tblCrawlStats表列表
// @Tags TblCrawlStats
// @Summary 分页获取tblCrawlStats表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgCrawlReportsReq.TblCrawlStatsSearch true "分页获取tblCrawlStats表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tblCrawlStats/getTblCrawlStatsList [get]
func (tblCrawlStatsApi *TblCrawlStatsApi) GetTblCrawlStatsList(c *gin.Context) {
	var pageInfo pkgCrawlReportsReq.TblCrawlStatsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := tblCrawlStatsService.GetTblCrawlStatsInfoList(pageInfo); err != nil {
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

// GetFirstCrawlReportsList 获取初次抓取统计数量
// @Router /tblCrawlStats/GetFirstCrawlReportsList [get]
func (tblCrawlStatsApi *TblCrawlStatsApi) GetFirstCrawlReportsList(c *gin.Context) {
	var query pkgCrawlReportsReq.TblCrawlStatsSearchQuery
	err := c.ShouldBindQuery(&query)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, err := tblCrawlStatsService.GetFirstCrawlCount(query); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

// GetTotalResourceInfo 获取初次抓取统计数量
// @Router /tblCrawlStats/getTotalResourceInfo [get]
func (tblCrawlStatsApi *TblCrawlStatsApi) GetTotalResourceInfo(c *gin.Context) {

	// mock
	data := gin.H{
		"total":     354710,
		"Tiktok":    177800,
		"Youtube":   84123,
		"Instagram": 22344,
	}
	response.OkWithDetailed(data, "获取成功", c)
}

// GetCrawlStatsPieData 获取饼图数据
// @Router /tblCrawlStats/getCrawlStatsPieData [get]
func (tblCrawlStatsApi *TblCrawlStatsApi) GetCrawlStatsPieData(c *gin.Context) {
	var query pkgCrawlReportsReq.TblCrawlStatsPieDataQuery
	err := c.ShouldBindQuery(&query)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var result struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data *struct {
			Platform []struct {
				Platform byte `json:"platform"`
				Count    int  `json:"count"`
			} `json:"platform"`
			Category []*struct {
				Name  string `json:"name"`
				Count int    `json:"count"`
			} `json:"category"`
			Country []struct {
				CountryName string `json:"country_name"`
				Count       int    `json:"count"`
			} `json:"country"`
		}
	}
	// 接口调用
	var values []byte
	values, _ = utils.Get("https://gmen.mediamz.com/Api/crawlCountryPlatformCategoryPieData", map[string]string{
		"st_time": query.StartTime + " 00:00:00",
		"ed_time": query.EndTime + " 23:59:59",
	})
	global.GVA_LOG.Info(string(values))
	if err := json.Unmarshal(values, &result); err != nil {
		global.GVA_LOG.Error("error in marshal:", zap.Error(err))
	} else {
		var categories []*struct {
			Name  string `json:"name"`
			Count int    `json:"count"`
		}
		var ms = map[string]*struct {
			Name  string `json:"name"`
			Count int    `json:"count"`
		}{}
		var categoriesMap = map[string]int{}

		for index, category := range result.Data.Category {
			if _, ok := ms[category.Name]; !ok {
				ms[category.Name] = category
				categoriesMap[category.Name] = index
				categories = append(categories, category)
			} else {
				lookupIndex := categoriesMap[category.Name]
				categories[lookupIndex].Count += category.Count
			}
		}
		result.Data.Category = categories
		response.OkWithDetailed(result.Data, "success", c)
	}
	//c.Header("Content-Type", "application/json;charset=utf-8")
	//c.String(http.StatusOK, string(values))
}

// GetSummaryCrawlInfo 获取所有统计数据
func (tblCrawlStatsApi *TblCrawlStatsApi) GetSummaryCrawlInfo(c *gin.Context) {
	var query pkgCrawlReportsReq.TblCrawlStatsTimeSearchQuery
	err := c.ShouldBindQuery(&query)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, err := tblCrawlStatsService.GetSummaryCrawlInfo(query); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

// GetCleanReportsList 获取清洗数据
func (tblCrawlStatsApi *TblCrawlStatsApi) GetCleanReportsList(c *gin.Context) {
	var query pkgCrawlReportsReq.TblCrawlStatsTimeSearchQuery
	err := c.ShouldBindQuery(&query)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, err := tblCrawlStatsService.GetCleanReportsList(query); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

// GetTotalResourceReportsList 获取总资源更新次数
func (tblCrawlStatsApi *TblCrawlStatsApi) GetTotalResourceReportsList(c *gin.Context) {
	var query pkgCrawlReportsReq.TblCrawlStatsTimeSearchQuery
	err := c.ShouldBindQuery(&query)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, err := tblCrawlStatsService.GetTotalResourceReportsList(query); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}
