package models

import (
	orm "go-admin-demo/database"
	"go-admin-demo/tools"
	_ "time"
)

type GridBasicHousing struct {
	HouseId    int     `json:"houseId" gorm:"column:houseId;type:bigint(20);primary_key"` // 房间ID
	BuildingId int     `json:"buildingId" gorm:"column:buildingId;type:bigint(20);"`      // 楼栋ID
	Cell       string  `json:"cell" gorm:"column:cell;type:varchar(20);"`                 // 单元号
	Number     string  `json:"number" gorm:"column:number;type:varchar(20);"`             // 房间号
	StbMac     string  `json:"stbMac" gorm:"column:stbMac;type:varchar(20);"`             // 机顶盒MAC地址
	Type       int     `json:"type" gorm:"column:type;type:tinyint(4);"`                  // 房屋类型,0-未知 1-正常 2-危房
	Purpose    int     `json:"purpose" gorm:"column:purpose;type:tinyint(4);"`            // 房屋用途，0-空置 1-自住 2-厂房 3-公寓 4-出租
	PhoneNum   string  `json:"phoneNum" gorm:"column:phoneNum;type:varchar(13);"`         // 联系电话
	OwnerId    int     `json:"ownerId" gorm:"column:ownerId;type:bigint(20);"`            // 产权人ID
	UserId     int     `json:"userId" gorm:"column:userId;type:bigint(20);"`              // 使用人
	Area       float32 `json:"area" gorm:"column:area;type:float;"`                       // 建筑面积
	PeopleNum  int     `json:"peopleNum" gorm:"column:peopleNum;type:tinyint(4);"`        // 居住人口
	Remarks    string  `json:"remarks" gorm:"column:remarks;type:tinytext;"`              // 备注
	CreateBy   string  `json:"createBy" gorm:"type:int(11);"`                             // 创建人
	UpdateBy   string  `json:"updateBy" gorm:"type:int(11);"`                             // 更新人
	DataScope  string  `json:"dataScope" gorm:"-"`
	Params     string  `json:"params"  gorm:"-"`
	BaseModel
	BuildingName string `json:"buildingName" gorm:"-"` // 楼栋名称-前端展示需要
	Owner        string `json:"owner" gorm:"-"`        // 产权人名称-前端展示需要
	User         string `json:"user" gorm:"-"`         // 使用人名称-前端展示需要
}

func (GridBasicHousing) TableName() string {
	return "grid_basic_housing"
}

// 创建GridBasicHousing
func (e *GridBasicHousing) Create() (GridBasicHousing, error) {
	var doc GridBasicHousing
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取GridBasicHousing
func (e *GridBasicHousing) Get() (GridBasicHousing, error) {
	var doc GridBasicHousing

	table := orm.Eloquent.Table(e.TableName()).Where("houseId = ?", e.HouseId)

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取GridBasicHousing带分页
func (e *GridBasicHousing) GetPage(pageSize int, pageIndex int) ([]GridBasicHousing, int, error) {
	var doc []GridBasicHousing

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

// 更新GridBasicHousing
func (e *GridBasicHousing) Update(id int) (update GridBasicHousing, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("houseId = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 删除GridBasicHousing
func (e *GridBasicHousing) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("houseId = ?", id).Delete(&GridBasicHousing{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *GridBasicHousing) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("houseId in (?)", id).Delete(&GridBasicHousing{}).Error; err != nil {
		return
	}
	Result = true
	return
}
