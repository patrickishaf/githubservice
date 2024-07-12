package services

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/patrickishaf/githubservice/pkg/models"
)

func GetRepoByName(orgName string, repoName string) (models.Repository, error) {
	url := fmt.Sprintf("/repos/%s/%s", orgName, repoName)
	data, err := get(url)

	var repo = models.Repository{}

	if err != nil {
		log.Println("error in github api service")
		return repo, err
	}

	if err := json.Unmarshal(data, &repo); err != nil {
		log.Println("failed to unmarshal json response")
		return repo, err
	}

	return repo, nil
}

func GetCommitsInRepo(orgName string, repoName string) ([]models.Commit, error) {
	url := fmt.Sprintf("/repos/%s/%s/commits", orgName, repoName)
	data, err := get(url)

	if err != nil {
		log.Println("failed to fetch commits for repo")
		return []models.Commit{}, err
	}

	var commits = []models.Commit{}

	if err := json.Unmarshal(data, &commits); err != nil {
		log.Println("failed to unmarshal json")
		return nil, err
	}

	return commits, nil
}

func GetRepoByLanguage(language string) {}

func GetRepositoriesByStarCount(limit int32) {}
