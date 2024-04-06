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

type StudentApi struct {
}

var GvaStudentService = service.ServiceGroupApp.PkgCrawlReportsServiceGroup.StudentService


// CreateStudent 创建测试表格
// @Tags Student
// @Summary 创建测试表格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgCrawlReports.Student true "创建测试表格"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /GvaStudent/createStudent [post]
func (GvaStudentApi *StudentApi) CreateStudent(c *gin.Context) {
	var GvaStudent pkgCrawlReports.Student
	err := c.ShouldBindJSON(&GvaStudent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := GvaStudentService.CreateStudent(&GvaStudent); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteStudent 删除测试表格
// @Tags Student
// @Summary 删除测试表格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgCrawlReports.Student true "删除测试表格"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /GvaStudent/deleteStudent [delete]
func (GvaStudentApi *StudentApi) DeleteStudent(c *gin.Context) {
	ID := c.Query("ID")
	if err := GvaStudentService.DeleteStudent(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteStudentByIds 批量删除测试表格
// @Tags Student
// @Summary 批量删除测试表格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /GvaStudent/deleteStudentByIds [delete]
func (GvaStudentApi *StudentApi) DeleteStudentByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := GvaStudentService.DeleteStudentByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateStudent 更新测试表格
// @Tags Student
// @Summary 更新测试表格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgCrawlReports.Student true "更新测试表格"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /GvaStudent/updateStudent [put]
func (GvaStudentApi *StudentApi) UpdateStudent(c *gin.Context) {
	var GvaStudent pkgCrawlReports.Student
	err := c.ShouldBindJSON(&GvaStudent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := GvaStudentService.UpdateStudent(GvaStudent); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindStudent 用id查询测试表格
// @Tags Student
// @Summary 用id查询测试表格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgCrawlReports.Student true "用id查询测试表格"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /GvaStudent/findStudent [get]
func (GvaStudentApi *StudentApi) FindStudent(c *gin.Context) {
	ID := c.Query("ID")
	if reGvaStudent, err := GvaStudentService.GetStudent(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reGvaStudent": reGvaStudent}, c)
	}
}

// GetStudentList 分页获取测试表格列表
// @Tags Student
// @Summary 分页获取测试表格列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgCrawlReportsReq.StudentSearch true "分页获取测试表格列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /GvaStudent/getStudentList [get]
func (GvaStudentApi *StudentApi) GetStudentList(c *gin.Context) {
	var pageInfo pkgCrawlReportsReq.StudentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := GvaStudentService.GetStudentInfoList(pageInfo); err != nil {
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
