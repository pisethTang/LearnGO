package main

import "github.com/gin-gonic/gin"


func main(){
	server := gin.Default()

	// trust only load balancer's IP
	// without the line below
	// we would get:
	// [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
	// Please check https://github.com/gin-gonic/gin/blob/master/docs/doc.md#dont-trust-all-proxies for details.
	// list of trusted proxies 
	var trusted_proxies = []string{"102.168.1.2"}
	server.SetTrustedProxies(trusted_proxies)
	server.GET("/test", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "OK!",
		})
	})
	
	server.Run(":8080")


}
