package projects

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type projectRepositoryPostgres struct {
	db *gorm.DB
}

func NewProjectRepositoryPostgres(db *gorm.DB) ProjectRepository {
	err := db.AutoMigrate(&UserProject{}, &Project{})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Successfully connected to postgres in projects service!")

	return &projectRepositoryPostgres{
		db: db,
	}
}
