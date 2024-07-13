package models

type Repository struct {
	Name            string `json:"name"`
	FullName        string `json:"full_name"`
	Description     string `json:"description"`
	URL             string `json:"html_url"`
	Language        string `json:"language"`
	ForksCount      int32  `json:"forks_count"`
	OpenIssuesCount int64  `json:"open_issues_count"`
	WatchersCount   int64  `json:"watchers_count"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	RepoId          int32  `gorm:"primarykey" json:"id"`
}

type RepositoryList struct {
	Items []Repository `json:"items"`
}
