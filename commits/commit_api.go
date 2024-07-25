package commits

import (
	"encoding/json"
	"fmt"
	"github.com/patrickishaf/githubservice/common"
	"github.com/patrickishaf/githubservice/db"
	"log"
)

func getCommitsInRepo(orgName string, repoName string) ([]db.CommitResponse, error) {
	url := fmt.Sprintf("/repos/%s/%s/commits", orgName, repoName)
	data, err := common.HttpGet(url)

	if err != nil {
		log.Println("failed to fetch commits for repo")
		return []db.CommitResponse{}, err
	}

	var commits = []db.CommitResponse{}

	if err := json.Unmarshal(data, &commits); err != nil {
		log.Println("failed to unmarshal json")
		return nil, err
	}

	return commits, nil
}
