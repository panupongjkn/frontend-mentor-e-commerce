package models

import "time"

type HomeProductSection struct {
	ID          int        `json:"-"`
	ProductID   int        `json:"-"`
	DisplayType int        `json:"display_type"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   *time.Time `json:"-"`
	DeletedAt   *time.Time `json:"-"`

	Product HomeProductSectionProduct `json:"product"`
}

type HomeProductSectionProduct struct {
	Slug     string `json:"slug"`
	SubTitle string `json:"sub_title"`
	Name     string `json:"name"`
	Image    string `json:"image"`
}
