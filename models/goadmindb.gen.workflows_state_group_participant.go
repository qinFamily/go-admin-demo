package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type _WorkflowsStateGroupParticipantMgr struct {
	*_BaseMgr
}

// WorkflowsStateGroupParticipantMgr open func
func WorkflowsStateGroupParticipantMgr(db *gorm.DB) *_WorkflowsStateGroupParticipantMgr {
	if db == nil {
		panic(fmt.Errorf("WorkflowsStateGroupParticipantMgr need init by db"))
	}
	return &_WorkflowsStateGroupParticipantMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WorkflowsStateGroupParticipantMgr) GetTableName() string {
	return "workflows_state_group_participant"
}

// Get 获取
func (obj *_WorkflowsStateGroupParticipantMgr) Get() (result WorkflowsStateGroupParticipant, err error) {
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
			var info SysDept //
			err = obj.DB.New().Table("sys_dept").Where("dept_id = ?", result.GroupID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysDept = info
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_WorkflowsStateGroupParticipantMgr) Gets() (results []*WorkflowsStateGroupParticipant, err error) {
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
				var info SysDept //
				err = obj.DB.New().Table("sys_dept").Where("dept_id = ?", results[i].GroupID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].SysDept = info
			}
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WorkflowsStateGroupParticipantMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithStateID state_id获取
func (obj *_WorkflowsStateGroupParticipantMgr) WithStateID(stateID int) Option {
	return optionFunc(func(o *options) { o.query["state_id"] = stateID })
}

// WithGroupID group_id获取
func (obj *_WorkflowsStateGroupParticipantMgr) WithGroupID(groupID int) Option {
	return optionFunc(func(o *options) { o.query["group_id"] = groupID })
}

// GetByOption 功能选项模式获取
func (obj *_WorkflowsStateGroupParticipantMgr) GetByOption(opts ...Option) (result WorkflowsStateGroupParticipant, err error) {
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
			var info SysDept //
			err = obj.DB.New().Table("sys_dept").Where("dept_id = ?", result.GroupID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysDept = info
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WorkflowsStateGroupParticipantMgr) GetByOptions(opts ...Option) (results []*WorkflowsStateGroupParticipant, err error) {
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
				var info SysDept //
				err = obj.DB.New().Table("sys_dept").Where("dept_id = ?", results[i].GroupID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].SysDept = info
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_WorkflowsStateGroupParticipantMgr) GetFromID(id int) (result WorkflowsStateGroupParticipant, err error) {
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
			var info SysDept //
			err = obj.DB.New().Table("sys_dept").Where("dept_id = ?", result.GroupID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysDept = info
		}
	}

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_WorkflowsStateGroupParticipantMgr) GetBatchFromID(ids []int) (results []*WorkflowsStateGroupParticipant, err error) {
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
				var info SysDept //
				err = obj.DB.New().Table("sys_dept").Where("dept_id = ?", results[i].GroupID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].SysDept = info
			}
		}
	}
	return
}

// GetFromStateID 通过state_id获取内容
func (obj *_WorkflowsStateGroupParticipantMgr) GetFromStateID(stateID int) (result WorkflowsStateGroupParticipant, err error) {
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
			var info SysDept //
			err = obj.DB.New().Table("sys_dept").Where("dept_id = ?", result.GroupID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysDept = info
		}
	}

	return
}

// GetBatchFromStateID 批量唯一主键查找
func (obj *_WorkflowsStateGroupParticipantMgr) GetBatchFromStateID(stateIDs []int) (results []*WorkflowsStateGroupParticipant, err error) {
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
				var info SysDept //
				err = obj.DB.New().Table("sys_dept").Where("dept_id = ?", results[i].GroupID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].SysDept = info
			}
		}
	}
	return
}

// GetFromGroupID 通过group_id获取内容
func (obj *_WorkflowsStateGroupParticipantMgr) GetFromGroupID(groupID int) (result WorkflowsStateGroupParticipant, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("group_id = ?", groupID).Find(&result).Error
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
			var info SysDept //
			err = obj.DB.New().Table("sys_dept").Where("dept_id = ?", result.GroupID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysDept = info
		}
	}

	return
}

// GetBatchFromGroupID 批量唯一主键查找
func (obj *_WorkflowsStateGroupParticipantMgr) GetBatchFromGroupID(groupIDs []int) (results []*WorkflowsStateGroupParticipant, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("group_id IN (?)", groupIDs).Find(&results).Error
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
				var info SysDept //
				err = obj.DB.New().Table("sys_dept").Where("dept_id = ?", results[i].GroupID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].SysDept = info
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_WorkflowsStateGroupParticipantMgr) FetchByPrimaryKey(id int) (result WorkflowsStateGroupParticipant, err error) {
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
			var info SysDept //
			err = obj.DB.New().Table("sys_dept").Where("dept_id = ?", result.GroupID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysDept = info
		}
	}

	return
}

// FetchByWorkflowsStateGroupParticipantStateIDGroupIDUniqUniqueIndex primay or index 获取唯一内容
func (obj *_WorkflowsStateGroupParticipantMgr) FetchByWorkflowsStateGroupParticipantStateIDGroupIDUniqUniqueIndex(stateID int, groupID int) (result WorkflowsStateGroupParticipant, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("state_id = ? AND group_id = ?", stateID, groupID).Find(&result).Error
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
			var info SysDept //
			err = obj.DB.New().Table("sys_dept").Where("dept_id = ?", result.GroupID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysDept = info
		}
	}

	return
}

// FetchByIndex  获取多个内容
func (obj *_WorkflowsStateGroupParticipantMgr) FetchByIndex(stateID int, groupID int) (results []*WorkflowsStateGroupParticipant, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("state_id = ? AND group_id = ?", stateID, groupID).Find(&results).Error
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
				var info SysDept //
				err = obj.DB.New().Table("sys_dept").Where("dept_id = ?", results[i].GroupID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].SysDept = info
			}
		}
	}
	return
}
