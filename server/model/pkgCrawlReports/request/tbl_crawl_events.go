package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type TblCrawlEventsSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	StartOccuredTime *time.Time `json:"startOccuredTime" form:"startOccuredTime"`
	EndOccuredTime   *time.Time `json:"endOccuredTime" form:"endOccuredTime"`
	request.PageInfo
}
