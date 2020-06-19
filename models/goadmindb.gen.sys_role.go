package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type _SysRoleMgr struct {
	*_BaseMgr
}

// SysRoleMgr open func
func SysRoleMgr(db *gorm.DB) *_SysRoleMgr {
	if db == nil {
		panic(fmt.Errorf("SysRoleMgr need init by db"))
	}
	return &_SysRoleMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysRoleMgr) GetTableName() string {
	return "sys_role"
}

// Get 获取
func (obj *_SysRoleMgr) Get() (result SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysRoleMgr) Gets() (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithRoleID role_id获取
func (obj *_SysRoleMgr) WithRoleID(roleID int) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithRoleName role_name获取
func (obj *_SysRoleMgr) WithRoleName(roleName string) Option {
	return optionFunc(func(o *options) { o.query["role_name"] = roleName })
}

// WithStatus status获取
func (obj *_SysRoleMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithRoleKey role_key获取
func (obj *_SysRoleMgr) WithRoleKey(roleKey string) Option {
	return optionFunc(func(o *options) { o.query["role_key"] = roleKey })
}

// WithRoleSort role_sort获取
func (obj *_SysRoleMgr) WithRoleSort(roleSort int) Option {
	return optionFunc(func(o *options) { o.query["role_sort"] = roleSort })
}

// WithFlag flag获取
func (obj *_SysRoleMgr) WithFlag(flag string) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithCreateBy create_by获取
func (obj *_SysRoleMgr) WithCreateBy(createBy string) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SysRoleMgr) WithUpdateBy(updateBy string) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithRemark remark获取
func (obj *_SysRoleMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithAdmin admin获取
func (obj *_SysRoleMgr) WithAdmin(admin string) Option {
	return optionFunc(func(o *options) { o.query["admin"] = admin })
}

// WithDataScope data_scope获取
func (obj *_SysRoleMgr) WithDataScope(dataScope string) Option {
	return optionFunc(func(o *options) { o.query["data_scope"] = dataScope })
}

// WithCreatedAt created_at获取
func (obj *_SysRoleMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SysRoleMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取
func (obj *_SysRoleMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// GetByOption 功能选项模式获取
func (obj *_SysRoleMgr) GetByOption(opts ...Option) (result SysRole, err error) {
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
func (obj *_SysRoleMgr) GetByOptions(opts ...Option) (results []*SysRole, err error) {
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

// GetFromRoleID 通过role_id获取内容
func (obj *_SysRoleMgr) GetFromRoleID(roleID int) (result SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("role_id = ?", roleID).Find(&result).Error

	return
}

// GetBatchFromRoleID 批量唯一主键查找
func (obj *_SysRoleMgr) GetBatchFromRoleID(roleIDs []int) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("role_id IN (?)", roleIDs).Find(&results).Error

	return
}

// GetFromRoleName 通过role_name获取内容
func (obj *_SysRoleMgr) GetFromRoleName(roleName string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("role_name = ?", roleName).Find(&results).Error

	return
}

// GetBatchFromRoleName 批量唯一主键查找
func (obj *_SysRoleMgr) GetBatchFromRoleName(roleNames []string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("role_name IN (?)", roleNames).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_SysRoleMgr) GetFromStatus(status int) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找
func (obj *_SysRoleMgr) GetBatchFromStatus(statuss []int) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromRoleKey 通过role_key获取内容
func (obj *_SysRoleMgr) GetFromRoleKey(roleKey string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("role_key = ?", roleKey).Find(&results).Error

	return
}

// GetBatchFromRoleKey 批量唯一主键查找
func (obj *_SysRoleMgr) GetBatchFromRoleKey(roleKeys []string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("role_key IN (?)", roleKeys).Find(&results).Error

	return
}

// GetFromRoleSort 通过role_sort获取内容
func (obj *_SysRoleMgr) GetFromRoleSort(roleSort int) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("role_sort = ?", roleSort).Find(&results).Error

	return
}

// GetBatchFromRoleSort 批量唯一主键查找
func (obj *_SysRoleMgr) GetBatchFromRoleSort(roleSorts []int) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("role_sort IN (?)", roleSorts).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_SysRoleMgr) GetFromFlag(flag string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("flag = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量唯一主键查找
func (obj *_SysRoleMgr) GetBatchFromFlag(flags []string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("flag IN (?)", flags).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SysRoleMgr) GetFromCreateBy(createBy string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_by = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量唯一主键查找
func (obj *_SysRoleMgr) GetBatchFromCreateBy(createBys []string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_by IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SysRoleMgr) GetFromUpdateBy(updateBy string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_by = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量唯一主键查找
func (obj *_SysRoleMgr) GetBatchFromUpdateBy(updateBys []string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_by IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容
func (obj *_SysRoleMgr) GetFromRemark(remark string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量唯一主键查找
func (obj *_SysRoleMgr) GetBatchFromRemark(remarks []string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

// GetFromAdmin 通过admin获取内容
func (obj *_SysRoleMgr) GetFromAdmin(admin string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("admin = ?", admin).Find(&results).Error

	return
}

// GetBatchFromAdmin 批量唯一主键查找
func (obj *_SysRoleMgr) GetBatchFromAdmin(admins []string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("admin IN (?)", admins).Find(&results).Error

	return
}

// GetFromDataScope 通过data_scope获取内容
func (obj *_SysRoleMgr) GetFromDataScope(dataScope string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("data_scope = ?", dataScope).Find(&results).Error

	return
}

// GetBatchFromDataScope 批量唯一主键查找
func (obj *_SysRoleMgr) GetBatchFromDataScope(dataScopes []string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("data_scope IN (?)", dataScopes).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SysRoleMgr) GetFromCreatedAt(createdAt time.Time) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找
func (obj *_SysRoleMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SysRoleMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找
func (obj *_SysRoleMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容
func (obj *_SysRoleMgr) GetFromDeletedAt(deletedAt time.Time) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("deleted_at = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量唯一主键查找
func (obj *_SysRoleMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("deleted_at IN (?)", deletedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysRoleMgr) FetchByPrimaryKey(roleID int) (result SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("role_id = ?", roleID).Find(&result).Error

	return
}

// FetchByIndex  获取多个内容
func (obj *_SysRoleMgr) FetchByIndex(deletedAt time.Time) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("deleted_at = ?", deletedAt).Find(&results).Error

	return
}
