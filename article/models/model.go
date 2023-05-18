package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

///自动生成user的数据库表
//登录用户表的是设计
type User struct {
	Id       int
	UserName string
	Passwd   string
}

type Article struct {
	Id2     int    `orm:"pk;auto"`
	Title   string `orm:"size(20)"`      //文章标题
	Content string `orm:"size(500)"`     //内容
	Img     string `orm:"size(50);null"` //图片路径
	// Type    string    //类型
	Time  time.Time `orm:"type(datetime);auto_now_add"` //发布时间
	Count int       `orm:"default(0)"`                  //阅读量
}

///插入数据库
func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456789@tcp(127.0.0.1:3306)/newsweb?charset=utf8")
	//注册表,在后台自动生成
	orm.RegisterModel(new(User), new(Article))
	orm.RunSyncdb("default", false, true)

}
