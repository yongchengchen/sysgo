package router

import "github.com/gin-gonic/gin"

var r *gin.Engine

func Factory() interface{} {
	r = gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	return r
}
