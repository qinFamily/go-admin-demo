package models

import (
	"fmt"
	"go-admin-demo/cache"
	orm "go-admin-demo/database"
)

// WorkflowsWorkflow [...]
type WorkflowsWorkflow struct {
	ID                    int         `gorm:"primary_key;column:id;type:int(11);not null" json:"id"`
	Memo                  string      `gorm:"column:memo;type:text;not null" json:"memo"`
	Name                  string      `gorm:"column:name;type:varchar(50);not null" json:"name"`
	TicketSnPrefix        string      `gorm:"column:ticket_sn_prefix;type:varchar(20);not null" json:"ticket_sn_prefix"`
	Status                bool        `gorm:"column:status;type:tinyint(4);not null" json:"status"`
	ViewPermissionCheck   bool        `gorm:"column:view_permission_check;type:tinyint(4);not null" json:"view_permission_check"`
	LimitExpression       string      `gorm:"column:limit_expression;type:text;not null" json:"limit_expression"`
	DisplayFormStr        string      `gorm:"column:display_form_str;type:text;not null" json:"display_form_str"`
	TitleTemplate         string      `gorm:"column:title_template;type:varchar(50)" json:"title_template"`
	TypeID                int         `gorm:"index;column:type_id;type:int(11);not null" json:"-"`
	WorkflowsWorkflowtype interface{} `gorm:"-" json:"type"`
	DataScope             string      `json:"-" gorm:"-"`
	Params                string      `json:"-"  gorm:"-"`
	BaseModel
}

// GetTableName get sql table name.获取数据库名字
func (WorkflowsWorkflow) TableName() string {
	return "workflows_workflow"
}

// 删除缓存
func (w *WorkflowsWorkflow) deleteKeys(){

	wfwKeyTrue := fmt.Sprintf("wfw:get:%+v:%d", true, w.ID)
	cache.LRU().Del(wfwKeyTrue)
	
	wfwKeyFalse := fmt.Sprintf("wfw:get:%+v:%d", false, w.ID)
	cache.LRU().Del(wfwKeyFalse)

	for i := 1; i < 11; i++ {
		wfwgetpKeyTrue := fmt.Sprintf("wfw:getp:%d:%d:%+v:%d", 20, i, true, w.TypeID)
		cache.LRU().Del(wfwgetpKeyTrue)
		wfwgetpKeyFalse := fmt.Sprintf("wfw:getp:%d:%d:%+v:%d", 20, i, false, w.TypeID)
		cache.LRU().Del(wfwgetpKeyFalse)
	}
	
}

func (w *WorkflowsWorkflow) Create() (WorkflowsWorkflow, error) {
	var doc WorkflowsWorkflow
	if w.TypeID == 0 {
		if f, ok := w.WorkflowsWorkflowtype.(float64); ok {
			w.TypeID = int(f)
		}
	}
	result := orm.Eloquent.Table(w.TableName()).Create(&w)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *w

	// 删除缓存
	w.deleteKeys()

	return doc, nil
}

// Get 获取
func (w *WorkflowsWorkflow) Get(isRelated bool) (result WorkflowsWorkflow, err error) {

	key := fmt.Sprintf("wfw:get:%+v:%d", isRelated, w.ID)

	getter := func() (interface{}, error) {
		table := orm.Eloquent.Table(w.TableName())
		if w.ID != 0 {
			table = table.Where("id = ?", w.ID)
		}
		if err = table.First(&result).Error; err != nil {
			return result, err
		}
		if isRelated {
			info := &WorkflowsWorkflowtype{
				ID: result.TypeID,
			}
			if wt, err := info.Get(); err == nil {
				result.WorkflowsWorkflowtype = wt
			}
		} else {
			result.WorkflowsWorkflowtype = result.TypeID
		}

		return result, err
	}
	val, err := cache.LRU().GetWithLoader(key, getter)
	// log.Println("***************************** cache.LRU().GetWithLoader", key, "val", val, "err", err)
	if val != nil {
		result = val.(WorkflowsWorkflow)
	}
	return
}

// Gets 获取批量结果
func (w *WorkflowsWorkflow) GetPage(pageSize int, pageIndex int, isRelated bool) (results []WorkflowsWorkflow, count int, err error) {

	key := fmt.Sprintf("wfw:getp:%d:%d:%+v:%d", pageSize, pageIndex, isRelated, w.TypeID)

	getter := func() (interface{}, error) {
		table := orm.Eloquent.Select("*").Table(w.TableName())
		if w.TypeID != 0 {
			table = table.Where("type_id = ?", w.TypeID)
		}

		// 数据权限控制(如果不需要数据权限请将此处去掉)
		//dataPermission := new(DataPermission)
		//dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
		//table = dataPermission.GetDataScope(e.TableName(), table)

		if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&results).Error; err != nil {
			return results, err
		}
		for i, r := range results {
			if isRelated {
				info := &WorkflowsWorkflowtype{
					ID: r.TypeID,
				}
				if wt, err := info.Get(); err == nil {
					results[i].WorkflowsWorkflowtype = wt
				}
			} else {
				results[i].WorkflowsWorkflowtype = r.TypeID
			}
		}
		table.Count(&count)
		return results, err
	}

	val, err := cache.LRU().GetWithLoader(key, getter)
	// log.Println("***************************** cache.LRU().GetWithLoader", key, "val", val, "err", err)
	if val != nil {
		results = val.([]WorkflowsWorkflow)
		count = len(results)
	}
	return

}

// 更新WorkflowsWorkflow
func (e *WorkflowsWorkflow) Update(id int) (update WorkflowsWorkflow, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	// 删除缓存
	e.deleteKeys()

	return
}

// 删除WorkflowsWorkflow
func (e *WorkflowsWorkflow) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&WorkflowsWorkflow{}).Error; err != nil {
		success = false
		return
	}
	success = true
	// 删除缓存
	e.deleteKeys()
	
	return
}

//批量删除
func (e *WorkflowsWorkflow) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&WorkflowsWorkflow{}).Error; err != nil {
		return
	}
	Result = true
	return
}
