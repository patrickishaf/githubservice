package main

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickishaf/githubservice/api"
	"github.com/patrickishaf/githubservice/config"
	"github.com/patrickishaf/githubservice/db"
)

func main() {
	config.LoadEnv()
	db.InitializeDb()

	router := gin.Default()

	router.GET("/repo/:org_name/:repo_name", api.GetRepoByName)
	router.GET("/commits/:org_name/:repo_name", api.GetCommitsByRepoName)
	router.GET("/repo/lang/:lang", api.SearchReposByLanguage)
	router.GET("/repo/stars", api.GetTopReposByStarCount)

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
