package models

import (
	"fmt"
	"go-admin-demo/cache"
	orm "go-admin-demo/database"
)

// WorkflowsStateUserParticipant [...]
type WorkflowsStateUserParticipant struct {
	ID             int            `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	StateID        int            `gorm:"unique_index:workflows_state_user_participant_state_id_user_id_uniq;index;column:state_id;type:int(11);not null" json:"state_id"`
	WorkflowsState WorkflowsState `gorm:"association_foreignkey:state_id;foreignkey:id" json:"workflows_state_list"`
	UserID         int            `gorm:"unique_index:workflows_state_user_participant_state_id_user_id_uniq;index;column:user_id;type:int(11);not null" json:"user_id"`
	SysUser        SysUserB       `gorm:"association_foreignkey:user_id;foreignkey:user_id" json:"sys_user_list"`
	DataScope      string         `json:"-" gorm:"-"`
	Params         string         `json:"-"  gorm:"-"`
	BaseModel
}

// TableName get sql table name.获取数据库名字
func (WorkflowsStateUserParticipant) TableName() string {
	return "workflows_state_user_participant"
}

func (w *WorkflowsStateUserParticipant) Create() (WorkflowsStateUserParticipant, error) {
	var doc WorkflowsStateUserParticipant
	result := orm.Eloquent.Table(w.TableName()).Create(&w)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *w

	// 更新缓存
	wfsupTrueKey := fmt.Sprintf("wfsup:get:%+v:%d:%d:%d", true, w.ID, w.StateID, w.UserID)
	if _, err1 := cache.LRU().Get(wfsupTrueKey); err1 == nil {

		info := &WorkflowsState{
			ID: doc.StateID,
		}
		if wt, err2 := info.Get(false, 2); err2 == nil {
			doc.WorkflowsState = wt
		}
		f := &SysUser{}
		f.UserId = doc.UserID
		if wt, err3 := f.Get(); err3 == nil {
			doc.SysUser = wt.SysUserB
		}
		cache.LRU().Set(wfsupTrueKey, doc)
	}
	wfsupFalseKey := fmt.Sprintf("wfsup:get:%+v:%d:%d:%d", false, w.ID, w.StateID, w.UserID)
	if _, err4 := cache.LRU().Get(wfsupTrueKey); err4 == nil {
		cache.LRU().Set(wfsupFalseKey, doc)
	}

	wfsupgetpkeyTrue := fmt.Sprintf("wfsup:getp:%d:%d:%+v:%d:%d", 20, 1, true, w.StateID, w.UserID)
	if result, err5 := cache.LRU().Get(wfsupgetpkeyTrue); err5 == nil {
		if resultA, ok := result.([]WorkflowsStateUserParticipant); ok {
			resultA = append(resultA, doc)
			cache.LRU().Set(wfsupgetpkeyTrue, resultA)
		}
	}

	wfsupgetpkeyFalse := fmt.Sprintf("wfsup:getp:%d:%d:%+v:%d:%d", 20, 1, false, w.StateID, w.UserID)
	if result, err6 := cache.LRU().Get(wfsupgetpkeyFalse); err6 == nil {
		if resultA, ok := result.([]WorkflowsStateUserParticipant); ok {
			resultA = append(resultA, doc)
			cache.LRU().Set(wfsupgetpkeyFalse, resultA)
		}
	}

	return doc, nil
}

// Get 获取
func (w *WorkflowsStateUserParticipant) Get(isRelated bool) (result WorkflowsStateUserParticipant, err error) {

	key := fmt.Sprintf("wfsup:get:%+v:%d:%d:%d", isRelated, w.ID, w.StateID, w.UserID)

	getter := func() (interface{}, error) {
		table := orm.Eloquent.Table(w.TableName())
		if w.ID != 0 {
			table = table.Where("id = ?", w.ID)
		}

		if w.StateID != 0 {
			table = table.Where("state_id = ?", w.StateID)
		}
		if w.UserID != 0 {
			table = table.Where("user_id = ?", w.UserID)
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
			f := &SysUser{}
			f.UserId = result.UserID
			if wt, err := f.Get(); err == nil {
				result.SysUser = wt.SysUserB
			}
		}

		return result, err
	}
	val, err := cache.LRU().GetWithLoader(key, getter)
	if val != nil {
		result = val.(WorkflowsStateUserParticipant)
	}
	return
}

// Gets 获取批量结果
func (w *WorkflowsStateUserParticipant) GetPage(pageSize int, pageIndex int, isRelated bool) (results []WorkflowsStateUserParticipant, count int, err error) {

	key := fmt.Sprintf("wfsup:getp:%d:%d:%+v:%d:%d", pageSize, pageIndex, isRelated, w.StateID, w.UserID)
	getter := func() (interface{}, error) {
		table := orm.Eloquent.Select("*").Table(w.TableName())

		if w.StateID != 0 {
			table = table.Where("state_id = ?", w.StateID)
		}
		if w.UserID != 0 {
			table = table.Where("user_id = ?", w.UserID)
		}

		// 数据权限控制(如果不需要数据权限请将此处去掉)
		//dataPermission := new(DataPermission)
		//dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
		//table = dataPermission.GetDataScope(e.TableName(), table)

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
				f := &SysUser{}
				f.UserId = r.UserID
				if wt, err := f.Get(); err == nil {
					results[i].SysUser = wt.SysUserB
				}
			}
		}
		table.Count(&count)
		return results, err
	}

	val, err := cache.LRU().GetWithLoader(key, getter)
	if val != nil {
		results = val.([]WorkflowsStateUserParticipant)
		count = len(results)
	}
	return

}

// 更新WorkflowsState
func (e *WorkflowsStateUserParticipant) Update(id int) (update WorkflowsStateUserParticipant, err error) {
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
func (e *WorkflowsStateUserParticipant) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&WorkflowsStateUserParticipant{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *WorkflowsStateUserParticipant) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&WorkflowsStateUserParticipant{}).Error; err != nil {
		return
	}
	Result = true
	return
}
