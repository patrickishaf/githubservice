package main

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickishaf/githubservice/pkg/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}

	router := gin.Default()

	router.GET("/repo", handlers.GetRepoByName)

	router.Run("localhost:8080")
}
