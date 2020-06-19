package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type _WorkflowsStateRoleParticipantMgr struct {
	*_BaseMgr
}

// WorkflowsStateRoleParticipantMgr open func
func WorkflowsStateRoleParticipantMgr(db *gorm.DB) *_WorkflowsStateRoleParticipantMgr {
	if db == nil {
		panic(fmt.Errorf("WorkflowsStateRoleParticipantMgr need init by db"))
	}
	return &_WorkflowsStateRoleParticipantMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WorkflowsStateRoleParticipantMgr) GetTableName() string {
	return "workflows_state_role_participant"
}

// Get 获取
func (obj *_WorkflowsStateRoleParticipantMgr) Get() (result WorkflowsStateRoleParticipant, err error) {
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
			var info SysRole //
			err = obj.DB.New().Table("sys_role").Where("role_id = ?", result.RoleID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysRole = info
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_WorkflowsStateRoleParticipantMgr) Gets() (results []*WorkflowsStateRoleParticipant, err error) {
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
				var info SysRole //
				err = obj.DB.New().Table("sys_role").Where("role_id = ?", results[i].RoleID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].SysRole = info
			}
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WorkflowsStateRoleParticipantMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithStateID state_id获取
func (obj *_WorkflowsStateRoleParticipantMgr) WithStateID(stateID int) Option {
	return optionFunc(func(o *options) { o.query["state_id"] = stateID })
}

// WithRoleID role_id获取
func (obj *_WorkflowsStateRoleParticipantMgr) WithRoleID(roleID int) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// GetByOption 功能选项模式获取
func (obj *_WorkflowsStateRoleParticipantMgr) GetByOption(opts ...Option) (result WorkflowsStateRoleParticipant, err error) {
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
			var info SysRole //
			err = obj.DB.New().Table("sys_role").Where("role_id = ?", result.RoleID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysRole = info
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WorkflowsStateRoleParticipantMgr) GetByOptions(opts ...Option) (results []*WorkflowsStateRoleParticipant, err error) {
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
				var info SysRole //
				err = obj.DB.New().Table("sys_role").Where("role_id = ?", results[i].RoleID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].SysRole = info
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_WorkflowsStateRoleParticipantMgr) GetFromID(id int) (result WorkflowsStateRoleParticipant, err error) {
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
			var info SysRole //
			err = obj.DB.New().Table("sys_role").Where("role_id = ?", result.RoleID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysRole = info
		}
	}

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_WorkflowsStateRoleParticipantMgr) GetBatchFromID(ids []int) (results []*WorkflowsStateRoleParticipant, err error) {
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
				var info SysRole //
				err = obj.DB.New().Table("sys_role").Where("role_id = ?", results[i].RoleID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].SysRole = info
			}
		}
	}
	return
}

// GetFromStateID 通过state_id获取内容
func (obj *_WorkflowsStateRoleParticipantMgr) GetFromStateID(stateID int) (result WorkflowsStateRoleParticipant, err error) {
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
			var info SysRole //
			err = obj.DB.New().Table("sys_role").Where("role_id = ?", result.RoleID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysRole = info
		}
	}

	return
}

// GetBatchFromStateID 批量唯一主键查找
func (obj *_WorkflowsStateRoleParticipantMgr) GetBatchFromStateID(stateIDs []int) (results []*WorkflowsStateRoleParticipant, err error) {
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
				var info SysRole //
				err = obj.DB.New().Table("sys_role").Where("role_id = ?", results[i].RoleID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].SysRole = info
			}
		}
	}
	return
}

// GetFromRoleID 通过role_id获取内容
func (obj *_WorkflowsStateRoleParticipantMgr) GetFromRoleID(roleID int) (result WorkflowsStateRoleParticipant, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("role_id = ?", roleID).Find(&result).Error
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
			var info SysRole //
			err = obj.DB.New().Table("sys_role").Where("role_id = ?", result.RoleID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysRole = info
		}
	}

	return
}

// GetBatchFromRoleID 批量唯一主键查找
func (obj *_WorkflowsStateRoleParticipantMgr) GetBatchFromRoleID(roleIDs []int) (results []*WorkflowsStateRoleParticipant, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("role_id IN (?)", roleIDs).Find(&results).Error
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
				var info SysRole //
				err = obj.DB.New().Table("sys_role").Where("role_id = ?", results[i].RoleID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].SysRole = info
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_WorkflowsStateRoleParticipantMgr) FetchByPrimaryKey(id int) (result WorkflowsStateRoleParticipant, err error) {
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
			var info SysRole //
			err = obj.DB.New().Table("sys_role").Where("role_id = ?", result.RoleID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysRole = info
		}
	}

	return
}

// FetchByWorkflowsStateRoleParticipantStateIDRoleIDUniqUniqueIndex primay or index 获取唯一内容
func (obj *_WorkflowsStateRoleParticipantMgr) FetchByWorkflowsStateRoleParticipantStateIDRoleIDUniqUniqueIndex(stateID int, roleID int) (result WorkflowsStateRoleParticipant, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("state_id = ? AND role_id = ?", stateID, roleID).Find(&result).Error
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
			var info SysRole //
			err = obj.DB.New().Table("sys_role").Where("role_id = ?", result.RoleID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysRole = info
		}
	}

	return
}

// FetchByIndex  获取多个内容
func (obj *_WorkflowsStateRoleParticipantMgr) FetchByIndex(stateID int, roleID int) (results []*WorkflowsStateRoleParticipant, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("state_id = ? AND role_id = ?", stateID, roleID).Find(&results).Error
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
				var info SysRole //
				err = obj.DB.New().Table("sys_role").Where("role_id = ?", results[i].RoleID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].SysRole = info
			}
		}
	}
	return
}
