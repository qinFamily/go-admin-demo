package model

import (
	"go-admin-demo/database"
	orm "go-admin-demo/database"

	"github.com/jinzhu/gorm"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

var db = orm.Eloquent

// Model 其它数据结构的公共部分
type Model struct {
	ID int `gorm:"primary_key" json:"id,omitempty"`
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	if db == nil {
		db = database.Eloquent
	}
	defer db.Close()
}

// GetDB getdb
func GetDB() *gorm.DB {
	if db == nil {
		db = database.Eloquent
	}
	return db
}

// GetTx GetTx
func GetTx() *gorm.DB {
	if db == nil {
		db = database.Eloquent
	}
	return db.Begin()
}
