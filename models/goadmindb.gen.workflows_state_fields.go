package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type _WorkflowsStateFieldsMgr struct {
	*_BaseMgr
}

// WorkflowsStateFieldsMgr open func
func WorkflowsStateFieldsMgr(db *gorm.DB) *_WorkflowsStateFieldsMgr {
	if db == nil {
		panic(fmt.Errorf("WorkflowsStateFieldsMgr need init by db"))
	}
	return &_WorkflowsStateFieldsMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WorkflowsStateFieldsMgr) GetTableName() string {
	return "workflows_state_fields"
}

// Get 获取
func (obj *_WorkflowsStateFieldsMgr) Get() (result WorkflowsStateFields, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info WorkflowsState //
			err = obj.DB.New().Table("workflows_state").Where("id = ?", result.StateID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsState = info
		}
		{
			var info WorkflowsCustomfield //
			err = obj.DB.New().Table("workflows_customfield").Where("id = ?", result.CustomfieldID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsCustomfield = info
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_WorkflowsStateFieldsMgr) Gets() (results []*WorkflowsStateFields, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].StateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsCustomfield //
				err = obj.DB.New().Table("workflows_customfield").Where("id = ?", results[i].CustomfieldID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsCustomfield = info
			}
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WorkflowsStateFieldsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithStateID state_id获取
func (obj *_WorkflowsStateFieldsMgr) WithStateID(stateID int) Option {
	return optionFunc(func(o *options) { o.query["state_id"] = stateID })
}

// WithCustomfieldID customfield_id获取
func (obj *_WorkflowsStateFieldsMgr) WithCustomfieldID(customfieldID int) Option {
	return optionFunc(func(o *options) { o.query["customfield_id"] = customfieldID })
}

// GetByOption 功能选项模式获取
func (obj *_WorkflowsStateFieldsMgr) GetByOption(opts ...Option) (result WorkflowsStateFields, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info WorkflowsState //
			err = obj.DB.New().Table("workflows_state").Where("id = ?", result.StateID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsState = info
		}
		{
			var info WorkflowsCustomfield //
			err = obj.DB.New().Table("workflows_customfield").Where("id = ?", result.CustomfieldID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsCustomfield = info
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WorkflowsStateFieldsMgr) GetByOptions(opts ...Option) (results []*WorkflowsStateFields, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].StateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsCustomfield //
				err = obj.DB.New().Table("workflows_customfield").Where("id = ?", results[i].CustomfieldID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsCustomfield = info
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_WorkflowsStateFieldsMgr) GetFromID(id int) (result WorkflowsStateFields, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info WorkflowsState //
			err = obj.DB.New().Table("workflows_state").Where("id = ?", result.StateID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsState = info
		}
		{
			var info WorkflowsCustomfield //
			err = obj.DB.New().Table("workflows_customfield").Where("id = ?", result.CustomfieldID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsCustomfield = info
		}
	}

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_WorkflowsStateFieldsMgr) GetBatchFromID(ids []int) (results []*WorkflowsStateFields, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].StateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsCustomfield //
				err = obj.DB.New().Table("workflows_customfield").Where("id = ?", results[i].CustomfieldID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsCustomfield = info
			}
		}
	}
	return
}

// GetFromStateID 通过state_id获取内容
func (obj *_WorkflowsStateFieldsMgr) GetFromStateID(stateID int) (result WorkflowsStateFields, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("state_id = ?", stateID).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info WorkflowsState //
			err = obj.DB.New().Table("workflows_state").Where("id = ?", result.StateID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsState = info
		}
		{
			var info WorkflowsCustomfield //
			err = obj.DB.New().Table("workflows_customfield").Where("id = ?", result.CustomfieldID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsCustomfield = info
		}
	}

	return
}

// GetBatchFromStateID 批量唯一主键查找
func (obj *_WorkflowsStateFieldsMgr) GetBatchFromStateID(stateIDs []int) (results []*WorkflowsStateFields, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("state_id IN (?)", stateIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].StateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsCustomfield //
				err = obj.DB.New().Table("workflows_customfield").Where("id = ?", results[i].CustomfieldID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsCustomfield = info
			}
		}
	}
	return
}

// GetFromCustomfieldID 通过customfield_id获取内容
func (obj *_WorkflowsStateFieldsMgr) GetFromCustomfieldID(customfieldID int) (result WorkflowsStateFields, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("customfield_id = ?", customfieldID).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info WorkflowsState //
			err = obj.DB.New().Table("workflows_state").Where("id = ?", result.StateID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsState = info
		}
		{
			var info WorkflowsCustomfield //
			err = obj.DB.New().Table("workflows_customfield").Where("id = ?", result.CustomfieldID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsCustomfield = info
		}
	}

	return
}

// GetBatchFromCustomfieldID 批量唯一主键查找
func (obj *_WorkflowsStateFieldsMgr) GetBatchFromCustomfieldID(customfieldIDs []int) (results []*WorkflowsStateFields, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("customfield_id IN (?)", customfieldIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].StateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsCustomfield //
				err = obj.DB.New().Table("workflows_customfield").Where("id = ?", results[i].CustomfieldID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsCustomfield = info
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_WorkflowsStateFieldsMgr) FetchByPrimaryKey(id int) (result WorkflowsStateFields, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info WorkflowsState //
			err = obj.DB.New().Table("workflows_state").Where("id = ?", result.StateID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsState = info
		}
		{
			var info WorkflowsCustomfield //
			err = obj.DB.New().Table("workflows_customfield").Where("id = ?", result.CustomfieldID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsCustomfield = info
		}
	}

	return
}

// FetchByWorkflowsStateFieldsStateIDCustomfieldIDUniqUniqueIndex primay or index 获取唯一内容
func (obj *_WorkflowsStateFieldsMgr) FetchByWorkflowsStateFieldsStateIDCustomfieldIDUniqUniqueIndex(stateID int, customfieldID int) (result WorkflowsStateFields, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("state_id = ? AND customfield_id = ?", stateID, customfieldID).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info WorkflowsState //
			err = obj.DB.New().Table("workflows_state").Where("id = ?", result.StateID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsState = info
		}
		{
			var info WorkflowsCustomfield //
			err = obj.DB.New().Table("workflows_customfield").Where("id = ?", result.CustomfieldID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsCustomfield = info
		}
	}

	return
}

// FetchByIndex  获取多个内容
func (obj *_WorkflowsStateFieldsMgr) FetchByIndex(stateID int, customfieldID int) (results []*WorkflowsStateFields, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("state_id = ? AND customfield_id = ?", stateID, customfieldID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].StateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsCustomfield //
				err = obj.DB.New().Table("workflows_customfield").Where("id = ?", results[i].CustomfieldID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsCustomfield = info
			}
		}
	}
	return
}
