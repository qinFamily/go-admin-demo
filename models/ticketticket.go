package models

import (
	"fmt"
	"go-admin-demo/cache"
	orm "go-admin-demo/database"
	"go-admin-demo/tools"
	"log"
	_ "time"
)

type TicketsTicket struct {
	Id            int                 `json:"id" gorm:"type:bigint(20);primary_key"` // 主键
	CreateTime    string              `json:"createTime" gorm:"type:timestamp;"`     // 创建时间
	UpdateTime    string              `json:"updateTime" gorm:"type:timestamp;"`     // 更新时间
	Memo          string              `json:"memo" gorm:"type:text;"`                // 备注
	Name          string              `json:"name" gorm:"type:varchar(112);"`        // 名称
	Sn            string              `json:"sn" gorm:"type:varchar(25);"`           // 序列
	Participant   string              `json:"participant" gorm:"type:varchar(50);"`  // 参与者
	Customfield   string              `json:"customfield" gorm:"type:text;"`         // 自定义字段
	CreatedAt     string              `json:"createdAt" gorm:"type:timestamp;"`      // 创建时间
	UpdatedAt     string              `json:"updatedAt" gorm:"type:timestamp;"`      // 更新时间
	DeletedAt     string              `json:"deletedAt" gorm:"type:timestamp;"`      // 删除时间
	WorkflowId    int                 `json:"workflowId" gorm:"type:bigint(20);"`    // 工作流
	TransitionId  int                 `json:"transitionId" gorm:"type:bigint(20);"`  // 工作流程
	StateId       int                 `json:"stateId" gorm:"type:bigint(20);"`       // 工作流状态
	CreateBy      int                 `json:"createBy" gorm:"type:tinyint(4);"`      // 创建者
	UpdateBy      int                 `json:"updateBy" gorm:"type:tinyint(4);"`      // 修改者
	CreateUser    SysUserB            `json:"create_user" gorm:"-"`                  // 用户
	Workflow      WorkflowsWorkflow   `json:"workflow" gorm:"-"`                     // 工作流
	WorkflowState WorkflowsState      `json:"state" gorm:"-"`                        // 当前状态
	Transition    WorkflowsTransition `json:"transition" gorm:"-"`                   // 流转
	DataScope     string              `json:"dataScope" gorm:"-"`
	Params        string              `json:"params"  gorm:"-"`
	BaseModel
}

func (TicketsTicket) TableName() string {
	return "tickets_ticket"
}

// 创建TicketsTicket
func (e *TicketsTicket) Create() (TicketsTicket, error) {
	var doc TicketsTicket
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取TicketsTicket
func (e *TicketsTicket) Get() (doc TicketsTicket, err error) {
	if e.CreateBy == 0 {
		e.CreateBy, err = tools.StringToInt(e.DataScope)
		if err != nil {
			return TicketsTicket{}, err
		}
	}
	key := fmt.Sprintf("tt:get:%d:%s", e.CreateBy, e.DataScope)
	getter := func() (interface{}, error) {
		table := orm.Eloquent.Table(e.TableName())
		if err = table.First(&doc).Error; err != nil {
			return doc, err
		}
		wfs := WorkflowsState{
			ID: doc.StateId,
		}
		wfs, err = wfs.Get(false, 1)
		if err != nil {
			return doc, err
		}
		doc.WorkflowState = wfs

		user := SysUser{}
		user.UserId = doc.CreateBy
		userv, err := user.Get()
		if err != nil {
			return doc, err
		}
		doc.CreateUser = userv.SysUserB

		wf := WorkflowsWorkflow{
			ID: doc.WorkflowId,
		}
		wf, err = wf.Get(false)
		if err != nil {
			return doc, err
		}
		doc.Workflow = wf

		wft := WorkflowsTransition{
			ID: doc.TransitionId,
		}
		wft, err = wft.Get(false)
		if err != nil {
			log.Println("==================== get WorkflowsTransition error", err)
			return doc, err
		}
		doc.Transition = wft
		return doc, nil
	}
	val, err := cache.LRU().GetWithLoader(key, getter)
	if val != nil {
		doc = val.(TicketsTicket)
	}
	return
}

// 获取TicketsTicket带分页
func (e *TicketsTicket) GetPage(pageSize int, pageIndex int, transitionLt string) (doc []TicketsTicket, count int, err error) {

	if e.CreateBy == 0 {
		e.CreateBy, err = tools.StringToInt(e.DataScope)
		if err != nil {
			return
		}
	}
	key := fmt.Sprintf("tt:getp:%d:%s:%s", e.CreateBy, e.DataScope, transitionLt)
	getter := func() (interface{}, error) {
		table := orm.Eloquent.Select("*").Table(e.TableName())

		// 数据权限控制(如果不需要数据权限请将此处去掉)
		dataPermission := new(DataPermission)
		dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
		table = dataPermission.GetDataScope(e.TableName(), table)
		if len(transitionLt) > 0 {
			table.Where("transition_id < " + transitionLt)
		}
		if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
			return nil, err
		}
		table.Count(&count)
		// log.Println("==================== get doc count", count)
		// log.Println("==================== get doc", fmt.Sprintf("%+v", doc))
		returnDoc := make([]TicketsTicket, 0)

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
			user := SysUser{}
			user.UserId = d.CreateBy
			userv, err := user.Get()
			if err != nil {
				log.Println("==================== get Sysuser error", err)
				continue
				// return doc, err
			}

			wf := WorkflowsWorkflow{
				ID: d.WorkflowId,
			}
			wf, err = wf.Get(false)
			if err != nil {
				log.Println("==================== get WorkflowsWorkflow error", err)
				continue
				// return doc, err
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
			doc[i].WorkflowState = wfs
			doc[i].CreateUser = userv.SysUserB
			doc[i].Workflow = wf
			doc[i].Transition = wft
			returnDoc = append(returnDoc, doc[i])
		}
		return returnDoc, nil
	}
	val, err := cache.LRU().GetWithLoader(key, getter)
	if val != nil {
		doc = val.([]TicketsTicket)
		count = len(doc)
	}
	return
}

// 更新TicketsTicket
func (e *TicketsTicket) Update(id int) (update TicketsTicket, err error) {
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

// 删除TicketsTicket
func (e *TicketsTicket) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&TicketsTicket{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *TicketsTicket) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&TicketsTicket{}).Error; err != nil {
		return
	}
	Result = true
	return
}
