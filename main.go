package main

import(
	"log"
	//"os"
	"github.com/gin-gonic/gin"
	routes "github.com/santoshsoren/gin-gonic/routes"
	models "github.com/santoshsoren/gin-gonic/config"
)

func main(){
	config.Connect()
	router := gin.Default()
	routes.Routes(router)
	log.Fatal(router.Run(":8081"))
}