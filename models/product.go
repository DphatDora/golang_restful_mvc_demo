package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string   `gorm:"type:varchar(200);not null" json:"name"`
	Description string   `gorm:"type:text" json:"description"`
	Price       float64  `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock       int      `gorm:"type:int;not null" json:"stock"`
	CategoryID  uint     `gorm:"not null" json:"category_id"`
	Category    Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}
