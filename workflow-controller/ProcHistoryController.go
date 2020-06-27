package controller

import (
	"fmt"

	"go-admin-demo/workflow-engine/model"

	"go-admin-demo/workflow-engine/service"

	"github.com/gin-gonic/gin"
	"github.com/mumushuiding/util"
)

// FindProcHistoryByToken 查看我审批的纪录
func FindProcHistoryByToken(c *gin.Context) {
	token, err := GetToken(c)
	if err != nil {
		util.ResponseErr(c.Writer, "获取token失败")
		return
	}

	var receiver = service.GetDefaultProcessPageReceiver()
	err = util.Body2Struct(c.Request, &receiver)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	result, err := service.FindProcHistoryByToken(token, receiver)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	fmt.Fprintf(c.Writer, result)
}

// FindProcHistory 查询我的审批纪录
func FindProcHistory(c *gin.Context) {
	if model.RedisOpen {
		util.ResponseErr(c.Writer, "已经连接 redis，请使用/workflow/procHistory/findTaskByToken")
		return
	}
	var receiver = service.GetDefaultProcessPageReceiver()
	err := util.Body2Struct(c.Request, &receiver)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	if len(receiver.UserID) == 0 {
		util.Response(c.Writer, "用户userID不能为空", false)
		return
	}
	if len(receiver.Company) == 0 {
		util.Response(c.Writer, "字段 company 不能为空", false)
		return
	}
	result, err := service.FindProcHistory(receiver)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	fmt.Fprintf(c.Writer, result)
}

// StartHistoryByMyself 查询我发起的流程
func StartHistoryByMyself(c *gin.Context) {
	var receiver = service.GetDefaultProcessPageReceiver()
	err := util.Body2Struct(c.Request, &receiver)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	if len(receiver.UserID) == 0 {
		util.Response(c.Writer, "用户userID不能为空", false)
		return
	}
	if len(receiver.Company) == 0 {
		util.Response(c.Writer, "字段 company 不能为空", false)
		return
	}
	result, err := service.StartHistoryByMyself(receiver)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	fmt.Fprintf(c.Writer, result)
}

// FindProcHistoryNotify 查询抄送我的流程
func FindProcHistoryNotify(c *gin.Context) {
	var receiver = service.GetDefaultProcessPageReceiver()
	err := util.Body2Struct(c.Request, &receiver)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	if len(receiver.UserID) == 0 {
		util.Response(c.Writer, "用户userID不能为空", false)
		return
	}
	if len(receiver.Company) == 0 {
		util.Response(c.Writer, "字段 company 不能为空", false)
		return
	}
	result, err := service.FindProcHistoryNotify(receiver)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	fmt.Fprintf(c.Writer, result)
}
