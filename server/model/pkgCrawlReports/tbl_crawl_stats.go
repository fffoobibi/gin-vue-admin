// 自动生成模板TblCrawlStats
package pkgCrawlReports

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// tblCrawlStats表 结构体  TblCrawlStats
type TblCrawlStats struct {
	global.GVA_MODEL
	Count      *int   `json:"count" form:"count" gorm:"type:int(11);column:count;comment:抓取次数;size:10;"`                 //抓取次数
	CreateTime *int   `json:"createTime" form:"createTime" gorm:"type:int(11);column:create_time;comment:创建时间;size:10;"` //创建时间
	Ip         string `json:"ip" form:"ip" gorm:"column:ip;comment:来源;size:30;"`                                         //来源
	Name       string `json:"name" form:"name" gorm:"column:name;comment:爬虫名称;size:40;"`                                 //爬虫名称
	Project    string `json:"project" form:"project" gorm:"column:project;comment:项目;size:30;"`                          //项目
}

// TableName tblCrawlStats表 TblCrawlStats自定义表名 tbl_crawl_stats
func (TblCrawlStats) TableName() string {
	return "tbl_crawl_stats"
}
