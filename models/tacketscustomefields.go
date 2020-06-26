package models

import (
	"fmt"
	"go-admin-demo/cache"
	orm "go-admin-demo/database"
	"go-admin-demo/tools"
	_ "time"
)

type TicketsTicketcustomfield struct {
	Id            int    `json:"id" gorm:"type:bigint(20);primary_key"` // 主键
	CreateTime    string `json:"createTime" gorm:"type:timestamp;"`     // 创建时间
	UpdateTime    string `json:"updateTime" gorm:"type:timestamp;"`     // 更新时间
	Memo          string `json:"memo" gorm:"type:text;"`                // 备注
	FieldValue    string `json:"fieldValue" gorm:"type:text;"`          // 字段值
	CreatedAt     string `json:"createdAt" gorm:"type:timestamp;"`      // 创建时间
	UpdatedAt     string `json:"updatedAt" gorm:"type:timestamp;"`      // 更新时间
	DeletedAt     string `json:"deletedAt" gorm:"type:timestamp;"`      // 删除时间
	CustomfieldId int    `json:"customfieldId" gorm:"type:bigint(20);"` // 字段
	TicketId      int    `json:"ticketId" gorm:"type:bigint(20);"`      // 工单
	CreateBy      string `json:"createBy" gorm:"-"`                     // 创建者
	UpdateBy      string `json:"updateBy" gorm:"-"`                     // 修改者
	DataScope     string `json:"dataScope" gorm:"-"`
	Params        string `json:"params"  gorm:"-"`
	BaseModel
}

func (TicketsTicketcustomfield) TableName() string {
	return "tickets_ticketcustomfield"
}

// 创建TicketsTicketcustomfield
func (e *TicketsTicketcustomfield) Create() (TicketsTicketcustomfield, error) {
	var doc TicketsTicketcustomfield
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取TicketsTicketcustomfield
func (e *TicketsTicketcustomfield) Get() (doc TicketsTicketcustomfield, err error) {
	key := fmt.Sprintf("ttcf:get:%s:%s", e.CreateBy, e.DataScope)
	getter := func() (interface{}, error) {
		table := orm.Eloquent.Table(e.TableName())

		if err = table.First(&doc).Error; err != nil {
			return doc, err
		}
		return doc, nil
	}
	val, err := cache.LRU().GetWithLoader(key, getter)
	if val != nil {
		doc = val.(TicketsTicketcustomfield)
	}
	return
}

// 获取TicketsTicketcustomfield带分页
func (e *TicketsTicketcustomfield) GetPage(pageSize int, pageIndex int) (doc []TicketsTicketcustomfield, count int, err error) {

	key := fmt.Sprintf("ttcf:getp:%s:%s", e.CreateBy, e.DataScope)
	getter := func() (interface{}, error) {
		table := orm.Eloquent.Select("*").Table(e.TableName())

		// 数据权限控制(如果不需要数据权限请将此处去掉)
		dataPermission := new(DataPermission)
		dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
		table = dataPermission.GetDataScope(e.TableName(), table)

		if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
			return nil, err
		}
		table.Count(&count)
		return doc, nil
	}
	val, err := cache.LRU().GetWithLoader(key, getter)
	if val != nil {
		doc = val.([]TicketsTicketcustomfield)
		count = len(doc)
	}
	return
}

// 更新TicketsTicketcustomfield
func (e *TicketsTicketcustomfield) Update(id int) (update TicketsTicketcustomfield, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 删除TicketsTicketcustomfield
func (e *TicketsTicketcustomfield) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&TicketsTicketcustomfield{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *TicketsTicketcustomfield) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&TicketsTicketcustomfield{}).Error; err != nil {
		return
	}
	Result = true
	return
}
