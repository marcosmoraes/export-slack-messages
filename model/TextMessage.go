package model

type TextMessage struct {
	Text        string       `json:"text"`
	UserProfile *UserProfile `json:"user_profile"`
}
