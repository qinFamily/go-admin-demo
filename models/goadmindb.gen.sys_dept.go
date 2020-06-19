package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type _SysDeptMgr struct {
	*_BaseMgr
}

// SysDeptMgr open func
func SysDeptMgr(db *gorm.DB) *_SysDeptMgr {
	if db == nil {
		panic(fmt.Errorf("SysDeptMgr need init by db"))
	}
	return &_SysDeptMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysDeptMgr) GetTableName() string {
	return "sys_dept"
}

// Get 获取
func (obj *_SysDeptMgr) Get() (result SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysDeptMgr) Gets() (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithDeptID dept_id获取
func (obj *_SysDeptMgr) WithDeptID(deptID int) Option {
	return optionFunc(func(o *options) { o.query["dept_id"] = deptID })
}

// WithParentID parent_id获取
func (obj *_SysDeptMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithDeptPath dept_path获取
func (obj *_SysDeptMgr) WithDeptPath(deptPath string) Option {
	return optionFunc(func(o *options) { o.query["dept_path"] = deptPath })
}

// WithDeptName dept_name获取
func (obj *_SysDeptMgr) WithDeptName(deptName string) Option {
	return optionFunc(func(o *options) { o.query["dept_name"] = deptName })
}

// WithSort sort获取
func (obj *_SysDeptMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithLeader leader获取
func (obj *_SysDeptMgr) WithLeader(leader string) Option {
	return optionFunc(func(o *options) { o.query["leader"] = leader })
}

// WithPhone phone获取
func (obj *_SysDeptMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithEmail email获取
func (obj *_SysDeptMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithStatus status获取
func (obj *_SysDeptMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateBy create_by获取
func (obj *_SysDeptMgr) WithCreateBy(createBy string) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SysDeptMgr) WithUpdateBy(updateBy string) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithCreatedAt created_at获取
func (obj *_SysDeptMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SysDeptMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取
func (obj *_SysDeptMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// GetByOption 功能选项模式获取
func (obj *_SysDeptMgr) GetByOption(opts ...Option) (result SysDept, err error) {
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
func (obj *_SysDeptMgr) GetByOptions(opts ...Option) (results []*SysDept, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromDeptID 通过dept_id获取内容
func (obj *_SysDeptMgr) GetFromDeptID(deptID int) (result SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("dept_id = ?", deptID).Find(&result).Error

	return
}

// GetBatchFromDeptID 批量唯一主键查找
func (obj *_SysDeptMgr) GetBatchFromDeptID(deptIDs []int) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("dept_id IN (?)", deptIDs).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容
func (obj *_SysDeptMgr) GetFromParentID(parentID int) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("parent_id = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量唯一主键查找
func (obj *_SysDeptMgr) GetBatchFromParentID(parentIDs []int) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("parent_id IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromDeptPath 通过dept_path获取内容
func (obj *_SysDeptMgr) GetFromDeptPath(deptPath string) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("dept_path = ?", deptPath).Find(&results).Error

	return
}

// GetBatchFromDeptPath 批量唯一主键查找
func (obj *_SysDeptMgr) GetBatchFromDeptPath(deptPaths []string) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("dept_path IN (?)", deptPaths).Find(&results).Error

	return
}

// GetFromDeptName 通过dept_name获取内容
func (obj *_SysDeptMgr) GetFromDeptName(deptName string) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("dept_name = ?", deptName).Find(&results).Error

	return
}

// GetBatchFromDeptName 批量唯一主键查找
func (obj *_SysDeptMgr) GetBatchFromDeptName(deptNames []string) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("dept_name IN (?)", deptNames).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容
func (obj *_SysDeptMgr) GetFromSort(sort int) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sort = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量唯一主键查找
func (obj *_SysDeptMgr) GetBatchFromSort(sorts []int) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sort IN (?)", sorts).Find(&results).Error

	return
}

// GetFromLeader 通过leader获取内容
func (obj *_SysDeptMgr) GetFromLeader(leader string) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("leader = ?", leader).Find(&results).Error

	return
}

// GetBatchFromLeader 批量唯一主键查找
func (obj *_SysDeptMgr) GetBatchFromLeader(leaders []string) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("leader IN (?)", leaders).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容
func (obj *_SysDeptMgr) GetFromPhone(phone string) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("phone = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量唯一主键查找
func (obj *_SysDeptMgr) GetBatchFromPhone(phones []string) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("phone IN (?)", phones).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容
func (obj *_SysDeptMgr) GetFromEmail(email string) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量唯一主键查找
func (obj *_SysDeptMgr) GetBatchFromEmail(emails []string) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_SysDeptMgr) GetFromStatus(status int) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找
func (obj *_SysDeptMgr) GetBatchFromStatus(statuss []int) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SysDeptMgr) GetFromCreateBy(createBy string) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_by = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量唯一主键查找
func (obj *_SysDeptMgr) GetBatchFromCreateBy(createBys []string) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_by IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SysDeptMgr) GetFromUpdateBy(updateBy string) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_by = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量唯一主键查找
func (obj *_SysDeptMgr) GetBatchFromUpdateBy(updateBys []string) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_by IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SysDeptMgr) GetFromCreatedAt(createdAt time.Time) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找
func (obj *_SysDeptMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SysDeptMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找
func (obj *_SysDeptMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容
func (obj *_SysDeptMgr) GetFromDeletedAt(deletedAt time.Time) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("deleted_at = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量唯一主键查找
func (obj *_SysDeptMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("deleted_at IN (?)", deletedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysDeptMgr) FetchByPrimaryKey(deptID int) (result SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("dept_id = ?", deptID).Find(&result).Error

	return
}

// FetchByIndex  获取多个内容
func (obj *_SysDeptMgr) FetchByIndex(deletedAt time.Time) (results []*SysDept, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("deleted_at = ?", deletedAt).Find(&results).Error

	return
}
