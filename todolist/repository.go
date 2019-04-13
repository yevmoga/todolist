package todolist

import (
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetAllForToday() (tasks []Task, err error)
	CreateNewTask(task *Task) error
}

func NewMysqlRepository(database *gorm.DB) Repository {
	return &MysqlRepository{
		database: database,
	}
}

type MysqlRepository struct {
	database *gorm.DB
}

func (m *MysqlRepository) GetAllForToday() (tasks []Task, err error) {
	err = m.database.Where("showing_date::timestamp >= TIMESTAMP 'today' AND showing_date::timestamp < TIMESTAMP 'tomorrow'").Find(&tasks).Error
	return
}

func (m *MysqlRepository) CreateNewTask(task *Task) error {
	err := m.database.Create(task).Error
	return err
}
