package user

import (
	"github.com/yongchengchen/sysgo/services/container"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID   uint64
	Name string
}

type UserIns struct {
}

func Factory() interface{} {
	u := UserIns{}
	return &u
}

func (u *UserIns) Boot() {
	r := container.Get("router").(*gin.Engine)
	users := []User{{ID: 123, Name: "张三"}, {ID: 456, Name: "李四"}}
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Blog":   "www.flysnow.org",
			"wechat": "flysnow_org",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"Blog":   "www.flysnow.org",
		// 	"wechat": "flysnow_org",
		// })
		c.JSON(200, users)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{"id": c.Param("id"), "user": users[0]})
	})
}
