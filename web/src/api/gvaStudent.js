import service from '@/utils/request'

// @Tags Student
// @Summary 创建测试表格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Student true "创建测试表格"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /GvaStudent/createStudent [post]
export const createStudent = (data) => {
  return service({
    url: '/GvaStudent/createStudent',
    method: 'post',
    data
  })
}

// @Tags Student
// @Summary 删除测试表格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Student true "删除测试表格"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /GvaStudent/deleteStudent [delete]
export const deleteStudent = (params) => {
  return service({
    url: '/GvaStudent/deleteStudent',
    method: 'delete',
    params
  })
}

// @Tags Student
// @Summary 批量删除测试表格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除测试表格"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /GvaStudent/deleteStudent [delete]
export const deleteStudentByIds = (params) => {
  return service({
    url: '/GvaStudent/deleteStudentByIds',
    method: 'delete',
    params
  })
}

// @Tags Student
// @Summary 更新测试表格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Student true "更新测试表格"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /GvaStudent/updateStudent [put]
export const updateStudent = (data) => {
  return service({
    url: '/GvaStudent/updateStudent',
    method: 'put',
    data
  })
}

// @Tags Student
// @Summary 用id查询测试表格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Student true "用id查询测试表格"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /GvaStudent/findStudent [get]
export const findStudent = (params) => {
  return service({
    url: '/GvaStudent/findStudent',
    method: 'get',
    params
  })
}

// @Tags Student
// @Summary 分页获取测试表格列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取测试表格列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /GvaStudent/getStudentList [get]
export const getStudentList = (params) => {
  return service({
    url: '/GvaStudent/getStudentList',
    method: 'get',
    params
  })
}
