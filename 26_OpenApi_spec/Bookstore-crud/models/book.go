package models

// TableName = strings.ToLower(structName) + "s"
type Book struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Author string `json:"author"`
}