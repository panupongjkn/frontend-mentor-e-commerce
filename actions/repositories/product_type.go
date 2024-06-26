package repositories

import (
	"database/sql"
	"e-commerce-api/models"
)

type productTypeRepositoryDB struct {
	db *sql.DB
}

func NewProductTypeRepositoryDB(db *sql.DB) productTypeRepositoryDB {
	return productTypeRepositoryDB{db: db}
}

func (r productTypeRepositoryDB) GetByHomeSection() ([]models.ProductType, error) {
	productTypes := []models.ProductType{}

	rows, err := r.db.Query("SELECT product_types.slug, product_types.name, product_types.image FROM home_product_type_sections JOIN product_types ON home_product_type_sections.product_type_id = product_types.id WHERE home_product_type_sections.deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		productType := models.ProductType{}
		if err := rows.Scan(&productType.Slug, &productType.Name, &productType.Image); err != nil {
			return nil, err
		}
		productTypes = append(productTypes, productType)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return productTypes, nil
}

func (r productTypeRepositoryDB) GetBySlug(slug string) (*models.ProductType, error) {
	productType := &models.ProductType{}

	rows, err := r.db.Query("SELECT slug, name, image FROM product_types WHERE slug = $1 AND deleted_at IS NULL", slug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		if err := rows.Scan(&productType.Slug, &productType.Name, &productType.Image); err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return productType, nil
}
