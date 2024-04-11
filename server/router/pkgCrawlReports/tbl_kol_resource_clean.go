package pkgCrawlReports

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TblKolResourceCleanRouter struct {
}

// InitTblKolResourceCleanRouter 初始化 tblKolResourceClean清洗表 路由信息
func (s *TblKolResourceCleanRouter) InitTblKolResourceCleanRouter(Router *gin.RouterGroup) {
	tblKolResourceCleanRouter := Router.Group("tblKolResourceClean").Use(middleware.OperationRecord())
	tblKolResourceCleanRouterWithoutRecord := Router.Group("tblKolResourceClean")
	var tblKolResourceCleanApi = v1.ApiGroupApp.PkgCrawlReportsApiGroup.TblKolResourceCleanApi
	{
		tblKolResourceCleanRouter.POST("createTblKolResourceClean", tblKolResourceCleanApi.CreateTblKolResourceClean)             // 新建tblKolResourceClean清洗表
		tblKolResourceCleanRouter.DELETE("deleteTblKolResourceClean", tblKolResourceCleanApi.DeleteTblKolResourceClean)           // 删除tblKolResourceClean清洗表
		tblKolResourceCleanRouter.DELETE("deleteTblKolResourceCleanByIds", tblKolResourceCleanApi.DeleteTblKolResourceCleanByIds) // 批量删除tblKolResourceClean清洗表
		tblKolResourceCleanRouter.PUT("updateTblKolResourceClean", tblKolResourceCleanApi.UpdateTblKolResourceClean)              // 更新tblKolResourceClean清洗表
	}
	{
		tblKolResourceCleanRouterWithoutRecord.GET("findTblKolResourceClean", tblKolResourceCleanApi.FindTblKolResourceClean)       // 根据ID获取tblKolResourceClean清洗表
		tblKolResourceCleanRouterWithoutRecord.GET("getTblKolResourceCleanList", tblKolResourceCleanApi.GetTblKolResourceCleanList) // 获取tblKolResourceClean清洗表列表
	}
}
