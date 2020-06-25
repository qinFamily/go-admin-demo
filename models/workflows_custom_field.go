package models

import (
	"fmt"
	"go-admin-demo/cache"
	orm "go-admin-demo/database"
	"time"
)

// WorkflowsCustomfield [...]
type WorkflowsCustomfield struct {
	ID                  int       `gorm:"primary_key;column:id;type:int(11);not null" json:"id"`
	CreateTime          time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`
	UpdateTime          time.Time `gorm:"column:update_time;type:datetime;not null" json:"update_time"`
	Memo                string    `gorm:"column:memo;type:text;not null" json:"memo"`
	FieldAttribute      bool      `gorm:"column:field_attribute;type:tinyint(4);not null" json:"field_attribute"`
	FieldType           int    `gorm:"column:field_type;type:varchar(1);not null" json:"field_type"`
	FieldKey            string    `gorm:"column:field_key;type:varchar(50);not null" json:"field_key"`
	FieldName           string    `gorm:"column:field_name;type:varchar(50);not null" json:"field_name"`
	OrderID             int       `gorm:"column:order_id;type:int(11);not null" json:"order_id"`
	DefaultValue        string    `gorm:"column:default_value;type:varchar(100)" json:"default_value"`
	FieldTemplate       string    `gorm:"column:field_template;type:text;not null" json:"field_template"`
	BooleanFieldDisplay string    `gorm:"column:boolean_field_display;type:varchar(100);not null" json:"boolean_field_display"`
	FieldChoice         string    `gorm:"column:field_choice;type:varchar(255);not null" json:"field_choice"`
	Label               string    `gorm:"column:label;type:varchar(100);not null" json:"label"`
	WorkflowID          int       `gorm:"index;column:workflow_id;type:int(11);not null" json:"workflow"`
	// WorkflowsWorkflow   WorkflowsWorkflow `gorm:"association_foreignkey:workflow_id;foreignkey:id" json:"workflow_set"`
	DataScope string `json:"dataScope" gorm:"-"`
	Params    string `json:"params"  gorm:"-"`
	BaseModel
}

// GetTableName get sql table name.获取数据库名字
func (WorkflowsCustomfield) TableName() string {
	return "workflows_customfield"
}

func (w *WorkflowsCustomfield) Create() (WorkflowsCustomfield, error) {
	var doc WorkflowsCustomfield
	result := orm.Eloquent.Table(w.TableName()).Create(&w)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *w
	return doc, nil
}

// Get 获取
func (w *WorkflowsCustomfield) Get(isRelated bool) (result WorkflowsCustomfield, err error) {

	key := fmt.Sprintf("wfc:get:%+v:%d", isRelated, w.ID)
	getter := func() (interface{}, error) {

		table := orm.Eloquent.Table(w.TableName()).Order("order_id")
		if w.ID != 0 {
			table = table.Where("id = ?", w.ID)
		}
		if err = table.First(&result).Error; err != nil {
			return result, err
		}
		// if isRelated {
		// 	info := &WorkflowsWorkflow{
		// 		ID: result.WorkflowID,
		// 	}
		// 	if wt, err := info.Get(false); err == nil {
		// 		result.WorkflowsWorkflow = wt
		// 	}
		// }

		return result, nil
	}
	val, err := cache.LRU().GetWithLoader(key, getter)
	if val != nil {
		result = val.(WorkflowsCustomfield)
	}
	return
}

// Gets 获取批量结果
func (w *WorkflowsCustomfield) GetPage(pageSize int, pageIndex int, isRelated bool) (results []WorkflowsCustomfield, count int, err error) {
	key := fmt.Sprintf("wfc:getp:%d:%d:%+v:%d", pageSize, pageIndex, isRelated, w.WorkflowID)

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
		// if isRelated {
		// 	for i, r := range results {
		// 		info := &WorkflowsWorkflow{
		// 			ID: r.WorkflowID,
		// 		}
		// 		if wt, err := info.Get(false); err == nil {
		// 			results[i].WorkflowsWorkflow = wt
		// 		}
		// 	}
		// }
		table.Count(&count)
		return results, err
	}
	val, err := cache.LRU().GetWithLoader(key, getter)
	if val != nil {
		results = val.([]WorkflowsCustomfield)
		count = len(results)
	}
	return
}

// 更新WorkflowsCustomfield
func (e *WorkflowsCustomfield) Update(id int) (update WorkflowsCustomfield, err error) {
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

// 删除WorkflowsCustomfield
func (e *WorkflowsCustomfield) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&WorkflowsCustomfield{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *WorkflowsCustomfield) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&WorkflowsCustomfield{}).Error; err != nil {
		return
	}
	Result = true
	return
}
