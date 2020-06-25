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
func GetWorkFlowTransition(c *gin.Context) {

	var data models.WorkflowsTransition
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

	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex, true, 1)
	tools.HasError(err, "抱歉未找到相关信息", -1)

	res := &app.WorkFlowResponse{
		Count:   count,
		Results: result,
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}

func UpdateWorkFlowTransition(c *gin.Context) {
	var data models.WorkflowsTransition
	err := c.ShouldBindJSON(&data)
	tools.HasError(err, "数据WorkFlow transition解析错误", 500)

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

func DeleteWorkflowsWorkFlowTransition(c *gin.Context) {
	var data models.WorkflowsTransition
	IDS := tools.IdsStrToIdsIntGroup("flowtypeId", c)
	_, err := data.BatchDelete(IDS)
	tools.HasError(err, msg.DeletedFail, 500)
	app.OK(c, nil, msg.DeletedSuccess)
}

/*
{
    "results":[
        {
            "id":1,
            "create_time":"2020-06-13 14:08:04",
            "update_time":"2020-06-13 14:08:04",
            "memo":"",
            "name":0,
            "transition_type":0,
            "timer":0,
            "condition_expression":"[]",
            "attribute_type":0,
            "alert_enable":false,
            "alert_text":"",
            "workflow":{
                "id":1,
                "create_time":"2020-06-13 14:08:00",
                "update_time":"2020-06-13 14:08:00",
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
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-14 18:54:40",
                    "memo":"mark",
                    "name":"行政",
                    "code":"ad",
                    "order_id":1
                }
            },
            "source_state":{
                "id":1,
                "create_time":"2020-06-13 14:08:02",
                "update_time":"2020-06-13 14:08:02",
                "memo":"",
                "name":"开始",
                "is_hidden":true,
                "order_id":1,
                "state_type":1,
                "enable_retreat":false,
                "participant_type":0,
                "workflow":{
                    "id":1,
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-13 14:08:00",
                    "memo":"",
                    "name":"请假单",
                    "ticket_sn_prefix":"leave",
                    "status":true,
                    "view_permission_check":true,
                    "limit_expression":"{}",
                    "display_form_str":"[]",
                    "title_template":"你有一个待办工单:{title}",
                    "type":1
                },
                "user_participant":[
                ],
                "group_participant":[
                ],
                "role_participant":[
                ],
                "fields":[
                ]
            },
            "dest_state":{
                "id":3,
                "create_time":"2020-06-13 14:08:02",
                "update_time":"2020-06-13 14:08:02",
                "memo":"",
                "name":"申请人-编辑中",
                "is_hidden":false,
                "order_id":2,
                "state_type":0,
                "enable_retreat":false,
                "participant_type":0,
                "workflow":{
                    "id":1,
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-13 14:08:00",
                    "memo":"",
                    "name":"请假单",
                    "ticket_sn_prefix":"leave",
                    "status":true,
                    "view_permission_check":true,
                    "limit_expression":"{}",
                    "display_form_str":"[]",
                    "title_template":"你有一个待办工单:{title}",
                    "type":1
                },
                "user_participant":[
                ],
                "group_participant":[
                ],
                "role_participant":[
                ],
                "fields":[
                    {
                        "id":5,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":7,
                        "field_key":"start_end_time",
                        "field_name":"请假时间",
                        "order_id":10,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{}",
                        "label":"{}",
                        "workflow":1
                    },
                    {
                        "id":6,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":9,
                        "field_key":"type",
                        "field_name":"请假类型",
                        "order_id":30,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{\"1\":\"病假\", \"2\":\"产假\"}",
                        "label":"{}",
                        "workflow":1
                    },
                    {
                        "id":7,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":8,
                        "field_key":"memo",
                        "field_name":"事由说明",
                        "order_id":50,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{}",
                        "label":"{}",
                        "workflow":1
                    }]
            }
        },
        {
            "id":2,
            "create_time":"2020-06-13 14:08:04",
            "update_time":"2020-06-13 14:08:04",
            "memo":"",
            "name":1,
            "transition_type":0,
            "timer":0,
            "condition_expression":"[]",
            "attribute_type":1,
            "alert_enable":false,
            "alert_text":"",
            "workflow":{
                "id":1,
                "create_time":"2020-06-13 14:08:00",
                "update_time":"2020-06-13 14:08:00",
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
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-14 18:54:40",
                    "memo":"mark",
                    "name":"行政",
                    "code":"ad",
                    "order_id":1
                }
            },
            "source_state":{
                "id":1,
                "create_time":"2020-06-13 14:08:02",
                "update_time":"2020-06-13 14:08:02",
                "memo":"",
                "name":"开始",
                "is_hidden":true,
                "order_id":1,
                "state_type":1,
                "enable_retreat":false,
                "participant_type":0,
                "workflow":{
                    "id":1,
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-13 14:08:00",
                    "memo":"",
                    "name":"请假单",
                    "ticket_sn_prefix":"leave",
                    "status":true,
                    "view_permission_check":true,
                    "limit_expression":"{}",
                    "display_form_str":"[]",
                    "title_template":"你有一个待办工单:{title}",
                    "type":1
                },
                "user_participant":[
                ],
                "group_participant":[
                ],
                "role_participant":[
                ],
                "fields":[
                ]
            },
            "dest_state":{
                "id":4,
                "create_time":"2020-06-13 14:08:02",
                "update_time":"2020-06-13 14:08:02",
                "memo":"",
                "name":"领导-审批中",
                "is_hidden":false,
                "order_id":3,
                "state_type":0,
                "enable_retreat":false,
                "participant_type":3,
                "workflow":{
                    "id":1,
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-13 14:08:00",
                    "memo":"",
                    "name":"请假单",
                    "ticket_sn_prefix":"leave",
                    "status":true,
                    "view_permission_check":true,
                    "limit_expression":"{}",
                    "display_form_str":"[]",
                    "title_template":"你有一个待办工单:{title}",
                    "type":1
                },
                "user_participant":[
                ],
                "group_participant":[
                ],
                "role_participant":[
                    {
                        "id":2,
                        "create_time":"2020-06-13 14:07:56",
                        "update_time":"2020-06-13 14:07:56",
                        "memo":"",
                        "name":"运维经理",
                        "code":"ops_tl",
                        "sequence":1,
                        "parent":1,
                        "menus":[
                        ],
                        "model_perms":[
                        ]
                    },
                    {
                        "id":4,
                        "create_time":"2020-06-13 14:07:56",
                        "update_time":"2020-06-13 14:07:56",
                        "memo":"",
                        "name":"开发经理",
                        "code":"dev_tl",
                        "sequence":2,
                        "parent":1,
                        "menus":[
                        ],
                        "model_perms":[
                        ]
                    },
                    {
                        "id":6,
                        "create_time":"2020-06-13 14:07:56",
                        "update_time":"2020-06-13 14:07:56",
                        "memo":"",
                        "name":"人事经理",
                        "code":"hr_tl",
                        "sequence":3,
                        "parent":1,
                        "menus":[
                        ],
                        "model_perms":[
                        ]
                    }],
                "fields":[
                    {
                        "id":8,
                        "create_time":"2020-06-13 14:08:02",
                        "update_time":"2020-06-13 14:08:02",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":9,
                        "field_key":"leader_radio",
                        "field_name":"领导审批",
                        "order_id":60,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{\"1\":\"同意\", \"2\":\"不同意\"}",
                        "label":"{}",
                        "workflow":1
                    }]
            }
        },
        {
            "id":3,
            "create_time":"2020-06-13 14:08:04",
            "update_time":"2020-06-13 14:08:04",
            "memo":"",
            "name":0,
            "transition_type":0,
            "timer":0,
            "condition_expression":"[]",
            "attribute_type":0,
            "alert_enable":false,
            "alert_text":"",
            "workflow":{
                "id":1,
                "create_time":"2020-06-13 14:08:00",
                "update_time":"2020-06-13 14:08:00",
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
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-14 18:54:40",
                    "memo":"mark",
                    "name":"行政",
                    "code":"ad",
                    "order_id":1
                }
            },
            "source_state":{
                "id":3,
                "create_time":"2020-06-13 14:08:02",
                "update_time":"2020-06-13 14:08:02",
                "memo":"",
                "name":"申请人-编辑中",
                "is_hidden":false,
                "order_id":2,
                "state_type":0,
                "enable_retreat":false,
                "participant_type":0,
                "workflow":{
                    "id":1,
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-13 14:08:00",
                    "memo":"",
                    "name":"请假单",
                    "ticket_sn_prefix":"leave",
                    "status":true,
                    "view_permission_check":true,
                    "limit_expression":"{}",
                    "display_form_str":"[]",
                    "title_template":"你有一个待办工单:{title}",
                    "type":1
                },
                "user_participant":[
                ],
                "group_participant":[
                ],
                "role_participant":[
                ],
                "fields":[
                    {
                        "id":5,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":7,
                        "field_key":"start_end_time",
                        "field_name":"请假时间",
                        "order_id":10,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{}",
                        "label":"{}",
                        "workflow":1
                    },
                    {
                        "id":6,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":9,
                        "field_key":"type",
                        "field_name":"请假类型",
                        "order_id":30,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{\"1\":\"病假\", \"2\":\"产假\"}",
                        "label":"{}",
                        "workflow":1
                    },
                    {
                        "id":7,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":8,
                        "field_key":"memo",
                        "field_name":"事由说明",
                        "order_id":50,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{}",
                        "label":"{}",
                        "workflow":1
                    }]
            },
            "dest_state":{
                "id":3,
                "create_time":"2020-06-13 14:08:02",
                "update_time":"2020-06-13 14:08:02",
                "memo":"",
                "name":"申请人-编辑中",
                "is_hidden":false,
                "order_id":2,
                "state_type":0,
                "enable_retreat":false,
                "participant_type":0,
                "workflow":{
                    "id":1,
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-13 14:08:00",
                    "memo":"",
                    "name":"请假单",
                    "ticket_sn_prefix":"leave",
                    "status":true,
                    "view_permission_check":true,
                    "limit_expression":"{}",
                    "display_form_str":"[]",
                    "title_template":"你有一个待办工单:{title}",
                    "type":1
                },
                "user_participant":[
                ],
                "group_participant":[
                ],
                "role_participant":[
                ],
                "fields":[
                    {
                        "id":5,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":7,
                        "field_key":"start_end_time",
                        "field_name":"请假时间",
                        "order_id":10,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{}",
                        "label":"{}",
                        "workflow":1
                    },
                    {
                        "id":6,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":9,
                        "field_key":"type",
                        "field_name":"请假类型",
                        "order_id":30,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{\"1\":\"病假\", \"2\":\"产假\"}",
                        "label":"{}",
                        "workflow":1
                    },
                    {
                        "id":7,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":8,
                        "field_key":"memo",
                        "field_name":"事由说明",
                        "order_id":50,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{}",
                        "label":"{}",
                        "workflow":1
                    }]
            }
        },
        {
            "id":4,
            "create_time":"2020-06-13 14:08:04",
            "update_time":"2020-06-13 14:08:04",
            "memo":"",
            "name":1,
            "transition_type":0,
            "timer":0,
            "condition_expression":"[]",
            "attribute_type":1,
            "alert_enable":false,
            "alert_text":"",
            "workflow":{
                "id":1,
                "create_time":"2020-06-13 14:08:00",
                "update_time":"2020-06-13 14:08:00",
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
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-14 18:54:40",
                    "memo":"mark",
                    "name":"行政",
                    "code":"ad",
                    "order_id":1
                }
            },
            "source_state":{
                "id":3,
                "create_time":"2020-06-13 14:08:02",
                "update_time":"2020-06-13 14:08:02",
                "memo":"",
                "name":"申请人-编辑中",
                "is_hidden":false,
                "order_id":2,
                "state_type":0,
                "enable_retreat":false,
                "participant_type":0,
                "workflow":{
                    "id":1,
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-13 14:08:00",
                    "memo":"",
                    "name":"请假单",
                    "ticket_sn_prefix":"leave",
                    "status":true,
                    "view_permission_check":true,
                    "limit_expression":"{}",
                    "display_form_str":"[]",
                    "title_template":"你有一个待办工单:{title}",
                    "type":1
                },
                "user_participant":[
                ],
                "group_participant":[
                ],
                "role_participant":[
                ],
                "fields":[
                    {
                        "id":5,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":7,
                        "field_key":"start_end_time",
                        "field_name":"请假时间",
                        "order_id":10,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{}",
                        "label":"{}",
                        "workflow":1
                    },
                    {
                        "id":6,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":9,
                        "field_key":"type",
                        "field_name":"请假类型",
                        "order_id":30,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{\"1\":\"病假\", \"2\":\"产假\"}",
                        "label":"{}",
                        "workflow":1
                    },
                    {
                        "id":7,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":8,
                        "field_key":"memo",
                        "field_name":"事由说明",
                        "order_id":50,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{}",
                        "label":"{}",
                        "workflow":1
                    }]
            },
            "dest_state":{
                "id":4,
                "create_time":"2020-06-13 14:08:02",
                "update_time":"2020-06-13 14:08:02",
                "memo":"",
                "name":"领导-审批中",
                "is_hidden":false,
                "order_id":3,
                "state_type":0,
                "enable_retreat":false,
                "participant_type":3,
                "workflow":{
                    "id":1,
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-13 14:08:00",
                    "memo":"",
                    "name":"请假单",
                    "ticket_sn_prefix":"leave",
                    "status":true,
                    "view_permission_check":true,
                    "limit_expression":"{}",
                    "display_form_str":"[]",
                    "title_template":"你有一个待办工单:{title}",
                    "type":1
                },
                "user_participant":[
                ],
                "group_participant":[
                ],
                "role_participant":[
                    {
                        "id":2,
                        "create_time":"2020-06-13 14:07:56",
                        "update_time":"2020-06-13 14:07:56",
                        "memo":"",
                        "name":"运维经理",
                        "code":"ops_tl",
                        "sequence":1,
                        "parent":1,
                        "menus":[
                        ],
                        "model_perms":[
                        ]
                    },
                    {
                        "id":4,
                        "create_time":"2020-06-13 14:07:56",
                        "update_time":"2020-06-13 14:07:56",
                        "memo":"",
                        "name":"开发经理",
                        "code":"dev_tl",
                        "sequence":2,
                        "parent":1,
                        "menus":[
                        ],
                        "model_perms":[
                        ]
                    },
                    {
                        "id":6,
                        "create_time":"2020-06-13 14:07:56",
                        "update_time":"2020-06-13 14:07:56",
                        "memo":"",
                        "name":"人事经理",
                        "code":"hr_tl",
                        "sequence":3,
                        "parent":1,
                        "menus":[
                        ],
                        "model_perms":[
                        ]
                    }],
                "fields":[
                    {
                        "id":8,
                        "create_time":"2020-06-13 14:08:02",
                        "update_time":"2020-06-13 14:08:02",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":9,
                        "field_key":"leader_radio",
                        "field_name":"领导审批",
                        "order_id":60,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{\"1\":\"同意\", \"2\":\"不同意\"}",
                        "label":"{}",
                        "workflow":1
                    }]
            }
        },
        {
            "id":5,
            "create_time":"2020-06-13 14:08:04",
            "update_time":"2020-06-13 14:08:04",
            "memo":"",
            "name":3,
            "transition_type":0,
            "timer":0,
            "condition_expression":"[]",
            "attribute_type":3,
            "alert_enable":false,
            "alert_text":"",
            "workflow":{
                "id":1,
                "create_time":"2020-06-13 14:08:00",
                "update_time":"2020-06-13 14:08:00",
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
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-14 18:54:40",
                    "memo":"mark",
                    "name":"行政",
                    "code":"ad",
                    "order_id":1
                }
            },
            "source_state":{
                "id":3,
                "create_time":"2020-06-13 14:08:02",
                "update_time":"2020-06-13 14:08:02",
                "memo":"",
                "name":"申请人-编辑中",
                "is_hidden":false,
                "order_id":2,
                "state_type":0,
                "enable_retreat":false,
                "participant_type":0,
                "workflow":{
                    "id":1,
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-13 14:08:00",
                    "memo":"",
                    "name":"请假单",
                    "ticket_sn_prefix":"leave",
                    "status":true,
                    "view_permission_check":true,
                    "limit_expression":"{}",
                    "display_form_str":"[]",
                    "title_template":"你有一个待办工单:{title}",
                    "type":1
                },
                "user_participant":[
                ],
                "group_participant":[
                ],
                "role_participant":[
                ],
                "fields":[
                    {
                        "id":5,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":7,
                        "field_key":"start_end_time",
                        "field_name":"请假时间",
                        "order_id":10,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{}",
                        "label":"{}",
                        "workflow":1
                    },
                    {
                        "id":6,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":9,
                        "field_key":"type",
                        "field_name":"请假类型",
                        "order_id":30,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{\"1\":\"病假\", \"2\":\"产假\"}",
                        "label":"{}",
                        "workflow":1
                    },
                    {
                        "id":7,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":8,
                        "field_key":"memo",
                        "field_name":"事由说明",
                        "order_id":50,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{}",
                        "label":"{}",
                        "workflow":1
                    }]
            },
            "dest_state":{
                "id":6,
                "create_time":"2020-06-13 14:08:03",
                "update_time":"2020-06-13 14:08:03",
                "memo":"",
                "name":"结束",
                "is_hidden":false,
                "order_id":98,
                "state_type":2,
                "enable_retreat":false,
                "participant_type":0,
                "workflow":{
                    "id":1,
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-13 14:08:00",
                    "memo":"",
                    "name":"请假单",
                    "ticket_sn_prefix":"leave",
                    "status":true,
                    "view_permission_check":true,
                    "limit_expression":"{}",
                    "display_form_str":"[]",
                    "title_template":"你有一个待办工单:{title}",
                    "type":1
                },
                "user_participant":[
                ],
                "group_participant":[
                ],
                "role_participant":[
                ],
                "fields":[
                ]
            }
        },
        {
            "id":6,
            "create_time":"2020-06-13 14:08:04",
            "update_time":"2020-06-13 14:08:04",
            "memo":"",
            "name":2,
            "transition_type":0,
            "timer":0,
            "condition_expression":"[]",
            "attribute_type":2,
            "alert_enable":false,
            "alert_text":"",
            "workflow":{
                "id":1,
                "create_time":"2020-06-13 14:08:00",
                "update_time":"2020-06-13 14:08:00",
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
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-14 18:54:40",
                    "memo":"mark",
                    "name":"行政",
                    "code":"ad",
                    "order_id":1
                }
            },
            "source_state":{
                "id":4,
                "create_time":"2020-06-13 14:08:02",
                "update_time":"2020-06-13 14:08:02",
                "memo":"",
                "name":"领导-审批中",
                "is_hidden":false,
                "order_id":3,
                "state_type":0,
                "enable_retreat":false,
                "participant_type":3,
                "workflow":{
                    "id":1,
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-13 14:08:00",
                    "memo":"",
                    "name":"请假单",
                    "ticket_sn_prefix":"leave",
                    "status":true,
                    "view_permission_check":true,
                    "limit_expression":"{}",
                    "display_form_str":"[]",
                    "title_template":"你有一个待办工单:{title}",
                    "type":1
                },
                "user_participant":[
                ],
                "group_participant":[
                ],
                "role_participant":[
                    {
                        "id":2,
                        "create_time":"2020-06-13 14:07:56",
                        "update_time":"2020-06-13 14:07:56",
                        "memo":"",
                        "name":"运维经理",
                        "code":"ops_tl",
                        "sequence":1,
                        "parent":1,
                        "menus":[
                        ],
                        "model_perms":[
                        ]
                    },
                    {
                        "id":4,
                        "create_time":"2020-06-13 14:07:56",
                        "update_time":"2020-06-13 14:07:56",
                        "memo":"",
                        "name":"开发经理",
                        "code":"dev_tl",
                        "sequence":2,
                        "parent":1,
                        "menus":[
                        ],
                        "model_perms":[
                        ]
                    },
                    {
                        "id":6,
                        "create_time":"2020-06-13 14:07:56",
                        "update_time":"2020-06-13 14:07:56",
                        "memo":"",
                        "name":"人事经理",
                        "code":"hr_tl",
                        "sequence":3,
                        "parent":1,
                        "menus":[
                        ],
                        "model_perms":[
                        ]
                    }],
                "fields":[
                    {
                        "id":8,
                        "create_time":"2020-06-13 14:08:02",
                        "update_time":"2020-06-13 14:08:02",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":9,
                        "field_key":"leader_radio",
                        "field_name":"领导审批",
                        "order_id":60,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{\"1\":\"同意\", \"2\":\"不同意\"}",
                        "label":"{}",
                        "workflow":1
                    }]
            },
            "dest_state":{
                "id":3,
                "create_time":"2020-06-13 14:08:02",
                "update_time":"2020-06-13 14:08:02",
                "memo":"",
                "name":"申请人-编辑中",
                "is_hidden":false,
                "order_id":2,
                "state_type":0,
                "enable_retreat":false,
                "participant_type":0,
                "workflow":{
                    "id":1,
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-13 14:08:00",
                    "memo":"",
                    "name":"请假单",
                    "ticket_sn_prefix":"leave",
                    "status":true,
                    "view_permission_check":true,
                    "limit_expression":"{}",
                    "display_form_str":"[]",
                    "title_template":"你有一个待办工单:{title}",
                    "type":1
                },
                "user_participant":[
                ],
                "group_participant":[
                ],
                "role_participant":[
                ],
                "fields":[
                    {
                        "id":5,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":7,
                        "field_key":"start_end_time",
                        "field_name":"请假时间",
                        "order_id":10,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{}",
                        "label":"{}",
                        "workflow":1
                    },
                    {
                        "id":6,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":9,
                        "field_key":"type",
                        "field_name":"请假类型",
                        "order_id":30,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{\"1\":\"病假\", \"2\":\"产假\"}",
                        "label":"{}",
                        "workflow":1
                    },
                    {
                        "id":7,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":8,
                        "field_key":"memo",
                        "field_name":"事由说明",
                        "order_id":50,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{}",
                        "label":"{}",
                        "workflow":1
                    }]
            }
        },
        {
            "id":7,
            "create_time":"2020-06-13 14:08:05",
            "update_time":"2020-06-13 14:08:05",
            "memo":"",
            "name":1,
            "transition_type":0,
            "timer":0,
            "condition_expression":"[]",
            "attribute_type":1,
            "alert_enable":false,
            "alert_text":"",
            "workflow":{
                "id":1,
                "create_time":"2020-06-13 14:08:00",
                "update_time":"2020-06-13 14:08:00",
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
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-14 18:54:40",
                    "memo":"mark",
                    "name":"行政",
                    "code":"ad",
                    "order_id":1
                }
            },
            "source_state":{
                "id":4,
                "create_time":"2020-06-13 14:08:02",
                "update_time":"2020-06-13 14:08:02",
                "memo":"",
                "name":"领导-审批中",
                "is_hidden":false,
                "order_id":3,
                "state_type":0,
                "enable_retreat":false,
                "participant_type":3,
                "workflow":{
                    "id":1,
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-13 14:08:00",
                    "memo":"",
                    "name":"请假单",
                    "ticket_sn_prefix":"leave",
                    "status":true,
                    "view_permission_check":true,
                    "limit_expression":"{}",
                    "display_form_str":"[]",
                    "title_template":"你有一个待办工单:{title}",
                    "type":1
                },
                "user_participant":[
                ],
                "group_participant":[
                ],
                "role_participant":[
                    {
                        "id":2,
                        "create_time":"2020-06-13 14:07:56",
                        "update_time":"2020-06-13 14:07:56",
                        "memo":"",
                        "name":"运维经理",
                        "code":"ops_tl",
                        "sequence":1,
                        "parent":1,
                        "menus":[
                        ],
                        "model_perms":[
                        ]
                    },
                    {
                        "id":4,
                        "create_time":"2020-06-13 14:07:56",
                        "update_time":"2020-06-13 14:07:56",
                        "memo":"",
                        "name":"开发经理",
                        "code":"dev_tl",
                        "sequence":2,
                        "parent":1,
                        "menus":[
                        ],
                        "model_perms":[
                        ]
                    },
                    {
                        "id":6,
                        "create_time":"2020-06-13 14:07:56",
                        "update_time":"2020-06-13 14:07:56",
                        "memo":"",
                        "name":"人事经理",
                        "code":"hr_tl",
                        "sequence":3,
                        "parent":1,
                        "menus":[
                        ],
                        "model_perms":[
                        ]
                    }],
                "fields":[
                    {
                        "id":8,
                        "create_time":"2020-06-13 14:08:02",
                        "update_time":"2020-06-13 14:08:02",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":9,
                        "field_key":"leader_radio",
                        "field_name":"领导审批",
                        "order_id":60,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{\"1\":\"同意\", \"2\":\"不同意\"}",
                        "label":"{}",
                        "workflow":1
                    }]
            },
            "dest_state":{
                "id":5,
                "create_time":"2020-06-13 14:08:03",
                "update_time":"2020-06-13 14:08:03",
                "memo":"",
                "name":"人事-审批中",
                "is_hidden":false,
                "order_id":4,
                "state_type":0,
                "enable_retreat":false,
                "participant_type":2,
                "workflow":{
                    "id":1,
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-13 14:08:00",
                    "memo":"",
                    "name":"请假单",
                    "ticket_sn_prefix":"leave",
                    "status":true,
                    "view_permission_check":true,
                    "limit_expression":"{}",
                    "display_form_str":"[]",
                    "title_template":"你有一个待办工单:{title}",
                    "type":1
                },
                "user_participant":[
                ],
                "group_participant":[
                    {
                        "id":4,
                        "name":"人事",
                        "create_time":"2020-06-13 14:07:55",
                        "update_time":"2020-06-13 14:07:55",
                        "memo":"",
                        "code":"hr",
                        "sequence":3,
                        "parent":1,
                        "permissions":[
                        ],
                        "roles":[
                        ]
                    }],
                "role_participant":[
                ],
                "fields":[
                    {
                        "id":9,
                        "create_time":"2020-06-13 14:08:02",
                        "update_time":"2020-06-13 14:08:02",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":9,
                        "field_key":"hr_radio",
                        "field_name":"人事审批",
                        "order_id":80,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{\"1\":\"同意\", \"2\":\"不同意\"}",
                        "label":"{}",
                        "workflow":1
                    }]
            }
        },
        {
            "id":8,
            "create_time":"2020-06-13 14:08:05",
            "update_time":"2020-06-13 14:08:05",
            "memo":"",
            "name":2,
            "transition_type":0,
            "timer":0,
            "condition_expression":"[]",
            "attribute_type":2,
            "alert_enable":false,
            "alert_text":"",
            "workflow":{
                "id":1,
                "create_time":"2020-06-13 14:08:00",
                "update_time":"2020-06-13 14:08:00",
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
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-14 18:54:40",
                    "memo":"mark",
                    "name":"行政",
                    "code":"ad",
                    "order_id":1
                }
            },
            "source_state":{
                "id":5,
                "create_time":"2020-06-13 14:08:03",
                "update_time":"2020-06-13 14:08:03",
                "memo":"",
                "name":"人事-审批中",
                "is_hidden":false,
                "order_id":4,
                "state_type":0,
                "enable_retreat":false,
                "participant_type":2,
                "workflow":{
                    "id":1,
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-13 14:08:00",
                    "memo":"",
                    "name":"请假单",
                    "ticket_sn_prefix":"leave",
                    "status":true,
                    "view_permission_check":true,
                    "limit_expression":"{}",
                    "display_form_str":"[]",
                    "title_template":"你有一个待办工单:{title}",
                    "type":1
                },
                "user_participant":[
                ],
                "group_participant":[
                    {
                        "id":4,
                        "name":"人事",
                        "create_time":"2020-06-13 14:07:55",
                        "update_time":"2020-06-13 14:07:55",
                        "memo":"",
                        "code":"hr",
                        "sequence":3,
                        "parent":1,
                        "permissions":[
                        ],
                        "roles":[
                        ]
                    }],
                "role_participant":[
                ],
                "fields":[
                    {
                        "id":9,
                        "create_time":"2020-06-13 14:08:02",
                        "update_time":"2020-06-13 14:08:02",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":9,
                        "field_key":"hr_radio",
                        "field_name":"人事审批",
                        "order_id":80,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{\"1\":\"同意\", \"2\":\"不同意\"}",
                        "label":"{}",
                        "workflow":1
                    }]
            },
            "dest_state":{
                "id":3,
                "create_time":"2020-06-13 14:08:02",
                "update_time":"2020-06-13 14:08:02",
                "memo":"",
                "name":"申请人-编辑中",
                "is_hidden":false,
                "order_id":2,
                "state_type":0,
                "enable_retreat":false,
                "participant_type":0,
                "workflow":{
                    "id":1,
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-13 14:08:00",
                    "memo":"",
                    "name":"请假单",
                    "ticket_sn_prefix":"leave",
                    "status":true,
                    "view_permission_check":true,
                    "limit_expression":"{}",
                    "display_form_str":"[]",
                    "title_template":"你有一个待办工单:{title}",
                    "type":1
                },
                "user_participant":[
                ],
                "group_participant":[
                ],
                "role_participant":[
                ],
                "fields":[
                    {
                        "id":5,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":7,
                        "field_key":"start_end_time",
                        "field_name":"请假时间",
                        "order_id":10,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{}",
                        "label":"{}",
                        "workflow":1
                    },
                    {
                        "id":6,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":9,
                        "field_key":"type",
                        "field_name":"请假类型",
                        "order_id":30,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{\"1\":\"病假\", \"2\":\"产假\"}",
                        "label":"{}",
                        "workflow":1
                    },
                    {
                        "id":7,
                        "create_time":"2020-06-13 14:08:01",
                        "update_time":"2020-06-13 14:08:01",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":8,
                        "field_key":"memo",
                        "field_name":"事由说明",
                        "order_id":50,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{}",
                        "label":"{}",
                        "workflow":1
                    }]
            }
        },
        {
            "id":9,
            "create_time":"2020-06-13 14:08:05",
            "update_time":"2020-06-13 14:08:05",
            "memo":"",
            "name":4,
            "transition_type":0,
            "timer":0,
            "condition_expression":"[]",
            "attribute_type":5,
            "alert_enable":false,
            "alert_text":"",
            "workflow":{
                "id":1,
                "create_time":"2020-06-13 14:08:00",
                "update_time":"2020-06-13 14:08:00",
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
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-14 18:54:40",
                    "memo":"mark",
                    "name":"行政",
                    "code":"ad",
                    "order_id":1
                }
            },
            "source_state":{
                "id":5,
                "create_time":"2020-06-13 14:08:03",
                "update_time":"2020-06-13 14:08:03",
                "memo":"",
                "name":"人事-审批中",
                "is_hidden":false,
                "order_id":4,
                "state_type":0,
                "enable_retreat":false,
                "participant_type":2,
                "workflow":{
                    "id":1,
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-13 14:08:00",
                    "memo":"",
                    "name":"请假单",
                    "ticket_sn_prefix":"leave",
                    "status":true,
                    "view_permission_check":true,
                    "limit_expression":"{}",
                    "display_form_str":"[]",
                    "title_template":"你有一个待办工单:{title}",
                    "type":1
                },
                "user_participant":[
                ],
                "group_participant":[
                    {
                        "id":4,
                        "name":"人事",
                        "create_time":"2020-06-13 14:07:55",
                        "update_time":"2020-06-13 14:07:55",
                        "memo":"",
                        "code":"hr",
                        "sequence":3,
                        "parent":1,
                        "permissions":[
                        ],
                        "roles":[
                        ]
                    }],
                "role_participant":[
                ],
                "fields":[
                    {
                        "id":9,
                        "create_time":"2020-06-13 14:08:02",
                        "update_time":"2020-06-13 14:08:02",
                        "memo":"",
                        "field_attribute":false,
                        "field_type":9,
                        "field_key":"hr_radio",
                        "field_name":"人事审批",
                        "order_id":80,
                        "default_value":null,
                        "field_template":"",
                        "boolean_field_display":"{}",
                        "field_choice":"{\"1\":\"同意\", \"2\":\"不同意\"}",
                        "label":"{}",
                        "workflow":1
                    }]
            },
            "dest_state":{
                "id":2,
                "create_time":"2020-06-13 14:08:02",
                "update_time":"2020-06-13 14:08:02",
                "memo":"",
                "name":"关闭",
                "is_hidden":true,
                "order_id":99,
                "state_type":2,
                "enable_retreat":false,
                "participant_type":0,
                "workflow":{
                    "id":1,
                    "create_time":"2020-06-13 14:08:00",
                    "update_time":"2020-06-13 14:08:00",
                    "memo":"",
                    "name":"请假单",
                    "ticket_sn_prefix":"leave",
                    "status":true,
                    "view_permission_check":true,
                    "limit_expression":"{}",
                    "display_form_str":"[]",
                    "title_template":"你有一个待办工单:{title}",
                    "type":1
                },
                "user_participant":[
                ],
                "group_participant":[
                ],
                "role_participant":[
                ],
                "fields":[
                ]
            }
        }],
    "code":20000
}

*/
