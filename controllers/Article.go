package controllers

import (
	"news/models"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ArticleController struct {
	beego.Controller
}

func (ac *ArticleController) ShowArticleList() {
	// 获取orm对象
	o := orm.NewOrm()
	// 查询
	qs := o.QueryTable("article")
	var articles []models.Article
	qs.All(&articles)

	// 向前端传递数据
	ac.Data["articles"] = articles
	ac.TplName = "index.html"
}

func (ac *ArticleController) ShowAddArticle() {
	ac.TplName = "add.html"
}

func (ac *ArticleController) HandleAddArticle() {

	// 获取数据
	articleName := ac.GetString("articleName")
	articleContent := ac.GetString("content")

	f, h, err := ac.GetFile("uploadname")
	if err != nil {
		beego.Info("上传文件失败")
		return
	}
	defer f.Close()
	// 判断文件格式
	ext := path.Ext(h.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		beego.Info("上传文件格式不正确")
		return
	}
	// 判断文件大小
	if h.Size > 5000000 {
		beego.Info("文件太大，不允许上传")
		return
	}
	// 不能重名
	fileName := time.Now().Format("2006-01-02-15:04:05")
	ac.SaveToFile("uploadname", "/static/img"+h.Filename)

	beego.Info(articleName, articleContent, fileName+ext)

	// 插入数据
	// 获取orm对象
	o := orm.NewOrm()
	// 创建一个插入对象
	article := models.Article{}
	// 赋值
	article.Title = articleName
	article.Content = articleContent
	article.Img = "static/img/" + fileName + ext
	// 插入
	_, err = o.Insert(&article)
	if err != nil {
		beego.Info("插入数据失败", err)
		return
	}

	beego.Info("insert article: ", article)
	ac.Redirect("showArticle", 302)
}

func (ac *ArticleController) ShowArticleDetail() {
	// 获取数据
	id, err := ac.GetInt("articleId")
	//数据校验
	if err != nil {
		beego.Info("传递的链接错误", err)
		return
	}
	// beego.Info("id: ", id)

	//操作数据
	o := orm.NewOrm()
	var article models.Article
	article.Id = id

	o.Read(&article)

	//修改阅读量
	article.Count += 1
	o.Update(&article)

	beego.Info("article：：", article)
	ac.Data["article"] = article
	ac.TplName = "content.html"

}

func (ac *ArticleController) DeleteArticleDetail() {
	// 获取数据
	id, err := ac.GetInt("articleId")
	//数据校验
	if err != nil {
		beego.Info("传递的链接错误", err)
		return
	}

	//操作数据
	o := orm.NewOrm()
	var article models.Article
	article.Id = id

	o.Delete(&article)

	ac.Redirect("showArticle", 302)

}

func (ac *ArticleController) ShowUpdate() {
	// 获取数据
	id := ac.GetString("articleId")
	if id == "" {
		beego.Info("查询失败")
		return
	}

	// 查询
	o := orm.NewOrm()
	article := models.Article{}
	// 类型转换
	id2, _ := strconv.Atoi(id)
	article.Id = id2

	err := o.Read(&article)
	if err != nil {
		beego.Info("查询失败", err)
		return
	}

	// 把数据传递给视图
	ac.Data["article"] = article
	ac.TplName = "update.html"
}

//封装上传文件函数
func UploadFile(this *beego.Controller, filePath string) string {
	//处理文件上传
	file, head, err := this.GetFile(filePath)
	if head.Filename == "" {
		return "NoImg"
	}

	if err != nil {
		this.Data["errmsg"] = "文件上传失败"
		this.TplName = "add.html"
		return ""
	}
	defer file.Close()

	//1.文件大小
	if head.Size > 5000000 {
		this.Data["errmsg"] = "文件太大，请重新上传"
		this.TplName = "add.html"
		return ""
	}

	//2.文件格式
	//a.jpg
	ext := path.Ext(head.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		this.Data["errmsg"] = "文件格式错误。请重新上传"
		this.TplName = "add.html"
		return ""
	}

	//3.防止重名
	fileName := time.Now().Format("2006-01-02-15:04:05") + ext
	//存储
	this.SaveToFile(filePath, "./static/img/"+fileName)
	return "/static/img/" + fileName
}

func (ac *ArticleController) HandleUpdate() {
	//获取数据
	id, err := ac.GetInt("articleId")
	articleName := ac.GetString("articleName")
	content := ac.GetString("content")
	filePath := UploadFile(&ac.Controller, "uploadname")
	//数据校验
	if err != nil || articleName == "" || content == "" || filePath == "" {
		beego.Info("请求错误")
		return
	}
	//数据处理
	o := orm.NewOrm()
	var article models.Article
	article.Id = id
	err = o.Read(&article)
	if err != nil {
		beego.Info("更新的文章不存在")
		return
	}
	article.Title = articleName
	article.Content = content
	if filePath != "NoImg" {
		article.Img = filePath
	}
	o.Update(&article)

	//返回视图
	ac.Redirect("/showArticleList", 302)
}
