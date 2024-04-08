package pkgCrawlReports

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgCrawlReports"
    pkgCrawlReportsReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgCrawlReports/request"
)

type TblCrawlEventsService struct {
}

// CreateTblCrawlEvents 创建时间表记录
// Author [piexlmax](https://github.com/piexlmax)
func (eService *TblCrawlEventsService) CreateTblCrawlEvents(e *pkgCrawlReports.TblCrawlEvents) (err error) {
	err = global.GVA_DB.Create(e).Error
	return err
}

// DeleteTblCrawlEvents 删除时间表记录
// Author [piexlmax](https://github.com/piexlmax)
func (eService *TblCrawlEventsService)DeleteTblCrawlEvents(ID string) (err error) {
	err = global.GVA_DB.Delete(&pkgCrawlReports.TblCrawlEvents{},"id = ?",ID).Error
	return err
}

// DeleteTblCrawlEventsByIds 批量删除时间表记录
// Author [piexlmax](https://github.com/piexlmax)
func (eService *TblCrawlEventsService)DeleteTblCrawlEventsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]pkgCrawlReports.TblCrawlEvents{},"id in ?",IDs).Error
	return err
}

// UpdateTblCrawlEvents 更新时间表记录
// Author [piexlmax](https://github.com/piexlmax)
func (eService *TblCrawlEventsService)UpdateTblCrawlEvents(e pkgCrawlReports.TblCrawlEvents) (err error) {
	err = global.GVA_DB.Save(&e).Error
	return err
}

// GetTblCrawlEvents 根据ID获取时间表记录
// Author [piexlmax](https://github.com/piexlmax)
func (eService *TblCrawlEventsService)GetTblCrawlEvents(ID string) (e pkgCrawlReports.TblCrawlEvents, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&e).Error
	return
}

// GetTblCrawlEventsInfoList 分页获取时间表记录
// Author [piexlmax](https://github.com/piexlmax)
func (eService *TblCrawlEventsService)GetTblCrawlEventsInfoList(info pkgCrawlReportsReq.TblCrawlEventsSearch) (list []pkgCrawlReports.TblCrawlEvents, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&pkgCrawlReports.TblCrawlEvents{})
    var es []pkgCrawlReports.TblCrawlEvents
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&es).Error
	return  es, total, err
}
