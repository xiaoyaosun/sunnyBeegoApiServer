package models

import (
	_ "errors"
	_ "strconv"
	_ "time"
	_ "github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "fmt"
)

type RuleArr struct {
	RuleCode string `json:"ruleCode"`
	RuleName string `json:"ruleName"`
	Id string `json:"id"`
}

type TaskDataArr struct {
	Tid string `json:"id"`
	Available string `json:"isAvailable"`
	Pv string `json:"pv"`
	Op string `json:"opType"`
	ProdId string `json:"prodId"`
	Url string `json:"url"`
	RuleCode string `json:"ruleCode"`
	TaskDesc string `json:"taskDesc"`
	TaskType string `json:"taskType"`
	Bal int `json:"balance"`
	Tn string `json:"taskName"`
	Rd string `json:"ruleId"`
	Rarr RuleArr `json:"rule"`
}

type JsonData struct {
	TaskData []TaskDataArr `json:"taskinfo"`
	Balance string `json:"balance"`
}

type Jsonstr struct {
	Code int `json:"error_code"`
	Msg string `json:"error_msg"`
	Data JsonData `json:"data"`
}

// 生成请求串
type Reqstr struct {
	User string `json:"userID"`
	Token string `json:"token"`
	Platform string `json:"platform"`
}

