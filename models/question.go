package models

type Question struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Question string `json:"question" gorm:"not null"`
	Upvotes  uint   `json:"upvotes" gorm:"default:0"`
	Downvotes uint   `json:"downvotes" gorm:"default:0"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
}
