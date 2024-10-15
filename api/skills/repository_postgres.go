package skills

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type skillRepositoryPostgres struct {
	db *gorm.DB
}

func NewSkillRepositoryPostgres(db *gorm.DB) SkillRepository {
	err := db.AutoMigrate(&UserSkill{}, &SkillCategory{}, &Skill{})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Successfully connected to postgres in skills service!")

	return &skillRepositoryPostgres{
		db: db,
	}
}

func (repo *skillRepositoryPostgres) createCategories(jsonData []SkillCategory) error {
	for _, cat := range jsonData {

		repo.db.Model(&cat).Where("id = ?", &cat.ID)
		dbRecord := SkillCategory{}
		if err := repo.db.Model(&cat).Where("id = ?", &cat.ID).First(&dbRecord).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				dbRecord = cat
				repo.db.Create(&dbRecord)
			}
		}
		if asSha256SkillCategory(dbRecord) != asSha256SkillCategory(cat) {
			repo.db.Model(&cat).Where("id = ?", &cat.ID).Updates(&cat)
		}
	}
	return nil
}
