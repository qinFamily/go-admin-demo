package controller

import (
	"log"
	"strconv"

	"go-admin-demo/workflow-engine/model"

	"go-admin-demo/workflow-engine/service"

	"github.com/gin-gonic/gin"
	"github.com/mumushuiding/util"
)

// WithDrawTask 撤回
func WithDrawTask(c *gin.Context) {
	if model.RedisOpen {
		util.ResponseErr(c.Writer, "已经连接redis缓存，请使用方法 /workflow/task/withdrawByToken ")
		return
	}
	var taskRe = service.TaskReceiver{}
	err := util.Body2Struct(c.Request, &taskRe)
	str, _ := util.ToJSONStr(taskRe)
	log.Println(str)
	if taskRe.TaskID == 0 {
		util.ResponseErr(c.Writer, "字段taskID不能为空,必须为数字！")
		return
	}
	if len(taskRe.UserID) == 0 {
		util.ResponseErr(c.Writer, "字段userID不能为空！")
		return
	}
	if taskRe.ProcInstID == 0 {
		util.ResponseErr(c.Writer, "字段 procInstID 不能为空,必须为数字！")
		return
	}
	if len(taskRe.Company) == 0 {
		util.ResponseErr(c.Writer, "字段company不能为空！")
		return
	}
	err = service.WithDrawTask(taskRe.TaskID, taskRe.ProcInstID, taskRe.UserID, taskRe.UserName, taskRe.Company, taskRe.Comment)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	util.ResponseOk(c.Writer)
}

// WithDrawTaskByToken 撤回
func WithDrawTaskByToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if len(token) == 0 {
		c.Request.ParseForm()
		if len(c.Request.Form["token"]) == 0 {
			util.ResponseErr(c.Writer, "header Authorization 没有保存 token, url参数也不存在 token， 访问失败 ！")
			return
		}
		token = c.Request.Form["token"][0]
	}
	var taskRe = service.TaskReceiver{}
	err := util.Body2Struct(c.Request, &taskRe)
	if taskRe.TaskID == 0 {
		util.ResponseErr(c.Writer, "字段taskID不能为空,必须为数字！")
		return
	}
	if taskRe.ProcInstID == 0 {
		util.ResponseErr(c.Writer, "字段 procInstID 不能为空,必须为数字！")
		return
	}
	err = service.WithDrawTaskByToken(token, &taskRe)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	util.ResponseOk(c.Writer)
}

// CompleteTaskByToken 使用redis缓存时使用当前方法，更安全
func CompleteTaskByToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if len(token) == 0 {
		c.Request.ParseForm()
		if len(c.Request.Form["token"]) == 0 {
			util.ResponseErr(c.Writer, "header Authorization 没有保存 token, url参数也不存在 token， 访问失败 ！")
			return
		}
		token = c.Request.Form["token"][0]
	}
	var taskRe = service.TaskReceiver{}
	err := util.Body2Struct(c.Request, &taskRe)
	// str, _ := util.ToJSONStr(taskRe)
	// log.Println(str)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	if len(taskRe.Comment) > 255 {
		util.ResponseErr(c.Writer, "字段comment 长度不能超过255")
		return
	}
	if len(taskRe.Pass) == 0 {
		util.ResponseErr(c.Writer, "字段pass不能为空！")
		return
	}
	if taskRe.TaskID == 0 {
		util.ResponseErr(c.Writer, "字段taskID不能为空！")
		return
	}
	err = service.CompleteByToken(token, &taskRe)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	util.ResponseOk(c.Writer)
}

// CompleteTask CompleteTask
// 审批
func CompleteTask(c *gin.Context) {
	if model.RedisOpen {
		util.ResponseErr(c.Writer, "已经连接redis缓存，请使用方法 /workflow/task/completeByToken")
		return
	}
	var taskRe = service.TaskReceiver{}
	err := util.Body2Struct(c.Request, &taskRe)
	// str, _ := util.ToJSONStr(taskRe)
	// log.Println(str)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	if len(taskRe.Pass) == 0 {
		util.ResponseErr(c.Writer, "字段pass不能为空！")
		return
	}
	pass, err := strconv.ParseBool(taskRe.Pass)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	if taskRe.TaskID == 0 {
		util.ResponseErr(c.Writer, "字段taskID不能为空！")
		return
	}
	if len(taskRe.UserID) == 0 {
		util.ResponseErr(c.Writer, "字段userID不能为空！")
		return
	}
	if len(taskRe.UserName) == 0 {
		util.ResponseErr(c.Writer, "字段username不能为空！")
		return
	}
	if len(taskRe.Company) == 0 {
		util.ResponseErr(c.Writer, "字段company不能为空！")
		return
	}
	err = service.Complete(taskRe.TaskID, taskRe.UserID, taskRe.UserName, taskRe.Company, taskRe.Comment, taskRe.Candidate, pass)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	util.ResponseOk(c.Writer)
}
