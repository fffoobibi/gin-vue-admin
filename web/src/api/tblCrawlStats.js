import service from '@/utils/request'

// @Tags TblCrawlStats
// @Summary 创建tblCrawlStats表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblCrawlStats true "创建tblCrawlStats表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tblCrawlStats/createTblCrawlStats [post]
export const createTblCrawlStats = (data) => {
  return service({
    url: '/tblCrawlStats/createTblCrawlStats',
    method: 'post',
    data
  })
}

// @Tags TblCrawlStats
// @Summary 删除tblCrawlStats表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblCrawlStats true "删除tblCrawlStats表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tblCrawlStats/deleteTblCrawlStats [delete]
export const deleteTblCrawlStats = (params) => {
  return service({
    url: '/tblCrawlStats/deleteTblCrawlStats',
    method: 'delete',
    params
  })
}

// @Tags TblCrawlStats
// @Summary 批量删除tblCrawlStats表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除tblCrawlStats表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tblCrawlStats/deleteTblCrawlStats [delete]
export const deleteTblCrawlStatsByIds = (params) => {
  return service({
    url: '/tblCrawlStats/deleteTblCrawlStatsByIds',
    method: 'delete',
    params
  })
}

// @Tags TblCrawlStats
// @Summary 更新tblCrawlStats表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblCrawlStats true "更新tblCrawlStats表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tblCrawlStats/updateTblCrawlStats [put]
export const updateTblCrawlStats = (data) => {
  return service({
    url: '/tblCrawlStats/updateTblCrawlStats',
    method: 'put',
    data
  })
}

// @Tags TblCrawlStats
// @Summary 用id查询tblCrawlStats表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TblCrawlStats true "用id查询tblCrawlStats表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tblCrawlStats/findTblCrawlStats [get]
export const findTblCrawlStats = (params) => {
  return service({
    url: '/tblCrawlStats/findTblCrawlStats',
    method: 'get',
    params
  })
}

// @Tags TblCrawlStats
// @Summary 分页获取tblCrawlStats表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取tblCrawlStats表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tblCrawlStats/getTblCrawlStatsList [get]
export const getTblCrawlStatsList = (params) => {
  return service({
    url: '/tblCrawlStats/getTblCrawlStatsList',
    method: 'get',
    params
  })
}

// @Tags TblCrawlStats
// @Summary 获取首次访问次数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tblCrawlStats/getFirstCrawlInfo [get]
export const getFirstCrawlInfo = (params) => {
  return service({
    url: '/tblCrawlStats/getFirstCrawlInfo',
    method: 'get',
    params
  })
}

// @Tags TblCrawlStats
// @Summary 获取总数量, 饼图数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tblCrawlStats/getTotalResourceInfo [get]
export const getTotalResourceInfo = (params) => {
  return service({
    url: '/tblCrawlStats/getTotalResourceInfo',
    method: 'get',
    params
  })
}
