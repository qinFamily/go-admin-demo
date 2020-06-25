package models

import (
	"fmt"
	"go-admin-demo/cache"
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
	IsHidden          bool        `gorm:"column:is_hidden;type:tinyint(4);not null" json:"is_hidden"`
	OrderID           int         `gorm:"column:order_id;type:int(11);not null" json:"order_id"`
	StateType         int         `gorm:"column:state_type;type:varchar(1);not null" json:"state_type"`
	EnableRetreat     bool        `gorm:"column:enable_retreat;type:tinyint(4);not null" json:"enable_retreat"`
	ParticipantType   int         `gorm:"column:participant_type;type:varchar(1);not null" json:"participant_type"`
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
func (w *WorkflowsState) Get(isRelated bool, depth int) (result WorkflowsState, err error) {

	key := fmt.Sprintf("wfs:get:%+v:%d:%d", isRelated, depth, w.ID)
	getter := func() (interface{}, error) {
		table := orm.Eloquent.Table(w.TableName()).Order("order_id")
		if w.ID != 0 {
			table = table.Where("id = ?", w.ID)
		}
		if err = table.First(&result).Error; err != nil {
			return result, err
		}

		wf := WorkflowsWorkflow{
			ID: result.WorkflowID,
		}

		if depth == 1 {
			wf, err = wf.Get(isRelated)
		} else {
			wf, err = wf.Get(false)
		}
		if err == nil {
			if isRelated {
				result.WorkflowsWorkflow = wf
			} else {
				result.WorkflowsWorkflow = wf.ID
			}
		}
		up := &WorkflowsStateUserParticipant{
			StateID: result.ID,
		}
		if wt, n, err := up.GetPage(200, 1, isRelated); err == nil {
			if n > 0 {
				if isRelated {
					users := make([]SysUserB, 0)
					for _, iwt := range wt {
						users = append(users, iwt.SysUser)
					}
					result.UserParticipant = users
				} else {
					wtIDs := make([]int, 0)
					for _, wti := range wt {
						wtIDs = append(wtIDs, wti.UserID)
					}
					result.UserParticipant = wtIDs
				}
			} else {
				result.UserParticipant = []int{}
			}
		}

		gp := &WorkflowsStateGroupParticipant{
			StateID: result.ID,
		}
		if wt, n, err := gp.GetPage(200, 1, isRelated); err == nil {
			wtIDs := make([]int, 0)
			if n > 0 {
				if isRelated {
					depts := make([]Dept, 0)
					for _, iwt := range wt {
						depts = append(depts, iwt.SysDept)
					}
					result.GroupParticipant = depts
				} else {
					for _, wti := range wt {
						wtIDs = append(wtIDs, wti.GroupID)
					}
					result.GroupParticipant = wtIDs
				}
			} else {
				result.GroupParticipant = wtIDs
			}
		}

		rp := &WorkflowsStateRoleParticipant{
			StateID: result.ID,
		}
		if wt, n, err := rp.GetPage(200, 1, isRelated); err == nil {
			wtIDs := make([]int, 0)
			if n > 0 {
				if isRelated {
					wfsr := make([]SysRole, 0)
					for _, iwt := range wt {
						wfsr = append(wfsr, iwt.SysRole)
					}
					result.RoleParticipant = wfsr
				} else {
					for _, wti := range wt {
						wtIDs = append(wtIDs, wti.RoleID)
					}
					result.RoleParticipant = wtIDs
				}
			} else {
				result.RoleParticipant = wtIDs
			}
		}

		f := &WorkflowsStateFields{
			StateID: result.ID,
		}
		if wt, n, err := f.GetPage(200, 1, isRelated); err == nil {
			wtIDs := make([]int, 0)
			if n > 0 {
				if isRelated {
					wff := make([]WorkflowsCustomfield, 0)
					for _, iwt := range wt {
						wff = append(wff, iwt.WorkflowsCustomfield)
					}
					result.Fields = wff
				} else {
					for _, wti := range wt {
						wtIDs = append(wtIDs, wti.CustomfieldID)
					}
					result.Fields = wtIDs
				}
			} else {
				result.Fields = wtIDs
			}
		}
		return result, nil
	}
	val, err := cache.LRU().GetWithLoader(key, getter)
	if val != nil {
		result = val.(WorkflowsState)
	}
	return
}

// Gets 获取批量结果
func (w *WorkflowsState) GetPage(pageSize int, pageIndex int, isRelated bool, depth int) (results []WorkflowsState, count int, err error) {

	key := fmt.Sprintf("wfs:getp:%d:%d:%+v:%d:%d", pageSize, pageIndex, isRelated, depth, w.WorkflowID)

	getter := func() (interface{}, error) {
		table := orm.Eloquent.Select("*").Table(w.TableName()).Order("order_id")
		if w.WorkflowID != 0 {
			table = table.Where("workflow_id = ?", w.WorkflowID)
		}

		// 数据权限控制(如果不需要数据权限请将此处去掉)
		//dataPermission := new(DataPermission)
		//dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
		//table = dataPermission.GetDataScope(e.TableName(), table)
		if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&results).Error; err != nil {
			return nil, err
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
			if wt, n, err := up.GetPage(200, 1, isRelated); err == nil {
				if n > 0 {
					if isRelated {
						users := make([]SysUserB, 0)
						for _, iwt := range wt {
							users = append(users, iwt.SysUser)
						}
						results[i].UserParticipant = users
					} else {
						wtIDs := make([]int, 0)
						for _, wti := range wt {
							wtIDs = append(wtIDs, wti.UserID)
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
			if wt, n, err := gp.GetPage(200, 1, isRelated); err == nil {
				wtIDs := make([]int, 0)
				if n > 0 {
					if isRelated {
						depts := make([]Dept, 0)
						for _, iwt := range wt {
							depts = append(depts, iwt.SysDept)
						}
						results[i].GroupParticipant = depts
					} else {
						for _, wti := range wt {
							wtIDs = append(wtIDs, wti.GroupID)
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
			if wt, n, err := rp.GetPage(200, 1, isRelated); err == nil {
				wtIDs := make([]int, 0)
				if n > 0 {
					if isRelated {
						wfsr := make([]SysRole, 0)
						for _, iwt := range wt {
							wfsr = append(wfsr, iwt.SysRole)
						}
						results[i].RoleParticipant = wfsr
					} else {
						for _, wti := range wt {
							wtIDs = append(wtIDs, wti.RoleID)
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
			if wt, n, err := f.GetPage(200, 1, isRelated); err == nil {
				wtIDs := make([]int, 0)
				if n > 0 {
					if isRelated {
						wff := make([]WorkflowsCustomfield, 0)
						for _, iwt := range wt {
							wff = append(wff, iwt.WorkflowsCustomfield)
						}
						results[i].Fields = wff
					} else {
						for _, wti := range wt {
							wtIDs = append(wtIDs, wti.CustomfieldID)
						}
						results[i].Fields = wtIDs
					}
				} else {
					results[i].Fields = wtIDs
				}
			}
		}
		table.Count(&count)
		return results, err
	}
	val, err := cache.LRU().GetWithLoader(key, getter)
	if val != nil {
		results = val.([]WorkflowsState)
		count = len(results)
	}
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
