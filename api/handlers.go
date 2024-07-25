package api

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickishaf/githubservice/commits"
	"github.com/patrickishaf/githubservice/events"
	"github.com/patrickishaf/githubservice/repositories"
	"net/http"
)

var eventEmitter events.EventEmitter = events.EventEmitter{}

func GetRepoByName(c *gin.Context) {
	orgName := c.Param("org_name")
	repoName := c.Param("repo_name")

	responseData, err := repositories.GetRepoByName(orgName, repoName)

	eventEmitter.AddEvent(&events.Event{
		Name: "monitor_repo",
		Args: []string{orgName, repoName},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to get repository information")
	} else {
		c.IndentedJSON(http.StatusOK, responseData)
	}
}

func GetCommitsByRepoName(c *gin.Context) {
	orgName := c.Param("org_name")
	repoName := c.Param("repo_name")

	responseData, err := commits.GetCommitsByRepoName(orgName, repoName)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "failed to get repository information")
	} else {
		c.IndentedJSON(http.StatusOK, responseData)
	}
}

func SearchReposByLanguage(c *gin.Context) {
	language := c.Param("lang")
	responseData, err := repositories.SearchReposByLanguage(language)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to get repository information")
	} else {
		c.IndentedJSON(http.StatusOK, responseData)
	}
}

func GetTopReposByStarCount(c *gin.Context) {
	starCount := c.Query("count")
	limit := c.Query("limit")

	responseData, err := repositories.GetTopReposByStarCount(starCount, limit)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "failed to get repository information")
	} else {
		c.IndentedJSON(http.StatusOK, responseData)
	}
}
