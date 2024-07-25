package repositories

import (
	"github.com/patrickishaf/githubservice/config"
	"github.com/patrickishaf/githubservice/db"
	"log"
	"time"
)

type RepositoryMonitor struct {
	ticker *time.Ticker
}

func (r *RepositoryMonitor) RefreshPeriodically(orgName, repoName string) {
	r.refreshRepository(orgName, repoName)
}

func (r *RepositoryMonitor) refreshRepository(orgName, repoName string) {
	duration := r.loadDurationFromEnv()
	r.ticker = time.NewTicker(duration * time.Second)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-r.ticker.C:
				r.getUpdatedRepositoryRecords(orgName, repoName, quit)
			case <-quit:
				r.ticker.Stop()
				return
			}
		}
	}()
}

func (r *RepositoryMonitor) getUpdatedRepositoryRecords(orgName, repoName string, c chan bool) {
	responseData, err := getRepoByName(orgName, repoName)

	if err != nil {
		log.Print("failed to get repository information")
	} else {
		db.InsertOrUpdateRepo(&responseData)
	}
}

func (r *RepositoryMonitor) loadDurationFromEnv() time.Duration {
	durationString, isPresent := config.GetEnvVariable("INTERVAL_MINUTES")

	if !isPresent {
		return time.Duration(10)
	}

	duration, err := time.ParseDuration(durationString)
	if err != nil {
		return time.Duration(10)
	}

	return duration
}
