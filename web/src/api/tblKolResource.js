import service from '@/utils/request'

// @Tags TblKolResource
// @Summary 创建tblKolResource表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblKolResource true "创建tblKolResource表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tblKolResource/createTblKolResource [post]
export const createTblKolResource = (data) => {
  return service({
    url: '/tblKolResource/createTblKolResource',
    method: 'post',
    data
  })
}

// @Tags TblKolResource
// @Summary 删除tblKolResource表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblKolResource true "删除tblKolResource表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tblKolResource/deleteTblKolResource [delete]
export const deleteTblKolResource = (params) => {
  return service({
    url: '/tblKolResource/deleteTblKolResource',
    method: 'delete',
    params
  })
}

// @Tags TblKolResource
// @Summary 批量删除tblKolResource表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除tblKolResource表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tblKolResource/deleteTblKolResource [delete]
export const deleteTblKolResourceByIds = (params) => {
  return service({
    url: '/tblKolResource/deleteTblKolResourceByIds',
    method: 'delete',
    params
  })
}

// @Tags TblKolResource
// @Summary 更新tblKolResource表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblKolResource true "更新tblKolResource表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tblKolResource/updateTblKolResource [put]
export const updateTblKolResource = (data) => {
  return service({
    url: '/tblKolResource/updateTblKolResource',
    method: 'put',
    data
  })
}

// @Tags TblKolResource
// @Summary 用id查询tblKolResource表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TblKolResource true "用id查询tblKolResource表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tblKolResource/findTblKolResource [get]
export const findTblKolResource = (params) => {
  return service({
    url: '/tblKolResource/findTblKolResource',
    method: 'get',
    params
  })
}

// @Tags TblKolResource
// @Summary 分页获取tblKolResource表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取tblKolResource表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tblKolResource/getTblKolResourceList [get]
export const getTblKolResourceList = (params) => {
  return service({
    url: '/tblKolResource/getTblKolResourceList',
    method: 'get',
    params
  })
}


// @Tags TblKolResource
// @Summary 获取邮箱资源统计
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tblCrawlStats/getFirstCrawlInfo [get]
export const getKolCrawlInfo = (params) => {
  return service({
    url: '/tblKolResource/getKolResourceCrawlInfo',
    method: 'get',
    params,
  })
}