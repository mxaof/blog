package login

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"log"
	"mxaof_blog/model/login"
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
