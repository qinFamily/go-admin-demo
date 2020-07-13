package controller

import (
	"fmt"
	"go-admin-demo/tools"

	"github.com/gin-gonic/gin"
)

type flowState struct {
	Code string          `json:"code"`
	Msg  string          `json:"msg"`
	Data []flowStateData `json:"data"`
}

type flowStateData struct {
	OrderNum int    `json:"orderNum"`
	Active   bool   `json:"active"`
	Title    string `json:"title"`
	NavTo    string `json:"navTo"`
}

// GetFlowState 获取流程状态
func GetFlowState(c *gin.Context) {
	configID := c.Request.FormValue("configId")
	if len(configID) == 0 {
		tools.HasError(fmt.Errorf("configId NOT EXISTED"), "", 500)
	}

	fmt.Fprintf(c.Writer, "%s", configID)
}
