package questions

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type questionRepositoryPostgres struct {
	db *gorm.DB
}

func NewQuestionRepositoryPostgres(db *gorm.DB) QuestionRepository {
	err := db.AutoMigrate(&QuestionOption{}, &Question{})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Successfully connected to postgres in questions service!")

	return &questionRepositoryPostgres{
		db: db,
	}
}
