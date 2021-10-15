package routers

import (
	"news/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 用户
	beego.Router("/", &controllers.LoginController{}, "get:ShowLogin;post:HanldeLogin")
	beego.Router("/register", &controllers.RegisterController{}, "get:ShowRegister;post:HandleRegister")

	// 文章
	beego.Router("/showArticle", &controllers.ArticleController{}, "get:ShowArticleList")
	beego.Router("/addArticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:HandleAddArticle")
	beego.Router("/showArticleDetail", &controllers.ArticleController{}, "get:ShowArticleDetail")
	beego.Router("/deleteArticleDetail", &controllers.ArticleController{}, "get:DeleteArticleDetail")
	beego.Router("/updateArticle", &controllers.ArticleController{}, "get:ShowUpdate;post:HandleUpdate")
}
