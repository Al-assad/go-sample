package main

// @desc go http rest server 示例（Gin）
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type user struct {
	Id   string `json:"id"`
	Name string `json:"name" binding:"required"`
	City string `json:"city"`
}

func main() {
	router := gin.Default()
	router.GET("/ping", ping)
	router.GET("/user/:uid", getUser)
	router.POST("/user/:uid", saveUser)
	router.Run(":8080")
}

func ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func getUser(ctx *gin.Context) {
	uid := ctx.Param("uid")
	if uid == "233" {
		user := user{uid, "assad", "guangzhou"}
		ctx.JSON(200, user)
	} else {
		ctx.JSON(200, gin.H{"message": "user is empty"})
	}

}

func saveUser(ctx *gin.Context) {
	// 绑定 request body json 到 struct
	var user user
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Id = ctx.Param("uid")
	ctx.JSON(200, user)
}
