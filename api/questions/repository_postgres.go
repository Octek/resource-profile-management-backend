package questions

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type questionRepositoryPostgres struct {
	db *gorm.DB
}

func NewQuestionRepositoryPostgres(db *gorm.DB) QuestionRepository {
	err := db.AutoMigrate(&Question{}, &QuestionOption{})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Successfully connected to postgres in questions service!")

	return &questionRepositoryPostgres{
		db: db,
	}
}

func (repo *questionRepositoryPostgres) AddQuestion(request AddQuestionRequestBulk) error {
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		for _, questionReq := range request.AddQuestionRequest {

			questionModel := Question{
				Questions:    questionReq.Questions,
				QuestionType: questionReq.QuestionType,
			}

			if err := tx.Create(&questionModel).Error; err != nil {
				return fmt.Errorf("failed to create question: %w", err)
			}

			var questionOptionArray []QuestionOption
			for _, name := range questionReq.OptionNames {
				questionOption := QuestionOption{
					QuestionID: questionModel.ID,
					Name:       name,
				}
				questionOptionArray = append(questionOptionArray, questionOption)
			}

			if len(questionOptionArray) > 0 {
				if err := tx.Create(&questionOptionArray).Error; err != nil {
					return fmt.Errorf("failed to create question options for question ID %d: %w", questionModel.ID, err)
				}
			}
		}

		fmt.Println("Questions and their options have been created successfully.")
		return nil
	})

	return err
}

func (repo *questionRepositoryPostgres) GetQuestionById(id uint) (*Question, error) {
	var question Question
	err := repo.db.Model(Question{}).First(&question, id).Preload("QuestionOptions").Error
	return &question, err
}

func (repo *questionRepositoryPostgres) UpdateQuestionByID(question *Question) error {
	err := repo.db.Save(&question).Error
	return err
}

func (repo *questionRepositoryPostgres) GetAllQuestions(limit int, offset int, orderBy string) ([]Question, int64, error) {
	var questions []Question
	var total int64

	query := repo.db.Model(Question{}).Where("deleted_at IS NULL")

	err := query.Count(&total).Error
	err = query.Order(orderBy).Limit(limit).Offset(offset).Find(&questions).Error

	return questions, total, err
}

func (repo *questionRepositoryPostgres) DeleteQuestionByID(id uint) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("question_id = ?", id).Delete(&QuestionOption{}).Error; err != nil {
			return err
		}
		result := tx.Delete(&Question{}, id)
		if result.RowsAffected == 0 {
			return fmt.Errorf("no experience record found for id: %d", id)
		}

		return result.Error
	})

	return nil
}
