package models

import "time"

type Banner struct {
	ID          int        `json:"-"`
	Title       string     `json:"title"`
	SubTitle    string     `json:"sub_title"`
	Description string     `json:"description"`
	ButtonText  string     `json:"button_text"`
	ButtonUrl   string     `json:"button_url"`
	Image       string     `json:"image"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   *time.Time `json:"-"`
	DeletedAt   *time.Time `json:"-"`
}
