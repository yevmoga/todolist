package todolist

import (
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetAllForToday() (array []Todo, err error)
}

func NewMysqlRepository(database *gorm.DB) Repository {
	return &MysqlRepository{
		database: database,
	}
}

type MysqlRepository struct {
	database *gorm.DB
}

func (m *MysqlRepository) GetAllForToday() (array []Todo, err error) {
	err = m.database.Where("showing_date::timestamp >= TIMESTAMP 'today' AND showing_date::timestamp < TIMESTAMP 'tomorrow'").Find(&array).Error
	return
}
