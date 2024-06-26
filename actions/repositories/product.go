package repositories

import (
	"database/sql"
	"e-commerce-api/models"
	"errors"
	"log"
)

type productRepositoryDB struct {
	db *sql.DB
}

func NewProductRepositoryDB(db *sql.DB) productRepositoryDB {
	return productRepositoryDB{db: db}
}

func (r productRepositoryDB) GetBySlug(slug string) (*models.Product, error) {
	product := &models.Product{}
	log.Print("workkkkk")
	rows, err := r.db.Query("SELECT id, slug, sub_title, name, image, description, price, quantity, feature, product_type_id FROM products WHERE products.slug = $1 AND products.deleted_at IS NULL", slug)
	if err != nil {
		log.Print("workkkkk 1")

		return nil, err
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Slug, &product.SubTitle, &product.Name, &product.Image, &product.Description, &product.Price, &product.Quantity, &product.Feature, &product.ProductTypeID); err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if product.ID == 0 {
		return nil, errors.New("Product not found.")
	}

	itemsInBox := []models.InTheBox{}
	inTheBoxRows, err := r.db.Query("SELECT name as item, quantity FROM item_in_product_boxs WHERE product_id = $1 AND deleted_at IS NULL", product.ID)
	if err != nil {
		log.Print("err 1")
		return nil, err
	}
	defer inTheBoxRows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for inTheBoxRows.Next() {
		log.Print("err 1")
		itemInBox := models.InTheBox{}
		if err := inTheBoxRows.Scan(&itemInBox.Item, &itemInBox.Quantity); err != nil {
			return nil, err
		}
		itemsInBox = append(itemsInBox, itemInBox)
	}
	if err := inTheBoxRows.Err(); err != nil {
		return nil, err
	}

	galleries := []string{}
	galleriesRows, err := r.db.Query("SELECT image FROM galleries WHERE product_id = $1 AND deleted_at IS NULL", product.ID)
	if err != nil {
		log.Print("err 2")
		return nil, err
	}
	defer galleriesRows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for galleriesRows.Next() {
		gallery := ""
		if err := galleriesRows.Scan(&gallery); err != nil {
			return nil, err
		}
		galleries = append(galleries, gallery)
	}
	if err := galleriesRows.Err(); err != nil {
		return nil, err
	}

	productsRelated := []models.ProductRelated{}
	productsRelatedRows, err := r.db.Query("SELECT slug, name, image, price FROM products WHERE products.id != $1 AND products.product_type_id = $2 AND products.deleted_at IS NULL LIMIT 3", product.ID, product.ProductTypeID)
	if err != nil {
		log.Print("err 3")
		return nil, err
	}
	defer productsRelatedRows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for productsRelatedRows.Next() {
		productRelated := models.ProductRelated{}
		if err := productsRelatedRows.Scan(&productRelated.Slug, &productRelated.Name, &productRelated.Image, &productRelated.Price); err != nil {
			return nil, err
		}
		productsRelated = append(productsRelated, productRelated)
	}
	if err := productsRelatedRows.Err(); err != nil {
		return nil, err
	}

	product.InTheBoxs = itemsInBox
	product.Gallery = galleries
	product.ProductRelated = productsRelated
	return product, nil
}

func (r productRepositoryDB) GetByProductType(productTypeSlug string) ([]models.ProductByProductType, error) {
	productType := &models.ProductType{}

	rows, err := r.db.Query("SELECT id FROM product_types WHERE slug = $1 AND deleted_at IS NULL", productTypeSlug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		if err := rows.Scan(&productType.ID); err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if productType.ID == 0 {
		return nil, errors.New("Product type not found.")
	}

	products := []models.ProductByProductType{}
	productsRows, err := r.db.Query("SELECT slug, sub_title, name, image, description FROM products WHERE products.product_type_id = $1 AND products.deleted_at IS NULL", productType.ID)
	if err != nil {
		return nil, err
	}
	defer productsRows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for productsRows.Next() {
		product := models.ProductByProductType{}
		if err := productsRows.Scan(&product.Slug, &product.SubTitle, &product.Name, &product.Image, &product.Description); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := productsRows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
