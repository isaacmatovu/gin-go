package models


type Book struct{
	ID string `json:"id"`
	Title  string `json:"title" binding:"required,min=1,max=200"`
	Author string `json:"author" binding:"required,min=2,max=100"`
	Year int `json:"year" binding:"required,min=1000,max=2100"`
}
