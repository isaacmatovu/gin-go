package models

import "gorm.io/gorm"


type Category struct{
	gorm.Model
	Name string `json:"name" binding:"required,min=2,max=100"`
	Books []Book `json:"books,omitempty" gorm:"many2many:book_categories;constraint:OnDelete:CASCADE"` //many to many with book
}


type CategorySimple struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name      string `json:"name"`
}

type CategoryResponse struct {
	ID uint `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name string `json:"name"`
    Books []BookResponse `json:"books,omitempty"` // Include books in category response
}