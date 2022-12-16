package storage

const (
	TaskTableName = "todos"
)

type Task struct {
	ID    int    `gorm:"column:id;primary_key"`
	Title string `gorm:"column:title"`
}

func (t Task) TableName() string {
	return TaskTableName
}
