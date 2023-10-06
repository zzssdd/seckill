package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"seckill/pkg/errmsg"
	"seckill/utils"
	"strings"
)

func JWT() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.Abort()
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": errmsg.NotLogin,
				"msg":  errmsg.GetMsg(errmsg.NotLogin),
			})
			return
		}
		checkToken := strings.Split(token, " ")
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			c.Abort()
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": errmsg.TokenError,
				"msg":  errmsg.GetMsg(errmsg.TokenError),
			})
			return
		}
		claim, err := utils.ParseToken(checkToken[1])
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": errmsg.LoginAgain,
				"msg":  errmsg.GetMsg(errmsg.LoginAgain),
			})
			return
		}
		c.Set("email", claim.Email)
		c.Set("id", claim.Id)
		c.Next(ctx)
	}
}
