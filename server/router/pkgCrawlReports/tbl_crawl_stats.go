package pkgCrawlReports

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TblCrawlStatsRouter struct {
}

// InitTblCrawlStatsRouter 初始化 tblCrawlStats表 路由信息
func (s *TblCrawlStatsRouter) InitTblCrawlStatsRouter(Router *gin.RouterGroup) {
	tblCrawlStatsRouter := Router.Group("tblCrawlStats").Use(middleware.OperationRecord())
	tblCrawlStatsRouterWithoutRecord := Router.Group("tblCrawlStats")
	var tblCrawlStatsApi = v1.ApiGroupApp.PkgCrawlReportsApiGroup.TblCrawlStatsApi
	{
		tblCrawlStatsRouter.POST("createTblCrawlStats", tblCrawlStatsApi.CreateTblCrawlStats)             // 新建tblCrawlStats表
		tblCrawlStatsRouter.DELETE("deleteTblCrawlStats", tblCrawlStatsApi.DeleteTblCrawlStats)           // 删除tblCrawlStats表
		tblCrawlStatsRouter.DELETE("deleteTblCrawlStatsByIds", tblCrawlStatsApi.DeleteTblCrawlStatsByIds) // 批量删除tblCrawlStats表
		tblCrawlStatsRouter.PUT("updateTblCrawlStats", tblCrawlStatsApi.UpdateTblCrawlStats)              // 更新tblCrawlStats表
	}
	{
		tblCrawlStatsRouterWithoutRecord.GET("findTblCrawlStats", tblCrawlStatsApi.FindTblCrawlStats)       // 根据ID获取tblCrawlStats表
		tblCrawlStatsRouterWithoutRecord.GET("getTblCrawlStatsList", tblCrawlStatsApi.GetTblCrawlStatsList) // 获取tblCrawlStats表列表
		tblCrawlStatsRouterWithoutRecord.GET("getFirstCrawlInfo", tblCrawlStatsApi.GetFirstCrawlInfo)       // 获取tblCrawlStats表列表
		tblCrawlStatsRouterWithoutRecord.GET("getTotalResourceInfo", tblCrawlStatsApi.GetTotalResourceInfo) // 获取总数量, 饼图数据
		tblCrawlStatsRouterWithoutRecord.GET("getCrawlStatsPieData", tblCrawlStatsApi.GetCrawlStatsPieData) // 获取总数量, 饼图数据
		tblCrawlStatsRouterWithoutRecord.GET("getSummaryCrawlInfo", tblCrawlStatsApi.GetSummaryCrawlInfo)   // 获取初次，有效，清洗，更新次数统计
	}
}
