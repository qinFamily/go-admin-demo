package models

import (
	"time"
)

// SysColumns [...]
type SysColumns struct {
	ColumnID      int       `gorm:"primary_key;column:column_id;type:int(11);not null" json:"column_id"`
	TableID       int       `gorm:"column:table_id;type:int(11)" json:"table_id"`
	ColumnName    string    `gorm:"column:column_name;type:varchar(128)" json:"column_name"`
	ColumnComment string    `gorm:"column:column_comment;type:varchar(128)" json:"column_comment"`
	ColumnType    string    `gorm:"column:column_type;type:varchar(128)" json:"column_type"`
	GoType        string    `gorm:"column:go_type;type:varchar(128)" json:"go_type"`
	GoField       string    `gorm:"column:go_field;type:varchar(128)" json:"go_field"`
	JSONField     string    `gorm:"column:json_field;type:varchar(128)" json:"json_field"`
	IsPk          string    `gorm:"column:is_pk;type:char(4)" json:"is_pk"`
	IsIncrement   string    `gorm:"column:is_increment;type:char(4)" json:"is_increment"`
	IsRequired    string    `gorm:"column:is_required;type:char(4)" json:"is_required"`
	IsInsert      string    `gorm:"column:is_insert;type:char(4)" json:"is_insert"`
	IsEdit        string    `gorm:"column:is_edit;type:char(4)" json:"is_edit"`
	IsList        string    `gorm:"column:is_list;type:char(4)" json:"is_list"`
	IsQuery       string    `gorm:"column:is_query;type:char(4)" json:"is_query"`
	QueryType     string    `gorm:"column:query_type;type:varchar(128)" json:"query_type"`
	HTMLType      string    `gorm:"column:html_type;type:varchar(128)" json:"html_type"`
	DictType      string    `gorm:"column:dict_type;type:varchar(128)" json:"dict_type"`
	Sort          int       `gorm:"column:sort;type:int(4)" json:"sort"`
	List          string    `gorm:"column:list;type:char(1)" json:"list"`
	Pk            string    `gorm:"column:pk;type:char(1)" json:"pk"`
	Required      string    `gorm:"column:required;type:char(1)" json:"required"`
	SuperColumn   string    `gorm:"column:super_column;type:char(1)" json:"super_column"`
	UsableColumn  string    `gorm:"column:usable_column;type:char(1)" json:"usable_column"`
	Increment     string    `gorm:"column:increment;type:char(1)" json:"increment"`
	Insert        string    `gorm:"column:insert;type:char(1)" json:"insert"`
	Edit          string    `gorm:"column:edit;type:char(1)" json:"edit"`
	Query         string    `gorm:"column:query;type:char(1)" json:"query"`
	Remark        string    `gorm:"column:remark;type:varchar(255)" json:"remark"`
	CreateBy      string    `gorm:"column:create_by;type:varchar(128)" json:"create_by"`
	UpdateBy      string    `gorm:"column:update_By;type:varchar(128)" json:"update_by"`
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt     time.Time `gorm:"index;column:deleted_at;type:timestamp" json:"deleted_at"`
}

// SysDept [...]
type SysDept struct {
	DeptID    int       `gorm:"primary_key;column:dept_id;type:int(11);not null" json:"dept_id"`
	ParentID  int       `gorm:"column:parent_id;type:int(11)" json:"parent_id"`
	DeptPath  string    `gorm:"column:dept_path;type:varchar(255)" json:"dept_path"`
	DeptName  string    `gorm:"column:dept_name;type:varchar(128)" json:"dept_name"`
	Sort      int       `gorm:"column:sort;type:int(4)" json:"sort"`
	Leader    string    `gorm:"column:leader;type:varchar(128)" json:"leader"`
	Phone     string    `gorm:"column:phone;type:varchar(11)" json:"phone"`
	Email     string    `gorm:"column:email;type:varchar(64)" json:"email"`
	Status    int       `gorm:"column:status;type:int(1)" json:"status"`
	CreateBy  string    `gorm:"column:create_by;type:varchar(64)" json:"create_by"`
	UpdateBy  string    `gorm:"column:update_by;type:varchar(64)" json:"update_by"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt time.Time `gorm:"index;column:deleted_at;type:timestamp" json:"deleted_at"`
}

// SysDictData [...]
type SysDictData struct {
	DictCode  int       `gorm:"primary_key;column:dict_code;type:int(11);not null" json:"dict_code"`
	DictSort  int       `gorm:"column:dict_sort;type:int(4)" json:"dict_sort"`
	DictLabel string    `gorm:"column:dict_label;type:varchar(128)" json:"dict_label"`
	DictValue string    `gorm:"column:dict_value;type:varchar(255)" json:"dict_value"`
	DictType  string    `gorm:"column:dict_type;type:varchar(64)" json:"dict_type"`
	CSSClass  string    `gorm:"column:css_class;type:varchar(128)" json:"css_class"`
	ListClass string    `gorm:"column:list_class;type:varchar(128)" json:"list_class"`
	IsDefault string    `gorm:"column:is_default;type:varchar(8)" json:"is_default"`
	Status    int       `gorm:"column:status;type:int(1)" json:"status"`
	Default   string    `gorm:"column:default;type:varchar(8)" json:"default"`
	CreateBy  string    `gorm:"column:create_by;type:varchar(64)" json:"create_by"`
	UpdateBy  string    `gorm:"column:update_by;type:varchar(64)" json:"update_by"`
	Remark    string    `gorm:"column:remark;type:varchar(255)" json:"remark"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt time.Time `gorm:"index;column:deleted_at;type:timestamp" json:"deleted_at"`
}

// SysDictType [...]
type SysDictType struct {
	DictID    int       `gorm:"primary_key;column:dict_id;type:int(11);not null" json:"dict_id"`
	DictName  string    `gorm:"column:dict_name;type:varchar(128)" json:"dict_name"`
	DictType  string    `gorm:"column:dict_type;type:varchar(128)" json:"dict_type"`
	Status    int       `gorm:"column:status;type:int(1)" json:"status"`
	CreateBy  string    `gorm:"column:create_by;type:varchar(11)" json:"create_by"`
	UpdateBy  string    `gorm:"column:update_by;type:varchar(11)" json:"update_by"`
	Remark    string    `gorm:"column:remark;type:varchar(255)" json:"remark"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt time.Time `gorm:"index;column:deleted_at;type:timestamp" json:"deleted_at"`
}

// SysLoginlog [...]
type SysLoginlog struct {
	InfoID        int       `gorm:"primary_key;column:info_id;type:int(11);not null" json:"info_id"`
	Username      string    `gorm:"column:username;type:varchar(128)" json:"username"`
	Status        int       `gorm:"column:status;type:int(1)" json:"status"`
	IPaddr        string    `gorm:"column:ipaddr;type:varchar(255)" json:"ipaddr"`
	LoginLocation string    `gorm:"column:login_location;type:varchar(255)" json:"login_location"`
	Browser       string    `gorm:"column:browser;type:varchar(255)" json:"browser"`
	Os            string    `gorm:"column:os;type:varchar(255)" json:"os"`
	Platform      string    `gorm:"column:platform;type:varchar(255)" json:"platform"`
	LoginTime     time.Time `gorm:"column:login_time;type:timestamp;not null" json:"login_time"`
	CreateBy      string    `gorm:"column:create_by;type:varchar(128)" json:"create_by"`
	UpdateBy      string    `gorm:"column:update_by;type:varchar(128)" json:"update_by"`
	Remark        string    `gorm:"column:remark;type:varchar(255)" json:"remark"`
	Msg           string    `gorm:"column:msg;type:varchar(255)" json:"msg"`
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt     time.Time `gorm:"index;column:deleted_at;type:timestamp" json:"deleted_at"`
}

// SysMenu [...]
type SysMenu struct {
	MenuID     int       `gorm:"primary_key;column:menu_id;type:int(11);not null" json:"menu_id"`
	MenuName   string    `gorm:"column:menu_name;type:varchar(11)" json:"menu_name"`
	Title      string    `gorm:"column:title;type:varchar(64)" json:"title"`
	Icon       string    `gorm:"column:icon;type:varchar(128)" json:"icon"`
	Path       string    `gorm:"column:path;type:varchar(128)" json:"path"`
	Paths      string    `gorm:"column:paths;type:varchar(128)" json:"paths"`
	MenuType   string    `gorm:"column:menu_type;type:varchar(1)" json:"menu_type"`
	Action     string    `gorm:"column:action;type:varchar(16)" json:"action"`
	Permission string    `gorm:"column:permission;type:varchar(32)" json:"permission"`
	ParentID   int       `gorm:"column:parent_id;type:int(11)" json:"parent_id"`
	NoCache    string    `gorm:"column:no_cache;type:char(1)" json:"no_cache"`
	Breadcrumb string    `gorm:"column:breadcrumb;type:varchar(255)" json:"breadcrumb"`
	Component  string    `gorm:"column:component;type:varchar(255)" json:"component"`
	Sort       int       `gorm:"column:sort;type:int(4)" json:"sort"`
	Visible    string    `gorm:"column:visible;type:char(1)" json:"visible"`
	CreateBy   string    `gorm:"column:create_by;type:varchar(128)" json:"create_by"`
	UpdateBy   string    `gorm:"column:update_by;type:varchar(128)" json:"update_by"`
	IsFrame    int       `gorm:"column:is_frame;type:int(1)" json:"is_frame"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt  time.Time `gorm:"index;column:deleted_at;type:timestamp" json:"deleted_at"`
}

// SysOperlog [...]
type SysOperlog struct {
	OperID        int       `gorm:"primary_key;column:oper_id;type:int(11);not null" json:"oper_id"`
	Title         string    `gorm:"column:title;type:varchar(255)" json:"title"`
	BusinessType  string    `gorm:"column:business_type;type:varchar(128)" json:"business_type"`
	BusinessTypes string    `gorm:"column:business_types;type:varchar(128)" json:"business_types"`
	Method        string    `gorm:"column:method;type:varchar(128)" json:"method"`
	RequestMethod string    `gorm:"column:request_method;type:varchar(128)" json:"request_method"`
	OperatorType  string    `gorm:"column:operator_type;type:varchar(128)" json:"operator_type"`
	OperName      string    `gorm:"column:oper_name;type:varchar(128)" json:"oper_name"`
	DeptName      string    `gorm:"column:dept_name;type:varchar(128)" json:"dept_name"`
	OperURL       string    `gorm:"column:oper_url;type:varchar(255)" json:"oper_url"`
	OperIP        string    `gorm:"column:oper_ip;type:varchar(128)" json:"oper_ip"`
	OperLocation  string    `gorm:"column:oper_location;type:varchar(128)" json:"oper_location"`
	OperParam     string    `gorm:"column:oper_param;type:varchar(255)" json:"oper_param"`
	Status        int       `gorm:"column:status;type:int(1)" json:"status"`
	OperTime      time.Time `gorm:"column:oper_time;type:timestamp;not null" json:"oper_time"`
	JSONResult    string    `gorm:"column:json_result;type:varchar(255)" json:"json_result"`
	CreateBy      string    `gorm:"column:create_by;type:varchar(128)" json:"create_by"`
	UpdateBy      string    `gorm:"column:update_by;type:varchar(128)" json:"update_by"`
	Remark        string    `gorm:"column:remark;type:varchar(255)" json:"remark"`
	LatencyTime   string    `gorm:"column:latency_time;type:varchar(128)" json:"latency_time"`
	UserAgent     string    `gorm:"column:user_agent;type:varchar(255)" json:"user_agent"`
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt     time.Time `gorm:"index;column:deleted_at;type:timestamp" json:"deleted_at"`
}

// SysPost [...]
type SysPost struct {
	PostID    int       `gorm:"primary_key;column:post_id;type:int(11);not null" json:"post_id"`
	PostName  string    `gorm:"column:post_name;type:varchar(128)" json:"post_name"`
	PostCode  string    `gorm:"column:post_code;type:varchar(128)" json:"post_code"`
	Sort      int       `gorm:"column:sort;type:int(4)" json:"sort"`
	Status    int       `gorm:"column:status;type:int(1)" json:"status"`
	Remark    string    `gorm:"column:remark;type:varchar(255)" json:"remark"`
	CreateBy  string    `gorm:"column:create_by;type:varchar(128)" json:"create_by"`
	UpdateBy  string    `gorm:"column:update_by;type:varchar(128)" json:"update_by"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt time.Time `gorm:"index;column:deleted_at;type:timestamp" json:"deleted_at"`
}

// SysRoleMenu [...]
type SysRoleMenu struct {
	RoleID   int    `gorm:"column:role_id;type:int(11)" json:"role_id"`
	MenuID   int    `gorm:"column:menu_id;type:int(11)" json:"menu_id"`
	RoleName string `gorm:"column:role_name;type:varchar(128)" json:"role_name"`
	CreateBy string `gorm:"column:create_by;type:varchar(128)" json:"create_by"`
	UpdateBy string `gorm:"column:update_by;type:varchar(128)" json:"update_by"`
}

// SysTables [...]
type SysTables struct {
	TableID             int       `gorm:"primary_key;column:table_id;type:int(11);not null" json:"table_id"`
	TableName           string    `gorm:"column:table_name;type:varchar(255)" json:"table_name"`
	TableComment        string    `gorm:"column:table_comment;type:varchar(255)" json:"table_comment"`
	ClassName           string    `gorm:"column:class_name;type:varchar(255)" json:"class_name"`
	TplCategory         string    `gorm:"column:tpl_category;type:varchar(255)" json:"tpl_category"`
	PackageName         string    `gorm:"column:package_name;type:varchar(255)" json:"package_name"`
	ModuleName          string    `gorm:"column:module_name;type:varchar(255)" json:"module_name"`
	BusinessName        string    `gorm:"column:business_name;type:varchar(255)" json:"business_name"`
	FunctionName        string    `gorm:"column:function_name;type:varchar(255)" json:"function_name"`
	FunctionAuthor      string    `gorm:"column:function_author;type:varchar(255)" json:"function_author"`
	PkColumn            string    `gorm:"column:pk_column;type:varchar(255)" json:"pk_column"`
	PkGoField           string    `gorm:"column:pk_go_field;type:varchar(255)" json:"pk_go_field"`
	PkJSONField         string    `gorm:"column:pk_json_field;type:varchar(255)" json:"pk_json_field"`
	Options             string    `gorm:"column:options;type:varchar(255)" json:"options"`
	TreeCode            string    `gorm:"column:tree_code;type:varchar(255)" json:"tree_code"`
	TreeParentCode      string    `gorm:"column:tree_parent_code;type:varchar(255)" json:"tree_parent_code"`
	TreeName            string    `gorm:"column:tree_name;type:varchar(255)" json:"tree_name"`
	Tree                string    `gorm:"column:tree;type:char(1)" json:"tree"`
	Crud                string    `gorm:"column:crud;type:char(1)" json:"crud"`
	Remark              string    `gorm:"column:remark;type:varchar(255)" json:"remark"`
	IsLogicalDelete     string    `gorm:"column:is_logical_delete;type:char(1)" json:"is_logical_delete"`
	LogicalDelete       string    `gorm:"column:logical_delete;type:char(1)" json:"logical_delete"`
	LogicalDeleteColumn string    `gorm:"column:logical_delete_column;type:varchar(128)" json:"logical_delete_column"`
	CreateBy            string    `gorm:"column:create_by;type:varchar(128)" json:"create_by"`
	UpdateBy            string    `gorm:"column:update_by;type:varchar(128)" json:"update_by"`
	CreatedAt           time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt           time.Time `gorm:"index;column:deleted_at;type:timestamp" json:"deleted_at"`
}

// WorkflowsStateFields [...]
type WorkflowsStateFields struct {
	ID                   int                  `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	StateID              int                  `gorm:"unique_index:workflows_state_fields_state_id_customfield_id_uniq;index;column:state_id;type:int(11);not null" json:"state_id"`
	WorkflowsState       WorkflowsState       `gorm:"association_foreignkey:state_id;foreignkey:id" json:"workflows_state_list"`
	CustomfieldID        int                  `gorm:"unique_index:workflows_state_fields_state_id_customfield_id_uniq;index;column:customfield_id;type:int(11);not null" json:"customfield_id"`
	WorkflowsCustomfield WorkflowsCustomfield `gorm:"association_foreignkey:customfield_id;foreignkey:id" json:"workflows_customfield_list"`
}

// WorkflowsStateGroupParticipant [...]
type WorkflowsStateGroupParticipant struct {
	ID             int            `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	StateID        int            `gorm:"unique_index:workflows_state_group_participant_state_id_group_id_uniq;index;column:state_id;type:int(11);not null" json:"state_id"`
	WorkflowsState WorkflowsState `gorm:"association_foreignkey:state_id;foreignkey:id" json:"workflows_state_list"`
	GroupID        int            `gorm:"unique_index:workflows_state_group_participant_state_id_group_id_uniq;index;column:group_id;type:int(11);not null" json:"group_id"`
	SysDept        SysDept        `gorm:"association_foreignkey:group_id;foreignkey:dept_id" json:"sys_dept_list"`
}

// WorkflowsStateRoleParticipant [...]
type WorkflowsStateRoleParticipant struct {
	ID             int            `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	StateID        int            `gorm:"unique_index:workflows_state_role_participant_state_id_role_id_uniq;index;column:state_id;type:int(11);not null" json:"state_id"`
	WorkflowsState WorkflowsState `gorm:"association_foreignkey:state_id;foreignkey:id" json:"workflows_state_list"`
	RoleID         int            `gorm:"unique_index:workflows_state_role_participant_state_id_role_id_uniq;index;column:role_id;type:int(11);not null" json:"role_id"`
	SysRole        SysRole        `gorm:"association_foreignkey:role_id;foreignkey:role_id" json:"sys_role_list"`
}

// WorkflowsStateUserParticipant [...]
type WorkflowsStateUserParticipant struct {
	ID             int            `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	StateID        int            `gorm:"unique_index:workflows_state_user_participant_state_id_user_id_uniq;index;column:state_id;type:int(11);not null" json:"state_id"`
	WorkflowsState WorkflowsState `gorm:"association_foreignkey:state_id;foreignkey:id" json:"workflows_state_list"`
	UserID         int            `gorm:"unique_index:workflows_state_user_participant_state_id_user_id_uniq;index;column:user_id;type:int(11);not null" json:"user_id"`
	SysUser        SysUser        `gorm:"association_foreignkey:user_id;foreignkey:user_id" json:"sys_user_list"`
}
