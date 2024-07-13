package main

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickishaf/githubservice/pkg/db"
	"github.com/patrickishaf/githubservice/pkg/handlers"
)

func main() {
	db.InitializeDb()

	router := gin.Default()

	router.GET("/repo/:org_name/:repo_name", handlers.GetRepoByName)
	router.GET("/commits/:org_name/:repo_name", handlers.GetCommitsByRepoName)
	router.GET("/repo/lang/:lang", handlers.SearchReposByLanguage)
	router.GET("/repo/stars", handlers.GetTopReposByStarCount)

	router.Run("localhost:8080")
}
