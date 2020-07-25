package router

import (
	"go-admin-demo/apis"
	"go-admin-demo/apis/grid"
	"go-admin-demo/apis/grid/base"
	"go-admin-demo/apis/grid/base/housing"
	log2 "go-admin-demo/apis/log"
	"go-admin-demo/apis/monitor"
	"go-admin-demo/apis/system"
	"go-admin-demo/apis/system/dict"
	"go-admin-demo/apis/tickets"
	. "go-admin-demo/apis/tools"
	"go-admin-demo/apis/workflow"
	_ "go-admin-demo/docs"
	"go-admin-demo/handler"
	"go-admin-demo/handler/sd"
	"go-admin-demo/middleware"
	_ "go-admin-demo/pkg/jwtauth"
	controller "go-admin-demo/workflow-controller"
	"log"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.LoggerToFile())
	r.Use(middleware.CustomError)
	r.Use(middleware.NoCache)
	r.Use(middleware.Options)
	r.Use(middleware.Secure)
	r.Use(middleware.RequestId())

	r.GET("/", system.HelloWorld)
	r.Static("/static", "./static")
	r.GET("/info", handler.Ping)

	// 监控信息
	svcd := r.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
		svcd.GET("/os", sd.OSCheck)
	}

	// the jwt middleware
	authMiddleware, err := middleware.AuthInit()

	if err != nil {
		log.Fatalln("JWT Error", err.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)

	// Refresh time can be longer than token timeout
	r.GET("/refresh_token", authMiddleware.RefreshHandler)
	r.GET("/routes", Dashboard)

	apiv1 := r.Group("/api/v1")
	{

		apiv1.GET("/monitor/server", monitor.ServerInfo)

		apiv1.GET("/getCaptcha", system.GenerateCaptchaHandler)
		apiv1.GET("/db/tables/page", GetDBTableList)
		apiv1.GET("/db/columns/page", GetDBColumnList)
		apiv1.GET("/sys/tables/page", GetSysTableList)
		apiv1.POST("/sys/tables/info", InsertSysTable)
		apiv1.PUT("/sys/tables/info", UpdateSysTable)
		apiv1.DELETE("/sys/tables/info/:tableId", DeleteSysTables)
		apiv1.GET("/sys/tables/info/:tableId", GetSysTables)
		apiv1.GET("/gen/preview/:tableId", Preview)
		apiv1.GET("/menuTreeselect", system.GetMenuTreeelect)
		apiv1.GET("/dict/databytype/:dictType", dict.GetDictDataByDictType)
		// apiv1.GET("/workflow/transition/", workflow.GetWorkFlowTransition)

	}

	auth := r.Group("/api/v1")
	auth.Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{

		auth.GET("/deptList", system.GetDeptList)
		auth.GET("/deptTree", system.GetDeptTree)
		auth.GET("/dept/:deptId", system.GetDept)
		auth.POST("/dept", system.InsertDept)
		auth.PUT("/dept", system.UpdateDept)
		auth.DELETE("/dept/:id", system.DeleteDept)

		auth.GET("/dict/datalist", dict.GetDictDataList)
		auth.GET("/dict/data/:dictCode", dict.GetDictData)
		auth.POST("/dict/data", dict.InsertDictData)
		auth.PUT("/dict/data/", dict.UpdateDictData)
		auth.DELETE("/dict/data/:dictCode", dict.DeleteDictData)

		auth.GET("/dict/typelist", dict.GetDictTypeList)
		auth.GET("/dict/type/:dictId", dict.GetDictType)
		auth.POST("/dict/type", dict.InsertDictType)
		auth.PUT("/dict/type", dict.UpdateDictType)
		auth.DELETE("/dict/type/:dictId", dict.DeleteDictType)

		auth.GET("/dict/typeoptionselect", dict.GetDictTypeOptionSelect)

		auth.GET("/sysUserList", system.GetSysUserList)
		auth.GET("/sysUser/:userId", system.GetSysUser)
		auth.GET("/sysUser/", system.GetSysUserInit)
		auth.POST("/sysUser", system.InsertSysUser)
		auth.PUT("/sysUser", system.UpdateSysUser)
		auth.DELETE("/sysUser/:userId", system.DeleteSysUser)

		auth.GET("/rolelist", system.GetRoleList)
		auth.GET("/role/:roleId", system.GetRole)
		auth.POST("/role", system.InsertRole)
		auth.PUT("/role", system.UpdateRole)
		auth.PUT("/roledatascope", system.UpdateRoleDataScope)
		auth.DELETE("/role/:roleId", system.DeleteRole)

		auth.GET("/configList", system.GetConfigList)
		auth.GET("/config/:configId", system.GetConfig)
		auth.POST("/config", system.InsertConfig)
		auth.PUT("/config", system.UpdateConfig)
		auth.DELETE("/config/:configId", system.DeleteConfig)

		auth.GET("/roleMenuTreeselect/:roleId", system.GetMenuTreeRoleselect)
		auth.GET("/roleDeptTreeselect/:roleId", system.GetDeptTreeRoleselect)

		auth.GET("/getinfo", system.GetInfo)
		auth.GET("/user/profile", system.GetSysUserProfile)
		auth.POST("/user/avatar", system.InsetSysUserAvatar)
		auth.PUT("/user/pwd", system.SysUserUpdatePwd)

		auth.GET("/postlist", system.GetPostList)
		auth.GET("/post/:postId", system.GetPost)
		auth.POST("/post", system.InsertPost)
		auth.PUT("/post", system.UpdatePost)
		auth.DELETE("/post/:postId", system.DeletePost)

		auth.GET("/menulist", system.GetMenuList)
		auth.GET("/menu/:id", system.GetMenu)
		auth.POST("/menu", system.InsertMenu)
		auth.PUT("/menu", system.UpdateMenu)
		auth.DELETE("/menu/:id", system.DeleteMenu)
		auth.GET("/menurole", system.GetMenuRole)

		auth.GET("/menuids", system.GetMenuIDS)

		auth.GET("/loginloglist", log2.GetLoginLogList)
		auth.GET("/loginlog/:infoId", log2.GetLoginLog)
		auth.POST("/loginlog", log2.InsertLoginLog)
		auth.PUT("/loginlog", log2.UpdateLoginLog)
		auth.DELETE("/loginlog/:infoId", log2.DeleteLoginLog)

		auth.GET("/operloglist", log2.GetOperLogList)
		auth.GET("/operlog/:operId", log2.GetOperLog)
		auth.DELETE("/operlog/:operId", log2.DeleteOperLog)

		auth.GET("/configKey/:configKey", system.GetConfigByConfigKey)

		auth.POST("/logout", handler.LogOut)

		auth.GET("/articleList", apis.GetArticleList)
		auth.GET("/article/:articleId", apis.GetArticle)
		auth.POST("/article", apis.InsertArticle)
		auth.PUT("/article", apis.UpdateArticle)
		auth.DELETE("/article/:articleId", apis.DeleteArticle)

		// 网格信息
		auth.GET("/grid/base/housingList", housing.GetGridBasicHousingList)
		auth.GET("/grid/base/housing/:housId", housing.GetGridBasicHousing)
		auth.POST("/grid/base/housing", housing.InsertGridBasicHousing)
		auth.PUT("/grid/base/housing", housing.UpdateGridBasicHousing)
		auth.DELETE("/grid/base/housing/:housId", housing.DeleteGridBasicHousing)

		auth.GET("/grid/base/buildingList", housing.GetGridBasicBuidingList)
		auth.GET("/grid/base/building/:buildingId", housing.GetGridBasicBuiding)
		auth.POST("/grid/base/building", housing.InsertGridBasicBuiding)
		auth.PUT("/grid/base/building", housing.UpdateGridBasicBuiding)
		auth.DELETE("/grid/base/building/:buildingId", housing.DeleteGridBasicBuiding)

		auth.GET("/grid/basic/peopleList", base.GetGridBasicPeopleList)
		auth.GET("/grid/basic/people/:peopleId", base.GetGridBasicPeople)
		auth.POST("/grid/basic/people", base.InsertGridBasicPeople)
		auth.PUT("/grid/basic/people", base.UpdateGridBasicPeople)
		auth.DELETE("/grid/basic/people/:peopleId", base.DeleteGridBasicPeople)

		auth.GET("/grid/gridList", grid.GetGridList)
		auth.GET("/grid/grid/:gridId", grid.GetGrid)
		auth.POST("/grid/grid", grid.InsertGrid)
		auth.PUT("/grid/grid", grid.UpdateGrid)
		auth.DELETE("/grid/grid/:gridId", grid.DeleteGrid)

		// 工作流
		auth.GET("/sys/auth/getmenubutons", system.GetMenubButons)
		auth.GET("/workflow/workflowtype/", workflow.GetWorkFlowType)
		auth.POST("/workflow/workflowtype/", workflow.InsertWorkFlowType)
		auth.PUT("/workflow/workflowtype/:flowtypeId/", workflow.UpdateWorkFlowType)
		auth.GET("/workflow/workflow/", workflow.GetWorkFlow)
		auth.POST("/workflow/workflow/", workflow.InsertWorkFlow)
		auth.PUT("/workflow/workflow/:flowId/", workflow.UpdateWorkFlow)
		auth.DELETE("/workflow/workflow/:flowId/", workflow.DeleteWorkflowsWorkflow)
		auth.DELETE("/workflow/workflowtype/:flowtypeId/", workflow.DeleteWorkflowsWorkflow)

		auth.GET("/workflow/customfield/", workflow.GetWorkFlowCustomField)
		auth.POST("/workflow/customfield/", workflow.InsertWorkFlowCustomField)
		auth.GET("/workflow/state/", workflow.GetWorkFlowState)
		auth.POST("/workflow/state/", workflow.InsertWorkFlowState)
		auth.GET("/workflow/transition/", workflow.GetWorkFlowTransition)
		auth.POST("/workflow/transition/", workflow.InsertWorkFlowTransition)
		auth.PUT("/workflow/wfconf/:flowId", workflow.UpdateWorkFlow)
		// auth.PUT("", workflow.UpdateWorkFlow)
		auth.GET("/sys/user/", system.GetSysUserWorkflow)
		auth.GET("/role/", system.GetRoleInit)
		auth.GET("/dept/", system.GetDeptInit)

		// 工单系统
		auth.GET("/ticket/ticket/", tickets.GetTicketsTicketList)
		auth.POST("/ticket/ticket/", tickets.InsertTicketsTicket)
		auth.DELETE("/ticket/ticket/bulk_delete/", tickets.DeleteTicketsTicket)
		auth.PUT("/ticket/ticket/:id/", tickets.UpdateTicketsTicket) // 驳回工单
		auth.GET("/ticket/ticketflowlog/", tickets.GetTicketsTicketflowlogList)
		auth.GET("/ticket/ticketcustomfield/", tickets.GetTicketsTicketcustomfieldList)
		auth.GET("/ticket/ticketuser/", tickets.GetTicketsTicketuserList)

		// 钉钉样式的工作流

		auth.GET("/workflow/", controller.Index)
		auth.GET("/workflow/procdef", controller.GetInitData)
		//-------------------------流程定义----------------------
		auth.POST("/workflow/procdef/save", controller.SaveProcdef)
		auth.POST("/workflow/procdef/saveByToken", controller.SaveProcdefByToken)
		auth.GET("/workflow/procdef/findAll", controller.FindAllProcdefPage)
		auth.GET("/workflow/procdef/delById", controller.DelProcdefByID)
		// -----------------------流程实例-----------------------
		auth.POST("/workflow/process/start", controller.StartProcessInstance)               // 启动流程
		auth.POST("/workflow/process/startByToken", controller.StartProcessInstanceByToken) // 启动流程
		auth.POST("/workflow/process/findTask", controller.FindMyProcInstPageAsJSON)        // 查询需要我审批的流程
		auth.POST("/workflow/process/findById", controller.FindProcInstByID)                // 根据id查询流程实例
		auth.POST("/workflow/process/findTaskByToken", controller.FindMyProcInstByToken)
		auth.POST("/workflow/process/startByMyself", controller.StartByMyself)   // 查询我启动的流程
		auth.POST("/workflow/process/FindProcNotify", controller.FindProcNotify) // 查询抄送我的流程
		// auth.GET("/workflow/process/moveToHistory", controller.MoveFinishedProcInstToHistory)
		// -----------------------任务--------------------------
		auth.POST("/workflow/task/complete", controller.CompleteTask)
		auth.POST("/workflow/task/completeByToken", controller.CompleteTaskByToken)
		auth.POST("/workflow/task/withdraw", controller.WithDrawTask)
		auth.POST("/workflow/task/withdrawByToken", controller.WithDrawTaskByToken)
		// ----------------------- 关系表 -------------------------
		auth.GET("/workflow/identitylink/findParticipant", controller.FindParticipantByProcInstID)

		// ******************************** 历史纪录 ***********************************
		// -------------------------- 流程实例 -------------------------------
		auth.POST("/workflow/procHistory/findTask", controller.FindProcHistory)
		auth.POST("/workflow/procHistory/findTaskByToken", controller.FindProcHistoryByToken)
		auth.POST("/workflow/procHistory/startByMyself", controller.StartHistoryByMyself)   // 查询我启动的流程
		auth.POST("/workflow/procHistory/FindProcNotify", controller.FindProcHistoryNotify) // 查询抄送我的流程
		// ----------------------- 关系表 -------------------------
		auth.GET("/workflow/identitylinkHistory/findParticipant", controller.FindParticipantHistoryByProcInstID)

	}
	//r.NoRoute(authMiddleware.MiddlewareFunc(), NoFound)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println("路由加载成功！")
	return r
}

func Dashboard(c *gin.Context) {

	var user = make(map[string]interface{})
	user["login_name"] = "admin"
	user["user_id"] = 1
	user["user_name"] = "管理员"
	user["dept_id"] = 1

	var cmenuList = make(map[string]interface{})
	cmenuList["children"] = nil
	cmenuList["parent_id"] = 1
	cmenuList["title"] = "用户管理"
	cmenuList["name"] = "Sysuser"
	cmenuList["icon"] = "user"
	cmenuList["order_num"] = 1
	cmenuList["id"] = 4
	cmenuList["path"] = "sysuser"
	cmenuList["component"] = "sysuser/index"

	var lista = make([]interface{}, 1)
	lista[0] = cmenuList

	var menuList = make(map[string]interface{})
	menuList["children"] = lista
	menuList["parent_id"] = 1
	menuList["name"] = "Upms"
	menuList["title"] = "权限管理"
	menuList["icon"] = "example"
	menuList["order_num"] = 1
	menuList["id"] = 4
	menuList["path"] = "/upms"
	menuList["component"] = "Layout"

	var list = make([]interface{}, 1)
	list[0] = menuList
	var data = make(map[string]interface{})
	data["user"] = user
	data["menuList"] = list

	var r = make(map[string]interface{})
	r["code"] = 200
	r["data"] = data

	c.JSON(200, r)
}
