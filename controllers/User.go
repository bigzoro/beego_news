package controllers

import (
	"news/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	beego.Controller
}

func (rc *RegisterController) ShowRegister() {
	rc.TplName = "register.html"
}

func (rc *RegisterController) HandleRegister() {
	// 获取前端传递的数据
	name := rc.GetString("userName")
	password := rc.GetString("password")
	beego.Info(name, password)

	// 数据处理
	if name == "" || password == "" {
		beego.Info("用户名或密码不能为空")
		rc.TplName = "register.html"
		return
	}

	// 插入数据
	//获取ORM对象
	o := orm.NewOrm()
	//获取插入对象
	var user models.User
	//给插入对象赋值
	user.UserName = name
	user.Password = password
	//插入
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入数据失败")
		return
	}

	//4.返回页面
	// rc.Ctx.WriteString("注册成功")
	rc.Redirect("/", 302)
	//this.TplName = "login.html"
}

type LoginController struct {
	beego.Controller
}

func (lc *LoginController) ShowLogin() {
	lc.TplName = "login.html"
}

func (lc *LoginController) HanldeLogin() {
	// 获取前端数据
	name := lc.GetString("userName")
	password := lc.GetString("password")
	beego.Info(name, password)

	// 数据处理
	if name == "" || password == "" {
		beego.Info("用户名或密码不能为空")
		lc.TplName = "login.html"
		return
	}

	// 查找数据
	// 获取orm对象
	o := orm.NewOrm()
	// 获取查询对象
	user := models.User{}
	// 查询
	user.UserName = name
	err := o.Read(&user, "UserName")
	if err != nil {
		beego.Info("用户名失败")
		lc.TplName = "login.html"
		return
	}

	// 判断是否正确
	if user.Password != password {
		beego.Info("密码失败")
		lc.TplName = "login.html"
		return
	}

	// 返回信息
	// lc.Ctx.WriteString("登录成功")
	lc.Redirect("/showArticle", 302)
}
