package pkgCrawlReports

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgCrawlReports"
	pkgCrawlReportsReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgCrawlReports/request"
	"go.uber.org/zap"
	"strings"
	"sync"
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

		var rs = map[string][]map[string]any{}

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
				global.GVA_LOG.Error("error in query", zap.Error(err))
			}
			rs[strings.Title(key)] = list
		}

		return rs, nil
	}
	return nil, nil
}

// GetSummaryCrawlInfo 获取访问统计次数
func (tblCrawlStatsService *TblCrawlStatsService) GetSummaryCrawlInfo(info pkgCrawlReportsReq.TblCrawlStatsTimeSearchQuery) (struct {
	CrawlCount  uint64           `json:"crawl_count"`
	ValidCount  uint64           `json:"valid_count"`
	CleanCount  uint64           `json:"clean_count"`
	UpdateCount uint64           `json:"update_count"`
	TotalCount  uint64           `json:"total_count"`
	Events      []map[string]any `json:"events"`
}, error) {
	s1, _ := time.Parse("2006-01-02", info.StartTime)
	s2, _ := time.Parse("2006-01-02", info.EndTime)
	location, _ := time.LoadLocation("Asia/Shanghai")
	midnight := time.Date(s1.Year(), s1.Month(), s1.Day(), 0, 0, 0, 0, location)
	endOfDay := time.Date(s2.Year(), s2.Month(), s2.Day(), 23, 59, 59, 0, location)

	global.GVA_LOG.Info("query time", zap.String("start", fmt.Sprintf("%s", midnight)),
		zap.String("end", fmt.Sprintf("%s", endOfDay)))

	midnightTimestamp := midnight.Unix()
	endOfDayTimestamp := endOfDay.Unix()

	var rs struct {
		CrawlCount  uint64           `json:"crawl_count"`
		ValidCount  uint64           `json:"valid_count"`
		CleanCount  uint64           `json:"clean_count"`
		UpdateCount uint64           `json:"update_count"`
		TotalCount  uint64           `json:"total_count"`
		Events      []map[string]any `json:"events"`
	}

	if err := global.GVA_DB.Model(&pkgCrawlReports.TblCrawlStats{}).Debug().
		Select("SUM(count) as `crawl_count`").
		Where("create_time >= ? AND create_time <= ?", midnightTimestamp, endOfDayTimestamp).
		Scan(&rs).Error; err != nil {
		global.GVA_LOG.Error("error in query crawl count", zap.Error(err))
	}
	if err := global.GetGlobalDBByDBName("mediamz").Model(&pkgCrawlReports.TblKolResource{}).Debug().
		Select("COUNT(*) as valid_count").
		Where("create_time >= ? AND create_time <= ?", midnightTimestamp, endOfDayTimestamp).Scan(&rs).Error; err != nil {
		global.GVA_LOG.Error("error in query valid count", zap.Error(err))
	}

	if err := global.GetGlobalDBByDBName("mediamz").Model(&pkgCrawlReports.TblKolResourceClean{}).Debug().
		Select("COUNT(*) as clean_count").
		Where("create_time >= ? AND create_time <= ?", midnightTimestamp, endOfDayTimestamp).Scan(&rs).Error; err != nil {
		global.GVA_LOG.Error("error in query clean count", zap.Error(err))
	}

	if err := global.GVA_DB.Model(&pkgCrawlReports.TblCrawlStats{}).Debug().
		Select("SUM(count) as update_count").
		Where("create_time >= ? AND create_time <= ? AND (name LIKE ? or name LIKE ? or name Like ?)", midnightTimestamp, endOfDayTimestamp, "%total-resource-tiktok%", "%total-resource-ytb%", "%total-resource-ins%").Scan(&rs).Error; err != nil {
		global.GVA_LOG.Error("error in query update count", zap.Error(err))
	}

	if err := global.GetGlobalDBByDBName("mediamz").Table("tbl_marketing_total_resource").Debug().
		Select("COUNT(*) as total_count").
		Scan(&rs).Error; err != nil {
		global.GVA_LOG.Error("error in query total count", zap.Error(err))
	}
	var infoList []map[string]any
	global.GVA_DB.Model(&pkgCrawlReports.TblCrawlEvents{}).Order("occurced_time DESC").Limit(3).Scan(&infoList)
	rs.Events = infoList

	return rs, nil
}

// GetCleanReportsList 获取清洗列表
func (tblCrawlStatsService *TblCrawlStatsService) GetCleanReportsList(info pkgCrawlReportsReq.TblCrawlStatsTimeSearchQuery) (map[string][]map[string]any, error) {
	s1, _ := time.Parse("2006-01-02", info.StartTime)
	s2, _ := time.Parse("2006-01-02", info.EndTime)
	location, _ := time.LoadLocation("Asia/Shanghai")
	midnight := time.Date(s1.Year(), s1.Month(), s1.Day(), 0, 0, 0, 0, location)
	endOfDay := time.Date(s2.Year(), s2.Month(), s2.Day(), 23, 59, 59, 0, location)

	global.GVA_LOG.Info("query time", zap.String("start", fmt.Sprintf("%s", midnight)),
		zap.String("end", fmt.Sprintf("%s", endOfDay)))

	midnightTimestamp := midnight.Unix()
	endOfDayTimestamp := endOfDay.Unix()
	conditions := map[string]uint8{"TikTok": 1, "Youtube": 2, "Instagram": 3}

	var rs = map[string][]map[string]any{}
	for k, v := range conditions {
		var list []map[string]any
		if err := global.GetGlobalDBByDBName("mediamz").Table("tbl_kol_resource_clean").
			Select("FROM_UNIXTIME(update_time, '%Y-%m-%d') as date, COUNT(*) AS count").
			Where("update_time >= ? AND update_time <= ? AND platform = ?",
				midnightTimestamp, endOfDayTimestamp, v,
			).
			Group("date").
			Order("date").Scan(&list).Error; err != nil {
			global.GVA_LOG.Error("error in query "+k, zap.Error(err))
		} else {
			rs[k] = list
		}
	}
	return rs, nil
}

// GetTotalResourceReportsList 获取总资源库更新次数
func (tblCrawlStatsService *TblCrawlStatsService) GetTotalResourceReportsList(info pkgCrawlReportsReq.TblCrawlStatsTimeSearchQuery) (map[string][]struct {
	Count int    `json:"count"`
	Date  string `json:"date"`
}, error) {
	s1, _ := time.Parse("2006-01-02", info.StartTime)
	s2, _ := time.Parse("2006-01-02", info.EndTime)
	location, _ := time.LoadLocation("Asia/Shanghai")
	midnight := time.Date(s1.Year(), s1.Month(), s1.Day(), 0, 0, 0, 0, location)
	endOfDay := time.Date(s2.Year(), s2.Month(), s2.Day(), 23, 59, 59, 0, location)

	global.GVA_LOG.Info("query time", zap.String("start", fmt.Sprintf("%s", midnight)),
		zap.String("end", fmt.Sprintf("%s", endOfDay)))

	midnightTimestamp := midnight.Unix()
	endOfDayTimestamp := endOfDay.Unix()
	conditions := map[string]string{"TikTok": "total-resource-tiktok", "Youtube": "total-resource-ytb", "Instagram": "total-resource-ins"}
	var rs = make(map[string][]struct {
		Count int    `json:"count"`
		Date  string `json:"date"`
	})
	var waiter sync.WaitGroup
	waiter.Add(3)
	for k, v := range conditions {
		go func(field, c string) {
			if rows, err := global.GVA_DB.Model(&pkgCrawlReports.TblCrawlStats{}).Debug().
				Select("FROM_UNIXTIME(create_time, '%Y-%m-%d') as date, SUM(count) AS count").
				Where("create_time >= ? AND create_time <= ? AND name = ?",
					midnightTimestamp, endOfDayTimestamp, c,
				).
				Group("date").
				Order("date").Rows(); err != nil {
				global.GVA_LOG.Error("error in query "+field, zap.Error(err))
			} else {
				defer rows.Close()
				var list []struct {
					Count int    `json:"count"`
					Date  string `json:"date"`
				}
				for rows.Next() {
					var info struct {
						Count int    `json:"count"`
						Date  string `json:"date"`
					}
					_ = global.GVA_DB.ScanRows(rows, &info)
					list = append(list, info)
				}
				rs[field] = list
			}
			waiter.Done()
		}(k, v)
	}
	waiter.Wait()
	return rs, nil
}

//// GetCrawlStatsPieData 获取饼图数据
//func (tblCrawlStatsService *TblCrawlStatsService) GetCrawlStatsPieData(query pkgCrawlReportsReq.TblCrawlStatsPieDataQuery) {
//	s1, _ := time.Parse("2006-01-02", query.StartTime)
//	s2, _ := time.Parse("2006-01-02", query.EndTime)
//	location, _ := time.LoadLocation("Asia/Shanghai")
//	midnight := time.Date(s1.Year(), s1.Month(), s1.Day(), 0, 0, 0, 0, location)
//	endOfDay := time.Date(s2.Year(), s2.Month(), s2.Day(), 23, 59, 59, 0, location)
//
//	global.GVA_LOG.Info("query time", zap.String("start", fmt.Sprintf("%s", midnight)),
//		zap.String("end", fmt.Sprintf("%s", endOfDay)))
//	//1 tt, 2 ytb 3 ins
//	midnightTimestamp := midnight.Unix()
//	endOfDayTimestamp := endOfDay.Unix()
//	var rs = make(map[string][]map[string]any)
//	switch query.Report {
//	case 0:
//		var platform []map[string]any
//		if err := global.GetGlobalDBByDBName("mediamz").Table("tbl_kol_resource").
//			Select("platform, count(*) as count").
//			Where("create_time >= ? AND create_time <= ?", midnightTimestamp, endOfDayTimestamp).
//			Group("platform").Scan(&platform).Error; err != nil {
//			global.GVA_LOG.Error("error in query", zap.Error(err))
//		} else {
//			rs["platform"] = platform
//		}
//	case 1:
//
//	case 2:
//
//	case 3:
//
//	}
//	if query.Report == 0 {
//		if err := global.GetGlobalDBByDBName("mediamz").Table("tbl_kol_resource").Select().
//			Where().
//			Group().Scan().Error; err != nil {
//
//		}
//	}
//}
