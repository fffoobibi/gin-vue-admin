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

type TblCrawlStatsService struct {
}

// CreateTblCrawlStats 创建tblCrawlStats表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tblCrawlStatsService *TblCrawlStatsService) CreateTblCrawlStats(tblCrawlStats *pkgCrawlReports.TblCrawlStats) (err error) {
	err = global.GVA_DB.Create(tblCrawlStats).Error
	return err
}

// DeleteTblCrawlStats 删除tblCrawlStats表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tblCrawlStatsService *TblCrawlStatsService) DeleteTblCrawlStats(ID string) (err error) {
	err = global.GVA_DB.Delete(&pkgCrawlReports.TblCrawlStats{}, "id = ?", ID).Error
	return err
}

// DeleteTblCrawlStatsByIds 批量删除tblCrawlStats表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tblCrawlStatsService *TblCrawlStatsService) DeleteTblCrawlStatsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]pkgCrawlReports.TblCrawlStats{}, "id in ?", IDs).Error
	return err
}

// UpdateTblCrawlStats 更新tblCrawlStats表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tblCrawlStatsService *TblCrawlStatsService) UpdateTblCrawlStats(tblCrawlStats pkgCrawlReports.TblCrawlStats) (err error) {
	err = global.GVA_DB.Save(&tblCrawlStats).Error
	return err
}

// GetTblCrawlStats 根据ID获取tblCrawlStats表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tblCrawlStatsService *TblCrawlStatsService) GetTblCrawlStats(ID string) (tblCrawlStats pkgCrawlReports.TblCrawlStats, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&tblCrawlStats).Error
	return
}

// GetTblCrawlStatsInfoList 分页获取tblCrawlStats表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tblCrawlStatsService *TblCrawlStatsService) GetTblCrawlStatsInfoList(info pkgCrawlReportsReq.TblCrawlStatsSearch) (list []pkgCrawlReports.TblCrawlStats, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&pkgCrawlReports.TblCrawlStats{})
	var tblCrawlStatss []pkgCrawlReports.TblCrawlStats
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tblCrawlStatss).Error
	return tblCrawlStatss, total, err
}

// GetFirstCrawlCount 获取初次访问次数
func (tblCrawlStatsService *TblCrawlStatsService) GetFirstCrawlCount(info pkgCrawlReportsReq.TblCrawlStatsSearchQuery) (map[string][]map[string]interface{}, error) {

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

		global.GVA_LOG.Debug("s ,t ", zap.Int64("start", midnightTimestamp), zap.Int64("end", endOfDayTimestamp))

		conditions := map[string][]string{
			"tiktok":    {"%tiktok%", "%tt%"},
			"youtube":   {"%youtube%", "%ytb%"},
			"instagram": {"%instagram%", "%ins%"},
		}

		var rs = map[string][]map[string]interface{}{}

		for key, value := range conditions {
			var list []map[string]interface{}
			db := global.GVA_DB.Model(&pkgCrawlReports.TblCrawlStats{})
			err := db.Select("ip, name, count, FROM_UNIXTIME(create_time, '%Y-%m-%d') as date").
				Where("(name LIKE ? or name LIKE ?) AND create_time >= ? AND create_time <= ?",
					value[0],
					value[1],
					midnightTimestamp, endOfDayTimestamp,
				).Order("create_time").Scan(&list).Error

			if err != nil {
				zap.Error(err)
			}
			rs[strings.Title(key)] = list
		}
		return rs, nil
	}
	return nil, nil
}
