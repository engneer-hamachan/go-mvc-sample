package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsLogin() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		customerId := session.Get("CustomerId")

		if customerId == nil {
			c.HTML(401, "login.html", gin.H{})
			c.Abort()
		} else {
			c.Set("CustomerId", customerId)
			c.Next()
		}
	}
}
