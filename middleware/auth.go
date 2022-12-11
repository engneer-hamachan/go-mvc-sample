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
			c.Status(401)
			c.Abort()
		} else {
			c.Set("CustomerId", customerId)
			c.Next()
		}
	}
}
