package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"golang-sample-xss/model"
	"net/http"
	"os"
)

/*
go get github.com/gin-gonic/gin
go get github.com/microcosm-cc/bluemonday

set API_HOST=localhost
set API_PORT=8888
*/

func NewUser(user model.User) *model.User {
	p := bluemonday.UGCPolicy()
	user.Username = p.Sanitize(user.Username)
	user.FirstName = p.Sanitize(user.FirstName)
	user.LastName = p.Sanitize(user.LastName)
	return &user
}

func main() {
	users := make([]model.User, 0)
	apiHost := os.Getenv("API_HOST")
	apiPort := os.Getenv("API_PORT")
	listenAddress := fmt.Sprintf("%s:%s", apiHost, apiPort)

	routerEngine := gin.Default()
	RouterGroup := routerEngine.Group("/api")
	RouterGroup.POST("/user", func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, x")
		ctx.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
		var newUser model.User
		if err := ctx.ShouldBindJSON(&newUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		users = append(users, newUser)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "SUCCESS",
			"data":    newUser,
		})
	})

	RouterGroup.GET("/user", func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, x")
		ctx.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "SUCCESS",
			"data":    users,
		})
	})

	err := routerEngine.Run(listenAddress)
	if err != nil {
		panic(err)
	}
}
