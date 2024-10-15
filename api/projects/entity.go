package projects

import (
	"gorm.io/gorm"
	"time"
)

type Project struct {
	ID           uint           `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE;"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Link         string         `json:"link"`
	Technologies string         `json:"technologies"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type UserProject struct {
	ID        uint      `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE;"`
	UserID    uint      `json:"user_id" gorm:"NOT NULL;index:user_id"`
	ProjectID uint      `json:"project_id" gorm:"NOT NULL;index:project_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
