package db

import (
	"log"

	"github.com/patrickishaf/githubservice/pkg/models"
)

func InsertOrUpdateRepo(repo *models.Repository) {
	matchingRepo := models.Repository{}
	result := db.Where("repo_id = ?", repo.RepoId).First(&matchingRepo)

	if result.Error != nil {
		log.Println("no repository with matching RepoId yet. Inserting new repository...")
		db.Create(repo)
	} else {
		log.Println("there is already a repo with the same RepoId. Updating record")
		db.Save(&models.Repository{
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
