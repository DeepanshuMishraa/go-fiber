package models

import "time"

type Order struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	ProductRefer int     `json:"product_refer"`
	Product      Product `gorm:foreignKey:ProductRefer`
	UserRefer    int     `json:"user_refer"`
	User         User    `gorm:foreignKey:UserRefer`
}
