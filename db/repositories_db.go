package db

import (
	"log"
)

func InsertOrUpdateRepo(repo *Repository) {
	matchingRepo := Repository{}
	result := GetDB().Where("repo_id = ?", repo.RepoId).First(&matchingRepo)

	if result.Error != nil {
		log.Println("no repository with matching RepoId yet. Inserting new repository...")
		GetDB().Create(repo)
	} else {
		log.Println("there is already a repo with the same RepoId. Updating record")
		GetDB().Save(&Repository{
			Name:            repo.Name,
			FullName:        repo.FullName,
			Description:     repo.Description,
			URL:             repo.URL,
			Language:        repo.Language,
			ForksCount:      repo.ForksCount,
			OpenIssuesCount: repo.OpenIssuesCount,
			WatchersCount:   repo.OpenIssuesCount,
			CreatedAt:       repo.CreatedAt,
			UpdatedAt:       repo.UpdatedAt,
			RepoId:          repo.RepoId,
		})
	}
}
