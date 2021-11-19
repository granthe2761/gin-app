package main

import(
	//"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/granthe2761/gin-app/handlers"
)

const(
	Port="8001"
)
/*
func init(){
	db,Connect()
}
*/



func main(){
	//Configure
	router := gin.Default()
	
	//Set html render options
	
	//Tell gin to use our html gender

	//Middlewares
	
	//Statics

	//Routes
	router.GET("/",func(c *gin.Context){
		c.String(http.StatusOK,"hello world!")
	})
	router.POST("/bpmonitor/uploadImage/",handlers.UploadImageHandler)

	//listen
	router.Run(":8001")		
}
