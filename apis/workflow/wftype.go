package workflow

import (
	orm "go-admin-demo/database"
	"go-admin-demo/models"
	"go-admin-demo/tools"
	"go-admin-demo/tools/app"
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
	wfwftmgr := models.WorkflowsWorkflowtypeMgr(orm.Eloquent)
	results, err := wfwftmgr.GetsWithWorkflowSet()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	res := &app.WorkFlowResponse{
		Count:   len(results),
		Results: results,
	}
	c.JSON(http.StatusOK, res.ReturnOK())
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
