package login

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"log"
	"mxaof_blog/model/login"
	"mxaof_blog/tools"
	"net/http"
)

var store = base64Captcha.DefaultMemStore

func GetCaptchaCode(c *gin.Context) {
	//key := c.ClientIP()
	//v, err := redis.GlobalRedis.Get(context.Background(), key).Result()
	//if err != nil {
	//	log.Fatal(err)
	//}
	captcha := &login.Captcha{
		KeyLong:            6,
		ImgWidth:           240,
		ImgHeight:          80,
		OpenCaptcha:        0,
		OpenCaptchaTimeOut: 3600,
	}
	driver := base64Captcha.NewDriverDigit(captcha.ImgHeight, captcha.ImgWidth, captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, login.SysCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: captcha.KeyLong,
		OpenCaptcha:   true,
	})
}
func Verify(c *gin.Context) {
	in := login.VerifyRequest{}
	err := c.ShouldBindJSON(&in)
	if err != nil {
		log.Fatal(err)
	}
	if store.Verify(in.CaptchaId, in.CaptchaAnswer, true) {
		c.JSON(http.StatusOK, tools.HttpCode{
			Msg:  "验证成功",
			Code: 0,
			Data: struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Msg:  "验证码错误",
		Code: 100,
		Data: struct{}{},
	})
}
