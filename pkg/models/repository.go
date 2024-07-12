package models

type Repository struct {
	Name            string `json:"name"`
	Description     string
	URL             string
	Language        string
	ForksCount      int32
	OpenIssuesCount int64
	WatchersCount   int64
	CreatedAt       string
	UpdatedAt       string
}
