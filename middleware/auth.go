package middleware

import (
	"github.com/gin-gonic/gin"
	"contract/response"
	"crypto/md5"
	"time"
	"strconv"
	"encoding/hex"
	"contract/config"
)

func AuthRequire(c *gin.Context) {
	return
	accessToken := c.Request.Header.Get("Authorization")
	if accessToken == "" {
		response.Error("缺少header头Authorization", c)
		c.Abort()
		return
	}
	h := md5.New()
	h.Write([]byte(config.Password + strconv.FormatInt(time.Now().Unix(), 10)))
	m := hex.EncodeToString(h.Sum(nil))
	if accessToken != string(m) {
		response.Error("验证不通过", c)
		c.Abort()
		return
	}
	return
}
