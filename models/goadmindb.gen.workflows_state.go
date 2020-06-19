package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type _WorkflowsStateMgr struct {
	*_BaseMgr
}

// WorkflowsStateMgr open func
func WorkflowsStateMgr(db *gorm.DB) *_WorkflowsStateMgr {
	if db == nil {
		panic(fmt.Errorf("WorkflowsStateMgr need init by db"))
	}
	return &_WorkflowsStateMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WorkflowsStateMgr) GetTableName() string {
	return "workflows_state"
}

// Get 获取
func (obj *_WorkflowsStateMgr) Get() (result WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info WorkflowsWorkflow //
			err = obj.DB.New().Table("workflows_workflow").Where("id = ?", result.WorkflowID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsWorkflow = info
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_WorkflowsStateMgr) Gets() (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WorkflowsStateMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCreateTime create_time获取
func (obj *_WorkflowsStateMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_WorkflowsStateMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithMemo memo获取
func (obj *_WorkflowsStateMgr) WithMemo(memo string) Option {
	return optionFunc(func(o *options) { o.query["memo"] = memo })
}

// WithName name获取
func (obj *_WorkflowsStateMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithIsHidden is_hidden获取
func (obj *_WorkflowsStateMgr) WithIsHidden(isHidden int8) Option {
	return optionFunc(func(o *options) { o.query["is_hidden"] = isHidden })
}

// WithOrderID order_id获取
func (obj *_WorkflowsStateMgr) WithOrderID(orderID int) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithStateType state_type获取
func (obj *_WorkflowsStateMgr) WithStateType(stateType string) Option {
	return optionFunc(func(o *options) { o.query["state_type"] = stateType })
}

// WithEnableRetreat enable_retreat获取
func (obj *_WorkflowsStateMgr) WithEnableRetreat(enableRetreat int8) Option {
	return optionFunc(func(o *options) { o.query["enable_retreat"] = enableRetreat })
}

// WithParticipantType participant_type获取
func (obj *_WorkflowsStateMgr) WithParticipantType(participantType string) Option {
	return optionFunc(func(o *options) { o.query["participant_type"] = participantType })
}

// WithWorkflowID workflow_id获取
func (obj *_WorkflowsStateMgr) WithWorkflowID(workflowID int) Option {
	return optionFunc(func(o *options) { o.query["workflow_id"] = workflowID })
}

// GetByOption 功能选项模式获取
func (obj *_WorkflowsStateMgr) GetByOption(opts ...Option) (result WorkflowsState, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info WorkflowsWorkflow //
			err = obj.DB.New().Table("workflows_workflow").Where("id = ?", result.WorkflowID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsWorkflow = info
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WorkflowsStateMgr) GetByOptions(opts ...Option) (results []*WorkflowsState, err error) {
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
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_WorkflowsStateMgr) GetFromID(id int) (result WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info WorkflowsWorkflow //
			err = obj.DB.New().Table("workflows_workflow").Where("id = ?", result.WorkflowID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsWorkflow = info
		}
	}

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_WorkflowsStateMgr) GetBatchFromID(ids []int) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_WorkflowsStateMgr) GetFromCreateTime(createTime time.Time) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量唯一主键查找
func (obj *_WorkflowsStateMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_WorkflowsStateMgr) GetFromUpdateTime(updateTime time.Time) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

// GetBatchFromUpdateTime 批量唯一主键查找
func (obj *_WorkflowsStateMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

// GetFromMemo 通过memo获取内容
func (obj *_WorkflowsStateMgr) GetFromMemo(memo string) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("memo = ?", memo).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

// GetBatchFromMemo 批量唯一主键查找
func (obj *_WorkflowsStateMgr) GetBatchFromMemo(memos []string) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("memo IN (?)", memos).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

// GetFromName 通过name获取内容
func (obj *_WorkflowsStateMgr) GetFromName(name string) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_WorkflowsStateMgr) GetBatchFromName(names []string) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

// GetFromIsHidden 通过is_hidden获取内容
func (obj *_WorkflowsStateMgr) GetFromIsHidden(isHidden int8) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_hidden = ?", isHidden).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

// GetBatchFromIsHidden 批量唯一主键查找
func (obj *_WorkflowsStateMgr) GetBatchFromIsHidden(isHiddens []int8) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_hidden IN (?)", isHiddens).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

// GetFromOrderID 通过order_id获取内容
func (obj *_WorkflowsStateMgr) GetFromOrderID(orderID int) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_id = ?", orderID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

// GetBatchFromOrderID 批量唯一主键查找
func (obj *_WorkflowsStateMgr) GetBatchFromOrderID(orderIDs []int) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_id IN (?)", orderIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

// GetFromStateType 通过state_type获取内容
func (obj *_WorkflowsStateMgr) GetFromStateType(stateType string) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("state_type = ?", stateType).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

// GetBatchFromStateType 批量唯一主键查找
func (obj *_WorkflowsStateMgr) GetBatchFromStateType(stateTypes []string) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("state_type IN (?)", stateTypes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

// GetFromEnableRetreat 通过enable_retreat获取内容
func (obj *_WorkflowsStateMgr) GetFromEnableRetreat(enableRetreat int8) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("enable_retreat = ?", enableRetreat).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

// GetBatchFromEnableRetreat 批量唯一主键查找
func (obj *_WorkflowsStateMgr) GetBatchFromEnableRetreat(enableRetreats []int8) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("enable_retreat IN (?)", enableRetreats).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

// GetFromParticipantType 通过participant_type获取内容
func (obj *_WorkflowsStateMgr) GetFromParticipantType(participantType string) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("participant_type = ?", participantType).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

// GetBatchFromParticipantType 批量唯一主键查找
func (obj *_WorkflowsStateMgr) GetBatchFromParticipantType(participantTypes []string) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("participant_type IN (?)", participantTypes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

// GetFromWorkflowID 通过workflow_id获取内容
func (obj *_WorkflowsStateMgr) GetFromWorkflowID(workflowID int) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("workflow_id = ?", workflowID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

// GetBatchFromWorkflowID 批量唯一主键查找
func (obj *_WorkflowsStateMgr) GetBatchFromWorkflowID(workflowIDs []int) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("workflow_id IN (?)", workflowIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_WorkflowsStateMgr) FetchByPrimaryKey(id int) (result WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info WorkflowsWorkflow //
			err = obj.DB.New().Table("workflows_workflow").Where("id = ?", result.WorkflowID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsWorkflow = info
		}
	}

	return
}

// FetchByIndex  获取多个内容
func (obj *_WorkflowsStateMgr) FetchByIndex(workflowID int) (results []*WorkflowsState, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("workflow_id = ?", workflowID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflow //
				err = obj.DB.New().Table("workflows_workflow").Where("id = ?", results[i].WorkflowID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflow = info
			}
		}
	}
	return
}
