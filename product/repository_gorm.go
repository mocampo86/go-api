package product

import (
	"context"

	"github.com/jinzhu/gorm"
)

type gormrepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) Repository {
	return &gormrepository{
		db: db,
	}
}

func (r gormrepository) GetAll(ctx context.Context) (products []DAOProduct, err error) {
	var result []DAOProduct
	if err := r.db.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r gormrepository) Find(ctx context.Context, filter ProductFilterRequest) (products []DAOProduct, err error) {
	var result []DAOProduct
	if filter.ProductId > 0 {
		if err := r.db.Where("ProductId=?", filter.ProductId).First(&result).Error; err != nil {
			return nil, err
		}
	} else {
		if filter.ProductName != "" {
			if err := r.db.Where("Name=?", filter.ProductName).First(&result).Error; err != nil {
				return nil, err
			}
		}
	}
	return result, err
}

func (r gormrepository) Insert(ctx context.Context, product DAOProduct) error {
	if dbc := r.db.Create(&product); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (r gormrepository) Update(ctx context.Context, product DAOProduct) error {
	var dbProduct DAOProduct
	if err := r.db.Where("ProductId=?", product.ProductId).First(&dbProduct).Error; err != nil {
		return err
	}
	if dbc := r.db.Save(&product); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (r gormrepository) Remove(ctx context.Context, id int64) error {
	var product DAOProduct
	if dbc := r.db.Where("ProductId=?", id).Delete(&product); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}
