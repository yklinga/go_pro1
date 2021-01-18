package api

import (
	"github.com/gin-gonic/gin"
	"go_pro1/model"
	"go_pro1/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register (ctx * gin.Context) {
	DB := utils.GetDB()
	// 获取参数 username telephone password
	username := ctx.PostForm("username")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位" })
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	//如果没有传 生成一个随机10位用户名
	if len(username) == 0 {
		username = utils.RandomUserName(10)
	}
	log.Println(username, telephone, password)

	// 判断手机号是否存在

	if utils.IsTelephoneExist(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}

	// 创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "用户已经存在"})
		return
	}

	newUser := model.User{
		Username: username,
		Password: string(hasedPassword),
		Telephone: telephone,
	}
	DB.Create(&newUser)
	//返回结果
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg": "注册成功",
	})
}

func Login (ctx * gin.Context) {
	db := utils.GetDB()
	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位" })
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	// 手机号是否存在
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	// 发放token
	token, err := utils.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Panicf("token generate error: %v", err)
		return
	}
	// 返回结果
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg": "登录成功",
	})
}