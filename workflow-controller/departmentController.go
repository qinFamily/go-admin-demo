package controller

import (
	"fmt"
	"go-admin-demo/tools"

	"github.com/gin-gonic/gin"
)

type department struct {
	Code string         `json:"code"`
	Msg  string         `json:"msg"`
	Data departmentData `json:"data"`
}

type departmentChildDepartments struct {
	DepartmentKey   string `json:"departmentKey"`
	DepartmentName  string `json:"departmentName"`
	ID              string `json:"id"`
	ParentID        string `json:"parentId"`
	DepartmentNames string `json:"departmentNames"`
}

type departmentEmployees struct {
	ID           string `json:"id"`
	EmployeeName string `json:"employeeName"`
	IsLeave      string `json:"isLeave"`
	Open         string `json:"open"`
}

type departmentData struct {
	ChildDepartments []departmentChildDepartments `json:"childDepartments"`
	Employees        []departmentEmployees        `json:"employees"`
	TitleDepartments []interface{}                `json:"titleDepartments"`
}

// GetDepartmentData 获取部门数据
func GetDepartmentData(c *gin.Context) {
	configID := c.Request.FormValue("configId")
	if len(configID) == 0 {
		tools.HasError(fmt.Errorf("configId NOT EXISTED"), "", 500)
	}

	fmt.Fprintf(c.Writer, "%s", configID)
}
