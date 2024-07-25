package events

import (
	"github.com/patrickishaf/githubservice/commits"
	"github.com/patrickishaf/githubservice/repositories"
)

type Event struct {
	Name string
	Args []string
}

type EventEmitter struct {
	repoMonitor   repositories.RepositoryMonitor
	commitMonitor commits.CommitMonitor
}

func (ep *EventEmitter) AddEvent(e *Event) {
	switch e.Name {
	case "monitor_repo":
		orgName := e.Args[0]
		repoName := e.Args[1]
		ep.repoMonitor.RefreshPeriodically(orgName, repoName)
	case "monitor_commits":
		orgName := e.Args[0]
		repoName := e.Args[1]
		ep.commitMonitor.RefreshPeriodically(orgName, repoName)
	}
}
