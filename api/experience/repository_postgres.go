package experience

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type experienceRepositoryPostgres struct {
	db *gorm.DB
}

func NewExperienceRepositoryPostgres(db *gorm.DB) ExperienceRepository {
	err := db.AutoMigrate(&UserExperience{}, &ExperienceSkill{}, &Experience{})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Successfully connected to postgres in experience service!")

	return &experienceRepositoryPostgres{
		db: db,
	}
}
