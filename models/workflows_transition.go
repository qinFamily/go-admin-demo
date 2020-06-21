package models

import (
	orm "go-admin-demo/database"
	"time"
)

// WorkflowsTransition [...]
type WorkflowsTransition struct {
	ID                  int               `gorm:"primary_key;column:id;type:int(11);not null" json:"id"`
	CreateTime          time.Time         `gorm:"column:create_time;type:datetime;not null" json:"create_time"`
	UpdateTime          time.Time         `gorm:"column:update_time;type:datetime;not null" json:"update_time"`
	Memo                string            `gorm:"column:memo;type:text;not null" json:"memo"`
	Name                string            `gorm:"column:name;type:varchar(1);not null" json:"name"`
	TransitionType      string            `gorm:"column:transition_type;type:varchar(1);not null" json:"transition_type"`
	Timer               int               `gorm:"column:timer;type:int(11);not null" json:"timer"`
	ConditionExpression string            `gorm:"column:condition_expression;type:text;not null" json:"condition_expression"`
	AttributeType       string            `gorm:"column:attribute_type;type:varchar(1);not null" json:"attribute_type"`
	AlertEnable         int8              `gorm:"column:alert_enable;type:tinyint(4);not null" json:"alert_enable"`
	AlertText           string            `gorm:"column:alert_text;type:varchar(100);not null" json:"alert_text"`
	DestStateID         int               `gorm:"index;column:dest_state_id;type:int(11)" json:"dest_state_id"`
	WorkflowsState      WorkflowsState    `gorm:"association_foreignkey:dest_state_id;foreignkey:id" json:"workflows_state_set"`
	SourceStateID       int               `gorm:"index;column:source_state_id;type:int(11)" json:"source_state_id"`
	WorkflowID          int               `gorm:"index;column:workflow_id;type:int(11);not null" json:"workflow_id"`
	WorkflowsWorkflow   WorkflowsWorkflow `gorm:"association_foreignkey:workflow_id;foreignkey:id" json:"workflow_set"`
	DataScope           string            `json:"dataScope" gorm:"-"`
	Params              string            `json:"params"  gorm:"-"`
	BaseModel
}

// GetTableName get sql table name.获取数据库名字
func (WorkflowsTransition) TableName() string {
	return "workflows_transition"
}

func (w *WorkflowsTransition) Create() (WorkflowsTransition, error) {
	var doc WorkflowsTransition
	result := orm.Eloquent.Table(w.TableName()).Create(&w)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *w
	return doc, nil
}

// Get 获取
func (w *WorkflowsTransition) Get(isRelated bool) (WorkflowsTransition, error) {

	var wft WorkflowsTransition

	table := orm.Eloquent.Table(w.TableName())
	if w.ID != 0 {
		table = table.Where("id = ?", w.ID)
	}
	if err := table.First(&wft).Error; err != nil {
		return wft, err
	}
	if isRelated {
		info := &WorkflowsWorkflow{
			ID: wft.WorkflowID,
		}
		if wt, err := info.Get(false); err == nil {
			wft.WorkflowsWorkflow = wt
		}
	}

	return wft, nil
}

// Gets 获取批量结果
func (w *WorkflowsTransition) GetPage(pageSize int, pageIndex int, isRelated bool) (results []WorkflowsTransition, count int, err error) {

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
	if isRelated {
		for i, r := range results {
			info := &WorkflowsWorkflow{
				ID: r.WorkflowID,
			}
			if wt, err := info.Get(false); err == nil {
				results[i].WorkflowsWorkflow = wt
			}
		}
	}
	table.Count(&count)
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
