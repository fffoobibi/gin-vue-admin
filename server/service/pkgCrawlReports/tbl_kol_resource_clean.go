package pkgCrawlReports

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgCrawlReports"
	pkgCrawlReportsReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgCrawlReports/request"
)

type TblKolResourceCleanService struct {
}

// CreateTblKolResourceClean 创建tblKolResourceClean清洗表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tblKolResourceCleanService *TblKolResourceCleanService) CreateTblKolResourceClean(tblKolResourceClean *pkgCrawlReports.TblKolResourceClean) (err error) {
	err = global.GVA_DB.Create(tblKolResourceClean).Error
	return err
}

// DeleteTblKolResourceClean 删除tblKolResourceClean清洗表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tblKolResourceCleanService *TblKolResourceCleanService) DeleteTblKolResourceClean(id string) (err error) {
	err = global.GVA_DB.Delete(&pkgCrawlReports.TblKolResourceClean{}, "id = ?", id).Error
	return err
}

// DeleteTblKolResourceCleanByIds 批量删除tblKolResourceClean清洗表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tblKolResourceCleanService *TblKolResourceCleanService) DeleteTblKolResourceCleanByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]pkgCrawlReports.TblKolResourceClean{}, "id in ?", ids).Error
	return err
}

// UpdateTblKolResourceClean 更新tblKolResourceClean清洗表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tblKolResourceCleanService *TblKolResourceCleanService) UpdateTblKolResourceClean(tblKolResourceClean pkgCrawlReports.TblKolResourceClean) (err error) {
	err = global.GVA_DB.Save(&tblKolResourceClean).Error
	return err
}

// GetTblKolResourceClean 根据id获取tblKolResourceClean清洗表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tblKolResourceCleanService *TblKolResourceCleanService) GetTblKolResourceClean(id string) (tblKolResourceClean pkgCrawlReports.TblKolResourceClean, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&tblKolResourceClean).Error
	return
}

// GetTblKolResourceCleanInfoList 分页获取tblKolResourceClean清洗表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tblKolResourceCleanService *TblKolResourceCleanService) GetTblKolResourceCleanInfoList(info pkgCrawlReportsReq.TblKolResourceCleanSearch) (list []pkgCrawlReports.TblKolResourceClean, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&pkgCrawlReports.TblKolResourceClean{})
	var tblKolResourceCleans []pkgCrawlReports.TblKolResourceClean
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tblKolResourceCleans).Error
	return tblKolResourceCleans, total, err
}
