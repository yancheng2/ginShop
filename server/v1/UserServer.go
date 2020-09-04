package v1

import (
	"ginShop/models"
	"ginShop/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// 登录注册
func Login(c *gin.Context) {
	// 接收post参数值
	mobile := c.PostForm("mobile") //手机号
	vCode := c.PostForm("code")    //验证码

	//使用beego 的验证参数s
	validations := validation.Validation{}
	validations.Required(mobile, "Mobile").Message("手机号有误")
	validations.Length(vCode, 4, "Code").Message("验证码格式错误")
	if isOk := checkValidation(&validations, c); isOk == false {
		return
	}

	//根据手机号查询用户信息
	userInfo, err := models.FindUserByMobile(mobile)
	if gorm.IsRecordNotFoundError(err) { //如果没有错误只是无数据则注册
		user, err := models.CreateUser(mobile)
		if err != nil {
			util.ResponseWithJson(9002, "", "", c)
			return
		}
		userInfo = user
	} else {
		if err != nil {
			util.ResponseWithJson(9002, "", "", c)
			return
		}
	}
	token, err := util.GeteraterToken(userInfo.Id, userInfo.Mobile)
	if err != nil {
		util.ResponseWithJson(9002, "", "", c)
		return
	}

	util.ResponseWithJson(200, gin.H{
		"User":  userInfo,
		"Token": token,
	}, "登录成功", c)
}

//检查错误
func checkValidation(vali *validation.Validation, c *gin.Context) bool {
	if vali.HasErrors() {
		var errs []string
		for _, val := range vali.Errors {
			errs = append(errs, val.Message)
		}
		util.ResponseWithJson(10000, errs, "", c)
		return false
	}
	return true
}
