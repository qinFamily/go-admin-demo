package models

import (
	orm "go-admin-demo/database"
	"go-admin-demo/tools"

	// "log"
	_ "time"
)

type GridBasicPeople struct {
	PeopleId                int    `json:"peopleId" gorm:"column:peopleId;type:bigint(20) unsigned;primary_key"`            // 居民ID
	GridId                  int    `json:"gridId" gorm:"column:gridId;type:bigint(20) unsigned;"`                           // 网格ID
	Name                    string `json:"name" gorm:"column:name;type:varchar(10);"`                                       // 姓名
	Alias                   string `json:"alias" gorm:"column:alias;type:varchar(10);"`                                     // 别名，曾用名
	Sex                     string `json:"sex" gorm:"column:sex;type:varchar(2);"`                                          // 性别
	CardType                int    `json:"cardType" gorm:"column:cardType;type:tinyint(4);"`                                // 证件类型，0-公民身份证,1-军官证,2-驾驶证,3-护照,9-其他
	CardNo                  string `json:"cardNo" gorm:"column:cardNo;type:varchar(20);"`                                   // 证件号码
	PhoneNo                 string `json:"phoneNo" gorm:"column:phoneNo;type:varchar(18);"`                                 // 联系电话
	ResidenceAddress        string `json:"residenceAddress" gorm:"column:residenceAddress;type:varchar(100);"`              // 现住址
	ReleatedStatus          int    `json:"releatedStatus" gorm:"column:releatedStatus;type:tinyint(4);"`                    // 关联状态，0-未关联， 1-已关联
	HouseId                 int    `json:"houseId" gorm:"column:houseId;type:bigint(20);"`                                  // 房号，跟房号关联
	BirthDate               string `json:"birthDate" gorm:"column:birthDate;type:date;"`                                    // 生日
	Nationality             string `json:"nationality" gorm:"column:nationality;type:varchar(10);"`                         // 民族
	PermanentAddress        string `json:"permanentAddress" gorm:"column:permanentAddress;type:varchar(100);"`              // 户籍地址
	HouseholderRelationship string `json:"householderRelationship" gorm:"column:householderRelationship;type:varchar(50);"` // 与户主关系
	HouseholdStatus         int    `json:"householdStatus" gorm:"column:householdStatus;type:tinyint(4);"`                  // 人户状态,0-未入户， 1-入户， 2-未知
	NativePlace             string `json:"nativePlace" gorm:"column:nativePlace;type:varchar(10);"`                         // 籍贯
	Education               int    `json:"education" gorm:"column:education;type:tinyint(4);"`                              // 文化程度，0-无，1-小学，2-初中，3-高中，4-大专，5-本科，6-重点本科，7-硕士研究生，8-博士研究生，9-其他
	Marital                 int    `json:"marital" gorm:"column:marital;type:tinyint(4);"`                                  // 婚姻状况	，0-未婚，1-已婚，2-未知
	EmploymentStatus        int    `json:"employmentStatus" gorm:"column:employmentStatus;type:tinyint(4);"`                // 从业状况, 0-待业，1-就业，3-其他
	IsEmployees             int    `json:"isEmployees" gorm:"column:isEmployees;type:tinyint(4);"`                          // 是否从业人员,0-是，1-否
	Work                    string `json:"work" gorm:"column:work;type:text;"`                                              // 服务处所
	Company                 string `json:"company" gorm:"column:company;type:text;"`                                        // 工作单位
	PoliticalStatus         int    `json:"politicalStatus" gorm:"column:politicalStatus;type:varchar(50);"`                 // 政治面貌
	Remarks                 string `json:"remarks" gorm:"column:remarks;type:text;"`                                        // 备注
	CreateBy                string `json:"createBy" gorm:"column:create_by;type:varchar(2);"`                               // 创建人
	UpdateBy                string `json:"updateBy" gorm:"column:update_by;type:varchar(2);"`                               // 更新人
	DataScope               string `json:"dataScope" gorm:"-"`
	Params                  string `json:"params"  gorm:"-"`
	BaseModel
}

func (GridBasicPeople) TableName() string {
	return "grid_basic_people"
}

// 创建GridBasicPeople
func (e *GridBasicPeople) Create() (GridBasicPeople, error) {
	var doc GridBasicPeople
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取GridBasicPeople
func (e *GridBasicPeople) Get() (doc GridBasicPeople, err error) {

	table := orm.Eloquent.Table(e.TableName()).Where("peopleId = ?", e.PeopleId)
	// log.Println("=============================", e.PeopleId)
	if err = table.First(&doc).Error; err != nil {
		return
	}
	return
}

// 获取GridBasicPeople带分页
func (e *GridBasicPeople) GetPage(pageSize int, pageIndex int) ([]GridBasicPeople, int, error) {
	var doc []GridBasicPeople

	table := orm.Eloquent.Select("*").Table(e.TableName())

	// 数据权限控制(如果不需要数据权限请将此处去掉)
	dataPermission := new(DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table = dataPermission.GetDataScope(e.TableName(), table)

	var count int

	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Count(&count)
	return doc, count, nil
}

// 更新GridBasicPeople
func (e *GridBasicPeople) Update(id int) (update GridBasicPeople, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("peopleId = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 删除GridBasicPeople
func (e *GridBasicPeople) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("peopleId = ?", id).Delete(e).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *GridBasicPeople) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("peopleId in (?)", id).Delete(e).Error; err != nil {
		return
	}
	Result = true
	return
}
