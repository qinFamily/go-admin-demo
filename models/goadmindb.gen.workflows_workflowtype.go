package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type _WorkflowsWorkflowtypeMgr struct {
	*_BaseMgr
}

// WorkflowsWorkflowtypeMgr open func
func WorkflowsWorkflowtypeMgr(db *gorm.DB) *_WorkflowsWorkflowtypeMgr {
	if db == nil {
		panic(fmt.Errorf("WorkflowsWorkflowtypeMgr need init by db"))
	}
	return &_WorkflowsWorkflowtypeMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WorkflowsWorkflowtypeMgr) GetTableName() string {
	return "workflows_workflowtype"
}

// Get 获取
func (obj *_WorkflowsWorkflowtypeMgr) Get() (result WorkflowsWorkflowtype, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WorkflowsWorkflowtypeMgr) Gets() (results []*WorkflowsWorkflowtype, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WorkflowsWorkflowtypeMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCreateTime create_time获取
func (obj *_WorkflowsWorkflowtypeMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_WorkflowsWorkflowtypeMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithMemo memo获取
func (obj *_WorkflowsWorkflowtypeMgr) WithMemo(memo string) Option {
	return optionFunc(func(o *options) { o.query["memo"] = memo })
}

// WithName name获取
func (obj *_WorkflowsWorkflowtypeMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCode code获取
func (obj *_WorkflowsWorkflowtypeMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithOrderID order_id获取
func (obj *_WorkflowsWorkflowtypeMgr) WithOrderID(orderID int) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// GetByOption 功能选项模式获取
func (obj *_WorkflowsWorkflowtypeMgr) GetByOption(opts ...Option) (result WorkflowsWorkflowtype, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WorkflowsWorkflowtypeMgr) GetByOptions(opts ...Option) (results []*WorkflowsWorkflowtype, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

// Gets 获取批量结果
func (obj *_WorkflowsWorkflowtypeMgr) GetsWithWorkflowSet() (results []*WorkflowtypeWorkflowsSet, err error) {
	resultWT := make([]*WorkflowsWorkflowtype, 0)
	err = obj.DB.Table(obj.GetTableName()).Find(&resultWT).Error
	if err != nil {
		return
	}
	wfwfmgr := WorkflowsWorkflowMgr(obj.DB)
	results = make([]*WorkflowtypeWorkflowsSet, 0)
	for _, rwt := range resultWT {
		wfs, err := wfwfmgr.GetFromTypeID(rwt.ID)
		if err != nil {
			continue
		}
		wfwfwt := make([]*WorkflowsWorkflowWioutType, 0)
		for _, w := range wfs {
			wfwfwt = append(wfwfwt, &WorkflowsWorkflowWioutType{
				ID:                  w.ID,
				CreateTime:          w.CreateTime,
				UpdateTime:          w.UpdateTime,
				Memo:                w.Memo,
				Name:                w.Name,
				TicketSnPrefix:      w.TicketSnPrefix,
				Status:              w.Status,
				ViewPermissionCheck: w.ViewPermissionCheck,
				LimitExpression:     w.LimitExpression,
				DisplayFormStr:      w.DisplayFormStr,
				TitleTemplate:       w.TitleTemplate,
				Workflowtype:        rwt.ID,
			})
		}
		wftwfs := &WorkflowtypeWorkflowsSet{rwt, wfwfwt}
		results = append(results, wftwfs)
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_WorkflowsWorkflowtypeMgr) GetFromID(id int) (result WorkflowsWorkflowtype, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_WorkflowsWorkflowtypeMgr) GetBatchFromID(ids []int) (results []*WorkflowsWorkflowtype, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_WorkflowsWorkflowtypeMgr) GetFromCreateTime(createTime time.Time) (results []*WorkflowsWorkflowtype, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找
func (obj *_WorkflowsWorkflowtypeMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WorkflowsWorkflowtype, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_WorkflowsWorkflowtypeMgr) GetFromUpdateTime(updateTime time.Time) (results []*WorkflowsWorkflowtype, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找
func (obj *_WorkflowsWorkflowtypeMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WorkflowsWorkflowtype, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromMemo 通过memo获取内容
func (obj *_WorkflowsWorkflowtypeMgr) GetFromMemo(memo string) (results []*WorkflowsWorkflowtype, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("memo = ?", memo).Find(&results).Error

	return
}

// GetBatchFromMemo 批量唯一主键查找
func (obj *_WorkflowsWorkflowtypeMgr) GetBatchFromMemo(memos []string) (results []*WorkflowsWorkflowtype, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("memo IN (?)", memos).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_WorkflowsWorkflowtypeMgr) GetFromName(name string) (results []*WorkflowsWorkflowtype, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_WorkflowsWorkflowtypeMgr) GetBatchFromName(names []string) (results []*WorkflowsWorkflowtype, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容
func (obj *_WorkflowsWorkflowtypeMgr) GetFromCode(code string) (result WorkflowsWorkflowtype, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("code = ?", code).Find(&result).Error

	return
}

// GetBatchFromCode 批量唯一主键查找
func (obj *_WorkflowsWorkflowtypeMgr) GetBatchFromCode(codes []string) (results []*WorkflowsWorkflowtype, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("code IN (?)", codes).Find(&results).Error

	return
}

// GetFromOrderID 通过order_id获取内容
func (obj *_WorkflowsWorkflowtypeMgr) GetFromOrderID(orderID int) (results []*WorkflowsWorkflowtype, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_id = ?", orderID).Find(&results).Error

	return
}

// GetBatchFromOrderID 批量唯一主键查找
func (obj *_WorkflowsWorkflowtypeMgr) GetBatchFromOrderID(orderIDs []int) (results []*WorkflowsWorkflowtype, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_id IN (?)", orderIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_WorkflowsWorkflowtypeMgr) FetchByPrimaryKey(id int) (result WorkflowsWorkflowtype, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchByUnique primay or index 获取唯一内容
func (obj *_WorkflowsWorkflowtypeMgr) FetchByUnique(code string) (result WorkflowsWorkflowtype, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("code = ?", code).Find(&result).Error

	return
}
