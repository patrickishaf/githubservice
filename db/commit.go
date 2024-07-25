package db

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}

type CommitData struct {
	Message string `json:"message"`
	Author  Author `json:"author"`
}

type CommitResponse struct {
	Hash   string     `json:"sha"`
	Commit CommitData `json:"commit"`
}

type Commit struct {
	Hash        string `gorm:"primarykey"`
	Message     string
	AuthorName  string
	AuthorEmail string
	Date        string
}

func (res *CommitResponse) ConvertToCommit() *Commit {
	return &Commit{
		Hash:        res.Hash,
		Message:     res.Commit.Message,
		AuthorName:  res.Commit.Author.Name,
		AuthorEmail: res.Commit.Author.Email,
		Date:        res.Commit.Author.Date,
	}
}

func (commit *Commit) ConvertToCommitResponse() *CommitResponse {
	return &CommitResponse{
		Hash: commit.Hash,
		Commit: CommitData{
			Message: commit.Message,
			Author: Author{
				Name:  commit.AuthorName,
				Email: commit.AuthorEmail,
				Date:  commit.Date,
			},
		},
	}
}
