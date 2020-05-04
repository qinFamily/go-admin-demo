package models

import (
	orm "go-admin-demo/database"
	"go-admin-demo/tools"
	_ "time"
)

type Article struct {
	ArticleId int    `json:"articleId" gorm:"type:int(11);primary_key"` // 编码
	Title     string `json:"title" gorm:"type:varchar(128);"`           // 标题
	Author    string `json:"author" gorm:"type:varchar(128);"`          // 作者
	Content   string `json:"content" gorm:"type:varchar(255);"`         // 内容
	CreatedAt string `json:"createdAt" gorm:"type:timestamp;"`          //
	UpdatedAt string `json:"updatedAt" gorm:"type:timestamp;"`          //
	DeletedAt string `json:"deletedAt" gorm:"type:timestamp;"`          //
	DataScope string `json:"dataScope" gorm:"-"`
	Params    string `json:"params"  gorm:"-"`
	BaseModel
}

func (Article) TableName() string {
	return "article"
}

// 创建Article
func (e *Article) Create() (Article, error) {
	var doc Article
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取Article
func (e *Article) Get() (Article, error) {
	var doc Article

	table := orm.Eloquent.Table(e.TableName())

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取Article带分页
func (e *Article) GetPage(pageSize int, pageIndex int) ([]Article, int, error) {
	var doc []Article

	table := orm.Eloquent.Select("*").Table(e.TableName())

	// 数据权限控制(如果不需要数据权限请将此处去掉)
	dataPermission := new(DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table = dataPermission.GetDataScope(e.TableName(), table)

	var count int
	table = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize)
	if err := table.Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Count(&count)
	return doc, count, nil
}

// 更新Article
func (e *Article) Update(id int) (update Article, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("article_id = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 删除Article
func (e *Article) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("article_id = ?", id).Delete(&Article{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}
