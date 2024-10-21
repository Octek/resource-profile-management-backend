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

func (repo *experienceRepositoryPostgres) AddExperienceWithUserAndSkills(userID, skillId uint, experience Experience) (Experience, error) {

	err := repo.db.Create(&experience).Error
	userExperience := UserExperience{
		UserID:       userID,
		ExperienceID: experience.ID,
	}

	experienceSkill := ExperienceSkill{
		SkillID:      skillId,
		ExperienceID: experience.ID,
	}

	err = repo.db.Create(&userExperience).Error
	err = repo.db.Create(&experienceSkill).Error

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
	err := repo.db.Delete(&Experience{}, id).Error

	return err
}
