package todolist

import "time"

type Todo struct {
	Id           uint64    `gorm:"primary_key;column:id"`
	Title        string    `gorm:"column:title"`
	UserID       uint64    `gorm:"column:user_id"`
	Done         bool      `gorm:"column:done"`
	ShowingDate  time.Time `gorm:"column:showing_date;default:''"`
	DateInserted time.Time `gorm:"column:date_inserted;default:''"`
	DateUpdated  time.Time `gorm:"column:date_updated;default:''"`
}

func (Todo) TableName() string {
	return "todos"
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
