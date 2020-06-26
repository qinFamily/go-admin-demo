package tickets

import (
	"go-admin-demo/models"
	"go-admin-demo/tools"
	"go-admin-demo/tools/app"
	"go-admin-demo/tools/app/msg"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetTicketsTicketList(c *gin.Context) {
	var data models.TicketsTicket
	var err error
	var pageSize = 10
	var pageIndex = 1

	id := c.Request.FormValue("id")
	if len(id) == 0 {
		if size := c.Request.FormValue("pageSize"); size != "" {
			pageSize = tools.StrToInt(err, size)
		}
		if index := c.Request.FormValue("pageIndex"); index != "" {
			pageIndex = tools.StrToInt(err, index)
		}
		transitionLt := c.Request.FormValue("transition__attribute_type__lt")
		// c.Request.FormValue("transition__attribute_type__lt")
		data.DataScope = tools.GetUserIdStr(c)
		result, count, err := data.GetPage(pageSize, pageIndex, transitionLt)
		tools.HasError(err, "", -1)

		app.PageOK(c, result, count, pageIndex, pageSize, "")
	} else {
		data.DataScope = tools.GetUserIdStr(c)
		data.Id, err = tools.StringToInt(id)
		tools.HasError(err, "抱歉id错误", -1)
		result, err := data.Get(true)
		tools.HasError(err, "抱歉未找到相关信息", -1)

		app.OK(c, result, "")
	}
}

func GetTicketsTicket(c *gin.Context) {
	var data models.TicketsTicket
	data.Id, _ = tools.StringToInt(c.Param("id"))
	result, err := data.Get(true)
	tools.HasError(err, "抱歉未找到相关信息", -1)

	app.OK(c, result, "")
}

func InsertTicketsTicket(c *gin.Context) {
	var data models.TicketsTicket
	err := c.ShouldBindJSON(&data)
	tools.HasError(err, "", 500)
	data.CreateBy = tools.GetUserId(c)
	result, err := data.Create()
	// 写 customefield
	customFileds := make([]models.TicketsTicketcustomfield, 0)
	if len(data.Customfield) > 0 {
		err = binding.JSON.BindBody([]byte(data.Customfield), customFileds)
		if err == nil {
			for _, cf := range customFileds {
				cf.TicketId = data.Id
				_, err = cf.Create()
				if err != nil {
					log.Println("ERROR!!!!!!!!!!!!!!!!!!!!!", err)
					tools.HasError(err, "", -1)
				}
			}
		} else {
			log.Println("bind customFileds ERROR!!!!!!!!!!!!!!!!!!!!!", err, data.Customfield)
		}
	}
	// 写 log
	// 写 user
	ttu := models.TicketsTicketuser{
		TicketId: data.Id,
		Username: tools.GetUserName(c),
	}
	_, err = ttu.Create()
	log.Println("create TicketsTicketuser ERROR!!!!!!!!!!!!!!!!!!!!!", err)
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

func UpdateTicketsTicket(c *gin.Context) {
	var data models.TicketsTicket
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "更新数据，数据解析失败", -1)
	// data.UpdateBy = tools.GetUserIdStr(c)
	IDS := tools.IdsStrToIdsIntGroup("flowId", c)
	if len(IDS) > 0 {
		data, err = data.Update(IDS[0])
		tools.HasError(err, "", -1)
	}

	app.OK(c, data, "")
}

func DeleteTicketsTicket(c *gin.Context) {
	var data models.TicketsTicket
	// data.UpdateBy = tools.GetUserIdStr(c)

	IDS := tools.IdsStrToIdsIntGroup("id", c)
	_, err := data.BatchDelete(IDS)
	tools.HasError(err, msg.DeletedFail, 500)
	app.OK(c, nil, msg.DeletedSuccess)
}

/*
{
create_user: 1
customfield: "[{"customfield":1,"field_key":"create_user","field_value":""},{"customfield":2,"field_key":"create_time","field_value":""},{"customfield":3,"field_key":"group","field_value":""},{"customfield":4,"field_key":"id","field_value":""},{"customfield":5,"field_key":"start_end_time","field_value":["2020-06-25","2020-07-24"]},{"customfield":6,"field_key":"type","field_value":"2"},{"customfield":7,"field_key":"memo","field_value":"建设银行"},{"customfield":8,"field_key":"leader_radio","field_value":""},{"customfield":9,"field_key":"hr_radio","field_value":""}]"
name: "请假单-2020-06-25-13-42-27-277"
participant: ""
state: 3
transition: 1
workflow: 1
}
{
    "results":{
        "id":2,
        "create_time":"2020-06-25 13:50:09",
        "update_time":"2020-06-25 13:50:09",
        "memo":"",
        "name":"请假单-2020-06-25-13-42-27-277",
        "sn":"leave_20200625135009552",
        "participant":"",
        "customfield":"[{\"customfield\":1,\"field_key\":\"create_user\",\"field_value\":\"\"},{\"customfield\":2,\"field_key\":\"create_time\",\"field_value\":\"\"},{\"customfield\":3,\"field_key\":\"group\",\"field_value\":\"\"},{\"customfield\":4,\"field_key\":\"id\",\"field_value\":\"\"},{\"customfield\":5,\"field_key\":\"start_end_time\",\"field_value\":[\"2020-06-25\",\"2020-07-24\"]},{\"customfield\":6,\"field_key\":\"type\",\"field_value\":\"2\"},{\"customfield\":7,\"field_key\":\"memo\",\"field_value\":\"建设银行\"},{\"customfield\":8,\"field_key\":\"leader_radio\",\"field_value\":\"\"},{\"customfield\":9,\"field_key\":\"hr_radio\",\"field_value\":\"\"}]",
        "create_user":1,
        "workflow":1,
        "state":3,
        "transition":1
    },
    "code":20000
}


#/todo_ticket
http://127.0.0.1:8080/api/ticket/ticket/?transition__attribute_type__lt=4
GET
{
    "results":[
        {
            "id":1,
            "create_time":"2020-06-13 18:33:36",
            "update_time":"2020-06-25 16:22:53",
            "memo":"",
            "name":"请假单-2020-06-13-18-32-48-202",
            "sn":"leave_20200613183336522",
            "participant":"",
            "customfield":"[{\"id\":1,\"ticket\":1,\"customfield\":1,\"field_value\":\"admin\"},{\"id\":2,\"ticket\":1,\"customfield\":2,\"field_value\":\"2020-06-13 10:33:36.882952+00:00\"},{\"id\":3,\"ticket\":1,\"customfield\":3,\"field_value\":\"top\"},{\"id\":4,\"ticket\":1,\"customfield\":4,\"field_value\":1},{\"id\":5,\"ticket\":1,\"customfield\":5,\"field_value\":\"['2020-06-16', '2020-07-16']\"},{\"id\":6,\"ticket\":1,\"customfield\":6,\"field_value\":\"1\"},{\"id\":7,\"ticket\":1,\"customfield\":7,\"field_value\":\"随便写写\"},{\"id\":8,\"ticket\":1,\"customfield\":8,\"field_value\":\"\"},{\"id\":9,\"ticket\":1,\"customfield\":9,\"field_value\":\"\"}]",
            "create_user":{
                "id":1,
                "password":"pbkdf2_sha256$180000$IyDdLTBtIQzI$tMajuC4dHeDh+JziEaJBMnDQmoFUfpvbq7VFtzH9Tks=",
                "last_login":null,
                "create_time":"2020-06-13 14:07:12",
                "update_time":"2020-06-13 14:07:12",
                "memo":"",
                "username":"admin",
                "realname":"图书馆管理员",
                "email":"itimor@126.com",
                "avatar":"http://m.imeitou.com/uploads/allimg/2017110610/b3c433vwhsk.jpg",
                "status":true,
                "is_admin":true,
                "group":1,
                "roles":[
                    1],
                "model_perms":[
                ]
            },
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
            "state":{
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
                "workflow":1,
                "user_participant":[
                ],
                "group_participant":[
                ],
                "role_participant":[
                ],
                "fields":[
                ]
            },
            "transition":{
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
                "workflow":1,
                "source_state":3,
                "dest_state":6
            }
        },
        {
            "id":2,
            "create_time":"2020-06-25 13:50:09",
            "update_time":"2020-06-25 13:50:09",
            "memo":"",
            "name":"请假单-2020-06-25-13-42-27-277",
            "sn":"leave_20200625135009552",
            "participant":"",
            "customfield":"[{\"customfield\":1,\"field_key\":\"create_user\",\"field_value\":\"\"},{\"customfield\":2,\"field_key\":\"create_time\",\"field_value\":\"\"},{\"customfield\":3,\"field_key\":\"group\",\"field_value\":\"\"},{\"customfield\":4,\"field_key\":\"id\",\"field_value\":\"\"},{\"customfield\":5,\"field_key\":\"start_end_time\",\"field_value\":[\"2020-06-25\",\"2020-07-24\"]},{\"customfield\":6,\"field_key\":\"type\",\"field_value\":\"2\"},{\"customfield\":7,\"field_key\":\"memo\",\"field_value\":\"建设银行\"},{\"customfield\":8,\"field_key\":\"leader_radio\",\"field_value\":\"\"},{\"customfield\":9,\"field_key\":\"hr_radio\",\"field_value\":\"\"}]",
            "create_user":{
                "id":1,
                "password":"pbkdf2_sha256$180000$IyDdLTBtIQzI$tMajuC4dHeDh+JziEaJBMnDQmoFUfpvbq7VFtzH9Tks=",
                "last_login":null,
                "create_time":"2020-06-13 14:07:12",
                "update_time":"2020-06-13 14:07:12",
                "memo":"",
                "username":"admin",
                "realname":"图书馆管理员",
                "email":"itimor@126.com",
                "avatar":"http://m.imeitou.com/uploads/allimg/2017110610/b3c433vwhsk.jpg",
                "status":true,
                "is_admin":true,
                "group":1,
                "roles":[
                    1],
                "model_perms":[
                ]
            },
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
            "state":{
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
                "workflow":1,
                "user_participant":[
                ],
                "group_participant":[
                ],
                "role_participant":[
                ],
                "fields":[
                    5,
                    6,
                    7]
            },
            "transition":{
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
                "workflow":1,
                "source_state":1,
                "dest_state":3
            }
        },
        {
            "id":3,
            "create_time":"2020-06-25 15:40:04",
            "update_time":"2020-06-25 15:40:04",
            "memo":"",
            "name":"请假单-2020-06-25-15-38-25-654",
            "sn":"leave_20200625154004503",
            "participant":"hr_tl",
            "customfield":"[{\"customfield\":1,\"field_key\":\"create_user\",\"field_value\":\"\"},{\"customfield\":2,\"field_key\":\"create_time\",\"field_value\":\"\"},{\"customfield\":3,\"field_key\":\"group\",\"field_value\":\"\"},{\"customfield\":4,\"field_key\":\"id\",\"field_value\":\"\"},{\"customfield\":5,\"field_key\":\"start_end_time\",\"field_value\":[\"2020-06-24\",\"2020-07-20\"]},{\"customfield\":6,\"field_key\":\"type\",\"field_value\":\"2\"},{\"customfield\":7,\"field_key\":\"memo\",\"field_value\":\"好像生产了\"},{\"customfield\":8,\"field_key\":\"leader_radio\",\"field_value\":\"\"},{\"customfield\":9,\"field_key\":\"hr_radio\",\"field_value\":\"\"}]",
            "create_user":{
                "id":1,
                "password":"pbkdf2_sha256$180000$IyDdLTBtIQzI$tMajuC4dHeDh+JziEaJBMnDQmoFUfpvbq7VFtzH9Tks=",
                "last_login":null,
                "create_time":"2020-06-13 14:07:12",
                "update_time":"2020-06-13 14:07:12",
                "memo":"",
                "username":"admin",
                "realname":"图书馆管理员",
                "email":"itimor@126.com",
                "avatar":"http://m.imeitou.com/uploads/allimg/2017110610/b3c433vwhsk.jpg",
                "status":true,
                "is_admin":true,
                "group":1,
                "roles":[
                    1],
                "model_perms":[
                ]
            },
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
            "state":{
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
                "workflow":1,
                "user_participant":[
                ],
                "group_participant":[
                ],
                "role_participant":[
                    2,
                    4,
                    6],
                "fields":[
                    8]
            },
            "transition":{
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
                "workflow":1,
                "source_state":1,
                "dest_state":4
            }
        }],
    "code":20000
}

#/s_ticket/1
http://127.0.0.1:8080/api/ticket/ticket/?id=1
{
    "results":[
        {
            "id":1,
            "create_time":"2020-06-13 18:33:36",
            "update_time":"2020-06-25 16:22:53",
            "memo":"",
            "name":"请假单-2020-06-13-18-32-48-202",
            "sn":"leave_20200613183336522",
            "participant":"",
            "customfield":"[{\"id\":1,\"ticket\":1,\"customfield\":1,\"field_value\":\"admin\"},{\"id\":2,\"ticket\":1,\"customfield\":2,\"field_value\":\"2020-06-13 10:33:36.882952+00:00\"},{\"id\":3,\"ticket\":1,\"customfield\":3,\"field_value\":\"top\"},{\"id\":4,\"ticket\":1,\"customfield\":4,\"field_value\":1},{\"id\":5,\"ticket\":1,\"customfield\":5,\"field_value\":\"['2020-06-16', '2020-07-16']\"},{\"id\":6,\"ticket\":1,\"customfield\":6,\"field_value\":\"1\"},{\"id\":7,\"ticket\":1,\"customfield\":7,\"field_value\":\"随便写写\"},{\"id\":8,\"ticket\":1,\"customfield\":8,\"field_value\":\"\"},{\"id\":9,\"ticket\":1,\"customfield\":9,\"field_value\":\"\"}]",
            "create_user":{
                "id":1,
                "password":"pbkdf2_sha256$180000$IyDdLTBtIQzI$tMajuC4dHeDh+JziEaJBMnDQmoFUfpvbq7VFtzH9Tks=",
                "last_login":null,
                "create_time":"2020-06-13 14:07:12",
                "update_time":"2020-06-13 14:07:12",
                "memo":"",
                "username":"admin",
                "realname":"图书馆管理员",
                "email":"itimor@126.com",
                "avatar":"http://m.imeitou.com/uploads/allimg/2017110610/b3c433vwhsk.jpg",
                "status":true,
                "is_admin":true,
                "group":1,
                "roles":[
                    1],
                "model_perms":[
                ]
            },
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
            "state":{
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
                "workflow":1,
                "user_participant":[
                ],
                "group_participant":[
                ],
                "role_participant":[
                ],
                "fields":[
                ]
            },
            "transition":{
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
                "workflow":1,
                "source_state":3,
                "dest_state":6
            }
        }],
    "code":20000
}


http://127.0.0.1:8080/api/ticket/ticket/?page=1&limit=20
GET
{
    "code":200,
    "data":{
        "list":[
            {
                "id":1,
                "createTime":"2020-06-13T10:33:37+08:00",
                "updateTime":"2020-06-13T10:33:37+08:00",
                "memo":"",
                "name":"请假单-2020-06-13-18-32-48-202",
                "sn":"leave_20200613183336522",
                "participant":"",
                "customfield":"[{`customfield`:1,`field_key`:`create_user`,`field_value`:``},{`customfield`:2,`field_key`:`create_time`,`field_value`:``},{`customfield`:3,`field_key`:`group`,`field_value`:``},{`customfield`:4,`field_key`:`id`,`field_value`:``},{`customfield`:5,`field_key`:`start_end_time`,`field_value`:[`2020-06-16`,`2020-07-16`]},{`customfield`:6,`field_key`:`type`,`field_value`:`1`},{`customfield`:7,`field_key`:`memo`,`field_value`:`随便写写`},{`customfield`:8,`field_key`:`leader_radio`,`field_value`:``},{`customfield`:9,`field_key`:`hr_radio`,`field_value`:``}]",
                "createdAt":"2020-06-25T15:26:51+08:00",
                "updatedAt":"2020-06-25T15:26:51+08:00",
                "deletedAt":"",
                "workflowId":1,
                "transitionId":1,
                "stateId":3,
                "createBy":1,
                "updateBy":0,
                "dataScope":"",
                "params":""
            },
            {
                "id":2,
                "createTime":"2020-06-25T05:50:10+08:00",
                "updateTime":"2020-06-25T05:50:10+08:00",
                "memo":"",
                "name":"请假单-2020-06-25-13-42-27-277",
                "sn":"leave_20200625135009552",
                "participant":"",
                "customfield":"[{`customfield`:1,`field_key`:`create_user`,`field_value`:``},{`customfield`:2,`field_key`:`create_time`,`field_value`:``},{`customfield`:3,`field_key`:`group`,`field_value`:``},{`customfield`:4,`field_key`:`id`,`field_value`:``},{`customfield`:5,`field_key`:`start_end_time`,`field_value`:[`2020-06-25`,`2020-07-24`]},{`customfield`:6,`field_key`:`type`,`field_value`:`2`},{`customfield`:7,`field_key`:`memo`,`field_value`:`建设银行`},{`customfield`:8,`field_key`:`leader_radio`,`field_value`:``},{`customfield`:9,`field_key`:`hr_radio`,`field_value`:``}]",
                "createdAt":"2020-06-25T15:26:51+08:00",
                "updatedAt":"2020-06-25T15:26:51+08:00",
                "deletedAt":"",
                "workflowId":1,
                "transitionId":1,
                "stateId":3,
                "createBy":1,
                "updateBy":0,
                "dataScope":"",
                "params":""
            }],
        "count":2,
        "pageIndex":1,
        "pageSize":10
    },
    "msg":""
}

*/
