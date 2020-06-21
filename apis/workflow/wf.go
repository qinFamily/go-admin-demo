package workflow

import (
	"go-admin-demo/models"
	"go-admin-demo/tools"
	"go-admin-demo/tools/app"
	"go-admin-demo/tools/app/msg"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*/workflow/workflowtype/?page=1&limit=20*/
func GetWorkFlow(c *gin.Context) {

	var data models.WorkflowsWorkflow
	var err error
	var pageSize = 20
	var pageIndex = 1

	if size := c.Request.FormValue("limit"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("page"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex, true)
	tools.HasError(err, "抱歉未找到相关信息", -1)

	res := app.WorkFlowResponse{
		Code:    200,
		Results: result,
		Count:   count,
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}

func UpdateWorkFlow(c *gin.Context) {
	var data models.WorkflowsWorkflow
	err := c.ShouldBindJSON(&data)
	tools.HasError(err, "数据WorkFlow解析错误", 500)

	IDS := tools.IdsStrToIdsIntGroup("flowId", c)
	if len(IDS) > 0 {
		_, err = data.Update(IDS[0])
		tools.HasError(err, "抱歉未找到相关信息", -1)
		app.OK(c, nil, msg.UpdatedSuccess)
		return
	}
	app.OK(c, nil, msg.NotFound)
}

func DeleteWorkflowsWorkflow(c *gin.Context) {
	var data models.WorkflowsWorkflow
	IDS := tools.IdsStrToIdsIntGroup("flowId", c)
	_, err := data.BatchDelete(IDS)
	tools.HasError(err, msg.DeletedFail, 500)
	app.OK(c, nil, msg.DeletedSuccess)
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
