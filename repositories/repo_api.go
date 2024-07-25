package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/patrickishaf/githubservice/common"
	"github.com/patrickishaf/githubservice/db"
	"log"
)

func getRepoByName(orgName string, repoName string) (db.Repository, error) {
	url := fmt.Sprintf("/repos/%s/%s", orgName, repoName)
	data, err := common.HttpGet(url)

	var repo = db.Repository{}

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

func getRepoByLanguage(language string) ([]db.Repository, error) {
	url := fmt.Sprintf("/search/repositories?q=language:%s", language)
	data, err := common.HttpGet(url)

	if err != nil {
		log.Println("failed to fetch commits for repo")
		return nil, err
	}

	var repos db.RepositoryList

	if err := json.Unmarshal(data, &repos); err != nil {
		log.Println("failed to unmarshal json")
		return nil, err
	}

	return repos.Items, nil
}

func getRepositoriesByStarCount(starCount, limit int) ([]db.Repository, error) {
	url := fmt.Sprintf("/search/repositories?q=stars:%d&sort=stars&per_page=%d", starCount, limit)
	data, err := common.HttpGet(url)

	if err != nil {
		log.Println("failed to fetch commits for repo")
		return nil, err
	}

	var repos db.RepositoryList

	if err := json.Unmarshal(data, &repos); err != nil {
		log.Println("failed to unmarshal json")
		return nil, err
	}

	return repos.Items, nil
}
