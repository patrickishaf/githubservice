package repositories

import (
	"github.com/patrickishaf/githubservice/db"
	"strconv"
)

func GetRepoByName(orgName, repoName string) (*db.Repository, error) {
	repo, err := getRepoByName(orgName, repoName)

	if err != nil {
		return nil, err
	}

	db.InsertOrUpdateRepo(&repo)
	return &repo, nil
}

func SearchReposByLanguage(language string) (*[]db.Repository, error) {
	repositories, err := getRepoByLanguage(language)

	if err != nil {
		return nil, err
	}

	return &repositories, nil
}

func GetTopReposByStarCount(starCount, limit string) (*[]db.Repository, error) {
	intStarCount, err := strconv.Atoi(starCount)
	if err != nil {
		return nil, err
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		return nil, err
	}

	repositories, err := getRepositoriesByStarCount(intStarCount, intLimit)

	if err != nil {
		return nil, err
	}

	return &repositories, nil
}
