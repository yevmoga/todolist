package todolist

import (
	"errors"
	"time"
)

const (
	validateById          = "id"
	validateByDone        = "done"
	validateByTitle       = "title"
	validateByShowingDate = "showing_date"
)

var validators = map[string]func(task *Task) error{
	validateById: func(task *Task) error {
		if task.Id == 0 {
			return errors.New("error: empty task id")
		}
		return nil
	},
	validateByDone: func(task *Task) error {
		if task.Done == 0 || task.Done == 1 {
			return nil
		}
		return errors.New("incorrect `done` value")
	},
	validateByTitle: func(task *Task) error {
		if task.Title == "" {
			return errors.New("title is empty")
		}
		return nil
	},
	validateByShowingDate: func(task *Task) error {
		if task.ShowingDate.IsZero() {
			return errors.New("showing date is empty")
		}
		return nil
	},
}

type Task struct {
	Id           uint64    `gorm:"primary_key;column:id" json:"id"`
	Title        string    `gorm:"column:title" json:"title"`
	UserID       uint64    `gorm:"column:user_id" json:"user_id"`
	Done         int       `gorm:"column:done" json:"done"`
	ShowingDate  time.Time `gorm:"column:showing_date;default:''" json:"showing_date" time_format:"2006-01-02"`
	DateInserted time.Time `gorm:"column:date_inserted;default:''" json:"date_inserted"`
	DateUpdated  time.Time `gorm:"column:date_updated;default:''" json:"date_updated"`
}

func (Task) TableName() string {
	return "tasks"
}

func (task Task) ValidateByFields(fields []string) (err error) {
	for _, field := range fields {
		validator, ok := validators[field]
		if !ok {
			panic("needs validator")
		}

		if err = validator(&task); err != nil {
			return err
		}
	}

	return nil
}

type User struct {
	Id       uint64
	Name     string
	Email    string
	Password string
}

func (User) TableName() string {
	return "users"
}
