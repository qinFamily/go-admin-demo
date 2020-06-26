package models

import (
	"fmt"
	"go-admin-demo/cache"
	orm "go-admin-demo/database"
	"time"
)

// WorkflowsTransition [...]
type WorkflowsTransition struct {
	ID                   int               `gorm:"primary_key;column:id;type:int(11);not null" json:"id"`
	CreateTime           time.Time         `gorm:"column:create_time;type:datetime;not null" json:"create_time"`
	UpdateTime           time.Time         `gorm:"column:update_time;type:datetime;not null" json:"update_time"`
	Memo                 string            `gorm:"column:memo;type:text;not null" json:"memo"`
	Name                 int               `gorm:"column:name;type:varchar(1);not null" json:"name"`
	TransitionType       int               `gorm:"column:transition_type;type:varchar(1);not null" json:"transition_type"`
	Timer                int               `gorm:"column:timer;type:int(11);not null" json:"timer"`
	ConditionExpression  string            `gorm:"column:condition_expression;type:text;not null" json:"condition_expression"`
	AttributeType        int               `gorm:"column:attribute_type;type:varchar(1);not null" json:"attribute_type"`
	AlertEnable          bool              `gorm:"column:alert_enable;type:tinyint(4);not null" json:"alert_enable"`
	AlertText            string            `gorm:"column:alert_text;type:varchar(100);not null" json:"alert_text"`
	DestStateID          int               `gorm:"index;column:dest_state_id;type:int(11)" json:"-"`
	WorkflowsDestState   WorkflowsState    `gorm:"association_foreignkey:dest_state_id;foreignkey:id" json:"dest_state"`
	SourceStateID        int               `gorm:"index;column:source_state_id;type:int(11)" json:"-"`
	WorkflowsSourceState WorkflowsState    `gorm:"association_foreignkey:dest_state_id;foreignkey:id" json:"source_state"`
	WorkflowID           int               `gorm:"index;column:workflow_id;type:int(11);not null" json:"-"`
	WorkflowsWorkflow    WorkflowsWorkflow `gorm:"association_foreignkey:workflow_id;foreignkey:id" json:"workflow"`
	DataScope            string            `json:"-" gorm:"-"`
	Params               string            `json:"-"  gorm:"-"`
	BaseModel
}

// GetTableName get sql table name.获取数据库名字
func (WorkflowsTransition) TableName() string {
	return "workflows_transition"
}

func (w *WorkflowsTransition) Create() (WorkflowsTransition, error) {
	var doc WorkflowsTransition
	w.CreateTime = time.Now()
	w.UpdateTime = w.CreateTime
	result := orm.Eloquent.Table(w.TableName()).Create(&w)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *w
	return doc, nil
}

// Get 获取
func (w *WorkflowsTransition) Get(isRelated bool) (result WorkflowsTransition, err error) {

	key := fmt.Sprintf("wft:get:%+v:%d", isRelated, w.ID)

	getter := func() (interface{}, error) {
		table := orm.Eloquent.Table(w.TableName())
		if w.ID != 0 {
			table = table.Where("id = ?", w.ID)
		}
		if w.SourceStateID != 0 {
			table = table.Where("source_state_id = ?", w.SourceStateID)
		}

		if err = table.First(&result).Error; err != nil {
			return result, err
		}
		if isRelated {
			info := &WorkflowsWorkflow{
				ID: result.WorkflowID,
			}
			if wt, err := info.Get(true); err == nil {
				result.WorkflowsWorkflow = wt
			}
			w.WorkflowsDestState.ID = result.DestStateID
			if stat, err := w.WorkflowsDestState.Get(true, 2); err == nil {
				result.WorkflowsDestState = stat
			}
			w.WorkflowsSourceState.ID = result.SourceStateID
			if stat, err := w.WorkflowsSourceState.Get(true, 2); err == nil {
				result.WorkflowsSourceState = stat
			}
		}

		return result, err
	}
	val, err := cache.LRU().GetWithLoader(key, getter)
	if val != nil {
		result = val.(WorkflowsTransition)
	}
	return
}

// Gets 获取批量结果
func (w *WorkflowsTransition) GetPage(pageSize int, pageIndex int, isRelated bool, depth int) (results []WorkflowsTransition, count int, err error) {

	key := fmt.Sprintf("wft:getp:%d:%d:%+v:%d:%d", pageSize, pageIndex, isRelated, depth, w.WorkflowID)

	getter := func() (interface{}, error) {
		table := orm.Eloquent.Select("*").Table(w.TableName())
		if w.WorkflowID != 0 {
			table = table.Where("workflow_id = ?", w.WorkflowID)
		}
		if w.SourceStateID != 0 {
			table = table.Where("source_state_id = ?", w.SourceStateID)
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
				info := &WorkflowsWorkflow{
					ID: r.WorkflowID,
				}
				if wt, err := info.Get(true); err == nil {
					results[i].WorkflowsWorkflow = wt
				}
				w.WorkflowsDestState.ID = r.DestStateID
				if stat, err := w.WorkflowsDestState.Get(true, 2); err == nil {
					results[i].WorkflowsDestState = stat
				}
				w.WorkflowsSourceState.ID = r.SourceStateID
				if stat, err := w.WorkflowsSourceState.Get(true, 2); err == nil {
					results[i].WorkflowsSourceState = stat
				}
			}
		}
		table.Count(&count)
		return results, err
	}

	val, err := cache.LRU().GetWithLoader(key, getter)
	if val != nil {
		results = val.([]WorkflowsTransition)
		count = len(results)
	}
	return
}

// 更新WorkflowsTransition
func (e *WorkflowsTransition) Update(id int) (update WorkflowsTransition, err error) {
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

// 删除WorkflowsTransition
func (e *WorkflowsTransition) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&WorkflowsTransition{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *WorkflowsTransition) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&WorkflowsTransition{}).Error; err != nil {
		return
	}
	Result = true
	return
}
