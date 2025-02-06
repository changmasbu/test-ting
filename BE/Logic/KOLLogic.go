package Logic

import (
	"context"
	"errors"
	"wan-api-kol-event/DTO"
	"wan-api-kol-event/Initializers"
	"wan-api-kol-event/Models"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

// * Get Kols from the database based on the range of pageIndex and pageSize
// ! USE GORM TO QUERY THE DATABASE
// ? There are some support function that can be access in Utils folder (/BE/Utils)
// --------------------------------------------------------------------------------
// @params: pageIndex
// @params: pageSize
// @return: List of KOLs and error message
func GetKolLogic(pageIndex, pageSize int) ([]*DTO.KolDTO, int64, error) {
	var kolEntities []Models.Kol
	var kolDTOs []*DTO.KolDTO
	var totalCount int64

	// Create database context
	db := Initializers.DB
	ctx := context.Background()

	// Count total records
	if err := db.WithContext(ctx).Model(&Models.Kol{}).Count(&totalCount).Error; err != nil {
		return nil, 0, errors.New("failed to count total KOLs")
	}

	// Fetch paginated KOLs
	if err := db.WithContext(ctx).
		Limit(pageSize).
		Offset((pageIndex - 1) * pageSize).
		Find(&kolEntities).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, errors.New("no KOLs found")
		}
		return nil, 0, err
	}

	// Convert models to DTOs
	err := copier.Copy(&kolDTOs, &kolEntities)
	if err != nil {
		return nil, 0, errors.New("failed to copy KOLs")
	}
	return kolDTOs, totalCount, nil
}
