package pkgCrawlReports

import (
	"database/sql"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgCrawlReports"
	pkgCrawlReportsReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgCrawlReports/request"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
	"sync"
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

	midnightTimestamp, endOfDayTimestamp := info.GetTimeStamps()

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

// GetSummaryCrawlInfo 获取访问统计次数
func (tblCrawlStatsService *TblCrawlStatsService) GetSummaryCrawlInfo(info pkgCrawlReportsReq.TblCrawlStatsTimeSearchQuery) (struct {
	CrawlCount  int              `json:"crawl_count"`
	ValidCount  int              `json:"valid_count"`
	CleanCount  int              `json:"clean_count"`
	UpdateCount int              `json:"update_count"`
	TotalCount  int              `json:"total_count"`
	Events      []map[string]any `json:"events" gorm:"-"`
}, error) {
	midnightTimestamp, endOfDayTimestamp := info.GetTimeStamps()

	var rs struct {
		CrawlCount  int              `json:"crawl_count"`
		ValidCount  int              `json:"valid_count"`
		CleanCount  int              `json:"clean_count"`
		UpdateCount int              `json:"update_count"`
		TotalCount  int              `json:"total_count"`
		Events      []map[string]any `json:"events" gorm:"-"`
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
		Where("create_time >= ? AND create_time <= ? AND (name LIKE ? or name LIKE ? or name Like ?)",
			midnightTimestamp, endOfDayTimestamp, "%total-resource-tiktok%", "%total-resource-ytb%", "%total-resource-ins%").Scan(&rs).Error; err != nil {
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
	midnightTimestamp, endOfDayTimestamp := info.GetTimeStamps()
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
	midnightTimestamp, endOfDayTimestamp := info.GetTimeStamps()
	conditions := map[string]string{"TikTok": "total-resource-tiktok", "Youtube": "total-resource-ytb", "Instagram": "total-resource-ins"}
	var rs = make(map[string][]struct {
		Count int    `json:"count"`
		Date  string `json:"date"`
	})
	var waiter sync.WaitGroup
	waiter.Add(len(conditions))
	for k, v := range conditions {
		go func(field, c string) {
			defer func() {
				if r := recover(); r != nil {
					// 处理 panic，例如打印错误日志等
					global.GVA_LOG.Error("panic in GetTotalResourceReportsList", zap.Any("error", r))
				}
				waiter.Done()
			}()
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
		}(k, v)
	}
	waiter.Wait()
	return rs, nil
}

// GetCrawlStatsPieData 获取饼图数据
func (tblCrawlStatsService *TblCrawlStatsService) GetCrawlStatsPieData(query pkgCrawlReportsReq.TblCrawlStatsPieDataQuery) (
	platform []*struct {
		Platform byte `json:"platform"`
		Count    int  `json:"count"`
	}, category []*struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	}, country []*struct {
		CountryName string `json:"country_name"`
		CountryCode int    `json:"country_code"`
		Count       int    `json:"count"`
	}) {

	stTime, edTime := query.GetTimeStamps()

	platform = make([]*struct {
		Platform byte `json:"platform"`
		Count    int  `json:"count"`
	}, 0)
	category = make([]*struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	}, 0)
	country = make([]*struct {
		CountryName string `json:"country_name"`
		CountryCode int    `json:"country_code"`
		Count       int    `json:"count"`
	}, 0)

	mediamzDb := global.MustGetGlobalDBByDBName("mediamz").Session(&gorm.Session{})

	categorySql := func(stTime, edTime int64, timeName, tableName string, limit int) string {
		return mediamzDb.ToSQL(func(tx *gorm.DB) *gorm.DB {
			return tx.Raw(
				fmt.Sprintf(
					`select r.name, COUNT(*) as count from %s t left join tbl_category r on t.category_id = r.id 
                	 where t.%s >= @st_time and t.%s <= @ed_time group by t.category_id order by count desc limit @limit`,
					tableName, timeName, timeName),
				sql.Named("st_time", stTime), sql.Named("ed_time", edTime), sql.Named("limit", limit),
			)
		})
	}
	countrySql := func(stTime, edTime int64, timeName, tableName, groupField string) string {
		return mediamzDb.ToSQL(func(tx *gorm.DB) *gorm.DB {
			return tx.Raw(
				fmt.Sprintf(
					`select t.%s As country_code, r.country_name, COUNT(*) as count from %s t left join tbl_country_geo r on t.%s = r.country_code 
					where t.%s >= ? and t.%s <= ? group by t.%s, r.country_name order by count`,
					groupField, tableName, groupField, timeName, timeName, groupField,
				), stTime, edTime)
		})
	}

	platformSql := func(stTime, edTime int64, timeName, tableName string) string {
		return mediamzDb.ToSQL(func(tx *gorm.DB) *gorm.DB {
			return tx.Raw(
				fmt.Sprintf(
					`select t.platform, COUNT(*) as count from %s t
					where t.%s >= @st_time and t.%s <= @ed_time group by t.platform`,
					tableName, timeName, timeName,
				),
				sql.Named("st_time", stTime), sql.Named("ed_time", edTime))
		})
	}
	tableName, categoryTimeName, platformTimeName, countryTimeName, groupField := "", "", "", "", ""
	limit := 10
	switch query.Report {
	case 0, 1:
		tableName = "tbl_kol_resource"
		groupField = "region_code"
		categoryTimeName = "create_time"
		countryTimeName = "create_time"
		platformTimeName = "create_time"
	case 2:
		tableName = "tbl_kol_resource_clean"
		groupField = "region_code"
		categoryTimeName = "create_time"
		countryTimeName = "create_time"
		platformTimeName = "create_time"
	case 3:
		tableName = "tbl_marketing_total_resource"
		groupField = "country_code"
		categoryTimeName = "last_crawler_time"
		countryTimeName = "last_crawler_time"
		platformTimeName = "last_crawler_time"
	}
	mediamzDb.Raw(countrySql(stTime, edTime, countryTimeName, tableName, groupField)).Scan(&country)
	mediamzDb.Raw(platformSql(stTime, edTime, platformTimeName, tableName)).Scan(&platform)
	mediamzDb.Raw(categorySql(stTime, edTime, categoryTimeName, tableName, limit)).Scan(&category)

	// 处理重复的类目数据
	categories := make([]*struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	}, 0)
	ms := map[string]*struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	}{}
	categoriesMap := map[string]int{}

	for index, category := range category {
		if _, ok := ms[category.Name]; !ok {
			ms[category.Name] = category
			categoriesMap[category.Name] = index
			categories = append(categories, category)
		} else {
			lookupIndex := categoriesMap[category.Name]
			categories[lookupIndex].Count += category.Count
		}
	}
	category = categories
	return platform, category, country
}
