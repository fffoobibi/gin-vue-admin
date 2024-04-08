package pkgCrawlReports

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TblCrawlEventsRouter struct {
}

// InitTblCrawlEventsRouter 初始化 爬虫大事迹 路由信息
func (s *TblCrawlEventsRouter) InitTblCrawlEventsRouter(Router *gin.RouterGroup) {
	eRouter := Router.Group("e").Use(middleware.OperationRecord())
	eRouterWithoutRecord := Router.Group("e")
	var eApi = v1.ApiGroupApp.PkgCrawlReportsApiGroup.TblCrawlEventsApi
	{
		eRouter.POST("createTblCrawlEvents", eApi.CreateTblCrawlEvents)             // 新建爬虫大事迹
		eRouter.DELETE("deleteTblCrawlEvents", eApi.DeleteTblCrawlEvents)           // 删除爬虫大事迹
		eRouter.DELETE("deleteTblCrawlEventsByIds", eApi.DeleteTblCrawlEventsByIds) // 批量删除爬虫大事迹
		eRouter.PUT("updateTblCrawlEvents", eApi.UpdateTblCrawlEvents)              // 更新爬虫大事迹
	}
	{
		eRouterWithoutRecord.GET("findTblCrawlEvents", eApi.FindTblCrawlEvents)       // 根据ID获取爬虫大事迹
		eRouterWithoutRecord.GET("getTblCrawlEventsList", eApi.GetTblCrawlEventsList) // 获取爬虫大事迹列表
	}
}
