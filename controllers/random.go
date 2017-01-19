package controllers

import (
	_ "fmt"
	"testgoapi/models"
	"testgoapi/myfun"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
	_ "encoding/json"
)

// Operations about Random
type RandomController struct {
	beego.Controller
}

// @router /get [get]
func (r *RandomController) Get() {

	
	//var temp []map[int]interface{}
	//s := make(map[int]interface{})
	//for i := 1; i <= 10; i++ {
	//	id := myfun.RandomAlphanumeric(10)
	//	s[i] = id
	//	fmt.Println(s)
	//	temp = append(temp, s)
	//}
	//cnnJson := make(map[string]interface{})
	//cnnJson["id"] = temp
	//t,_ := json.Marshal(cnnJson)
	//fmt.Println("%s", string(t))
	
	var slice1 []string
	var id string
	for i := 1; i <= 1998; i++ {
		id = myfun.RandomAlphanumeric(10)
		slice1 = append(slice1, id)
		models.InsertAdminInfo(id)
		//fmt.Println(slice1)
	}
	
	r.Data["json"] = map[string]interface{}{"id":slice1}
	
	r.ServeJSON()
}


