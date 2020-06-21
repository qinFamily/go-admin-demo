package models

import (
	_ "time"

	orm "go-admin-demo/database"
)

// WorkflowsWorkflowtype [...]
type WorkflowsWorkflowtype struct {
	ID        int    `gorm:"primary_key;column:id;type:int(11);not null" json:"id"`
	Memo      string `gorm:"column:memo;type:text;not null" json:"memo"`
	Name      string `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Code      string `gorm:"unique;column:code;type:varchar(32);not null" json:"code"`
	OrderID   int    `gorm:"column:order_id;type:int(11);not null" json:"order_id"`
	DataScope string `json:"dataScope" gorm:"-"`
	Params    string `json:"params"  gorm:"-"`
	BaseModel
}

// WorkflowtypeWorkflowsSet [...]
type WorkflowtypeWorkflowsSet struct {
	*WorkflowsWorkflowtype
	WorkflowSet []WorkflowsWorkflow `gorm:"-" json:"workflow_set"`
}

// GetTableName get sql table name.获取数据库名字
func (WorkflowsWorkflowtype) TableName() string {
	return "workflows_workflowtype"
}

func (w *WorkflowsWorkflowtype) Create() (wt WorkflowsWorkflowtype, err error) {

	result := orm.Eloquent.Table(w.TableName()).Create(&wt)
	if result.Error != nil {
		err = result.Error
	}
	return
}

func (w *WorkflowsWorkflowtype) Get() (wt WorkflowtypeWorkflowsSet, err error) {
	table := orm.Eloquent.Table(w.TableName())
	if w.ID != 0 {
		table = table.Where("id = ?", w.ID)
	}
	err = table.First(&wt).Error
	if err == nil {
		wf := &WorkflowsWorkflow{}
		res, _, err := wf.GetPage(1, 20, false)
		if err == nil {
			wt.WorkflowSet = res
		}
	}
	return
}

// Gets 获取批量结果
func (w *WorkflowsWorkflowtype) GetPage(pageSize int, pageIndex int) (results []WorkflowtypeWorkflowsSet, count int, err error) {

	table := orm.Eloquent.Select("*").Table(w.TableName())

	// 数据权限控制(如果不需要数据权限请将此处去掉)
	//dataPermission := new(DataPermission)
	//dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	//table = dataPermission.GetDataScope(e.TableName(), table)

	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&results).Error; err != nil {
		return nil, 0, err
	}
	table.Count(&count)
	for i := range results {
		wf := &WorkflowsWorkflow{}
		res, _, err := wf.GetPage(1, 20, false)
		if err == nil {
			results[i].WorkflowSet = res
		}
	}
	return
}

// 更新WorkflowsWorkflow
func (e *WorkflowsWorkflowtype) Update(id int) (update WorkflowsWorkflowtype, err error) {
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

// 删除WorkflowsWorkflow
func (e *WorkflowsWorkflowtype) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&WorkflowsWorkflowtype{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *WorkflowsWorkflowtype) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&WorkflowsWorkflowtype{}).Error; err != nil {
		return
	}
	Result = true
	return
}
