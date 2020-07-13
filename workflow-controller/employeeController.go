package controller

import (
	"fmt"
	"go-admin-demo/tools"

	"github.com/gin-gonic/gin"
)

type employee struct {
	Code string       `json:"code"`
	Msg  string       `json:"msg"`
	Data employeeData `json:"data"`
}

type employeeList struct {
	DepartmentName       string `json:"departmentName"`
	EmployeePhone        string `json:"employeePhone"`
	EmployeeName         string `json:"employeeName"`
	IsLeave              string `json:"isLeave"`
	EmployeeDepartmentID string `json:"employeeDepartmentId"`
	ID                   string `json:"id"`
	Open                 string `json:"open"`
	DepartmentNames      string `json:"departmentNames"`
}

type employeeData struct {
	NavigatepageNums  []string       `json:"navigatepageNums"`
	StartRow          string         `json:"startRow"`
	HasNextPage       string         `json:"hasNextPage"`
	PrePage           string         `json:"prePage"`
	NextPage          string         `json:"nextPage"`
	EndRow            string         `json:"endRow"`
	PageSize          string         `json:"pageSize"`
	List              []employeeList `json:"list"`
	PageNum           string         `json:"pageNum"`
	NavigatePages     string         `json:"navigatePages"`
	NavigateFirstPage string         `json:"navigateFirstPage"`
	Total             string         `json:"total"`
	Pages             string         `json:"pages"`
	Size              string         `json:"size"`
	IsLastPage        string         `json:"isLastPage"`
	HasPreviousPage   string         `json:"hasPreviousPage"`
	NavigateLastPage  string         `json:"navigateLastPage"`
	IsFirstPage       string         `json:"isFirstPage"`
}

// GetEmployeeData 获取雇员数据
func GetEmployeeData(c *gin.Context) {
	configID := c.Request.FormValue("configId")
	if len(configID) == 0 {
		tools.HasError(fmt.Errorf("configId NOT EXISTED"), "", 500)
	}

	fmt.Fprintf(c.Writer, "%s", configID)
}
