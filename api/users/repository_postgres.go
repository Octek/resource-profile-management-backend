package user

import (
	"errors"
	"fmt"
	"github.com/Octek/resource-profile-management-backend.git/api/experience"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
)

type userRepositoryPostgres struct {
	db *gorm.DB
}

func NewUserRepositoryPostgres(db *gorm.DB) UserRepository {
	err := db.AutoMigrate(&Role{}, &UserRole{}, &UserCategory{}, &User{}, &Education{})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Successfully connected to postgres in users service!")

	return &userRepositoryPostgres{
		db: db,
	}
}

func (repo *userRepositoryPostgres) CreateUser(user *User) (*User, error) {
	err := repo.db.Create(user).Error
	return user, err
}

func (repo *userRepositoryPostgres) GetAllUser(keyword string, limit int, offset int, orderBy string) ([]User, uint, error) {
	var users []User
	var total int64

	query := repo.db.Model(&User{}).Where("deleted_at IS NULL")
	if keyword != "" {
		query = query.Where("LOWER(first_name) LIKE ?", "%"+strings.ToLower(keyword)+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Order(orderBy).Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, uint(total), err
	}

	return users, uint(total), nil
}

func (repo *userRepositoryPostgres) GetUserDetailsByUserId(id uint) (*User, error) {
	var user User
	var experienceObj experience.Experience
	err := repo.db.Model(&User{}).Where("id = ? AND deleted_at IS NULL", id).
		Preload("Educations").Preload("Bookings").Preload("Roles").Preload("Skills").
		Preload("Skills.SkillCategory").
		Preload("Experiences").Preload("Projects").Preload("UserCategory").
		First(&user).Error

	experienceObj.ParseResponsibilities()

	return &user, err
}

func (repo *userRepositoryPostgres) DeleteUserByUserID(id uint) error {
	err := repo.db.Delete(&User{}, id).Error

	return err
}

func (repo *userRepositoryPostgres) UpdateUserByUserID(user *User) (*User, error) {
	if err := repo.db.Model(&User{}).Where("id = ?", user.ID).Updates(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &User{}, fmt.Errorf("user with ID %d not found", user.ID)
		}
		return &User{}, err
	}

	return user, nil
}

func (repo *userRepositoryPostgres) createCategories(jsonData []UserCategory) error {
	for _, cat := range jsonData {

		repo.db.Model(&cat).Where("id = ?", &cat.ID)
		dbRecord := UserCategory{}
		if err := repo.db.Model(&cat).Where("id = ?", &cat.ID).First(&dbRecord).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				dbRecord = cat
				repo.db.Create(&dbRecord)
			}
		}
		if asSha256Category(dbRecord) != asSha256Category(cat) {
			repo.db.Model(&cat).Where("id = ?", &cat.ID).Updates(&cat)
		}
	}
	return nil
}

func (repo *userRepositoryPostgres) createRoles(jsonData []Role) error {
	for _, role := range jsonData {

		repo.db.Model(&role).Where("id = ?", &role.ID)
		dbRecord := Role{}
		if err := repo.db.Model(&role).Where("id = ?", &role.ID).First(&dbRecord).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				dbRecord = role
				repo.db.Create(&dbRecord)
			}
		}
		if asSha256Role(dbRecord) != asSha256Role(role) {
			repo.db.Model(&role).Where("id = ?", &role.ID).Updates(&role)
		}
	}
	return nil
}

func (repo *userRepositoryPostgres) AddUserEducation(education Education) (Education, error) {

	err := repo.db.Model(Education{}).Create(&education).Error

	return education, err
}

func (repo *userRepositoryPostgres) GetEducationById(id uint) (*Education, error) {
	var education Education
	err := repo.db.Model(Education{}).First(&education, id).Error
	return &education, err
}

func (repo *userRepositoryPostgres) GetUserEducationByUserAndEducationId(userId, id uint) (*Education, error) {
	var education Education
	err := repo.db.Model(Education{}).Where("user_id = ? AND id = ?", userId, id).First(&education).Error
	return &education, err
}

func (repo *userRepositoryPostgres) UpdateEducation(education *Education) error {
	if err := repo.db.Model(&Education{}).Where("id = ?", education.ID).Updates(education).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("education with ID %d not found", education.UserID)
		}
		return err
	}
	return nil
}

func (repo *userRepositoryPostgres) GetUserEducationByUserId(userId uint) (*Education, error) {
	var education Education
	err := repo.db.Model(Education{}).Where("user_id = ? ", userId).First(&education).Error
	return &education, err
}

func (repo *userRepositoryPostgres) DeleteUserEducationByID(userId uint) error {
	err := repo.db.Model(Education{}).Where("user_id = ?", userId).Delete(&Education{}).Error

	return err
}

func (repo *userRepositoryPostgres) GetAllUserEducation(userId uint, limit int, offset int, orderBy string) ([]Education, uint, error) {
	var educations []Education
	var total int64

	query := repo.db.Model(Education{})
	query = query.Where("user_id = ?", userId)

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Order(orderBy).Limit(limit).Offset(offset).Find(&educations).Error
	if err != nil {
		return nil, uint(total), err
	}

	return educations, uint(total), nil
}

func (repo *userRepositoryPostgres) GetAllUserCategories(keyword string, limit int, offset int, orderBy string) ([]UserCategory, int64, error) {
	users := make([]UserCategory, 0)
	var count int64
	if keyword == "" {
		_ = repo.db.Model(&UserCategory{}).Count(&count)
		results := repo.db.Model(&UserCategory{}).Select("id,name").Limit(limit).Offset(offset).Order(orderBy).Find(&users)
		if err := results.Error; err != nil {
			return users, count, err
		}
	} else {
		keyword = "%" + strings.ToLower(keyword) + "%"
		_ = repo.db.Model(&UserCategory{}).Where("Lower(name) LIKE ?", keyword).Count(&count)
		results := repo.db.Model(&UserCategory{}).Select("id,name").Where("LOWER(name) LIKE ?", keyword).Limit(limit).Offset(offset).Order(orderBy).Find(&users)
		if err := results.Error; err != nil {
			return users, count, err
		}
	}
	return users, count, nil
}
