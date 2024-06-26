package repositories

import (
	"database/sql"
	"e-commerce-api/models"
)

type bannerRepositoryDB struct {
	db *sql.DB
}

func NewBannerRepositoryDB(db *sql.DB) bannerRepositoryDB {
	return bannerRepositoryDB{db: db}
}

func (r bannerRepositoryDB) GetAll() ([]models.Banner, error) {
	banners := []models.Banner{}

	rows, err := r.db.Query("SELECT * FROM banners WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		banner := models.Banner{}
		if err := rows.Scan(&banner.ID, &banner.Title, &banner.SubTitle, &banner.Description, &banner.ButtonText, &banner.ButtonUrl, &banner.Image, &banner.CreatedAt, &banner.UpdatedAt, &banner.DeletedAt); err != nil {
			return nil, err
		}
		banners = append(banners, banner)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return banners, nil
}
