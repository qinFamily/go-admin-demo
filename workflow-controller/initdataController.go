package controller

import (
	"fmt"
	"go-admin-demo/tools"
	"go-admin-demo/tools/app"

	"github.com/gin-gonic/gin"
)

type initData struct {
	Code string       `json:"code"`
	Msg  string       `json:"msg"`
	Data initDataData `json:"data"`
}

type initDataWorkFlowDef struct {
	Name                 string `json:"name"`
	PublicFlag           int    `json:"publicFlag"`
	SortNo               int    `json:"sortNo"`
	DuplicateRemovelFlag int    `json:"duplicateRemovelFlag"`
	OptionTip            string `json:"optionTip"`
	OptionNotNull        int    `json:"optionNotNull"`
	Status               int    `json:"status"`
}

type initDataChildNode1 struct {
	NodeName         string        `json:"nodeName"`
	Type             int           `json:"type"`
	CcSelfSelectFlag int           `json:"ccSelfSelectFlag"`
	ChildNode        interface{}   `json:"childNode"`
	NodeUserList     []interface{} `json:"nodeUserList"`
	Error            bool          `json:"error"`
}

type initDataNodeUserList struct {
	TargetID int    `json:"targetId"`
	Type     int    `json:"type"`
	Name     string `json:"name"`
}

type initDataChildNode2 struct {
	NodeName                string                 `json:"nodeName"`
	Type                    int                    `json:"type"`
	PriorityLevel           int                    `json:"priorityLevel"`
	Settype                 int                    `json:"settype"`
	SelectMode              int                    `json:"selectMode"`
	SelectRange             int                    `json:"selectRange"`
	ExamineRoleID           int                    `json:"examineRoleId"`
	DirectorLevel           int                    `json:"directorLevel"`
	ReplaceByUp             int                    `json:"replaceByUp"`
	ExamineMode             int                    `json:"examineMode"`
	NoHanderAction          int                    `json:"noHanderAction"`
	ExamineEndType          int                    `json:"examineEndType"`
	ExamineEndRoleID        int                    `json:"examineEndRoleId"`
	ExamineEndDirectorLevel int                    `json:"examineEndDirectorLevel"`
	CcSelfSelectFlag        int                    `json:"ccSelfSelectFlag"`
	ConditionList           []interface{}          `json:"conditionList"`
	NodeUserList            []initDataNodeUserList `json:"nodeUserList"`
	ChildNode               interface{}            `json:"childNode"`
	ConditionNodes          []interface{}          `json:"conditionNodes"`
	Error                   bool                   `json:"error"`
}

type initDataConditionNodes struct {
	NodeName                string             `json:"nodeName"`
	Type                    int                `json:"type"`
	PriorityLevel           int                `json:"priorityLevel"`
	Settype                 int                `json:"settype"`
	SelectMode              int                `json:"selectMode"`
	SelectRange             int                `json:"selectRange"`
	ExamineRoleID           int                `json:"examineRoleId"`
	DirectorLevel           int                `json:"directorLevel"`
	ReplaceByUp             int                `json:"replaceByUp"`
	ExamineMode             int                `json:"examineMode"`
	NoHanderAction          int                `json:"noHanderAction"`
	ExamineEndType          int                `json:"examineEndType"`
	ExamineEndRoleID        int                `json:"examineEndRoleId"`
	ExamineEndDirectorLevel int                `json:"examineEndDirectorLevel"`
	CcSelfSelectFlag        int                `json:"ccSelfSelectFlag"`
	ConditionList           []interface{}      `json:"conditionList"`
	NodeUserList            []interface{}      `json:"nodeUserList"`
	ChildNode               initDataChildNode2 `json:"childNode"`
	ConditionNodes          []interface{}      `json:"conditionNodes"`
	Error                   bool               `json:"error"`
}

type initDataChildNode0 struct {
	NodeName                string                   `json:"nodeName"`
	Type                    int                      `json:"type"`
	PriorityLevel           int                      `json:"priorityLevel"`
	Settype                 int                      `json:"settype"`
	SelectMode              int                      `json:"selectMode"`
	SelectRange             int                      `json:"selectRange"`
	ExamineRoleID           int                      `json:"examineRoleId"`
	DirectorLevel           int                      `json:"directorLevel"`
	ReplaceByUp             int                      `json:"replaceByUp"`
	ExamineMode             int                      `json:"examineMode"`
	NoHanderAction          int                      `json:"noHanderAction"`
	ExamineEndType          int                      `json:"examineEndType"`
	ExamineEndRoleID        int                      `json:"examineEndRoleId"`
	ExamineEndDirectorLevel int                      `json:"examineEndDirectorLevel"`
	CcSelfSelectFlag        int                      `json:"ccSelfSelectFlag"`
	ConditionList           []interface{}            `json:"conditionList"`
	NodeUserList            []interface{}            `json:"nodeUserList"`
	ChildNode               initDataChildNode1       `json:"childNode"`
	ConditionNodes          []initDataConditionNodes `json:"conditionNodes"`
}

type initDataChildNode struct {
	NodeName                string             `json:"nodeName"`
	Error                   bool               `json:"error"`
	Type                    int                `json:"type"`
	Settype                 int                `json:"settype"`
	SelectMode              int                `json:"selectMode"`
	SelectRange             int                `json:"selectRange"`
	DirectorLevel           int                `json:"directorLevel"`
	ReplaceByUp             int                `json:"replaceByUp"`
	ExamineMode             int                `json:"examineMode"`
	NoHanderAction          int                `json:"noHanderAction"`
	ExamineEndDirectorLevel int                `json:"examineEndDirectorLevel"`
	ChildNode               initDataChildNode0 `json:"childNode"`
	NodeUserList            []interface{}      `json:"nodeUserList"`
}

type initDataNodeConfig struct {
	PkID                    string            `json:"pkId"`
	NodeName                string            `json:"nodeName"`
	Type                    int               `json:"type"`
	PriorityLevel           string            `json:"priorityLevel"`
	Settype                 string            `json:"settype"`
	SelectMode              string            `json:"selectMode"`
	SelectRange             string            `json:"selectRange"`
	ExamineRoleID           string            `json:"examineRoleId"`
	DirectorLevel           string            `json:"directorLevel"`
	ReplaceByUp             string            `json:"replaceByUp"`
	ExamineMode             string            `json:"examineMode"`
	NoHanderAction          string            `json:"noHanderAction"`
	ExamineEndType          string            `json:"examineEndType"`
	ExamineEndRoleID        string            `json:"examineEndRoleId"`
	ExamineEndDirectorLevel string            `json:"examineEndDirectorLevel"`
	CcSelfSelectFlag        string            `json:"ccSelfSelectFlag"`
	ConditionList           []interface{}     `json:"conditionList"`
	NodeUserList            []interface{}     `json:"nodeUserList"`
	ChildNode               initDataChildNode `json:"childNode"`
	ConditionNodes          []interface{}     `json:"conditionNodes"`
}

type initDataData struct {
	TableID           int                 `json:"tableId"`
	WorkFlowVersionID string              `json:"workFlowVersionId"`
	WorkFlowDef       initDataWorkFlowDef `json:"workFlowDef"`
	DirectorMaxLevel  int                 `json:"directorMaxLevel"`
	FlowPermission    []interface{}       `json:"flowPermission"`
	NodeConfig        initDataNodeConfig  `json:"nodeConfig"`
}

// GetInitData 获取初始化数据
func GetInitData(c *gin.Context) {
	workFlowDefID := c.Request.FormValue("workFlowDefId")
	if len(workFlowDefID) == 0 {
		tools.HasError(fmt.Errorf("configId NOT EXISTED"), "", 500)
	}
	app.OK(c, initDataData{}, "🆗")
	// fmt.Fprintf(c.Writer, "%s", workFlowDefID)
}
