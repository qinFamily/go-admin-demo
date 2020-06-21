package models

import (
	orm "go-admin-demo/database"
	"time"
)

// WorkflowsState [...]
type WorkflowsState struct {
	ID                int               `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	CreateTime        time.Time         `gorm:"column:create_time;type:datetime;not null" json:"create_time"`
	UpdateTime        time.Time         `gorm:"column:update_time;type:datetime;not null" json:"update_time"`
	Memo              string            `gorm:"column:memo;type:text;not null" json:"memo"`
	Name              string            `gorm:"column:name;type:varchar(50);not null" json:"name"`
	IsHidden          int8              `gorm:"column:is_hidden;type:tinyint(4);not null" json:"is_hidden"`
	OrderID           int               `gorm:"column:order_id;type:int(11);not null" json:"order_id"`
	StateType         string            `gorm:"column:state_type;type:varchar(1);not null" json:"state_type"`
	EnableRetreat     int8              `gorm:"column:enable_retreat;type:tinyint(4);not null" json:"enable_retreat"`
	ParticipantType   string            `gorm:"column:participant_type;type:varchar(1);not null" json:"participant_type"`
	WorkflowID        int               `gorm:"index;column:workflow_id;type:int(11);not null" json:"workflow_id"`
	WorkflowsWorkflow WorkflowsWorkflow `gorm:"association_foreignkey:workflow_id;foreignkey:id" json:"workflow_set"`
	DataScope         string            `json:"dataScope" gorm:"-"`
	Params            string            `json:"params"  gorm:"-"`
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

	var wft WorkflowsState

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
