package commits

import (
	"github.com/patrickishaf/githubservice/config"
	"github.com/patrickishaf/githubservice/db"
	"log"
	"time"
)

type CommitMonitor struct {
	ticker *time.Ticker
}

func (cm *CommitMonitor) RefreshPeriodically(orgName, repoName string) {
	cm.refreshCommits(orgName, repoName)
}

func (cm *CommitMonitor) refreshCommits(orgName, repoName string) {
	duration := cm.loadDurationFromEnv()
	cm.ticker = time.NewTicker(duration * time.Second)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-cm.ticker.C:
				cm.getUpdatedCommitRecords(orgName, repoName, quit)
			case <-quit:
				cm.ticker.Stop()
				return
			}
		}
	}()
}

func (cm *CommitMonitor) getUpdatedCommitRecords(orgName, repoName string, c chan bool) {
	commitResponses, err := getCommitsInRepo(orgName, repoName)
	commitData := []db.Commit{}

	if err != nil {
		log.Print("failed to get commit information")
	} else {
		for _, c := range commitResponses {
			commitData = append(commitData, *c.ConvertToCommit())
		}
		db.InsertCommits(&commitData)
	}
}

func (cm *CommitMonitor) loadDurationFromEnv() time.Duration {
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
