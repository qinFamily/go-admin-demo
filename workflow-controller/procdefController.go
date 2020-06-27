package controller

import (
	"fmt"
	"strconv"

	"go-admin-demo/tools"
	"go-admin-demo/workflow-engine/service"

	"github.com/gin-gonic/gin"
	"github.com/mumushuiding/util"
)

// SaveProcdefByToken SaveProcdefByToken
func SaveProcdefByToken(c *gin.Context) {
	var procdef service.Procdef
	err := c.ShouldBindJSON(&procdef)
	tools.HasError(err, "", 500)

	token, err := GetToken(c)
	if err != nil {
		util.ResponseErr(c.Writer, "获取token失败")
		return
	}

	if len(procdef.Name) == 0 {
		util.ResponseErr(c.Writer, "流程名称 name 不能为空")
		return
	}
	if procdef.Resource == nil || len(procdef.Resource.Name) == 0 {
		util.ResponseErr(c.Writer, "字段 resource 不能为空")
		return
	}
	id, err := procdef.SaveProcdefByToken(token)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	util.Response(c.Writer, fmt.Sprintf("%d", id), true)
}

// SaveProcdef save new procdefnition
// 保存流程定义
func SaveProcdef(c *gin.Context) {
	var procdef service.Procdef
	err := c.ShouldBindJSON(&procdef)
	tools.HasError(err, "", 500)

	if len(procdef.Userid) == 0 {
		util.ResponseErr(c.Writer, "字段 userid 不能为空")
		return
	}
	if len(procdef.Company) == 0 {
		util.ResponseErr(c.Writer, "字段 company 不能为空")
		return
	}
	if len(procdef.Name) == 0 {
		util.ResponseErr(c.Writer, "流程名称 name 不能为空")
		return
	}
	if procdef.Resource == nil || len(procdef.Resource.Name) == 0 {
		util.ResponseErr(c.Writer, "字段 resource 不能为空")
		return
	}
	id, err := procdef.SaveProcdef()
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	util.Response(c.Writer, fmt.Sprintf("%d", id), true)
}

// FindAllProcdefPage find by page
// 分页查询
func FindAllProcdefPage(c *gin.Context) {
	var procdef = service.Procdef{PageIndex: 1, PageSize: 10}
	err := c.ShouldBindJSON(&procdef)
	tools.HasError(err, "", 500)

	var pageSize = 10
	var pageIndex = 1
	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}
	procdef.PageIndex = pageIndex
	procdef.PageSize = pageSize

	datas, err := procdef.FindAllPageAsJSON()
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	fmt.Fprintf(c.Writer, "%s", datas)
}

// DelProcdefByID del by id
// 根据 id 删除
func DelProcdefByID(c *gin.Context) {
	c.Request.ParseForm()
	var ids []string = c.Request.Form["id"]
	if len(ids) == 0 {
		util.ResponseErr(c.Writer, "request param 【id】 is not valid , id 不存在 ")
		return
	}
	id, err := strconv.Atoi(ids[0])
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	err = service.DelProcdefByID(id)
	if err != nil {
		util.ResponseErr(c.Writer, err)
		return
	}
	util.ResponseOk(c.Writer)
}
