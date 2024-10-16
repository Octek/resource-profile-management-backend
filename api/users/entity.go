package user

import (
	"crypto/sha256"
	"fmt"
	"github.com/Octek/resource-profile-management-backend.git/api/bookings"
	"github.com/Octek/resource-profile-management-backend.git/api/experience"
	"github.com/Octek/resource-profile-management-backend.git/api/projects"
	"github.com/Octek/resource-profile-management-backend.git/api/skills"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID             uint                    `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE;"`
	FirstName      string                  `json:"first_name" validate:"required"`
	LastName       string                  `json:"last_name" validate:"required"`
	Email          string                  `json:"email" validate:"required"`
	MobileNumber   string                  `json:"mobile_number" validate:"required"`
	Bio            string                  `json:"bio"`
	Location       string                  `json:"location"`
	VideoUrl       string                  `json:"video_url"`
	Certifications string                  `json:"certifications"`
	UserCategoryID uint                    `json:"user_category_id" gorm:"NOT NULL;index:user_category_id"`
	UserCategory   *UserCategory           `json:"user_category" gorm:"foreignKey:UserCategoryID;references:ID"`
	Educations     []Education             `json:"educations" gorm:"foreignKey:UserID"`
	Bookings       []bookings.Booking      `json:"bookings" gorm:"foreignKey:UserID"`
	Roles          []Role                  `json:"roles" gorm:"many2many:user_roles;"`
	Skills         []skills.Skill          `json:"skills" gorm:"many2many:user_skills;"`
	Experiences    []experience.Experience `json:"experiences" gorm:"many2many:user_experiences;"`
	Projects       []projects.Project      `json:"projects" gorm:"many2many:user_projects;"`
	DeletedAt      gorm.DeletedAt          `json:"deleted_at"`
	CreatedAt      time.Time               `json:"created_at"`
	UpdatedAt      time.Time               `json:"updated_at"`
}

type Education struct {
	ID              uint      `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE;"`
	UserID          uint      `json:"user_id" gorm:"NOT NULL;index"`
	InstitutionName string    `json:"institution_name"`
	Degree          string    `json:"degree"`
	FieldOfStudy    string    `json:"field_of_study"`
	Achievements    string    `json:"achievements"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type UserCategory struct {
	ID        uint      `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE;"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Role struct {
	ID        uint      `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE;"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRole struct {
	ID        uint      `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE;"`
	UserID    uint      `json:"user_id" gorm:"NOT NULL;index:user_id"`
	RoleID    uint      `json:"role_id" gorm:"NOT NULL;index:role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func asSha256Category(category UserCategory) string {
	org := UserCategory{
		ID:   category.ID,
		Name: category.Name,
	}
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%v", org)))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func asSha256Role(role Role) string {
	org := Role{
		ID:   role.ID,
		Name: role.Name,
	}
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%v", org)))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
