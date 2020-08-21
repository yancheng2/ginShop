package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()
	r.GET("/test",func(context *gin.Context){
		context.JSON(200,gin.H{"data":"success"})
	})
	r.Run()
}
