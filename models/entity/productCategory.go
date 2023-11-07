package entity

import "time"

type ProductCategory struct {
	ID               int       `gorm:"primaryKey"`
	ParentCategoryId int       `gorm:"column:parent_category_id"`
	Name             string    `gorm:"column:name;uniqueIndex"`
	Products         []Product `gorm:"foreignKey:product_category_id;references:id"`
	CreatedAt        time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}