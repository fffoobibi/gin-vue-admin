package pkgCrawlReports

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgCrawlReports"
    pkgCrawlReportsReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgCrawlReports/request"
)

type StudentService struct {
}

// CreateStudent 创建测试表格记录
// Author [piexlmax](https://github.com/piexlmax)
func (GvaStudentService *StudentService) CreateStudent(GvaStudent *pkgCrawlReports.Student) (err error) {
	err = global.GVA_DB.Create(GvaStudent).Error
	return err
}

// DeleteStudent 删除测试表格记录
// Author [piexlmax](https://github.com/piexlmax)
func (GvaStudentService *StudentService)DeleteStudent(ID string) (err error) {
	err = global.GVA_DB.Delete(&pkgCrawlReports.Student{},"id = ?",ID).Error
	return err
}

// DeleteStudentByIds 批量删除测试表格记录
// Author [piexlmax](https://github.com/piexlmax)
func (GvaStudentService *StudentService)DeleteStudentByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]pkgCrawlReports.Student{},"id in ?",IDs).Error
	return err
}

// UpdateStudent 更新测试表格记录
// Author [piexlmax](https://github.com/piexlmax)
func (GvaStudentService *StudentService)UpdateStudent(GvaStudent pkgCrawlReports.Student) (err error) {
	err = global.GVA_DB.Save(&GvaStudent).Error
	return err
}

// GetStudent 根据ID获取测试表格记录
// Author [piexlmax](https://github.com/piexlmax)
func (GvaStudentService *StudentService)GetStudent(ID string) (GvaStudent pkgCrawlReports.Student, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&GvaStudent).Error
	return
}

// GetStudentInfoList 分页获取测试表格记录
// Author [piexlmax](https://github.com/piexlmax)
func (GvaStudentService *StudentService)GetStudentInfoList(info pkgCrawlReportsReq.StudentSearch) (list []pkgCrawlReports.Student, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&pkgCrawlReports.Student{})
    var GvaStudents []pkgCrawlReports.Student
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&GvaStudents).Error
	return  GvaStudents, total, err
}
