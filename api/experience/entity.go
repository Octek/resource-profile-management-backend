package experience

import (
	"github.com/Octek/resource-profile-management-backend.git/api/skills"
	"gorm.io/gorm"
	"time"
)

type Experience struct {
	ID                 uint           `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE;"`
	Position           string         `json:"position"`
	Company            string         `json:"company"`
	Description        string         `json:"description"`
	StartDate          time.Time      `json:"start_date"`
	EndDate            time.Time      `json:"end_date"`
	IsCurrentlyWorking bool           `json:"is_currently_working"`
	Responsibilities   string         `json:"responsibilities"`
	Skills             []skills.Skill `json:"skills" gorm:"many2many:experience_skills;"`
	DeletedAt          gorm.DeletedAt `json:"deleted_at"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
}

type UserExperience struct {
	ID           uint      `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE;"`
	UserID       uint      `json:"user_id" gorm:"NOT NULL;index:user_id"`
	ExperienceID uint      `json:"experience_id" gorm:"NOT NULL;index:skill_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ExperienceSkill struct {
	ID           uint      `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE;"`
	SkillID      uint      `json:"skill_id" gorm:"NOT NULL;index:user_id"`
	ExperienceID uint      `json:"experience_id" gorm:"NOT NULL;index:experience_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ExperienceResponse struct {
	Experience       Experience `json:"experience"`
	Responsibilities []string   `json:"responsibilities"`
}
