package models

import "time"

type Product struct {
	ID            int        `json:"id"`
	Slug          string     `json:"slug"`
	SubTitle      string     `json:"sub_title"`
	Name          string     `json:"name"`
	Image         string     `json:"image"`
	Description   string     `json:"description"`
	Price         float64    `json:"price"`
	Quantity      int        `json:"quantity"`
	Feature       string     `json:"feature"`
	ProductTypeID int        `json:"-"`
	CreatedAt     time.Time  `json:"-"`
	UpdatedAt     *time.Time `json:"-"`
	DeletedAt     *time.Time `json:"-"`

	InTheBoxs      []InTheBox       `json:"in_the_box"`
	Gallery        []string         `json:"gallery"`
	ProductRelated []ProductRelated `json:"product_related"`
}

type InTheBox struct {
	Item     string `json:"item"`
	Quantity int    `json:"quantity"`
}

type ProductRelated struct {
	Slug  string  `json:"slug"`
	Name  string  `json:"name"`
	Image string  `json:"image"`
	Price float64 `json:"price"`
}

type ProductByProductType struct {
	Slug        string `json:"slug"`
	SubTitle    string `json:"sub_title"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
}
