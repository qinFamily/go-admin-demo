package service

import (
	"go-admin-demo/workflow-engine/model"

	"github.com/mumushuiding/util"
)

// FindParticipantHistoryByProcInstID 历史纪录查询
func FindParticipantHistoryByProcInstID(procInstID int) (string, error) {
	datas, err := model.FindParticipantHistoryByProcInstID(procInstID)
	if err != nil {
		return "", err
	}
	str, err := util.ToJSONStr(datas)
	if err != nil {
		return "", err
	}
	return str, nil
}
