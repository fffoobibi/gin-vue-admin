package pkgCrawlReports

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StudentRouter struct {
}

// InitStudentRouter 初始化 测试表格 路由信息
func (s *StudentRouter) InitStudentRouter(Router *gin.RouterGroup) {
	GvaStudentRouter := Router.Group("GvaStudent").Use(middleware.OperationRecord())
	GvaStudentRouterWithoutRecord := Router.Group("GvaStudent")
	var GvaStudentApi = v1.ApiGroupApp.PkgCrawlReportsApiGroup.StudentApi
	{
		GvaStudentRouter.POST("createStudent", GvaStudentApi.CreateStudent)   // 新建测试表格
		GvaStudentRouter.DELETE("deleteStudent", GvaStudentApi.DeleteStudent) // 删除测试表格
		GvaStudentRouter.DELETE("deleteStudentByIds", GvaStudentApi.DeleteStudentByIds) // 批量删除测试表格
		GvaStudentRouter.PUT("updateStudent", GvaStudentApi.UpdateStudent)    // 更新测试表格
	}
	{
		GvaStudentRouterWithoutRecord.GET("findStudent", GvaStudentApi.FindStudent)        // 根据ID获取测试表格
		GvaStudentRouterWithoutRecord.GET("getStudentList", GvaStudentApi.GetStudentList)  // 获取测试表格列表
	}
}
