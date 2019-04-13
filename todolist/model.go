package todolist

import "time"

type Task struct {
	Id           uint64    `gorm:"primary_key;column:id" json:"id"`
	Title        string    `gorm:"column:title" json:"title" binding:"required"`
	UserID       uint64    `gorm:"column:user_id" json:"user_id" binding:"required"`
	Done         bool      `gorm:"column:done" json:"done"`
	ShowingDate  time.Time `gorm:"column:showing_date;default:''" json:"showing_date" time_format:"2006-01-02" binding:"required"`
	DateInserted time.Time `gorm:"column:date_inserted;default:''" json:"date_inserted"`
	DateUpdated  time.Time `gorm:"column:date_updated;default:''" json:"date_updated"`
}

func (Task) TableName() string {
	return "tasks"
}

func (Task) Validate() error {
	// todo-moga: implement validation (or gin-validation)
	return nil
}

type User struct {
	id       uint64
	name     string
	email    string
	password string
}

func (User) TableName() string {
	return "users"
}
