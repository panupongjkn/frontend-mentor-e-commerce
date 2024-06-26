package models

import "time"

type HomeProductTypeSection struct {
	ID            int        `json:"-"`
	ProductTypeID int        `json:"-"`
	CreatedAt     time.Time  `json:"-"`
	UpdatedAt     *time.Time `json:"-"`
	DeletedAt     *time.Time `json:"-"`
}
