package user

import (
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
	var user []User
	var total int64

	err := repo.db.Model(&User{}).Where("LOWER(first_name) LIKE ? AND deleted_at IS NULL", "%"+strings.ToLower(keyword)+"%").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = repo.db.Model(&User{}).Where("LOWER(first_name) LIKE ? AND deleted_at IS NULL", "%"+strings.ToLower(keyword)+"%").Order(orderBy).Limit(limit).Offset(offset).Find(&user).Error

	if err != nil {
		return nil, uint(total), err
	}

	return user, uint(total), nil
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
