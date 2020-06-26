package models

import (
	"fmt"
	"go-admin-demo/cache"
	orm "go-admin-demo/database"
	"go-admin-demo/tools"
	_ "time"
)

type TicketsTicketuser struct {
	Id         int    `json:"id" gorm:"type:bigint(20);primary_key"` // 主键
	CreateTime string `json:"createTime" gorm:"type:timestamp;"`     // 创建时间
	UpdateTime string `json:"updateTime" gorm:"type:timestamp;"`     // 更新时间
	Memo       string `json:"memo" gorm:"type:text;"`                // 备注
	Username   string `json:"username" gorm:"type:varchar(100);"`    // 关系人
	InProcess  int    `json:"inProcess" gorm:"type:tinyint(1);"`     // 待处理中
	Worked     int    `json:"worked" gorm:"type:tinyint(1);"`        // 处理过
	CreatedAt  string `json:"createdAt" gorm:"type:timestamp;"`      // 创建时间
	UpdatedAt  string `json:"updatedAt" gorm:"type:timestamp;"`      // 更新时间
	DeletedAt  string `json:"deletedAt" gorm:"type:timestamp;"`      // 删除时间
	TicketId   int    `json:"ticketId" gorm:"type:bigint(20);"`      // 工单
	CreateBy   string `json:"createBy" gorm:"-"`                     // 创建者
	UpdateBy   string `json:"updateBy" gorm:"-"`                     // 修改者
	DataScope  string `json:"dataScope" gorm:"-"`
	Params     string `json:"params"  gorm:"-"`
	BaseModel
}

func (TicketsTicketuser) TableName() string {
	return "tickets_ticketuser"
}

// 创建TicketsTicketuser
func (e *TicketsTicketuser) Create() (TicketsTicketuser, error) {
	var doc TicketsTicketuser
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取TicketsTicketuser
func (e *TicketsTicketuser) Get() (doc TicketsTicketuser, err error) {

	key := fmt.Sprintf("ttu:get:%s:%s", e.Username, e.DataScope)
	getter := func() (interface{}, error) {
		table := orm.Eloquent.Table(e.TableName())
		if err = table.First(&doc).Error; err != nil {
			return doc, err
		}
		return doc, nil
	}
	val, err := cache.LRU().GetWithLoader(key, getter)
	if val != nil {
		doc = val.(TicketsTicketuser)
	}
	return
}

// 获取TicketsTicketuser带分页
func (e *TicketsTicketuser) GetPage(pageSize int, pageIndex int) (doc []TicketsTicketuser, count int, err error) {

	key := fmt.Sprintf("ttu:getp:%s:%s", e.Username, e.DataScope)
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
		doc = val.([]TicketsTicketuser)
		count = len(doc)
	}
	return
}

// 更新TicketsTicketuser
func (e *TicketsTicketuser) Update(id int) (update TicketsTicketuser, err error) {
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

// 删除TicketsTicketuser
func (e *TicketsTicketuser) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&TicketsTicketuser{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *TicketsTicketuser) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&TicketsTicketuser{}).Error; err != nil {
		return
	}
	Result = true
	return
}
