package main

type Book struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Author string `json:"author"`
} // helps in creating book table

