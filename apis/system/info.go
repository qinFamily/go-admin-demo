package system

import (
	"go-admin-demo/models"
	"go-admin-demo/tools"
	"go-admin-demo/tools/app"

	"github.com/gin-gonic/gin"
)

func GetInfo(c *gin.Context) {

	var roles = make([]string, 1)
	roles[0] = tools.GetRoleName(c)

	var permissions = make([]string, 1)
	permissions[0] = "*:*:*"
	RoleMenu := models.RoleMenu{}
	RoleMenu.RoleId = tools.GetRoleId(c)

	var mp = make(map[string]interface{})
	mp["roles"] = roles
	if tools.GetRoleName(c) == "admin" || tools.GetRoleName(c) == "系统管理员" {
		mp["permissions"] = permissions
	} else {
		list, _ := RoleMenu.GetPermis()
		mp["permissions"] = list
	}

	sysuser := models.SysUser{}
	sysuser.UserId = tools.GetUserId(c)
	user, err := sysuser.Get()
	tools.HasError(err, "", 500)

	mp["introduction"] = " am a super administrator"

	mp["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	// mp["avatar"] = "http://m.imeitou.com/uploads/allimg/2017110610/b3c433vwhsk.jpg"
	// mp["avatar"] = "https://wpimg.wallstcn.com/577965b9-bb9e-4e02-9f0c-095b41417191"
	// mp["avatar"] = "https://wpimg.wallstcn.com/e7d23d71-cf19-4b90-a1cc-f56af8c0903d.png"
	// mp["avatar"] = "https://wpimg.wallstcn.com/007ef517-bafd-4066-aae4-6883632d9646"
	if user.Avatar != "" {
		mp["avatar"] = user.Avatar
	}
	mp["name"] = user.NickName
	mp["ip"] = c.ClientIP()
	mp["userId"] = user.UserId
	app.OK(c, mp, "")
}
