package questions

import (
	"gorm.io/gorm"
	"time"
)

type Question struct {
	ID              uint             `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE;"`
	Questions       string           `json:"questions"`
	QuestionType    string           `json:"question_type"`
	QuestionOptions []QuestionOption `json:"question_options"  gorm:"foreignKey:QuestionID"`
	DeletedAt       gorm.DeletedAt   `json:"deleted_at"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
}

type QuestionOption struct {
	ID         uint      `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE;"`
	QuestionID uint      `json:"question_id" gorm:"NOT NULL;index:question_id"`
	Name       string    `json:"name" gorm:"NOT NULL"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
