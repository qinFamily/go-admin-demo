/*
 * @Author: xiaoxu@mgtv.com
 * @Date: 2020-07-23 21:59:21
 * @Jira:
 * @Wiki:
 * @LastEditTime: 2020-07-23 22:00:33
 * @LastEditors: xiaoxu
 * @Description:
 * @FilePath: \go-admin-ui-vuef:\project\work\go\src\go-admin-demo\models\grid.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package models

import (
	orm "go-admin-demo/database"
	"go-admin-demo/tools"
	_ "time"
)

type Grid struct {
	GridId      int    `json:"gridId" gorm:"column:gridId;type:bigint(20);primary_key"`  // 网格ID
	GridName    string `json:"gridName" gorm:"column:gridName;type:varchar(50);"`        // 网格名称
	Description string `json:"description" gorm:"column:description;type:varchar(100);"` // 辖区描述
	Area        string `json:"area" gorm:"column:area;type:float;"`                      // 面积
	ChairmanId  int    `json:"chairmanId" gorm:"column:chairmanId;type:bigint(20);"`     // 负责人信息
	SecretaryId int    `json:"secretaryId" gorm:"column:secretaryId;type:bigint(20);"`   // 干事
	Remarks     string `json:"remarks" gorm:"column:remarks;type:text;"`                 // 备注
	CreateBy    string `json:"createBy" gorm:"column:create_by;type:varchar(2);"`        // 创建人
	UpdateBy    string `json:"updateBy" gorm:"column:update_by;type:varchar(2);"`        // 更新人
	DataScope   string `json:"dataScope" gorm:"-"`
	Params      string `json:"params"  gorm:"-"`
	BaseModel
}

func (Grid) TableName() string {
	return "grid"
}

// 创建Grid
func (e *Grid) Create() (Grid, error) {
	var doc Grid
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取Grid
func (e *Grid) Get() (Grid, error) {
	var doc Grid

	table := orm.Eloquent.Table(e.TableName())

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取Grid带分页
func (e *Grid) GetPage(pageSize int, pageIndex int) ([]Grid, int, error) {
	var doc []Grid

	table := orm.Eloquent.Select("*").Table(e.TableName())

	// 数据权限控制(如果不需要数据权限请将此处去掉)
	dataPermission := new(DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table = dataPermission.GetDataScope(e.TableName(), table)

	var count int

	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Count(&count)
	return doc, count, nil
}

// 更新Grid
func (e *Grid) Update(id int) (update Grid, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("gridId = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 删除Grid
func (e *Grid) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("gridId = ?", id).Delete(&Grid{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *Grid) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("gridId in (?)", id).Delete(&Grid{}).Error; err != nil {
		return
	}
	Result = true
	return
}
