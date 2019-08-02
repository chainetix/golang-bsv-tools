package models

import "gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models"

type ProjectStats struct {
	models.Proto
	Project string `json:"project"`
	Title string `json:"title"`
	TxCount int `json:"txCount"`
	Currencies int `json:"currencies"`
	Users int `json:"users"`
	Streams int `json:"streams"`
	Exchanges int `json:"exchanges"`
	Wallets int `json:"exchanges"`
	BytesConsumed int64 `json:"bytesConsumed"`
}

func NewProjectStats(project *Project) *ProjectStats {
	return &ProjectStats{
		Project: project.UID,
		Title: project.Title,
	}
}
