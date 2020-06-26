package models

import (
	"fmt"
	"go-admin-demo/cache"
	orm "go-admin-demo/database"
	"go-admin-demo/tools"
	"log"
	"time"
)

type TicketsTicketcustomfield struct {
	Id            int                  `json:"id" gorm:"type:bigint(20);primary_key"` // 主键
	CreateTime    time.Time            `json:"create_time" gorm:"type:timestamp;"`    // 创建时间
	UpdateTime    time.Time            `json:"update_time" gorm:"type:timestamp;"`    // 更新时间
	Memo          string               `json:"memo" gorm:"type:text;"`                // 备注
	FieldValue    string               `json:"field_value" gorm:"type:text;"`         // 字段值
	CustomfieldId int                  `json:"-" gorm:"type:bigint(20);"`             // 字段
	TicketId      int                  `json:"-" gorm:"type:bigint(20);"`             // 工单
	CreateBy      string               `json:"createBy" gorm:"-"`                     // 创建者
	UpdateBy      string               `json:"updateBy" gorm:"-"`                     // 修改者
	WorkflowId    int                  `json:"-" gorm:"-"`                            // 工作流ID
	SrcStateId    int                  `json:"-" gorm:"-"`                            // source_stateID
	Ticket        TicketsTicket        `json:"ticket" gorm:"-"`                       // 工单
	Customfield   WorkflowsCustomfield `json:"customfield" gorm:"-"`                  // 可编辑字段
	DataScope     string               `json:"dataScope" gorm:"-"`
	Params        string               `json:"params"  gorm:"-"`
	BaseModel
}

func (TicketsTicketcustomfield) TableName() string {
	return "tickets_ticketcustomfield"
}

// 创建TicketsTicketcustomfield
func (e *TicketsTicketcustomfield) Create() (TicketsTicketcustomfield, error) {
	var doc TicketsTicketcustomfield
	e.CreateTime = time.Now()
	e.UpdateTime = e.CreateTime
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
		createBy := 1
		if len(e.DataScope) > 0 {
			createBy, err = tools.StringToInt(e.DataScope)
			if err != nil {
				log.Println("========", err, "e.DataScope", e.DataScope)
				return doc, err
			}
		}

		tt := TicketsTicket{
			Id:        doc.TicketId,
			CreateBy:  createBy,
			DataScope: e.DataScope,
		}
		tt, err = tt.Get(false)
		if err != nil {
			log.Println("==================== get TicketsTicket error", err)
			return doc, err
		}

		wfc := WorkflowsCustomfield{
			ID: doc.CustomfieldId,
		}
		wfc, err = wfc.Get(false)
		if err != nil {
			log.Println("==================== get WorkflowsCustomfield error", err)
			return doc, err
		}

		doc.Customfield = wfc
		doc.Ticket = tt

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
		if e.TicketId > 0 {
			table = table.Where("ticket_id = ?", e.TicketId)
		}
		// log.Println("==============", table.QueryExpr())
		// 数据权限控制(如果不需要数据权限请将此处去掉)
		// dataPermission := new(DataPermission)
		// dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
		// table = dataPermission.GetDataScope(e.TableName(), table)

		if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
			return nil, err
		}

		returnDoc := make([]TicketsTicketcustomfield, 0)
		for i, d := range doc {

			createBy := 1
			if len(e.DataScope) > 0 {
				createBy, err = tools.StringToInt(e.DataScope)
				if err != nil {
					log.Println("========", err, "e.DataScope", e.DataScope)
					return doc, err
				}
			}

			tt := TicketsTicket{
				Id:        d.TicketId,
				CreateBy:  createBy,
				DataScope: e.DataScope,
			}
			tt, err = tt.Get(false)
			if err != nil {
				log.Println("==================== get TicketsTicket error", err)
				continue
			}

			wfc := WorkflowsCustomfield{
				ID: d.CustomfieldId,
			}
			wfc, err = wfc.Get(false)
			if err != nil {
				log.Println("==================== get WorkflowsCustomfield error", err)
				continue
			}

			doc[i].Customfield = wfc
			doc[i].Ticket = tt
			returnDoc = append(returnDoc, doc[i])
		}
		table.Count(&count)
		return returnDoc, nil
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
