package model

import "go-admin-demo/database"

// ProcdefHistory 历史流程定义
type ProcdefHistory struct {
	Procdef
}

// Save Save
func (p *ProcdefHistory) Save() (ID int, err error) {
	if db == nil {
		db = database.Eloquent
	}
	err = db.Create(p).Error
	if err != nil {
		return 0, err
	}
	return p.ID, nil
}
