package controllers

import (
	_ "fmt"
	"testgoapi/models"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
)

// Operations about Db
type DbController struct {
	beego.Controller
}


// @Title Get
// @Description get user by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (u *DbController) Get() {
	id := u.GetString(":id")
	if id != "" {
		user, err := models.GetAdmin(id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}



