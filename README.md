# sunnyBeegoApiServer
According to the beego framework, try to deploy a API server

### Configure
You can custom your configure in conf/app.conf<br>
For the details you can see the [Beego.me](https://beego.me/docs/mvc/controller/config.md)

### Method to Request

#### Get content from DB

```
// @router /:id [get]
http://xxx:8080/v1/db/1

// @router /admin/:id [get]
http://xxx:8080/v1/db/admin/1

// @router /admin [get]
http://xxx:8080/v1/db/admin?uid=1
```

#### Get content from HTTP

```
// @router /get [get]
http://xxx:8080/v1/http/get
```
#### Get content from astaxie demo

```
// @router /login [get]
http://xxx:8080/v1/user/login?username=astaxie&password=11111
```

### Contact me
xiaoyaosun@qq.com
