package workflow

import (
	orm "go-admin-demo/database"
	"go-admin-demo/models"
	"go-admin-demo/tools"
	"go-admin-demo/tools/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*/workflow/workflowtype/?page=1&limit=20*/
func GetWorkFlow(c *gin.Context) {
	// var data models.Menu
	// var err error
	// var pageSize = 10
	// var pageIndex = 1

	// if size := c.Request.FormValue("limit"); size != "" {
	// 	pageSize = tools.StrToInt(err, size)
	// }

	// if index := c.Request.FormValue("page"); index != "" {
	// 	pageIndex = tools.StrToInt(err, index)
	// }

	// data.ConfigKey = c.Request.FormValue("configKey")
	// data.ConfigName = c.Request.FormValue("configName")
	// data.ConfigType = c.Request.FormValue("configType")
	// data.DataScope = tools.GetUserIdStr(c)
	// result, count, err := data.GetPage(pageSize, pageIndex)
	// tools.HasError(err, "", -1)

	// var mp = make(map[string]interface{}, 3)
	// mp["list"] = result
	// mp["count"] = count
	// mp["pageIndex"] = pageIndex
	// mp["pageSize"] = pageSize

	wfwfmgr := models.WorkflowsWorkflowMgr(orm.Eloquent)
	results, err := wfwfmgr.Gets()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	res := app.WorkFlowResponse{
		Code:     20000,
		Next:     nil,
		Previous: nil,
		Results:  results,
		Count:    len(results),
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}

/*
{
    "code":20000,
    "count":3,
    "next":null,
    "previous":null,
    "results":[
        {
            "id":1,
            "create_time":"2020-06-13T06:08:01+08:00",
            "update_time":"2020-06-13T06:08:01+08:00",
            "memo":"",
            "name":"请假单",
            "ticket_sn_prefix":"leave",
            "status":true,
            "view_permission_check":true,
            "limit_expression":"{}",
            "display_form_str":"[]",
            "title_template":"你有一个待办工单:{title}",
            "type":{
                "id":1,
                "create_time":"2020-06-13T06:08:00+08:00",
                "update_time":"2020-06-14T10:54:40+08:00",
                "memo":"mark",
                "name":"行政",
                "code":"ad",
                "order_id":1
            }
        },
        {
            "id":2,
            "create_time":"2020-06-13T06:08:06+08:00",
            "update_time":"2020-06-13T06:08:06+08:00",
            "memo":"",
            "name":"发布单",
            "ticket_sn_prefix":"deploy",
            "status":true,
            "view_permission_check":true,
            "limit_expression":"{}",
            "display_form_str":"[]",
            "title_template":"你有一个待办工单:{title}",
            "type":{
                "id":2,
                "create_time":"2020-06-13T06:08:06+08:00",
                "update_time":"2020-06-14T11:06:15+08:00",
                "memo":"mark",
                "name":"技术",
                "code":"it",
                "order_id":2
            }
        },
        {
            "id":3,
            "create_time":"2020-06-13T15:09:09+08:00",
            "update_time":"2020-06-13T15:09:09+08:00",
            "memo":"盖章签呈",
            "name":"盖章签呈",
            "ticket_sn_prefix":"gzqc",
            "status":true,
            "view_permission_check":true,
            "limit_expression":"",
            "display_form_str":"",
            "title_template":"",
            "type":{
                "id":3,
                "create_time":"2020-06-13T14:53:07+08:00",
                "update_time":"2020-06-13T14:53:07+08:00",
                "memo":"公共事物中心",
                "name":"盖章签呈",
                "code":"gzqc",
                "order_id":3
            }
        }]
}

*/
