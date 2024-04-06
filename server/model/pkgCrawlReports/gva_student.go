// 自动生成模板Student
package pkgCrawlReports

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 测试表格 结构体  Student
type Student struct {
 global.GVA_MODEL 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:255;"`  //姓名 
}


// TableName 测试表格 Student自定义表名 gva_students
func (Student) TableName() string {
  return "gva_students"
}

