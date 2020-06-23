package models

import (
	orm "go-admin-demo/database"
	"time"
)

// WorkflowsState [...]
type WorkflowsState struct {
	ID                int         `gorm:"primary_key;column:id;type:int(11);not null" json:"id"`
	CreateTime        time.Time   `gorm:"column:create_time;type:datetime;not null" json:"create_time"`
	UpdateTime        time.Time   `gorm:"column:update_time;type:datetime;not null" json:"update_time"`
	Memo              string      `gorm:"column:memo;type:text;not null" json:"memo"`
	Name              string      `gorm:"column:name;type:varchar(50);not null" json:"name"`
	IsHidden          int8        `gorm:"column:is_hidden;type:tinyint(4);not null" json:"is_hidden"`
	OrderID           int         `gorm:"column:order_id;type:int(11);not null" json:"order_id"`
	StateType         string      `gorm:"column:state_type;type:varchar(1);not null" json:"state_type"`
	EnableRetreat     int8        `gorm:"column:enable_retreat;type:tinyint(4);not null" json:"enable_retreat"`
	ParticipantType   string      `gorm:"column:participant_type;type:varchar(1);not null" json:"participant_type"`
	WorkflowID        int         `gorm:"index;column:workflow_id;type:int(11);not null" json:"-"`
	WorkflowsWorkflow interface{} `gorm:"association_foreignkey:workflow_id;foreignkey:id" json:"workflow"`
	UserParticipant   interface{} `gorm:"-" json:"user_participant"`
	GroupParticipant  interface{} `gorm:"-" json:"group_participant"`
	RoleParticipant   interface{} `gorm:"-" json:"role_participant"`
	Fields            interface{} `gorm:"-" json:"fields"`
	DataScope         string      `json:"-" gorm:"-"`
	Params            string      `json:"-"  gorm:"-"`
	BaseModel
}

// GetTableName get sql table name.获取数据库名字
func (WorkflowsState) TableName() string {
	return "workflows_state"
}

func (w *WorkflowsState) Create() (WorkflowsState, error) {
	var doc WorkflowsState
	result := orm.Eloquent.Table(w.TableName()).Create(&w)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *w
	return doc, nil
}

// Get 获取
func (w *WorkflowsState) Get(isRelated bool) (WorkflowsState, error) {

	var wfs WorkflowsState

	table := orm.Eloquent.Table(w.TableName())
	if w.ID != 0 {
		table = table.Where("id = ?", w.ID)
	}
	if err := table.First(&wfs).Error; err != nil {
		return wfs, err
	}

	wf := &WorkflowsWorkflow{
		ID: wfs.WorkflowID,
	}
	if wt, err := wf.Get(isRelated); err == nil {
		if isRelated {
			wfs.WorkflowsWorkflow = wt
		} else {
			wfs.WorkflowsWorkflow = wt.ID
		}
	}
	up := &WorkflowsStateUserParticipant{
		StateID: wfs.ID,
	}
	if wt, n, err := up.GetPage(1, 200, isRelated); err == nil {
		if n > 0 {
			if isRelated {
				wfs.UserParticipant = wt
			} else {
				wtIDs := make([]int, 0)
				for _, wti := range wt {
					wtIDs = append(wtIDs, wti.ID)
				}
				wfs.UserParticipant = wtIDs
			}
		} else {
			wfs.UserParticipant = []int{}
		}
	}

	gp := &WorkflowsStateGroupParticipant{
		StateID: wfs.ID,
	}
	if wt, n, err := gp.GetPage(1, 200, isRelated); err == nil {
		wtIDs := make([]int, 0)
		if n > 0 {
			if isRelated {
				wfs.GroupParticipant = wt
			} else {
				for _, wti := range wt {
					wtIDs = append(wtIDs, wti.ID)
				}
				wfs.GroupParticipant = wtIDs
			}
		} else {
			wfs.GroupParticipant = wtIDs
		}
	}

	rp := &WorkflowsStateRoleParticipant{
		StateID: wfs.ID,
	}
	if wt, n, err := rp.GetPage(1, 200, isRelated); err == nil {
		wtIDs := make([]int, 0)
		if n > 0 {
			if isRelated {
				wfs.RoleParticipant = wt
			} else {
				for _, wti := range wt {
					wtIDs = append(wtIDs, wti.ID)
				}
				wfs.RoleParticipant = wtIDs
			}
		} else {
			wfs.RoleParticipant = wtIDs
		}
	}

	f := &WorkflowsStateFields{
		StateID: wfs.ID,
	}
	if wt, n, err := f.GetPage(1, 200, isRelated); err == nil {
		wtIDs := make([]int, 0)
		if n > 0 {
			if isRelated {
				wfs.Fields = wt
			} else {
				for _, wti := range wt {
					wtIDs = append(wtIDs, wti.ID)
				}
				wfs.Fields = wtIDs
			}
		} else {
			wfs.Fields = wtIDs
		}
	}

	return wfs, nil
}

// Gets 获取批量结果
func (w *WorkflowsState) GetPage(pageSize int, pageIndex int, isRelated bool) (results []WorkflowsState, count int, err error) {

	table := orm.Eloquent.Select("*").Table(w.TableName())
	if w.WorkflowID != 0 {
		table = table.Where("workflow_id = ?", w.WorkflowID)
	}

	// 数据权限控制(如果不需要数据权限请将此处去掉)
	//dataPermission := new(DataPermission)
	//dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	//table = dataPermission.GetDataScope(e.TableName(), table)

	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&results).Error; err != nil {
		return nil, 0, err
	}
	for i, r := range results {
		info := &WorkflowsWorkflow{
			ID: r.WorkflowID,
		}
		if wt, err := info.Get(isRelated); err == nil {
			if isRelated {
				results[i].WorkflowsWorkflow = wt
			} else {
				results[i].WorkflowsWorkflow = wt.ID
			}
		}
		up := &WorkflowsStateUserParticipant{
			StateID: r.ID,
		}
		if wt, n, err := up.GetPage(1, 200, isRelated); err == nil {
			if n > 0 {
				if isRelated {
					results[i].UserParticipant = wt
				} else {
					wtIDs := make([]int, 0)
					for _, wti := range wt {
						wtIDs = append(wtIDs, wti.ID)
					}
					results[i].UserParticipant = wtIDs
				}
			} else {
				results[i].UserParticipant = []int{}
			}
		}

		gp := &WorkflowsStateGroupParticipant{
			StateID: r.ID,
		}
		if wt, n, err := gp.GetPage(1, 200, isRelated); err == nil {
			wtIDs := make([]int, 0)
			if n > 0 {
				if isRelated {
					results[i].GroupParticipant = wt
				} else {
					for _, wti := range wt {
						wtIDs = append(wtIDs, wti.ID)
					}
					results[i].GroupParticipant = wtIDs
				}
			} else {
				results[i].GroupParticipant = wtIDs
			}
		}

		rp := &WorkflowsStateRoleParticipant{
			StateID: r.ID,
		}
		if wt, n, err := rp.GetPage(1, 200, isRelated); err == nil {
			wtIDs := make([]int, 0)
			if n > 0 {
				if isRelated {
					results[i].RoleParticipant = wt
				} else {
					for _, wti := range wt {
						wtIDs = append(wtIDs, wti.ID)
					}
					results[i].RoleParticipant = wtIDs
				}
			} else {
				results[i].RoleParticipant = wtIDs
			}
		}

		f := &WorkflowsStateFields{
			StateID: r.ID,
		}
		if wt, n, err := f.GetPage(1, 200, isRelated); err == nil {
			wtIDs := make([]int, 0)
			if n > 0 {
				if isRelated {
					results[i].Fields = wt
				} else {
					for _, wti := range wt {
						wtIDs = append(wtIDs, wti.ID)
					}
					results[i].Fields = wtIDs
				}
			} else {
				results[i].Fields = wtIDs
			}
		}
	}
	table.Count(&count)
	return

}

// 更新WorkflowsState
func (e *WorkflowsState) Update(id int) (update WorkflowsState, err error) {
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

// 删除WorkflowsState
func (e *WorkflowsState) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&WorkflowsState{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *WorkflowsState) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&WorkflowsState{}).Error; err != nil {
		return
	}
	Result = true
	return
}
