package models

type UserMeta struct {
	ID        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	AvatarURL string `json:"avatarUrl,omitempty"`
}

type User struct {
	URL  string   `json:"url,omitempty"`
	Meta UserMeta `json:"meta,omitempty"`
}
