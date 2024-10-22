package projects

import (
	"fmt"
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

func (repo *projectRepositoryPostgres) AddUserProject(userID uint, project *Project) (*Project, error) {

	err := repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&project).Error; err != nil {
			return err
		}

		userProject := UserProject{
			UserID:    userID,
			ProjectID: project.ID,
		}
		if err := tx.Create(&userProject).Error; err != nil {
			return err
		}

		fmt.Println("Experience, UserExperience, and ExperienceSkill have been created successfully.")
		return nil
	})

	return project, err
}

func (repo *projectRepositoryPostgres) GetProjectById(id uint) (*Project, error) {
	var project Project
	if err := repo.db.First(&project, id).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (repo *projectRepositoryPostgres) GetUserProjectByUserAndProjectId(userId, projId uint) (*UserProject, error) {
	var userProject UserProject
	if err := repo.db.Where("user_id = ? AND project_id = ?", userId, projId).First(&userProject).Error; err != nil {
		return nil, err
	}
	return &userProject, nil
}

func (repo *projectRepositoryPostgres) UpdateProject(project *Project) error {
	return repo.db.Save(project).Error
}

func (repo *projectRepositoryPostgres) GetUserProjectByUserId(userID, projId uint) (*Project, error) {
	var project Project

	err := repo.db.Model(&Project{}).
		Joins("JOIN user_projects up ON up.project_id = projects.id").
		Where("up.user_id = ? and up.project_id = ?", userID, projId).
		First(&project).
		Error

	return &project, err
}

func (repo *projectRepositoryPostgres) DeleteUserProjectByID(userId uint) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		var projIDs []uint
		if err := tx.Model(&UserProject{}).Where("user_id = ?", userId).
			Pluck("project_id", &projIDs).Error; err != nil {
			return err
		}
		if len(projIDs) == 0 {
			return fmt.Errorf("no experience records found for user_id: %d", userId)
		}
		err := tx.Where("user_id = ?", userId).Delete(&UserProject{}).Error
		err = tx.Where("id IN (?)", projIDs).Delete(&Project{}).Error

		return err
	})
}

func (repo *projectRepositoryPostgres) GetAllUserProject(userId uint, limit int, offset int, orderBy string) ([]Project, uint, error) {
	var projIDs []uint
	var projects []Project
	var total int64

	err := repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&UserProject{}).
			Where("user_id = ?", userId).
			Pluck("project_id", &projIDs).Error; err != nil {
			return err
		}
		if len(projIDs) == 0 {
			return fmt.Errorf("no experience records found for user_id: %d", userId)
		}

		query := tx.Model(&Project{}).Where("deleted_at IS NULL").Where("id IN (?)", projIDs)
		err := query.Count(&total).Error
		err = query.Order(orderBy).Limit(limit).Offset(offset).Find(&projects).Error

		return err
	})

	return projects, uint(total), err
}
