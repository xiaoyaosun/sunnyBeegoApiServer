package controllers

import (
	"fmt"
	"testgoapi/models"
	_ "testgoapi/myfun"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/httplib"
)

// Operations about Http
type HttpController struct {
	beego.Controller
}


// @router /get [get]
func (r *HttpController) Get() {

	//req := httplib.Get("http://xxx:8082/api/web/v1/user/gettrafficlist?param=xxx")
	
	req := httplib.Get("http://xxx:8082/api/web/v1/user/gettrafficlist")
	req.Header("Accept-Encoding","gzip,deflate,sdch")
	
	var req_param models.Reqstr
	req_param.User = "xxx"
	req_param.Token = "xxx"
	req_param.Platform = "xxx"
	
	fmt.Println(req_param.User)
	fmt.Println(req_param.Token)
	fmt.Println(req_param.Platform)
	
	json_req, _ := json.Marshal(req_param)
	
	fmt.Println("json_req: ", string(json_req))
	
	//req.Param("param", "{\"userID\":\"xxx\",\"token\":\"xxx\",\"platform\":\"xxx\"}")
	
	req.Param("param", string(json_req))
	
	req.Debug(true)
	
	str1, _  := req.String()
	fmt.Println("str1: ", str1)
	
	var json_str models.Jsonstr
	var json_str2 models.Jsonstr
	var json_str3 map[string]interface{}
	
	json.Unmarshal([]byte(str1), &json_str2)
	
	fmt.Println(json_str2)
	
	json.Unmarshal([]byte(str1), &json_str3)
	
	fmt.Println(json_str3)
	
	//var result string
	//
	req.ToJSON(&json_str)
	//
	fmt.Println(json_str.Code)
	fmt.Println(json_str.Msg)
	fmt.Println(json_str.Data)
	
	r.Data["json"] = json_str
	
	r.ServeJSON()
}


