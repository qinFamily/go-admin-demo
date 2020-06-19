package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type _WorkflowsWorkflowMgr struct {
	*_BaseMgr
}

// WorkflowsWorkflowMgr open func
func WorkflowsWorkflowMgr(db *gorm.DB) *_WorkflowsWorkflowMgr {
	if db == nil {
		panic(fmt.Errorf("WorkflowsWorkflowMgr need init by db"))
	}
	return &_WorkflowsWorkflowMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WorkflowsWorkflowMgr) GetTableName() string {
	return "workflows_workflow"
}

// Get 获取
func (obj *_WorkflowsWorkflowMgr) Get() (result WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info WorkflowsWorkflowtype //
			err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", result.TypeID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsWorkflowtype = info
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_WorkflowsWorkflowMgr) Gets() (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX:obj.isRelated", obj.isRelated)
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				fmt.Println(fmt.Sprintf("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX:%+v", results[i]))
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX:error", err)
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WorkflowsWorkflowMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCreateTime create_time获取
func (obj *_WorkflowsWorkflowMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_WorkflowsWorkflowMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithMemo memo获取
func (obj *_WorkflowsWorkflowMgr) WithMemo(memo string) Option {
	return optionFunc(func(o *options) { o.query["memo"] = memo })
}

// WithName name获取
func (obj *_WorkflowsWorkflowMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithTicketSnPrefix ticket_sn_prefix获取
func (obj *_WorkflowsWorkflowMgr) WithTicketSnPrefix(ticketSnPrefix string) Option {
	return optionFunc(func(o *options) { o.query["ticket_sn_prefix"] = ticketSnPrefix })
}

// WithStatus status获取
func (obj *_WorkflowsWorkflowMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithViewPermissionCheck view_permission_check获取
func (obj *_WorkflowsWorkflowMgr) WithViewPermissionCheck(viewPermissionCheck int8) Option {
	return optionFunc(func(o *options) { o.query["view_permission_check"] = viewPermissionCheck })
}

// WithLimitExpression limit_expression获取
func (obj *_WorkflowsWorkflowMgr) WithLimitExpression(limitExpression string) Option {
	return optionFunc(func(o *options) { o.query["limit_expression"] = limitExpression })
}

// WithDisplayFormStr display_form_str获取
func (obj *_WorkflowsWorkflowMgr) WithDisplayFormStr(displayFormStr string) Option {
	return optionFunc(func(o *options) { o.query["display_form_str"] = displayFormStr })
}

// WithTitleTemplate title_template获取
func (obj *_WorkflowsWorkflowMgr) WithTitleTemplate(titleTemplate string) Option {
	return optionFunc(func(o *options) { o.query["title_template"] = titleTemplate })
}

// WithTypeID type_id获取
func (obj *_WorkflowsWorkflowMgr) WithTypeID(typeID int) Option {
	return optionFunc(func(o *options) { o.query["type_id"] = typeID })
}

// GetByOption 功能选项模式获取
func (obj *_WorkflowsWorkflowMgr) GetByOption(opts ...Option) (result WorkflowsWorkflow, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info WorkflowsWorkflowtype //
			err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", result.TypeID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsWorkflowtype = info
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WorkflowsWorkflowMgr) GetByOptions(opts ...Option) (results []*WorkflowsWorkflow, err error) {
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
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_WorkflowsWorkflowMgr) GetFromID(id int) (result WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info WorkflowsWorkflowtype //
			err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", result.TypeID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsWorkflowtype = info
		}
	}

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_WorkflowsWorkflowMgr) GetBatchFromID(ids []int) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_WorkflowsWorkflowMgr) GetFromCreateTime(createTime time.Time) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量唯一主键查找
func (obj *_WorkflowsWorkflowMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_WorkflowsWorkflowMgr) GetFromUpdateTime(updateTime time.Time) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetBatchFromUpdateTime 批量唯一主键查找
func (obj *_WorkflowsWorkflowMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetFromMemo 通过memo获取内容
func (obj *_WorkflowsWorkflowMgr) GetFromMemo(memo string) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("memo = ?", memo).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetBatchFromMemo 批量唯一主键查找
func (obj *_WorkflowsWorkflowMgr) GetBatchFromMemo(memos []string) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("memo IN (?)", memos).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetFromName 通过name获取内容
func (obj *_WorkflowsWorkflowMgr) GetFromName(name string) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_WorkflowsWorkflowMgr) GetBatchFromName(names []string) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetFromTicketSnPrefix 通过ticket_sn_prefix获取内容
func (obj *_WorkflowsWorkflowMgr) GetFromTicketSnPrefix(ticketSnPrefix string) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("ticket_sn_prefix = ?", ticketSnPrefix).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetBatchFromTicketSnPrefix 批量唯一主键查找
func (obj *_WorkflowsWorkflowMgr) GetBatchFromTicketSnPrefix(ticketSnPrefixs []string) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("ticket_sn_prefix IN (?)", ticketSnPrefixs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetFromStatus 通过status获取内容
func (obj *_WorkflowsWorkflowMgr) GetFromStatus(status int8) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetBatchFromStatus 批量唯一主键查找
func (obj *_WorkflowsWorkflowMgr) GetBatchFromStatus(statuss []int8) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetFromViewPermissionCheck 通过view_permission_check获取内容
func (obj *_WorkflowsWorkflowMgr) GetFromViewPermissionCheck(viewPermissionCheck int8) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("view_permission_check = ?", viewPermissionCheck).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetBatchFromViewPermissionCheck 批量唯一主键查找
func (obj *_WorkflowsWorkflowMgr) GetBatchFromViewPermissionCheck(viewPermissionChecks []int8) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("view_permission_check IN (?)", viewPermissionChecks).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetFromLimitExpression 通过limit_expression获取内容
func (obj *_WorkflowsWorkflowMgr) GetFromLimitExpression(limitExpression string) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("limit_expression = ?", limitExpression).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetBatchFromLimitExpression 批量唯一主键查找
func (obj *_WorkflowsWorkflowMgr) GetBatchFromLimitExpression(limitExpressions []string) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("limit_expression IN (?)", limitExpressions).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetFromDisplayFormStr 通过display_form_str获取内容
func (obj *_WorkflowsWorkflowMgr) GetFromDisplayFormStr(displayFormStr string) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("display_form_str = ?", displayFormStr).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetBatchFromDisplayFormStr 批量唯一主键查找
func (obj *_WorkflowsWorkflowMgr) GetBatchFromDisplayFormStr(displayFormStrs []string) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("display_form_str IN (?)", displayFormStrs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetFromTitleTemplate 通过title_template获取内容
func (obj *_WorkflowsWorkflowMgr) GetFromTitleTemplate(titleTemplate string) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title_template = ?", titleTemplate).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetBatchFromTitleTemplate 批量唯一主键查找
func (obj *_WorkflowsWorkflowMgr) GetBatchFromTitleTemplate(titleTemplates []string) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title_template IN (?)", titleTemplates).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetFromTypeID 通过type_id获取内容
func (obj *_WorkflowsWorkflowMgr) GetFromTypeID(typeID int) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type_id = ?", typeID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

// GetBatchFromTypeID 批量唯一主键查找
func (obj *_WorkflowsWorkflowMgr) GetBatchFromTypeID(typeIDs []int) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type_id IN (?)", typeIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_WorkflowsWorkflowMgr) FetchByPrimaryKey(id int) (result WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info WorkflowsWorkflowtype //
			err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", result.TypeID).Find(&info).Error
			if err != nil {
				return
			}
			result.WorkflowsWorkflowtype = info
		}
	}

	return
}

// FetchByIndex  获取多个内容
func (obj *_WorkflowsWorkflowMgr) FetchByIndex(typeID int) (results []*WorkflowsWorkflow, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type_id = ?", typeID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info WorkflowsWorkflowtype //
				err = obj.DB.New().Table("workflows_workflowtype").Where("id = ?", results[i].TypeID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].WorkflowsWorkflowtype = info
			}
		}
	}
	return
}
