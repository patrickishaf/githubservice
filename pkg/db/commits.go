package db

import (
	"github.com/patrickishaf/githubservice/pkg/models"
	"log"
)

func insertOrUpdateCommit(currentCommit *models.Commit) {
	matchingCommit := models.Commit{}
	result := db.Where("hash = ?", currentCommit.Hash).First(&matchingCommit)

	if result.Error != nil {
		log.Println("no commit with this commit hash. writing this commit to the db")
		db.Create(currentCommit)
	} else {
		// Update this entry with the new commit details
		db.Save(&models.Commit{
			Hash:        currentCommit.Hash,
			Message:     currentCommit.Message,
			AuthorName:  currentCommit.AuthorName,
			AuthorEmail: currentCommit.AuthorEmail,
			Date:        currentCommit.Date,
		})
	}
}

func InsertCommits(commits *[]models.Commit) {
	for _, currentCommit := range *commits {
		insertOrUpdateCommit(&currentCommit)
	}
}
