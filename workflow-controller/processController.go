package controller

import (
	"fmt"
	"strconv"

	"go-admin-demo/workflow-engine/model"

	"github.com/gin-gonic/gin"
	"github.com/mumushuiding/util"

	"go-admin-demo/workflow-engine/service"
)

// StartProcessInstanceByToken 启动流程
func StartProcessInstanceByToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if len(token) == 0 {
		c.Request.ParseForm()
		if len(c.Request.Form["token"]) == 0 {
			util.ResponseErr(c.Writer, "header Authorization 没有保存 token, url参数也不存在 token， 访问失败 ！")
			return
		}
		token = c.Request.Form["token"][0]
	}
	var proc = service.ProcessReceiver{}
	err := util.Body2Struct(c.Request, &proc)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	if len(proc.ProcName) == 0 {
		util.Response(c.Writer, "流程定义名procName不能为空", false)
		return
	}
	id, err := service.StartProcessInstanceByToken(token, &proc)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	util.Response(c.Writer, fmt.Sprintf("%d", id), true)
}

// StartProcessInstance 启动流程
func StartProcessInstance(c *gin.Context) {
	if model.RedisOpen {
		util.ResponseErr(c.Writer, "已经连接 redis，请使用/workflow/process/startByToken 路径访问")
		return
	}
	var proc = service.ProcessReceiver{}
	err := util.Body2Struct(c.Request, &proc)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	if len(proc.ProcName) == 0 {
		util.Response(c.Writer, "流程定义名procName不能为空", false)
		return
	}
	if len(proc.Company) == 0 {
		util.Response(c.Writer, "用户所在的公司company不能为空", false)
		return
	}
	if len(proc.UserID) == 0 {
		util.Response(c.Writer, "启动流程的用户userId不能为空", false)
		return
	}
	if len(proc.Username) == 0 {
		util.Response(c.Writer, "启动流程的用户username不能为空", false)
		return
	}
	if len(proc.Department) == 0 {
		util.Response(c.Writer, "用户所在部门department不能为空", false)
		return
	}
	id, err := proc.StartProcessInstanceByID(proc.Var)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	util.Response(c.Writer, fmt.Sprintf("%d", id), true)
}

// FindMyProcInstPageAsJSON FindMyProcInstPageAsJSON
// 查询到我审批的流程实例
func FindMyProcInstPageAsJSON(c *gin.Context) {
	if model.RedisOpen {
		util.ResponseErr(c.Writer, "已经连接 redis，请使用/workflow/process/findTaskByToken 路径访问")
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
	result, err := service.FindAllPageAsJSON(receiver)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	fmt.Fprintf(c.Writer, result)
}

// FindMyProcInstByToken FindMyProcInstByToken
// 查询待办的流程
func FindMyProcInstByToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if len(token) == 0 {
		c.Request.ParseForm()
		if len(c.Request.Form["token"]) == 0 {
			util.ResponseErr(c.Writer, "header Authorization 没有保存 token, url参数也不存在 token， 访问失败 ！")
			return
		}
		token = c.Request.Form["token"][0]
	}
	// fmt.Printf("token:%s\n", token)
	var receiver = service.GetDefaultProcessPageReceiver()
	err := util.Body2Struct(c.Request, &receiver)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	result, err := service.FindMyProcInstByToken(token, receiver)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	fmt.Fprintf(c.Writer, result)
}

// StartByMyself 我启动的流程
func StartByMyself(c *gin.Context) {
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
	result, err := service.StartByMyself(receiver)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	fmt.Fprintf(c.Writer, result)
}

// FindProcNotify 查询抄送我的流程
func FindProcNotify(c *gin.Context) {
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
	result, err := service.FindProcNotify(receiver)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	fmt.Fprintf(c.Writer, result)
}

// MoveFinishedProcInstToHistory MoveFinishedProcInstToHistory
// 将已经结束的流程实例移动到历史数据库
func MoveFinishedProcInstToHistory(c *gin.Context) {
	err := service.MoveFinishedProcInstToHistory()
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	util.ResponseOk(c.Writer)
}

// FindProcInstByID 根据流程ID查询流程
func FindProcInstByID(c *gin.Context) {
	c.Request.ParseForm()
	if len(c.Request.Form["id"]) == 0 {
		util.ResponseErr(c.Writer, "字段 id 不能为空")
		return
	}
	id, err := strconv.Atoi(c.Request.Form["id"][0])
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	data, err := service.FindProcInstByID(id)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	util.Response(c.Writer, data, true)
}
