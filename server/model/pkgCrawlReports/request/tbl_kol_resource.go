package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type TblKolResourceSearch struct {
	request.PageInfo
}

type TblKolResourceSearchQuery struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	StartTime string `json:"st_time" form:"st_time"`
	EndTime   string `json:"ed_time" form:"ed_time"`
}
