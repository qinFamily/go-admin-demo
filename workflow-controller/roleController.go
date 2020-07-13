package controller

import (
	"fmt"
	"go-admin-demo/tools"

	"github.com/gin-gonic/gin"
)

type role struct {
	Code string   `json:"code"`
	Msg  string   `json:"msg"`
	Data roleData `json:"data"`
}

type roleList struct {
	Code        string `json:"code"`
	RoleID      string `json:"roleId"`
	Scope       string `json:"scope"`
	RoleName    string `json:"roleName"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type roleData struct {
	NavigatepageNums  []interface{} `json:"navigatepageNums"`
	StartRow          string        `json:"startRow"`
	HasNextPage       string        `json:"hasNextPage"`
	PrePage           string        `json:"prePage"`
	NextPage          string        `json:"nextPage"`
	EndRow            string        `json:"endRow"`
	PageSize          string        `json:"pageSize"`
	List              []roleList    `json:"list"`
	PageNum           string        `json:"pageNum"`
	NavigatePages     string        `json:"navigatePages"`
	NavigateFirstPage string        `json:"navigateFirstPage"`
	Total             string        `json:"total"`
	Pages             string        `json:"pages"`
	Size              string        `json:"size"`
	IsLastPage        string        `json:"isLastPage"`
	HasPreviousPage   string        `json:"hasPreviousPage"`
	NavigateLastPage  string        `json:"navigateLastPage"`
	IsFirstPage       string        `json:"isFirstPage"`
}

// GetRoleData 获取初始化数据
func GetRoleData(c *gin.Context) {
	configID := c.Request.FormValue("configId")
	if len(configID) == 0 {
		tools.HasError(fmt.Errorf("configId NOT EXISTED"), "", 500)
	}

	fmt.Fprintf(c.Writer, "%s", configID)
}
