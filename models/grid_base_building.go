package models

import (
	orm "go-admin-demo/database"
	"go-admin-demo/tools"
	_ "time"
)

type GridBasicBuiding struct {
	BuildingId  int    `json:"buildingId" gorm:"column:buildingId;type:bigint(20);primary_key"` // 建筑ID
	GridId      int    `json:"gridId" gorm:"column:gridId;type:bigint(20);"`                    // 网格ID
	SecretaryId int    `json:"secretaryId" gorm:"column:secretaryId;type:bigint(20);"`          // 楼栋长ID
	Sits        string `json:"sits" gorm:"column:sits;type:tinytext;"`                          // 坐落地
	Structure   int    `json:"structure" gorm:"column:structure;type:tinyint(4);"`              // 建筑结构 0-未知 1-钢混 2-木 3-砖 4-钢架构
	Layers      int    `json:"layers" gorm:"column:layers;type:tinyint(4);"`                    // 建筑层数
	Years       string `json:"years" gorm:"column:years;type:date;"`                            // 建筑年代
	RiskLevel   int    `json:"riskLevel" gorm:"column:riskLevel;type:tinyint(4);"`              // 建筑危险等级 0-未知 1-正常 2-轻度 3-中度 4-重度 5-危房 6-拆除
	Developers  string `json:"developers" gorm:"column:developers;type:varchar(50);"`           // 开发商
	Remarks     string `json:"remarks" gorm:"column:remarks;type:tinytext;"`                    // 备注
	CreateBy    string `json:"createBy" gorm:"type:varchar(2);"`                                // 创建人
	UpdateBy    string `json:"updateBy" gorm:"type:varchar(2);"`                                // 更新人
	DataScope   string `json:"dataScope" gorm:"-"`
	Params      string `json:"params"  gorm:"-"`
	BaseModel
	GridName  string `json:"gridName"  gorm:"-"`  // 网格名称 - 前端展示需要
	Secretary string `json:"secretary"  gorm:"-"` // 楼栋长姓名 - 前端展示需要
}

func (GridBasicBuiding) TableName() string {
	return "grid_basic_buiding"
}

// 创建GridBasicBuiding
func (e *GridBasicBuiding) Create() (GridBasicBuiding, error) {
	var doc GridBasicBuiding
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取GridBasicBuiding
func (e *GridBasicBuiding) Get() (GridBasicBuiding, error) {
	var doc GridBasicBuiding

	table := orm.Eloquent.Table(e.TableName())

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取GridBasicBuiding带分页
func (e *GridBasicBuiding) GetPage(pageSize int, pageIndex int) ([]GridBasicBuiding, int, error) {
	var doc []GridBasicBuiding

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

// 更新GridBasicBuiding
func (e *GridBasicBuiding) Update(id int) (update GridBasicBuiding, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("buildingId = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 删除GridBasicBuiding
func (e *GridBasicBuiding) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("buildingId = ?", id).Delete(&GridBasicBuiding{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *GridBasicBuiding) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("buildingId in (?)", id).Delete(&GridBasicBuiding{}).Error; err != nil {
		return
	}
	Result = true
	return
}
