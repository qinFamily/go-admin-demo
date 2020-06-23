package models

import (
	orm "go-admin-demo/database"
)

// WorkflowsStateFields [...]
type WorkflowsStateFields struct {
	ID                   int                  `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	StateID              int                  `gorm:"unique_index:workflows_state_fields_state_id_customfield_id_uniq;index;column:state_id;type:int(11);not null" json:"state_id"`
	WorkflowsState       WorkflowsState       `gorm:"association_foreignkey:state_id;foreignkey:id" json:"workflows_state_list"`
	CustomfieldID        int                  `gorm:"unique_index:workflows_state_fields_state_id_customfield_id_uniq;index;column:customfield_id;type:int(11);not null" json:"customfield_id"`
	WorkflowsCustomfield WorkflowsCustomfield `gorm:"association_foreignkey:customfield_id;foreignkey:id" json:"workflows_customfield_list"`
	DataScope            string               `json:"-" gorm:"-"`
	Params               string               `json:"-"  gorm:"-"`
	BaseModel
}

// TableName get sql table name.获取数据库名字
func (WorkflowsStateFields) TableName() string {
	return "workflows_state_fields"
}

func (w *WorkflowsStateFields) Create() (WorkflowsStateFields, error) {
	var doc WorkflowsStateFields
	result := orm.Eloquent.Table(w.TableName()).Create(&w)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *w
	return doc, nil
}

// Get 获取
func (w *WorkflowsStateFields) Get(isRelated bool) (WorkflowsStateFields, error) {

	var wft WorkflowsStateFields

	table := orm.Eloquent.Table(w.TableName())
	if w.ID != 0 {
		table = table.Where("id = ?", w.ID)
	}

	if w.StateID != 0 {
		table = table.Where("state_id = ?", w.StateID)
	}
	if w.CustomfieldID != 0 {
		table = table.Where("customfield_id = ?", w.StateID)
	}

	if err := table.First(&wft).Error; err != nil {
		return wft, err
	}
	if isRelated {
		info := &WorkflowsState{
			ID: wft.StateID,
		}
		if wt, err := info.Get(false); err == nil {
			wft.WorkflowsState = wt
		}
		f := &WorkflowsCustomfield{
			ID: wft.CustomfieldID,
		}
		if wt, err := f.Get(false); err == nil {
			wft.WorkflowsCustomfield = wt
		}
	}

	return wft, nil
}

// Gets 获取批量结果
func (w *WorkflowsStateFields) GetPage(pageSize int, pageIndex int, isRelated bool) (results []WorkflowsStateFields, count int, err error) {

	table := orm.Eloquent.Select("*").Table(w.TableName())

	if w.StateID != 0 {
		table = table.Where("state_id = ?", w.StateID)
	}
	if w.CustomfieldID != 0 {
		table = table.Where("customfield_id = ?", w.StateID)
	}

	// 数据权限控制(如果不需要数据权限请将此处去掉)
	//dataPermission := new(DataPermission)
	//dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	//table = dataPermission.GetDataScope(e.TableName(), table)

	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&results).Error; err != nil {
		return nil, 0, err
	}
	if isRelated {
		for i, r := range results {
			info := &WorkflowsState{
				ID: r.StateID,
			}
			if wt, err := info.Get(false); err == nil {
				results[i].WorkflowsState = wt
			}
			f := &WorkflowsCustomfield{
				ID: r.CustomfieldID,
			}
			if wt, err := f.Get(false); err == nil {
				results[i].WorkflowsCustomfield = wt
			}
		}
	}
	table.Count(&count)
	return

}

// 更新WorkflowsState
func (e *WorkflowsStateFields) Update(id int) (update WorkflowsStateFields, err error) {
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
func (e *WorkflowsStateFields) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&WorkflowsStateFields{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *WorkflowsStateFields) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&WorkflowsStateFields{}).Error; err != nil {
		return
	}
	Result = true
	return
}
