package models

import "time"

type Order struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	ProductRefer uint      `json:"product_refer"` 
	Product      Product   `gorm:"foreignKey:ProductRefer;references:ID"`
	UserRefer    uint      `json:"user_refer"`
	User         User      `gorm:"foreignKey:UserRefer;references:ID"`
}
