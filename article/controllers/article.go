package controllers

import (
	"article/models"
	"math"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ArticleControler struct {
	beego.Controller
}

///展示文章列表页
func (this *ArticleControler) ShowArticleList1() {
	//获取数据
	//高级查询
	//指定表
	o := orm.NewOrm()
	qs := o.QueryTable("Article") //queryseter
	var articles []models.Article
	//_,err := qs.All(&articles)
	//if err != nil{
	//	beego.Info("查询数据错误")
	//}
	//查询总记录数
	count, _ := qs.Count()
	//获取总页数
	pageSize := 2

	pageCount := math.Ceil(float64(count) / float64(pageSize))
	//获取页码
	pageIndex, err := this.GetInt("pageIndex")
	if err != nil {
		pageIndex = 1
	}

	//获取数据
	//作用就是获取数据库部分数据,第一个参数，获取几条,第二个参数，从那条数据开始获取,返回值还是querySeter
	//起始位置计算
	start := (pageIndex - 1) * pageSize

	//qs.Limit(pageSize,start).RelatedSel("ArticleType").All(&articles)

	//获取文章类型
	// var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)
	// this.Data["types"] = types

	//根据选中的类型查询相应类型文章
	typeName := this.GetString("select")
	beego.Info(typeName)
	qs.Limit(pageSize, start).RelatedSel("ArticleType").Filter("ArticleType__TypeName", typeName).All(&articles)

	//传递数据
	this.Data["pageIndex"] = pageIndex
	this.Data["pageCount"] = int(pageCount)
	this.Data["count"] = count
	this.Data["articles"] = articles
	this.TplName = "index.html"
}

func (c *ArticleControler) ShowArticleList() {
	//查询
	o := orm.NewOrm()
	qs := o.QueryTable("Article")
	var articles []models.Article
	qs.All(&articles)
	c.Data["articles"] = articles
	beego.Info(articles[0].Title)
	c.TplName = "index.html"

}

func (c *ArticleControler) ShowAddArticle() {
	///查询所有类型数据,并展示
}

/*
 1.拿数据
 2.判断数据
 3.插入是数据
 4.返回视图
*/

func (c *ArticleControler) HandleAddArticle() {
	artiName := c.GetString("articleName")
	artiContent := c.GetString("content")
	f, h, err := c.GetFile("uploadname")
	if err != nil {
		beego.Info(err.Error())
		return
	}
	defer f.Close()
	//1.判断文件格式
	ext := path.Ext(h.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		beego.Info("上传文件格式不正确")
		return
	}
	//2.文件大小
	if h.Size > 500000 {
		beego.Info("文件太大,不允许上传")
		return
	}
	//3.不能重名
	fileName := time.Now().Format("2006-01-02 15:04:05")

	err = c.SaveToFile("uploadname", "../static/img/"+ext+fileName)
	if err != nil {
		beego.Info("上传文件失败")
		return
	}
	// beego.Info(artiName, artiCount)
	// c.TplName = "register.html"

	//3.插入数据
	//获取orm对象
	o := orm.NewOrm()

	//2.创建一个插入对象
	article := models.Article{}
	///3.赋值
	article.Title = artiName
	article.Content = artiContent
	article.Img = "/static/img/" + ext + fileName

	//插入
	_, err = o.Insert(&article)
	if err != nil {
		beego.Info("插入数据失败")
		return
	}
	//返回视图
	c.Redirect("/ShowArticle", 302)

}

///删除文章
func (c *ArticleControler) HandleDelete() {
	// id, _ := c.GetInt("id")
	//1.orm对象
	// o := orm.NewOrm()

	// //要有删除对象
	// article := models.Article{}

	// //类型转换
	// id2, _ := strconv.Atoi(id)
}

//显示更新页面
func (c *ArticleControler) ShowUpdate() {
	id := c.GetString("id")
	if id == "" {
		beego.Info("连接错误")
		return
	}
	beego.Info(id)

	//2.查询数据
	//1.获取orm
	o := orm.NewOrm()
	//2.获取查询对象
	// id2, _ := strconv.Atoi(id)
	article := models.Article{}
	// id2 := strconv.Atoi(id)
	id2, _ := strconv.Atoi(id)

	article.Id2 = id2

	//3.查询
	err := o.Read(&article)
	if err != nil {
		beego.Info("查询数据为空")
		return
	}
	article.Count += 1
	o.Update(&article)

	//3.传递数据给视图
	c.Data["article"] = article
	c.TplName = "content.html"

}
