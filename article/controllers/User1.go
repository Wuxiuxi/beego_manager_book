package controllers

import (
	"article/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegControler struct {
	beego.Controller
}

func (c *RegControler) ShowReg() {
	c.TplName = "register.html"
}

/*
 1.拿到浏览器传递的数据
 2.数据处理
 3.插入数据库(数据User)
 4.返回视图
*/
//注册业务
func (c *RegControler) HandReg() {
	// 1.拿到浏览器传递的数据
	name := c.GetString("userName")
	passwd := c.GetString("password")
	// beego.Info(name, passwd)

	///2.数据处理
	if name == "" || passwd == "" {
		beego.Info("用户名和密码不为空")
		c.TplName = "register.html"
		return
	} else {
		//插入数据库
		//1.获取ORM对象
		o := orm.NewOrm()
		user := models.User{}
		user.UserName = name
		user.Passwd = passwd
		//2.获取插入对象
		_, err := o.Insert(&user)
		if err != nil {
			beego.Info("插入数据失败")
			return
		}
		///4.返回登录
		// c.Ctx.WriteString("注册成功")
		c.TplName = "login.html"

		//重定向
		c.Redirect("", 302)

		//
	}
}

type LoginControler struct {
	beego.Controller
}

func (c *LoginControler) ShowLogin() {
	c.TplName = "login.html"
}

func (c *LoginControler) HandleLogin() {
	//1.拿数据
	name := c.GetString("userName")
	passwd := c.GetString("password")
	//1.拿到数据
	// beego.Info(name, passwd)

	if name == "" || passwd == "" {
		beego.Info("用户名和密码不能为空")
		c.TplName = "login.html"
		return
	}
	//3.查询数据库
	///1.获取orm对象
	o := orm.NewOrm()
	///2.获取查询对象
	user := models.User{}

	///3.查询
	user.UserName = name
	err := o.Read(&user, "UserName")
	if err != nil {
		beego.Info("登录失败")
		c.TplName = "login.html"
		return
	}
	///4.判断密码是否一致
	if user.Passwd != passwd {
		beego.Info("密码失败")
		c.TplName = "login.html"
		c.Ctx.WriteString("密码错误")
		return
	}
	//新网页显示密码错误
	c.Ctx.WriteString("登录成功")
}
