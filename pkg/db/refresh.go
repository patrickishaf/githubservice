package db

import (
	"github.com/patrickishaf/githubservice/pkg/services"
	"log"
	"time"
)

func RefreshRepository(orgName, repoName string) {
	ticker := time.NewTicker(time.Hour)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				getUpdatedRepositoryRecords(orgName, repoName, quit)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func getUpdatedRepositoryRecords(orgName, repoName string, c chan bool) {
	responseData, err := services.GetRepoByName(orgName, repoName)

	if err != nil {
		log.Print("failed to get repository information")
	} else {
		InsertOrUpdateRepo(&responseData)
	}
}
