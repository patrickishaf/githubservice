package db

import (
	"log"
)

func insertOrUpdateCommit(currentCommit *Commit) {
	matchingCommit := Commit{}
	result := GetDB().Where("hash = ?", currentCommit.Hash).First(&matchingCommit)

	if result.Error != nil {
		log.Println("no commit with this commit hash. writing this commit to the db")
		GetDB().Create(currentCommit)
	} else {
		// Update this entry with the new commit details
		GetDB().Save(&Commit{
			Hash:        currentCommit.Hash,
			Message:     currentCommit.Message,
			AuthorName:  currentCommit.AuthorName,
			AuthorEmail: currentCommit.AuthorEmail,
			Date:        currentCommit.Date,
		})
	}
}

func InsertCommits(commits *[]Commit) {
	for _, currentCommit := range *commits {
		insertOrUpdateCommit(&currentCommit)
	}
}
