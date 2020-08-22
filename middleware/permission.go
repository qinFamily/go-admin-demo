/*
 * @Author: xiaoxu@mgtv.com
 * @Date: 2020-05-24 17:32:46
 * @Jira:
 * @Wiki:
 * @LastEditTime: 2020-08-22 19:50:11
 * @LastEditors: xiaoxu
 * @Description:
 * @FilePath: \go-admin-ui-vuef:\project\work\go\src\go-admin-demo\middleware\permission.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package middleware

import (
	mycasbin "go-admin-demo/pkg/casbin"
	"go-admin-demo/pkg/jwtauth"
	_ "go-admin-demo/pkg/jwtauth"
	"go-admin-demo/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

//权限检查中间件
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.Get("JWT_PAYLOAD")
		v := data.(jwtauth.MapClaims)
		e, err := mycasbin.Casbin()
		tools.HasError(err, "", 500)
		//检查权限
		res, err := e.Enforce(v["rolekey"], c.Request.URL.Path, c.Request.Method)
		// log.Println("----------------", v["rolekey"], c.Request.URL.Path, c.Request.Method, "error", err)

		tools.HasError(err, "", 500)

		if res {
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "对不起，您没有该接口访问权限，请联系管理员",
			})
			c.Abort()
			return
		}
	}
}
