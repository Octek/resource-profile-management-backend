package experience

import (
	"fmt"
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

func (repo *experienceRepositoryPostgres) AddExperienceWithUserAndSkills(userID, skillId uint, experience *Experience) (*Experience, error) {

	err := repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&experience).Error; err != nil {
			return err
		}

		experience.ParseResponsibilities()
		userExperience := UserExperience{
			UserID:       userID,
			ExperienceID: experience.ID,
		}
		if err := tx.Create(&userExperience).Error; err != nil {
			return err
		}

		experienceSkill := ExperienceSkill{
			SkillID:      skillId,
			ExperienceID: experience.ID,
		}
		if err := tx.Create(&experienceSkill).Error; err != nil {
			return err
		}
		fmt.Println("Experience, UserExperience, and ExperienceSkill have been created successfully.")
		return nil
	})

	return experience, err
}

func (repo *experienceRepositoryPostgres) GetExperienceById(id uint) (*Experience, error) {
	var experience Experience
	if err := repo.db.First(&experience, id).Error; err != nil {
		return nil, err
	}
	return &experience, nil
}

func (repo *experienceRepositoryPostgres) GetUserExperienceByUserIdAndExperienceId(userId, experienceId uint) (*UserExperience, error) {
	var userExperience UserExperience
	if err := repo.db.Where("user_id = ? AND experience_id = ?", userId, experienceId).First(&userExperience).Error; err != nil {
		return nil, err
	}
	return &userExperience, nil
}

func (repo *experienceRepositoryPostgres) UpdateExperience(experience *Experience) error {
	return repo.db.Save(experience).Error
}

func (repo *experienceRepositoryPostgres) GetAllUserExperienceList(expID, userID uint) (Experience, error) {
	var experience Experience

	err := repo.db.Model(&Experience{}).
		Joins("JOIN user_experiences ue ON ue.experience_id = experiences.id").
		Where("ue.user_id = ? AND ue.experience_id = ?", userID, expID).
		Preload("Skills").
		Preload("Skills.SkillCategory").
		First(&experience).
		Error

	experience.ParseResponsibilities()

	return experience, err
}

func (repo *experienceRepositoryPostgres) DeleteUserExperienceByID(id uint) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("experience_id = ?", id).Delete(&ExperienceSkill{}).Error; err != nil {
			return err
		}
		if err := tx.Where("experience_id = ?", id).Delete(&UserExperience{}).Error; err != nil {
			return err
		}
		result := tx.Delete(&Experience{}, id)
		if result.RowsAffected == 0 {
			return fmt.Errorf("no experience record found for id: %d", id)
		}

		return result.Error
	})
}

func (repo *experienceRepositoryPostgres) DeleteUserExperienceByUserID(userId uint) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		var experienceIDs []uint
		if err := tx.Model(&UserExperience{}).Where("user_id = ?", userId).
			Pluck("experience_id", &experienceIDs).Error; err != nil {
			return err
		}
		if len(experienceIDs) == 0 {
			return fmt.Errorf("no experience records found for user_id: %d", userId)
		}
		err := tx.Where("user_id = ?", userId).Delete(&UserExperience{}).Error
		err = tx.Where("experience_id IN (?)", experienceIDs).Delete(&ExperienceSkill{}).Error

		err = tx.Where("id IN (?)", experienceIDs).Delete(&Experience{}).Error

		return err
	})
}

func (repo *experienceRepositoryPostgres) GetAllUserExperience(userId uint, limit int, offset int, orderBy string) ([]Experience, uint, error) {
	var experienceIDs []uint
	var exp []Experience
	var total int64

	err := repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&UserExperience{}).
			Where("user_id = ?", userId).
			Pluck("experience_id", &experienceIDs).Error; err != nil {
			return err
		}
		if len(experienceIDs) == 0 {
			return fmt.Errorf("no experience records found for user_id: %d", userId)
		}

		query := tx.Model(&Experience{}).Where("deleted_at IS NULL").Where("id IN (?)", experienceIDs)
		err := query.Count(&total).Error
		err = query.Order(orderBy).Limit(limit).Offset(offset).Find(&exp).Error
		for i := range exp {
			exp[i].ParseResponsibilities()
		}

		return err
	})

	return exp, uint(total), err
}
