package main

import (
	"log"
	"one_test_case/DBConfig"
	routes "one_test_case/Routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBConfig.DB, err = gorm.Open("mysql", DBConfig.DbURL(DBConfig.BuildDBConfig()))

	if err != nil {
		panic(err)
	}

	defer DBConfig.DB.Close()

	router := gin.Default()

	routes.ProductRoutes(router)
	routes.CampaignRoutes(router)
	routes.OrderRoutes(router)

	router.Run(os.Getenv("RUN_PORT"))
}
