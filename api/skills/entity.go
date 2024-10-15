package skills

import (
	"crypto/sha256"
	"fmt"
	"github.com/Octek/resource-profile-management-backend.git/api/bookings"
	"gorm.io/gorm"
	"time"
)

type Skill struct {
	ID              uint               `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE;"`
	Name            string             `json:"name"`
	Icon            string             `json:"icon"`
	SkillCategoryID uint               `json:"skill_category_id" gorm:"NOT NULL;index:skill_category_id"`
	SkillCategory   *SkillCategory     `json:"skill_category" gorm:"foreignKey:SkillCategoryID;references:ID"`
	Bookings        []bookings.Booking `json:"bookings" gorm:"many2many:booking_skills;"`
	DeletedAt       gorm.DeletedAt     `json:"deleted_at"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
}

type SkillCategory struct {
	ID        uint      `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE;"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserSkill struct {
	ID         uint      `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE;"`
	UserID     uint      `json:"user_id" gorm:"NOT NULL;index:user_id"`
	SkillID    uint      `json:"skill_id" gorm:"NOT NULL;index:skill_id"`
	SkillLevel string    `json:"skill_level"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func asSha256SkillCategory(category SkillCategory) string {
	newSkillCategory := SkillCategory{
		ID:   category.ID,
		Name: category.Name,
	}
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%v", newSkillCategory)))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
