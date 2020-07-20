package base

import (
	"go-admin-demo/models"
	"go-admin-demo/tools"
	"go-admin-demo/tools/app"
	"go-admin-demo/tools/app/msg"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// GetGridBasicPeopleList 列表
func GetGridBasicPeopleList(c *gin.Context) {
	var data models.GridBasicPeople
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

// GetGridBasicPeople 详情
func GetGridBasicPeople(c *gin.Context) {
	var data models.GridBasicPeople
	data.PeopleId, _ = tools.StringToInt(c.Param("peopleId"))
	result, err := data.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	app.OK(c, result, "")
}

// InsertGridBasicPeople 详情
func InsertGridBasicPeople(c *gin.Context) {
	var data models.GridBasicPeople
	err := c.ShouldBindJSON(&data)
	data.CreateBy = tools.GetUserIdStr(c)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

func UpdateGridBasicPeople(c *gin.Context) {
	var data models.GridBasicPeople
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)
	data.UpdateBy = tools.GetUserIdStr(c)
	result, err := data.Update(data.PeopleId)
	tools.HasError(err, "", -1)

	app.OK(c, result, "")
}

func DeleteGridBasicPeople(c *gin.Context) {
	var data models.GridBasicPeople
	data.UpdateBy = tools.GetUserIdStr(c)

	IDS := tools.IdsStrToIdsIntGroup("peopleId", c)
	_, err := data.BatchDelete(IDS)
	tools.HasError(err, msg.DeletedFail, 500)
	app.OK(c, nil, msg.DeletedSuccess)
}
