/*
 * @Author: xiaoxu@mgtv.com
 * @Date: 2020-07-23 21:58:29
 * @Jira:
 * @Wiki:
 * @LastEditTime: 2020-07-23 22:00:14
 * @LastEditors: xiaoxu
 * @Description:
 * @FilePath: \go-admin-ui-vuef:\project\work\go\src\go-admin-demo\apis\grid\grid.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package grid

import (
	"go-admin-demo/models"
	"go-admin-demo/tools"
	"go-admin-demo/tools/app"
	"go-admin-demo/tools/app/msg"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetGridList(c *gin.Context) {
	var data models.Grid
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

func GetGrid(c *gin.Context) {
	var data models.Grid
	data.GridId, _ = tools.StringToInt(c.Param("gridId"))
	result, err := data.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	app.OK(c, result, "")
}

func InsertGrid(c *gin.Context) {
	var data models.Grid
	err := c.ShouldBindJSON(&data)
	data.CreateBy = tools.GetUserIdStr(c)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

func UpdateGrid(c *gin.Context) {
	var data models.Grid
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)
	data.UpdateBy = tools.GetUserIdStr(c)
	result, err := data.Update(data.GridId)
	tools.HasError(err, "", -1)

	app.OK(c, result, "")
}

func DeleteGrid(c *gin.Context) {
	var data models.Grid
	data.UpdateBy = tools.GetUserIdStr(c)

	IDS := tools.IdsStrToIdsIntGroup("gridId", c)
	_, err := data.BatchDelete(IDS)
	tools.HasError(err, msg.DeletedFail, 500)
	app.OK(c, nil, msg.DeletedSuccess)
}
