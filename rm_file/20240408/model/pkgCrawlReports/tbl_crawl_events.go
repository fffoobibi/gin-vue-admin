// 自动生成模板TblCrawlEvents
package pkgCrawlReports

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 时间表 结构体  TblCrawlEvents
type TblCrawlEvents struct {
 global.GVA_MODEL 
      Title  string `json:"title" form:"title" gorm:"column:title;comment:名称;size:255;"`  //名称 
      Details  string `json:"details" form:"details" gorm:"column:details;comment:详细;size:1024;"`  //详细 
}


// TableName 时间表 TblCrawlEvents自定义表名 tbl_crawl_events
func (TblCrawlEvents) TableName() string {
  return "tbl_crawl_events"
}

