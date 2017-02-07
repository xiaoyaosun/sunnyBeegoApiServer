package models

import (
	_ "errors"
	"strconv"
	_ "time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type Admin struct {
	Uid int `orm:"pk"`
	Mid int
	Username string
	Userpass string
	Useremail string
	Addtime int
	Logintime int
	Loginip string	
}

type AdminInfo struct {
	Id string `orm:"pk"`
}

func init() {
	
	dbhost := beego.AppConfig.String("mysql.host")
    dbport := beego.AppConfig.String("mysql.port")
    dbuser := beego.AppConfig.String("mysql.user")
    dbpassword := beego.AppConfig.String("mysql.password")
    db := beego.AppConfig.String("mysql.schema")
	
	
    //注册mysql Driver
    orm.RegisterDriver("mysql", orm.DRMySQL)
    //构造conn连接
    conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + db + "?charset=utf8&loc=Asia%2FShanghai"
	
	// 设置最大空闲连接
	maxIdle := beego.AppConfig.DefaultInt("mysql.idlecon", 100)
	//maxIdle := 100
	
	// 设置最大数据库连接
	maxConn := beego.AppConfig.DefaultInt("mysql.maxcon", 100)
	//maxConn := 100
	
    //注册数据库连接
    orm.RegisterDataBase("default", "mysql", conn, maxIdle, maxConn)

	//orm.RegisterModel(new(Admin))
	
	// 带前缀的
	orm.RegisterModelWithPrefix("xiaolong_", new(Admin), new(AdminInfo))
	orm.RunSyncdb("default", false, true)
    fmt.Printf("数据库连接成功！%s\n", conn)
	
	// 设置为 UTC 时间
	//orm.DefaultTimeLoc = time.UTC
}


func GetAdmin(uid string) (u *Admin, err error) {
	fmt.Println("uid=", uid)
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	
	intid, err := strconv.Atoi(uid)
	fmt.Println("intid=", intid)
	
	admin := Admin{Uid: intid}
	
	o.Read(&admin)
	fmt.Println(admin)
	
	return &admin, nil
}

func InsertAdminInfo(id string) () {
	orm.Debug = true
	o := orm.NewOrm()
	
	random := new(AdminInfo)
	random.Id = id
	fmt.Println(o.Insert(random))

	//return true	
}



