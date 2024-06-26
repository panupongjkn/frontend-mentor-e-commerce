package repositories

import (
	"database/sql"
	"e-commerce-api/models"
)

type homeProductSectionRepositoryDB struct {
	db *sql.DB
}

func NewHomeProductSectionRepositoryDB(db *sql.DB) homeProductSectionRepositoryDB {
	return homeProductSectionRepositoryDB{db: db}
}

func (r homeProductSectionRepositoryDB) GetAll() ([]models.HomeProductSection, error) {
	homeProductSections := []models.HomeProductSection{}

	rows, err := r.db.Query("SELECT home_product_sections.display_type, products.slug, products.name, products.image FROM home_product_sections JOIN products ON home_product_sections.product_id = products.id WHERE home_product_sections.deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		homeProductSection := models.HomeProductSection{}
		if err := rows.Scan(&homeProductSection.DisplayType, &homeProductSection.Product.Slug, &homeProductSection.Product.Name, &homeProductSection.Product.Image); err != nil {
			return nil, err
		}
		homeProductSections = append(homeProductSections, homeProductSection)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return homeProductSections, nil
}
