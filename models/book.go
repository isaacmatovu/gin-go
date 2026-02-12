package models

import "gorm.io/gorm"


type Book struct{
	gorm.Model    // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	Title  string `json:"title" binding:"required,min=1,max=200"`
	Year int `json:"year" binding:"required,min=1000,max=2100"`
	AuthorID uint `json:"author_id" binding:"required"` //foreign_key
    Author   Author `json:"author,omitempty" binding:"-"` // Belongs to Author
    Categories []Category `json:"categories,omitempty" gorm:"many2many:book_categories;constraint:OnDelete:CASCADE"` //many to many with category
}

//bookresponse for when you want to include author
type BookResponse struct {
	ID uint    	`json:"id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	Title string       `json:"title"`
	Year int `json:"year"`
	AuthorID  uint `json:"author_id"`
	Author *Author    `json:"author,omitempty"`
}

//BookSimple for when you DONT want to include author (nested in Author response)
type BookSimple struct {
	ID uint `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Title string `json:"title"`
	Year int `json:"year"`
	AuthorID uint `json:"author_id"`
}

type BookCategory struct{
	ID uint `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Title string `json:"title"`
	Year int `json:"year"`
	AuthorID uint `json:"author_id"`
	Author *Author `json:"author,omitempty"`
	Categories []CategorySimple `json:"categories,omitempty" gorm:"many2many:book_categories;constraint:OnDelete:CASCADE"`
}