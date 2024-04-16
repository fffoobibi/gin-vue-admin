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

type TblCrawlStatsSearchQuery struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	StartTime string `json:"st_time" form:"st_time" binding:"required"`
	EndTime   string `json:"ed_time" form:"ed_time" binding:"required"`
}

type TblCrawlStatsTimeSearchQuery struct {
	StartTime string `json:"st_time" form:"st_time" binding:"required"`
	EndTime   string `json:"ed_time" form:"ed_time" binding:"required"`
}

type TblCrawlStatsPieDataQuery struct {
	Report uint8 `json:"report" form:"report"`
	*TblCrawlStatsTimeSearchQuery
}
