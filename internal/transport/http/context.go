package http

import "github.com/gin-gonic/gin"

type Context struct {
	ctx  *gin.Context
	err  error
	data any
}

func NewContext(ctx *gin.Context) *Context {
	return &Context{
		ctx: ctx,
	}
}
