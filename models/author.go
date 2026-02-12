package models

import "gorm.io/gorm"



type  Author struct{
 gorm.Model
    Name    string `json:"name" binding:"required,min=2,max=100"`
    Country string `json:"country" binding:"required,min=2,max=100"`
	Books []Book `json:"-" gorm:"foreignKey:AuthorID"` //HAS MANY BOOKS
}

//AUthorResponse - for API responses with books
type AuthorResponse struct {
    ID  uint `json:"id"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
    Name string `json:"name"`
    Country string `json:"country"`
    Books []BookSimple `json:"books,omitempty"` // Use BookSimple to avoid nested author details
}