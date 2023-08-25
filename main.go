package main

import (
	"log"
	"one_test_case/DBConfig"
	ProductRoutes "one_test_case/Routes"

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

	ProductRoutes.Routes(router)

	router.Run(":9000")
}
