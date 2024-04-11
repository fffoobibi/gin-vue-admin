package pkgCrawlReports

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgCrawlReports"
	pkgCrawlReportsReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgCrawlReports/request"
	"go.uber.org/zap"
	"strings"
	"time"
)

type TblKolResourceService struct {
}

// CreateTblKolResource 创建tblKolResource表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tblKolResourceService *TblKolResourceService) CreateTblKolResource(tblKolResource *pkgCrawlReports.TblKolResource) (err error) {
	err = global.GetGlobalDBByDBName("mediamz").Create(tblKolResource).Error
	return err
}

// DeleteTblKolResource 删除tblKolResource表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tblKolResourceService *TblKolResourceService) DeleteTblKolResource(id string) (err error) {
	err = global.GetGlobalDBByDBName("mediamz").Delete(&pkgCrawlReports.TblKolResource{}, "id = ?", id).Error
	return err
}

// DeleteTblKolResourceByIds 批量删除tblKolResource表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tblKolResourceService *TblKolResourceService) DeleteTblKolResourceByIds(ids []string) (err error) {
	err = global.GetGlobalDBByDBName("mediamz").Delete(&[]pkgCrawlReports.TblKolResource{}, "id in ?", ids).Error
	return err
}

// UpdateTblKolResource 更新tblKolResource表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tblKolResourceService *TblKolResourceService) UpdateTblKolResource(tblKolResource pkgCrawlReports.TblKolResource) (err error) {
	err = global.GetGlobalDBByDBName("mediamz").Save(&tblKolResource).Error
	return err
}

// GetTblKolResource 根据id获取tblKolResource表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tblKolResourceService *TblKolResourceService) GetTblKolResource(id string) (tblKolResource pkgCrawlReports.TblKolResource, err error) {
	err = global.GetGlobalDBByDBName("mediamz").Where("id = ?", id).First(&tblKolResource).Error
	return
}

// GetTblKolResourceInfoList 分页获取tblKolResource表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tblKolResourceService *TblKolResourceService) GetTblKolResourceInfoList(info pkgCrawlReportsReq.TblKolResourceSearch) (list []pkgCrawlReports.TblKolResource, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GetGlobalDBByDBName("mediamz").Model(&pkgCrawlReports.TblKolResource{})
	var tblKolResources []pkgCrawlReports.TblKolResource
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tblKolResources).Error
	return tblKolResources, total, err
}

// GetKolResourceCrawlInfo 获取有效资源访问资源列表
func (tblKolResourceService *TblKolResourceService) GetKolResourceCrawlInfo(info pkgCrawlReportsReq.TblKolResourceSearchQuery) (map[string][]map[string]any, error) {

	if info.StartTime != "" {
		t, _ := time.Parse("2006-01-02", info.StartTime)
		info.StartCreatedAt = &t
	}

	if info.EndTime != "" {
		t, _ := time.Parse("2006-01-02", info.EndTime)
		info.EndCreatedAt = &t
	}

	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		s1 := info.StartCreatedAt
		s2 := info.EndCreatedAt
		location, _ := time.LoadLocation("Asia/Shanghai")
		midnight := time.Date(s1.Year(), s1.Month(), s1.Day(), 0, 0, 0, 0, location)
		endOfDay := time.Date(s2.Year(), s2.Month(), s2.Day(), 23, 59, 59, 0, location)

		global.GVA_LOG.Info("start time", zap.String("end", fmt.Sprintf("%s", endOfDay)))

		midnightTimestamp := midnight.Unix()
		endOfDayTimestamp := endOfDay.Unix()

		//global.GVA_LOG.Debug("s ,t ", zap.Int64("start", midnightTimestamp), zap.Int64("end", endOfDayTimestamp))

		conditions := map[string]uint8{
			"tiktok":    1,
			"youtube":   2,
			"instagram": 3,
		}
		//平台：1 - TikTok 2 - YouTube 3 - Instagram

		var rs = map[string][]map[string]interface{}{}

		for key, value := range conditions {
			var list []map[string]interface{}
			db := global.GetGlobalDBByDBName("mediamz").Model(&pkgCrawlReports.TblKolResource{})
			err := db.Select("COUNT(*) as `count`, FROM_UNIXTIME(create_time, '%Y-%m-%d') as date").
				Where("platform = ? AND create_time >= ? AND create_time <= ?",
					value, midnightTimestamp, endOfDayTimestamp,
				).Group("date").Order("date").Scan(&list).Error

			if err != nil {
				global.GVA_LOG.Error("error in query", zap.Error(err))
			}
			rs[strings.Title(key)] = list
		}
		return rs, nil
	}
	return nil, nil
}
