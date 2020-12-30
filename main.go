package main

import(
	"log"
	//"os"
	"github.com/gin-gonic/gin"
	config "github.com/santoshsoren/gin-gonic/configs"
	routes "github.com/santoshsoren/gin-gonic/routes"
)

func main(){
	config.Connect()
	router := gin.Default()
	routes.Routes(router)
	log.Fatal(router.Run(":8081"))
}