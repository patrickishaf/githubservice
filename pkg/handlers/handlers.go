package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRepoByName(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "this is the base route")
}

func GetCommitsByRepoName(c *gin.Context) {}

func SearchReposByLanguage(c *gin.Context) {}

func GetTopReposByStarCount(c *gin.Context) {}
