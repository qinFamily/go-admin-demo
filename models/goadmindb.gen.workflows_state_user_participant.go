package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type _WorkflowsStateUserParticipantMgr struct {
	*_BaseMgr
}

// WorkflowsStateUserParticipantMgr open func
func WorkflowsStateUserParticipantMgr(db *gorm.DB) *_WorkflowsStateUserParticipantMgr {
	if db == nil {
		panic(fmt.Errorf("WorkflowsStateUserParticipantMgr need init by db"))
	}
	return &_WorkflowsStateUserParticipantMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WorkflowsStateUserParticipantMgr) GetTableName() string {
	return "workflows_state_user_participant"
}

// Get 获取
func (obj *_WorkflowsStateUserParticipantMgr) Get() (result WorkflowsStateUserParticipant, err error) {
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
			var info SysUser //
			err = obj.DB.New().Table("sys_user").Where("user_id = ?", result.UserID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysUser = info
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_WorkflowsStateUserParticipantMgr) Gets() (results []*WorkflowsStateUserParticipant, err error) {
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
				var info SysUser //
				err = obj.DB.New().Table("sys_user").Where("user_id = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].SysUser = info
			}
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WorkflowsStateUserParticipantMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithStateID state_id获取
func (obj *_WorkflowsStateUserParticipantMgr) WithStateID(stateID int) Option {
	return optionFunc(func(o *options) { o.query["state_id"] = stateID })
}

// WithUserID user_id获取
func (obj *_WorkflowsStateUserParticipantMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// GetByOption 功能选项模式获取
func (obj *_WorkflowsStateUserParticipantMgr) GetByOption(opts ...Option) (result WorkflowsStateUserParticipant, err error) {
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
			var info SysUser //
			err = obj.DB.New().Table("sys_user").Where("user_id = ?", result.UserID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysUser = info
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WorkflowsStateUserParticipantMgr) GetByOptions(opts ...Option) (results []*WorkflowsStateUserParticipant, err error) {
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
				var info SysUser //
				err = obj.DB.New().Table("sys_user").Where("user_id = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].SysUser = info
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_WorkflowsStateUserParticipantMgr) GetFromID(id int) (result WorkflowsStateUserParticipant, err error) {
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
			var info SysUser //
			err = obj.DB.New().Table("sys_user").Where("user_id = ?", result.UserID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysUser = info
		}
	}

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_WorkflowsStateUserParticipantMgr) GetBatchFromID(ids []int) (results []*WorkflowsStateUserParticipant, err error) {
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
				var info SysUser //
				err = obj.DB.New().Table("sys_user").Where("user_id = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].SysUser = info
			}
		}
	}
	return
}

// GetFromStateID 通过state_id获取内容
func (obj *_WorkflowsStateUserParticipantMgr) GetFromStateID(stateID int) (result WorkflowsStateUserParticipant, err error) {
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
			var info SysUser //
			err = obj.DB.New().Table("sys_user").Where("user_id = ?", result.UserID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysUser = info
		}
	}

	return
}

// GetBatchFromStateID 批量唯一主键查找
func (obj *_WorkflowsStateUserParticipantMgr) GetBatchFromStateID(stateIDs []int) (results []*WorkflowsStateUserParticipant, err error) {
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
				var info SysUser //
				err = obj.DB.New().Table("sys_user").Where("user_id = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].SysUser = info
			}
		}
	}
	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_WorkflowsStateUserParticipantMgr) GetFromUserID(userID int) (result WorkflowsStateUserParticipant, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&result).Error
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
			var info SysUser //
			err = obj.DB.New().Table("sys_user").Where("user_id = ?", result.UserID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysUser = info
		}
	}

	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_WorkflowsStateUserParticipantMgr) GetBatchFromUserID(userIDs []int) (results []*WorkflowsStateUserParticipant, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error
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
				var info SysUser //
				err = obj.DB.New().Table("sys_user").Where("user_id = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].SysUser = info
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_WorkflowsStateUserParticipantMgr) FetchByPrimaryKey(id int) (result WorkflowsStateUserParticipant, err error) {
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
			var info SysUser //
			err = obj.DB.New().Table("sys_user").Where("user_id = ?", result.UserID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysUser = info
		}
	}

	return
}

// FetchByWorkflowsStateUserParticipantStateIDUserIDUniqUniqueIndex primay or index 获取唯一内容
func (obj *_WorkflowsStateUserParticipantMgr) FetchByWorkflowsStateUserParticipantStateIDUserIDUniqUniqueIndex(stateID int, userID int) (result WorkflowsStateUserParticipant, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("state_id = ? AND user_id = ?", stateID, userID).Find(&result).Error
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
			var info SysUser //
			err = obj.DB.New().Table("sys_user").Where("user_id = ?", result.UserID).Find(&info).Error
			if err != nil {
				return
			}
			result.SysUser = info
		}
	}

	return
}

// FetchByIndex  获取多个内容
func (obj *_WorkflowsStateUserParticipantMgr) FetchByIndex(stateID int, userID int) (results []*WorkflowsStateUserParticipant, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("state_id = ? AND user_id = ?", stateID, userID).Find(&results).Error
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
				var info SysUser //
				err = obj.DB.New().Table("sys_user").Where("user_id = ?", results[i].UserID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].SysUser = info
			}
		}
	}
	return
}
