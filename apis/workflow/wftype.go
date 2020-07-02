package workflow

import (
	"go-admin-demo/models"
	"go-admin-demo/tools"
	"go-admin-demo/tools/app"
	"go-admin-demo/tools/app/msg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type wftResponseResult struct {
	models.WorkflowsWorkflowtype
	WorkflowSet []struct {
		CreateTime          string `json:"create_time"`
		DisplayFormStr      string `json:"display_form_str"`
		ID                  int64  `json:"id"`
		LimitExpression     string `json:"limit_expression"`
		Memo                string `json:"memo"`
		Name                string `json:"name"`
		Status              bool   `json:"status"`
		TicketSnPrefix      string `json:"ticket_sn_prefix"`
		TitleTemplate       string `json:"title_template"`
		Type                int64  `json:"type"`
		UpdateTime          string `json:"update_time"`
		ViewPermissionCheck bool   `json:"view_permission_check"`
	} `json:"workflow_set"`
}

/*/workflow/workflowtype/?page=1&limit=20*/
func GetWorkFlowType(c *gin.Context) {

	var data models.WorkflowsWorkflowtype
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
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "抱歉未找到相关信息", -1)

	res := &app.WorkFlowResponse{
		Count:   count,
		Results: result,
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}

func UpdateWorkFlowType(c *gin.Context) {
	var data models.WorkflowsWorkflowtype
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

func InsertWorkFlowType(c *gin.Context) {
	var data models.WorkflowsWorkflowtype
	err := c.ShouldBindJSON(&data)
	tools.HasError(err, "", 500) // 数据WorkFlow解析错误
	result, err := data.Create()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

func DeleteWorkflowsWorkflowType(c *gin.Context) {
	var data models.WorkflowsWorkflowtype
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
