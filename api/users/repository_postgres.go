package user

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
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
