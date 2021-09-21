package models

// TikTokVideo structure of each tiktok video to be shared
type TikTokVideo struct {
	MediaID int64  `json:"mediaId"`
	Owner   string `json:"owner"`
}
