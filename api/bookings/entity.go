package bookings

import (
	"github.com/Octek/resource-profile-management-backend.git/api/questions"
	"gorm.io/gorm"
	"time"
)

type Booking struct {
	ID              uint                 `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE;"`
	UserID          uint                 `json:"user_id" gorm:"NOT NULL;index:user_id"`
	BookingDateTime time.Time            `json:"booking_date_time"`
	MeetingLink     string               `json:"meeting_link"`
	Questions       []questions.Question `json:"questions" gorm:"many2many:booking_questions;"`
	DeletedAt       gorm.DeletedAt       `json:"deleted_at"`
	CreatedAt       time.Time            `json:"created_at"`
	UpdatedAt       time.Time            `json:"updated_at"`
}

type BookingSkill struct {
	ID        uint      `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE;"`
	SkillID   uint      `json:"skill_id" gorm:"NOT NULL;index:skill_id"`
	BookingID uint      `json:"booking_id" gorm:"NOT NULL;index:booking_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookingQuestion struct {
	ID         uint      `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE;"`
	QuestionID uint      `json:"question_id" gorm:"NOT NULL;index:question_id"`
	BookingID  uint      `json:"booking_id" gorm:"NOT NULL;index:booking_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
