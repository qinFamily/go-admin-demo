package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type _WorkflowsTransitionMgr struct {
	*_BaseMgr
}

// WorkflowsTransitionMgr open func
func WorkflowsTransitionMgr(db *gorm.DB) *_WorkflowsTransitionMgr {
	if db == nil {
		panic(fmt.Errorf("WorkflowsTransitionMgr need init by db"))
	}
	return &_WorkflowsTransitionMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WorkflowsTransitionMgr) GetTableName() string {
	return "workflows_transition"
}

// Get 获取
func (obj *_WorkflowsTransitionMgr) Get() (result WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info WorkflowsState //
			err = obj.DB.New().Table("workflows_state").Where("id = ?", result.DestStateID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsState = info
		}
		{
			var info WorkflowsState //
			err = obj.DB.New().Table("workflows_state").Where("id = ?", result.SourceStateID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsState = info
		}
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
func (obj *_WorkflowsTransitionMgr) Gets() (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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
func (obj *_WorkflowsTransitionMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCreateTime create_time获取
func (obj *_WorkflowsTransitionMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_WorkflowsTransitionMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithMemo memo获取
func (obj *_WorkflowsTransitionMgr) WithMemo(memo string) Option {
	return optionFunc(func(o *options) { o.query["memo"] = memo })
}

// WithName name获取
func (obj *_WorkflowsTransitionMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithTransitionType transition_type获取
func (obj *_WorkflowsTransitionMgr) WithTransitionType(transitionType string) Option {
	return optionFunc(func(o *options) { o.query["transition_type"] = transitionType })
}

// WithTimer timer获取
func (obj *_WorkflowsTransitionMgr) WithTimer(timer int) Option {
	return optionFunc(func(o *options) { o.query["timer"] = timer })
}

// WithConditionExpression condition_expression获取
func (obj *_WorkflowsTransitionMgr) WithConditionExpression(conditionExpression string) Option {
	return optionFunc(func(o *options) { o.query["condition_expression"] = conditionExpression })
}

// WithAttributeType attribute_type获取
func (obj *_WorkflowsTransitionMgr) WithAttributeType(attributeType string) Option {
	return optionFunc(func(o *options) { o.query["attribute_type"] = attributeType })
}

// WithAlertEnable alert_enable获取
func (obj *_WorkflowsTransitionMgr) WithAlertEnable(alertEnable int8) Option {
	return optionFunc(func(o *options) { o.query["alert_enable"] = alertEnable })
}

// WithAlertText alert_text获取
func (obj *_WorkflowsTransitionMgr) WithAlertText(alertText string) Option {
	return optionFunc(func(o *options) { o.query["alert_text"] = alertText })
}

// WithDestStateID dest_state_id获取
func (obj *_WorkflowsTransitionMgr) WithDestStateID(destStateID int) Option {
	return optionFunc(func(o *options) { o.query["dest_state_id"] = destStateID })
}

// WithSourceStateID source_state_id获取
func (obj *_WorkflowsTransitionMgr) WithSourceStateID(sourceStateID int) Option {
	return optionFunc(func(o *options) { o.query["source_state_id"] = sourceStateID })
}

// WithWorkflowID workflow_id获取
func (obj *_WorkflowsTransitionMgr) WithWorkflowID(workflowID int) Option {
	return optionFunc(func(o *options) { o.query["workflow_id"] = workflowID })
}

// GetByOption 功能选项模式获取
func (obj *_WorkflowsTransitionMgr) GetByOption(opts ...Option) (result WorkflowsTransition, err error) {
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
			err = obj.DB.New().Table("workflows_state").Where("id = ?", result.DestStateID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsState = info
		}
		{
			var info WorkflowsState //
			err = obj.DB.New().Table("workflows_state").Where("id = ?", result.SourceStateID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsState = info
		}
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
func (obj *_WorkflowsTransitionMgr) GetByOptions(opts ...Option) (results []*WorkflowsTransition, err error) {
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
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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
func (obj *_WorkflowsTransitionMgr) GetFromID(id int) (result WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info WorkflowsState //
			err = obj.DB.New().Table("workflows_state").Where("id = ?", result.DestStateID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsState = info
		}
		{
			var info WorkflowsState //
			err = obj.DB.New().Table("workflows_state").Where("id = ?", result.SourceStateID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsState = info
		}
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
func (obj *_WorkflowsTransitionMgr) GetBatchFromID(ids []int) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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
func (obj *_WorkflowsTransitionMgr) GetFromCreateTime(createTime time.Time) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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
func (obj *_WorkflowsTransitionMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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
func (obj *_WorkflowsTransitionMgr) GetFromUpdateTime(updateTime time.Time) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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
func (obj *_WorkflowsTransitionMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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
func (obj *_WorkflowsTransitionMgr) GetFromMemo(memo string) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("memo = ?", memo).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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
func (obj *_WorkflowsTransitionMgr) GetBatchFromMemo(memos []string) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("memo IN (?)", memos).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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
func (obj *_WorkflowsTransitionMgr) GetFromName(name string) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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
func (obj *_WorkflowsTransitionMgr) GetBatchFromName(names []string) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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

// GetFromTransitionType 通过transition_type获取内容
func (obj *_WorkflowsTransitionMgr) GetFromTransitionType(transitionType string) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("transition_type = ?", transitionType).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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

// GetBatchFromTransitionType 批量唯一主键查找
func (obj *_WorkflowsTransitionMgr) GetBatchFromTransitionType(transitionTypes []string) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("transition_type IN (?)", transitionTypes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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

// GetFromTimer 通过timer获取内容
func (obj *_WorkflowsTransitionMgr) GetFromTimer(timer int) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("timer = ?", timer).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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

// GetBatchFromTimer 批量唯一主键查找
func (obj *_WorkflowsTransitionMgr) GetBatchFromTimer(timers []int) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("timer IN (?)", timers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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

// GetFromConditionExpression 通过condition_expression获取内容
func (obj *_WorkflowsTransitionMgr) GetFromConditionExpression(conditionExpression string) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("condition_expression = ?", conditionExpression).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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

// GetBatchFromConditionExpression 批量唯一主键查找
func (obj *_WorkflowsTransitionMgr) GetBatchFromConditionExpression(conditionExpressions []string) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("condition_expression IN (?)", conditionExpressions).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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

// GetFromAttributeType 通过attribute_type获取内容
func (obj *_WorkflowsTransitionMgr) GetFromAttributeType(attributeType string) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("attribute_type = ?", attributeType).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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

// GetBatchFromAttributeType 批量唯一主键查找
func (obj *_WorkflowsTransitionMgr) GetBatchFromAttributeType(attributeTypes []string) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("attribute_type IN (?)", attributeTypes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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

// GetFromAlertEnable 通过alert_enable获取内容
func (obj *_WorkflowsTransitionMgr) GetFromAlertEnable(alertEnable int8) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("alert_enable = ?", alertEnable).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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

// GetBatchFromAlertEnable 批量唯一主键查找
func (obj *_WorkflowsTransitionMgr) GetBatchFromAlertEnable(alertEnables []int8) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("alert_enable IN (?)", alertEnables).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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

// GetFromAlertText 通过alert_text获取内容
func (obj *_WorkflowsTransitionMgr) GetFromAlertText(alertText string) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("alert_text = ?", alertText).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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

// GetBatchFromAlertText 批量唯一主键查找
func (obj *_WorkflowsTransitionMgr) GetBatchFromAlertText(alertTexts []string) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("alert_text IN (?)", alertTexts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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

// GetFromDestStateID 通过dest_state_id获取内容
func (obj *_WorkflowsTransitionMgr) GetFromDestStateID(destStateID int) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("dest_state_id = ?", destStateID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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

// GetBatchFromDestStateID 批量唯一主键查找
func (obj *_WorkflowsTransitionMgr) GetBatchFromDestStateID(destStateIDs []int) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("dest_state_id IN (?)", destStateIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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

// GetFromSourceStateID 通过source_state_id获取内容
func (obj *_WorkflowsTransitionMgr) GetFromSourceStateID(sourceStateID int) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("source_state_id = ?", sourceStateID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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

// GetBatchFromSourceStateID 批量唯一主键查找
func (obj *_WorkflowsTransitionMgr) GetBatchFromSourceStateID(sourceStateIDs []int) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("source_state_id IN (?)", sourceStateIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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
func (obj *_WorkflowsTransitionMgr) GetFromWorkflowID(workflowID int) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("workflow_id = ?", workflowID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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
func (obj *_WorkflowsTransitionMgr) GetBatchFromWorkflowID(workflowIDs []int) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("workflow_id IN (?)", workflowIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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
func (obj *_WorkflowsTransitionMgr) FetchByPrimaryKey(id int) (result WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info WorkflowsState //
			err = obj.DB.New().Table("workflows_state").Where("id = ?", result.DestStateID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsState = info
		}
		{
			var info WorkflowsState //
			err = obj.DB.New().Table("workflows_state").Where("id = ?", result.SourceStateID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsState = info
		}
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
func (obj *_WorkflowsTransitionMgr) FetchByIndex(destStateID int, sourceStateID int, workflowID int) (results []*WorkflowsTransition, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("dest_state_id = ? AND source_state_id = ? AND workflow_id = ?", destStateID, sourceStateID, workflowID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].DestStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
			{
				var info WorkflowsState //
				err = obj.DB.New().Table("workflows_state").Where("id = ?", results[i].SourceStateID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsState = info
			}
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
