package storage

type Task struct {
	ID    int    `gorm:"column:id;primary_key"`
	Title string `gorm:"column:title"`
}
