package models

import "gorm.io/gorm"

type Repository struct {
	gorm.Model
	Name            string `json:"name"`
	Description     string `json:"description"`
	URL             string `json:"html_url"`
	Language        string `json:"language"`
	ForksCount      int32  `json:"forks_count"`
	OpenIssuesCount int64  `json:"open_issues_count"`
	WatchersCount   int64  `json:"watchers_count"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
