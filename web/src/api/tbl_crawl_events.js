import service from '@/utils/request'

// @Tags TblCrawlEvents
// @Summary 创建爬虫大事迹
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblCrawlEvents true "创建爬虫大事迹"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /e/createTblCrawlEvents [post]
export const createTblCrawlEvents = (data) => {
  return service({
    url: '/e/createTblCrawlEvents',
    method: 'post',
    data
  })
}

// @Tags TblCrawlEvents
// @Summary 删除爬虫大事迹
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblCrawlEvents true "删除爬虫大事迹"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /e/deleteTblCrawlEvents [delete]
export const deleteTblCrawlEvents = (params) => {
  return service({
    url: '/e/deleteTblCrawlEvents',
    method: 'delete',
    params
  })
}

// @Tags TblCrawlEvents
// @Summary 批量删除爬虫大事迹
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除爬虫大事迹"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /e/deleteTblCrawlEvents [delete]
export const deleteTblCrawlEventsByIds = (params) => {
  return service({
    url: '/e/deleteTblCrawlEventsByIds',
    method: 'delete',
    params
  })
}

// @Tags TblCrawlEvents
// @Summary 更新爬虫大事迹
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblCrawlEvents true "更新爬虫大事迹"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /e/updateTblCrawlEvents [put]
export const updateTblCrawlEvents = (data) => {
  return service({
    url: '/e/updateTblCrawlEvents',
    method: 'put',
    data
  })
}

// @Tags TblCrawlEvents
// @Summary 用id查询爬虫大事迹
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TblCrawlEvents true "用id查询爬虫大事迹"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /e/findTblCrawlEvents [get]
export const findTblCrawlEvents = (params) => {
  return service({
    url: '/e/findTblCrawlEvents',
    method: 'get',
    params
  })
}

// @Tags TblCrawlEvents
// @Summary 分页获取爬虫大事迹列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取爬虫大事迹列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /e/getTblCrawlEventsList [get]
export const getTblCrawlEventsList = (params) => {
  return service({
    url: '/e/getTblCrawlEventsList',
    method: 'get',
    params
  })
}
