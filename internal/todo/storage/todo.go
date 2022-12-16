package storage

import "gorm.io/gorm"

type TaskStorage struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) TaskStorage {
	return TaskStorage{db: db}
}

func (ts TaskStorage) Create(t Task) error {
	return ts.db.Create(t).Error
}

func (ts TaskStorage) Get(t Task) (*Task, error) {
	err := ts.db.Model(&t).Where(t).Take(&t).Error

	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (ts TaskStorage) Update(t Task) error {
	return ts.db.Model(t).Where("id = ?", t.ID).Save(t).Error
}

func (ts TaskStorage) Delete(t Task) error {
	err := ts.db.Where("id = ?", t.ID).Delete(Task{}).Error
	if err != nil {
		return err
	}

	return ts.db.Delete(&t).Error
}
