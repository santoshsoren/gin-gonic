package main

import(
	"log"
	//"os"
	"github.com/gin-gonic/gin"
	routes "github.com/santoshsoren/go_db/routes"
	models "github.com/santoshsoren/go_db/config"
)

func main(){
	config.Connect()
	router := gin.Default()
	routes.Routes(router)
	log.Fatal(router.Run(":8081"))
}