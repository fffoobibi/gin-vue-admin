import service from '@/utils/request'

// @Tags TblKolResourceClean
// @Summary 创建tblKolResourceClean清洗表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblKolResourceClean true "创建tblKolResourceClean清洗表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tblKolResourceClean/createTblKolResourceClean [post]
export const createTblKolResourceClean = (data) => {
  return service({
    url: '/tblKolResourceClean/createTblKolResourceClean',
    method: 'post',
    data
  })
}

// @Tags TblKolResourceClean
// @Summary 删除tblKolResourceClean清洗表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblKolResourceClean true "删除tblKolResourceClean清洗表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tblKolResourceClean/deleteTblKolResourceClean [delete]
export const deleteTblKolResourceClean = (params) => {
  return service({
    url: '/tblKolResourceClean/deleteTblKolResourceClean',
    method: 'delete',
    params
  })
}

// @Tags TblKolResourceClean
// @Summary 批量删除tblKolResourceClean清洗表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除tblKolResourceClean清洗表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tblKolResourceClean/deleteTblKolResourceClean [delete]
export const deleteTblKolResourceCleanByIds = (params) => {
  return service({
    url: '/tblKolResourceClean/deleteTblKolResourceCleanByIds',
    method: 'delete',
    params
  })
}

// @Tags TblKolResourceClean
// @Summary 更新tblKolResourceClean清洗表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblKolResourceClean true "更新tblKolResourceClean清洗表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tblKolResourceClean/updateTblKolResourceClean [put]
export const updateTblKolResourceClean = (data) => {
  return service({
    url: '/tblKolResourceClean/updateTblKolResourceClean',
    method: 'put',
    data
  })
}

// @Tags TblKolResourceClean
// @Summary 用id查询tblKolResourceClean清洗表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TblKolResourceClean true "用id查询tblKolResourceClean清洗表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tblKolResourceClean/findTblKolResourceClean [get]
export const findTblKolResourceClean = (params) => {
  return service({
    url: '/tblKolResourceClean/findTblKolResourceClean',
    method: 'get',
    params
  })
}

// @Tags TblKolResourceClean
// @Summary 分页获取tblKolResourceClean清洗表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取tblKolResourceClean清洗表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tblKolResourceClean/getTblKolResourceCleanList [get]
export const getTblKolResourceCleanList = (params) => {
  return service({
    url: '/tblKolResourceClean/getTblKolResourceCleanList',
    method: 'get',
    params
  })
}
