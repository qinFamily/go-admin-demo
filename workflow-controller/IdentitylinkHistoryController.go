package controller

import (
	"fmt"
	"strconv"

	"go-admin-demo/workflow-engine/service"

	"github.com/gin-gonic/gin"
	"github.com/mumushuiding/util"
)

// FindParticipantHistoryByProcInstID 历史纪录查询
func FindParticipantHistoryByProcInstID(c *gin.Context) {
	c.Request.ParseForm()
	if len(c.Request.Form["procInstID"]) == 0 {
		util.ResponseErr(c.Writer, "流程 procInstID 不能为空")
		return
	}
	procInstID, err := strconv.Atoi(c.Request.Form["procInstID"][0])
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	result, err := service.FindParticipantHistoryByProcInstID(procInstID)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	fmt.Fprintf(c.Writer, result)

}
