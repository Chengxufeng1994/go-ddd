package middleware

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func AuthorizeMiddleware(e *casbin.Enforcer) gin.HandlerFunc {

	return func(c *gin.Context) {
		// Load policy from Database
		err := e.LoadPolicy()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "failed to load policy from DB"})
			return
		}

		sub := "system-admin" // the role that wants to access a resource.
		obj := "accounts"     // the resource that is going to be accessed.
		act := "write"        // the operation that the user performs on the resource.

		ok, err := e.Enforce(sub, obj, act)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "failed to authorize user"})
			return
		}

		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"msg": "unauthorized"})
			return
		}

		c.Next()
	}
}
