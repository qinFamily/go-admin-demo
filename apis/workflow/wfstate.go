package workflow

import (
	"go-admin-demo/models"
	"go-admin-demo/tools"
	"go-admin-demo/tools/app"
	"go-admin-demo/tools/app/msg"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*/workflow/state/?page=1&limit=20*/
func GetWorkFlowState(c *gin.Context) {

	var data models.WorkflowsState
	var err error
	var pageSize = 20
	var pageIndex = 1

	if size := c.Request.FormValue("limit"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("page"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	var workflowID = 0
	if id := c.Request.FormValue("workflow"); id != "" {
		workflowID = tools.StrToInt(err, id)
	}
	data.WorkflowID = workflowID
	// /api/workflow/state/?is_hidden=false&workflow=1
	isHidden := c.Request.FormValue("is_hidden")
	checkHiden := false
	if len(isHidden) > 0 {
		data.IsHidden = isHidden == "true"
		checkHiden = true
	}
	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex, false, 1, checkHiden)
	tools.HasError(err, "抱歉未找到相关信息", -1)

	res := &app.WorkFlowResponse{
		Count:   count,
		Results: result,
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}

func UpdateWorkFlowState(c *gin.Context) {
	var data models.WorkflowsState
	err := c.ShouldBindJSON(&data)
	tools.HasError(err, "数据WorkFlow解析错误", 500)

	IDS := tools.IdsStrToIdsIntGroup("flowtypeId", c)
	if len(IDS) > 0 {
		_, err = data.Update(IDS[0])
		tools.HasError(err, "抱歉未找到相关信息", -1)
		app.OK(c, nil, msg.UpdatedSuccess)
		return
	}
	app.OK(c, nil, msg.NotFound)
	// res := &app.WorkFlowResponse{
	// 	Count:   len(results),
	// 	Results: results,
	// }
	// c.JSON(http.StatusOK, res.ReturnOK())
}

func DeleteWorkflowsWorkFlowState(c *gin.Context) {
	var data models.WorkflowsState
	IDS := tools.IdsStrToIdsIntGroup("flowtypeId", c)
	_, err := data.BatchDelete(IDS)
	tools.HasError(err, msg.DeletedFail, 500)
	app.OK(c, nil, msg.DeletedSuccess)
}

/*
	{
	    "results":{
	        "id":2,
	        "workflow_set":[
	            {
	                "id":2,
	                "create_time":"2020-06-13 14:08:05",
	                "update_time":"2020-06-13 14:08:05",
	                "memo":"",
	                "name":"发布单",
	                "ticket_sn_prefix":"deploy",
	                "status":true,
	                "view_permission_check":true,
	                "limit_expression":"{}",
	                "display_form_str":"[]",
	                "title_template":"你有一个待办工单:{title}",
	                "type":2
	            }],
	        "create_time":"2020-06-13 14:08:05",
	        "update_time":"2020-06-14 19:06:14",
	        "memo":"mark",
	        "name":"技术",
	        "code":"it",
	        "order_id":2
	    },
	    "code":20000
	}

*/
