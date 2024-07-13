package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/patrickishaf/githubservice/pkg/db"
	"github.com/patrickishaf/githubservice/pkg/models"
	"github.com/patrickishaf/githubservice/pkg/services"
)

func GetRepoByName(c *gin.Context) {
	orgName := c.Param("org_name")
	repoName := c.Param("repo_name")

	responseData, err := services.GetRepoByName(orgName, repoName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to get repository information")
		return
	}

	db.InsertOrUpdateRepo(&responseData)
	db.RefreshRepository(orgName, repoName)
	c.IndentedJSON(http.StatusOK, responseData)
}

func GetCommitsByRepoName(c *gin.Context) {
	orgName := c.Param("org_name")
	repoName := c.Param("repo_name")

	responseData, err := services.GetCommitsInRepo(orgName, repoName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to get repository information")
		return
	}

	commits := []models.Commit{}

	for _, v := range responseData {
		commit := v.ConvertToCommit()
		commits = append(commits, *commit)
	}

	db.InsertCommits(&commits)
	c.IndentedJSON(http.StatusOK, responseData)
}

func SearchReposByLanguage(c *gin.Context) {
	language := c.Param("lang")
	responseData, err := services.GetRepoByLanguage(language)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to get repository information")
		return
	}

	c.IndentedJSON(http.StatusOK, responseData)
}

func GetTopReposByStarCount(c *gin.Context) {
	starCount := c.Query("count")
	limit := c.Query("limit")

	log.Println("the value of limit is ", limit)

	intStarCount, err := strconv.Atoi(starCount)
	if err != nil {
		c.JSON(http.StatusBadRequest, "you entered an invalid star count")
		return
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, "you entered an invalid limit")
		return
	}

	responseData, err := services.GetRepositoriesByStarCount(intStarCount, intLimit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to get repository information")
		return
	}

	c.IndentedJSON(http.StatusOK, responseData)
}
