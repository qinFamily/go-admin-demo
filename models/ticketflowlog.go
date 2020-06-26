package models

import (
	"fmt"
	"go-admin-demo/cache"
	orm "go-admin-demo/database"
	"log"
	"time"
)

type TicketsTicketflowlog struct {
	Id            int                 `json:"id" gorm:"type:bigint(20);primary_key"` // 主键
	CreateTime    time.Time           `json:"create_time" gorm:"type:timestamp;"`    // 创建时间
	UpdateTime    time.Time           `json:"update_time" gorm:"type:timestamp;"`    // 更新时间
	Memo          string              `json:"memo" gorm:"type:text;"`                // 备注
	Suggestion    string              `json:"suggestion" gorm:"type:varchar(140);"`  // 审批意见
	Participant   string              `json:"participant" gorm:"type:varchar(50);"`  // 处理人
	InterveneType int                 `json:"interveneType" gorm:"type:tinyint(1);"` // 干预类型.0: '转交操作',1: '接单操作',2: '评论操作',3: '删除操作',4: '强制关闭操作',5: '强制修改状态操作',6: '撤回',
	StateId       int                 `json:"stateId" gorm:"type:bigint(20);"`       // 当前状态
	TicketId      int                 `json:"ticketId" gorm:"type:bigint(20);"`      // 工单
	TransitionId  int                 `json:"transitionId" gorm:"type:bigint(20);"`  // 流转
	CreateBy      string              `json:"createBy" gorm:"-"`                     // 创建者
	UpdateBy      string              `json:"updateBy" gorm:"-"`                     // 修改者
	Ticket        TicketsTicket       `json:"ticket" gorm:"-"`                       // 工单
	WorkflowState WorkflowsState      `json:"state" gorm:"-"`                        // 当前状态
	Transition    WorkflowsTransition `json:"transition" gorm:"-"`                   // 流转
	DataScope     string              `json:"dataScope" gorm:"-"`
	Params        string              `json:"params"  gorm:"-"`
	BaseModel
}

func (TicketsTicketflowlog) TableName() string {
	return "tickets_ticketflowlog"
}

// 创建TicketsTicketflowlog
func (e *TicketsTicketflowlog) Create() (TicketsTicketflowlog, error) {
	var doc TicketsTicketflowlog
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

// 获取TicketsTicketflowlog
func (e *TicketsTicketflowlog) Get() (doc TicketsTicketflowlog, err error) {
	key := fmt.Sprintf("ttfl:get:%s:%s", e.CreateBy, e.DataScope)
	getter := func() (interface{}, error) {
		table := orm.Eloquent.Table(e.TableName())

		if err := table.First(&doc).Error; err != nil {
			return doc, err
		}

		wfs := WorkflowsState{
			ID: doc.StateId,
		}
		wfs, err = wfs.Get(false, 1)
		if err != nil {
			log.Println("==================== get WorkflowsState error", err)
			return doc, err
		}

		wft := WorkflowsTransition{
			ID: doc.TransitionId,
		}
		wft, err = wft.Get(false)
		if err != nil {
			log.Println("==================== get WorkflowsTransition error", err)
			return doc, err
		}
		tt := TicketsTicket{
			Id: doc.TicketId,
		}
		tt, err = tt.Get(true)
		if err != nil {
			log.Println("==================== get TicketsTicket error", err)
			return doc, err
		}
		doc.WorkflowState = wfs
		doc.Transition = wft
		doc.Ticket = tt
		return doc, nil
	}
	val, err := cache.LRU().GetWithLoader(key, getter)
	if val != nil {
		doc = val.(TicketsTicketflowlog)
	}
	return
}

// 获取TicketsTicketflowlog带分页
func (e *TicketsTicketflowlog) GetPage(pageSize int, pageIndex int) (doc []TicketsTicketflowlog, count int, err error) {

	key := fmt.Sprintf("ttfl:getp:%s:%s", e.CreateBy, e.DataScope)
	getter := func() (interface{}, error) {
		table := orm.Eloquent.Select("*").Table(e.TableName())
		if e.TicketId > 0 {
			table = table.Where("ticket_id = ?", e.TicketId)
		}
		// 数据权限控制(如果不需要数据权限请将此处去掉)
		// dataPermission := new(DataPermission)
		// dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
		// table = dataPermission.GetDataScope(e.TableName(), table)

		if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
			return nil, err
		}
		returnDoc := make([]TicketsTicketflowlog, 0)

		for i, d := range doc {
			wfs := WorkflowsState{
				ID: d.StateId,
			}
			wfs, err = wfs.Get(false, 1)
			if err != nil {
				log.Println("==================== get WorkflowsState error", err)
				//return doc, err
				continue
			}

			wft := WorkflowsTransition{
				ID: d.TransitionId,
			}
			wft, err = wft.Get(false)
			if err != nil {
				log.Println("==================== get WorkflowsTransition error", err)
				continue
				// return doc, err
			}
			tt := TicketsTicket{
				Id: d.TicketId,
			}
			tt, err = tt.Get(true)
			if err != nil {
				log.Println("==================== get TicketsTicket error", err)
				continue
				// return doc, err
			}
			doc[i].Ticket = tt
			doc[i].WorkflowState = wfs
			doc[i].Transition = wft
			returnDoc = append(returnDoc, doc[i])
		}
		table.Count(&count)
		return returnDoc, nil
	}
	val, err := cache.LRU().GetWithLoader(key, getter)
	if val != nil {
		doc = val.([]TicketsTicketflowlog)
		count = len(doc)
	}
	return
}

// 更新TicketsTicketflowlog
func (e *TicketsTicketflowlog) Update(id int) (update TicketsTicketflowlog, err error) {
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

// 删除TicketsTicketflowlog
func (e *TicketsTicketflowlog) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&TicketsTicketflowlog{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *TicketsTicketflowlog) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&TicketsTicketflowlog{}).Error; err != nil {
		return
	}
	Result = true
	return
}
