package models

import "time"

type ProductType struct {
	ID        int        `json:"-"`
	Slug      string     `json:"slug"`
	Name      string     `json:"name"`
	Image     string     `json:"image"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
