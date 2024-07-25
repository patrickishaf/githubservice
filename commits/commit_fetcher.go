package commits

import (
	"github.com/patrickishaf/githubservice/db"
)

func GetCommitsByRepoName(orgName, repoName string) (*[]db.Commit, error) {
	responseData, err := getCommitsInRepo(orgName, repoName)

	if err != nil {
		return nil, err
	}

	commits := []db.Commit{}

	for _, v := range responseData {
		commit := v.ConvertToCommit()
		commits = append(commits, *commit)
	}

	db.InsertCommits(&commits)
	return &commits, nil
}
