package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type TblCrawlStatsSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	request.PageInfo
}
type TblCrawlStatsTimeSearchQuery struct {
	StartTime string `json:"st_time" form:"st_time" binding:"required"`
	EndTime   string `json:"ed_time" form:"ed_time" binding:"required"`
}

type TblCrawlStatsSearchQuery struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	*TblCrawlStatsTimeSearchQuery
}

type TblCrawlStatsPieDataQuery struct {
	Report uint8 `json:"report" form:"report"`
	*TblCrawlStatsTimeSearchQuery
}

func (c TblCrawlStatsTimeSearchQuery) GetTimeStamps() (int64, int64) {
	s1, _ := time.Parse("2006-01-02", c.StartTime)
	s2, _ := time.Parse("2006-01-02", c.EndTime)
	location, _ := time.LoadLocation("Asia/Shanghai")
	midnight := time.Date(s1.Year(), s1.Month(), s1.Day(), 0, 0, 0, 0, location)
	endOfDay := time.Date(s2.Year(), s2.Month(), s2.Day(), 23, 59, 59, 0, location)
	return midnight.Unix(), endOfDay.Unix()
}
