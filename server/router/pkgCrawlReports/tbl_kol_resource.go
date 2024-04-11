package pkgCrawlReports

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TblKolResourceRouter struct {
}

// InitTblKolResourceRouter 初始化 tblKolResource表 路由信息
func (s *TblKolResourceRouter) InitTblKolResourceRouter(Router *gin.RouterGroup) {
	tblKolResourceRouter := Router.Group("tblKolResource").Use(middleware.OperationRecord())
	tblKolResourceRouterWithoutRecord := Router.Group("tblKolResource")
	var tblKolResourceApi = v1.ApiGroupApp.PkgCrawlReportsApiGroup.TblKolResourceApi
	{
		tblKolResourceRouter.POST("createTblKolResource", tblKolResourceApi.CreateTblKolResource)             // 新建tblKolResource表
		tblKolResourceRouter.DELETE("deleteTblKolResource", tblKolResourceApi.DeleteTblKolResource)           // 删除tblKolResource表
		tblKolResourceRouter.DELETE("deleteTblKolResourceByIds", tblKolResourceApi.DeleteTblKolResourceByIds) // 批量删除tblKolResource表
		tblKolResourceRouter.PUT("updateTblKolResource", tblKolResourceApi.UpdateTblKolResource)              // 更新tblKolResource表
	}
	{
		tblKolResourceRouterWithoutRecord.GET("findTblKolResource", tblKolResourceApi.FindTblKolResource)       // 根据ID获取tblKolResource表
		tblKolResourceRouterWithoutRecord.GET("getTblKolResourceList", tblKolResourceApi.GetTblKolResourceList) // 获取tblKolResource表列表
		tblKolResourceRouter.GET("getKolResourceCrawlInfo", tblKolResourceApi.GetKolResourceCrawlInfo)          // 获取初次抓取数量
	}
}
