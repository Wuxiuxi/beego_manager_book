package routers

import (
	"article/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegControler{}, "get:ShowReg;post:HandReg")
	beego.Router("/", &controllers.LoginControler{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/ShowArticle", &controllers.ArticleControler{}, "get:ShowArticleList")
	beego.Router("/AddArticle", &controllers.ArticleControler{}, "get:ShowAddArticle;post:HandleAddArticle")
	beego.Router("/HandleAddArticle", &controllers.ArticleControler{}, "get:HandleAddArticle")
	beego.Router("/HandleAddArticle", &controllers.ArticleControler{}, "get:ShowArticleList")

}
