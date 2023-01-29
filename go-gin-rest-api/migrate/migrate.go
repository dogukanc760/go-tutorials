package main

import (
	"fmt"
	"log"
	"rest-api/initializers"
	"rest-api/models"
	// import models and intiliaziers
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}) //, &models.Post{})
	fmt.Println("👍 Migration complete")
}
