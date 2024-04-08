package pkgCrawlReports

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TblCrawlEventsRouter struct {
}

// InitTblCrawlEventsRouter 初始化 时间表 路由信息
func (s *TblCrawlEventsRouter) InitTblCrawlEventsRouter(Router *gin.RouterGroup) {
	eRouter := Router.Group("e").Use(middleware.OperationRecord())
	eRouterWithoutRecord := Router.Group("e")
	var eApi = v1.ApiGroupApp.PkgCrawlReportsApiGroup.TblCrawlEventsApi
	{
		eRouter.POST("createTblCrawlEvents", eApi.CreateTblCrawlEvents)   // 新建时间表
		eRouter.DELETE("deleteTblCrawlEvents", eApi.DeleteTblCrawlEvents) // 删除时间表
		eRouter.DELETE("deleteTblCrawlEventsByIds", eApi.DeleteTblCrawlEventsByIds) // 批量删除时间表
		eRouter.PUT("updateTblCrawlEvents", eApi.UpdateTblCrawlEvents)    // 更新时间表
	}
	{
		eRouterWithoutRecord.GET("findTblCrawlEvents", eApi.FindTblCrawlEvents)        // 根据ID获取时间表
		eRouterWithoutRecord.GET("getTblCrawlEventsList", eApi.GetTblCrawlEventsList)  // 获取时间表列表
	}
}
