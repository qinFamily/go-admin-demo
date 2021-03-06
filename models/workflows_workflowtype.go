package models

import (
	"fmt"
	"go-admin-demo/cache"

	orm "go-admin-demo/database"
)

// WorkflowsWorkflowtype [...]
type WorkflowsWorkflowtype struct {
	ID        int    `gorm:"primary_key;column:id;type:int(11);not null" json:"id"`
	Memo      string `gorm:"column:memo;type:text;not null" json:"memo"`
	Name      string `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Code      string `gorm:"unique;column:code;type:varchar(32);not null" json:"code"`
	OrderID   int    `gorm:"column:order_id;type:int(11);not null" json:"order_id"`
	DataScope string `json:"-" gorm:"-"`
	Params    string `json:"-"  gorm:"-"`
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

// 删除缓存
func (w *WorkflowsWorkflowtype) deleteKeys() {
	
	keyGet := fmt.Sprintf("wfwt:get:%d", w.ID)
	cache.LRU().Del(keyGet)
	for i := 1; i < 11; i++ {
		keyGetp := fmt.Sprintf("wfwt:getp:20:%d", i)
		cache.LRU().Del(keyGetp)
	}
}

func (w *WorkflowsWorkflowtype) Create() (wt WorkflowsWorkflowtype, err error) {

	result := orm.Eloquent.Table(w.TableName()).Create(&w)
	if result.Error != nil {
		err = result.Error
	}
	wt = *w

	// 删除缓存【笑哭】
	w.deleteKeys()

	return
}

func (w *WorkflowsWorkflowtype) Get() (result WorkflowtypeWorkflowsSet, err error) {
	key := fmt.Sprintf("wfwt:get:%d", w.ID)

	getter := func() (interface{}, error) {
		table := orm.Eloquent.Table(w.TableName()).Order("order_id")
		if w.ID != 0 {
			table = table.Where("id = ?", w.ID)
		}
		err = table.First(&result).Error
		if err == nil {
			wf := &WorkflowsWorkflow{
				TypeID: result.ID,
			}
			res, _, err := wf.GetPage(200, 1, false)
			if err == nil {
				result.WorkflowSet = res
			}
		}
		return result, err
	}
	val, err := cache.LRU().GetWithLoader(key, getter)
	// log.Println("***************************** cache.LRU().GetWithLoader", key, "val", val, "err", err)
	if val != nil {
		result = val.(WorkflowtypeWorkflowsSet)
	}
	return
}

// Gets 获取批量结果
func (w *WorkflowsWorkflowtype) GetPage(pageSize int, pageIndex int) (results []WorkflowtypeWorkflowsSet, count int, err error) {

	key := fmt.Sprintf("wfwt:getp:%d:%d", pageSize, pageIndex)

	getter := func() (interface{}, error) {
		table := orm.Eloquent.Select("*").Table(w.TableName()).Order("order_id")

		// 数据权限控制(如果不需要数据权限请将此处去掉)
		//dataPermission := new(DataPermission)
		//dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
		//table = dataPermission.GetDataScope(e.TableName(), table)

		if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&results).Error; err != nil {
			return results, err
		}
		table.Count(&count)
		for i := range results {
			wf := &WorkflowsWorkflow{
				TypeID: results[i].ID,
			}
			res, _, err := wf.GetPage(200, 1, false)
			if err == nil {
				results[i].WorkflowSet = res
			}
		}
		return results, err
	}

	val, err := cache.LRU().GetWithLoader(key, getter)
	// log.Println("***************************** cache.LRU().GetWithLoader", key, "val", val, "err", err)
	if val != nil {
		results = val.([]WorkflowtypeWorkflowsSet)
		count = len(results)
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
	// 删除缓存【笑哭】
	e.deleteKeys()

	return
}

// 删除WorkflowsWorkflow
func (e *WorkflowsWorkflowtype) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&WorkflowsWorkflowtype{}).Error; err != nil {
		success = false
		return
	}
	success = true
	// 删除缓存【笑哭】
	e.deleteKeys()

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
