package models

import (
	"fmt"
	"go-admin-demo/cache"
	orm "go-admin-demo/database"
)

// WorkflowsStateRoleParticipant [...]
type WorkflowsStateRoleParticipant struct {
	ID             int            `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	StateID        int            `gorm:"unique_index:workflows_state_role_participant_state_id_role_id_uniq;index;column:state_id;type:int(11);not null" json:"state_id"`
	WorkflowsState WorkflowsState `gorm:"association_foreignkey:state_id;foreignkey:id" json:"workflows_state_list"`
	RoleID         int            `gorm:"unique_index:workflows_state_role_participant_state_id_role_id_uniq;index;column:role_id;type:int(11);not null" json:"role_id"`
	SysRole        SysRole        `gorm:"association_foreignkey:role_id;foreignkey:role_id" json:"sys_role_list"`
	DataScope      string         `json:"-" gorm:"-"`
	Params         string         `json:"-"  gorm:"-"`
	BaseModel
}

// TableName get sql table name.获取数据库名字
func (WorkflowsStateRoleParticipant) TableName() string {
	return "workflows_state_role_participant"
}

func (w *WorkflowsStateRoleParticipant) Create() (WorkflowsStateRoleParticipant, error) {
	var doc WorkflowsStateRoleParticipant
	result := orm.Eloquent.Table(w.TableName()).Create(&w)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *w
	return doc, nil
}

// Get 获取
func (w *WorkflowsStateRoleParticipant) Get(isRelated bool) (result WorkflowsStateRoleParticipant, err error) {

	key := fmt.Sprintf("wfsrp:get:%+v:%d:%d:%d", isRelated, w.ID, w.StateID, w.RoleID)

	getter := func() (interface{}, error) {
		table := orm.Eloquent.Table(w.TableName())
		if w.ID != 0 {
			table = table.Where("id = ?", w.ID)
		}

		if w.StateID != 0 {
			table = table.Where("state_id = ?", w.StateID)
		}
		if w.RoleID != 0 {
			table = table.Where("role_id = ?", w.RoleID)
		}

		if err = table.First(&result).Error; err != nil {
			return result, err
		}
		if isRelated {
			info := &WorkflowsState{
				ID: result.StateID,
			}
			if wt, err := info.Get(false, 2); err == nil {
				result.WorkflowsState = wt
			}
			f := &SysRole{
				RoleId: result.RoleID,
			}
			if wt, err := f.Get(); err == nil {
				result.SysRole = wt
			}
		}

		return result, err
	}
	val, err := cache.LRU().GetWithLoader(key, getter)
	if val != nil {
		result = val.(WorkflowsStateRoleParticipant)
	}
	return
}

// Gets 获取批量结果
func (w *WorkflowsStateRoleParticipant) GetPage(pageSize int, pageIndex int, isRelated bool) (results []WorkflowsStateRoleParticipant, count int, err error) {

	key := fmt.Sprintf("wfsrp:getp:%d:%d:%+v:%d:%d", pageSize, pageIndex, isRelated, w.StateID, w.RoleID)
	getter := func() (interface{}, error) {
		table := orm.Eloquent.Select("*").Table(w.TableName())

		if w.StateID != 0 {
			table = table.Where("state_id = ?", w.StateID)
		}
		if w.RoleID != 0 {
			table = table.Where("role_id = ?", w.RoleID)
		}

		// 数据权限控制(如果不需要数据权限请将此处去掉)
		//dataPermission := new(DataPermission)
		//dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
		//table = dataPermission.GetDataScope(e.TableName(), table)
		// log.Println("pageIndex", pageIndex, "pageSize", pageSize, "Offset", (pageIndex-1)*pageSize)

		if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&results).Error; err != nil {
			return results, err
		}
		if isRelated {
			for i, r := range results {
				info := &WorkflowsState{
					ID: r.StateID,
				}
				if wt, err := info.Get(false, 2); err == nil {
					results[i].WorkflowsState = wt
				}
				f := &SysRole{
					RoleId: r.RoleID,
				}
				if wt, err := f.Get(); err == nil {
					results[i].SysRole = wt
				}
			}
		}
		table.Count(&count)
		return results, err
	}

	val, err := cache.LRU().GetWithLoader(key, getter)
	if val != nil {
		results = val.([]WorkflowsStateRoleParticipant)
		count = len(results)
	}
	return

}

// 更新WorkflowsState
func (e *WorkflowsStateRoleParticipant) Update(id int) (update WorkflowsStateRoleParticipant, err error) {
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
func (e *WorkflowsStateRoleParticipant) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&WorkflowsStateRoleParticipant{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *WorkflowsStateRoleParticipant) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&WorkflowsStateRoleParticipant{}).Error; err != nil {
		return
	}
	Result = true
	return
}
