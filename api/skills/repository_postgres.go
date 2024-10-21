package skills

import (
	"fmt"
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
func (repo *skillRepositoryPostgres) createSkill(skillObj *Skill) error {
	if err := repo.db.Create(&skillObj).Error; err != nil {
		return err
	}
	fmt.Println("Skill object has been stored")
	return nil
}

func (repo *skillRepositoryPostgres) createSkillCategories(skillCategories []SkillCategory) error {
	if err := repo.db.CreateInBatches(skillCategories, len(skillCategories)).Error; err != nil {
		return err
	}
	fmt.Println("Skill categories  has been stored")
	return nil
}

func (repo *skillRepositoryPostgres) getSkillCategoryById(id uint) (SkillCategory, error) {
	var skillCategoryObj SkillCategory
	if err := repo.db.Where("id = ?", id).First(&skillCategoryObj).Error; err != nil {
		return SkillCategory{}, err
	}

	fmt.Printf("Skill category by id %d has been fetched\n", id)
	return skillCategoryObj, nil
}

func (repo *skillRepositoryPostgres) deleteSkillCategoryById(id uint) error {
	result := repo.db.Where("id = ?", id).Delete(&SkillCategory{})
	if result.RowsAffected == 0 {
		return fmt.Errorf("skill category with id %d not found", id)
	}
	if result.Error != nil {
		return result.Error
	}

	fmt.Printf("Skill category by id %d has been deleted\n", id)
	return nil
}

func (repo *skillRepositoryPostgres) updateSkillCategory(skillCategoryObj SkillCategory) error {
	if err := repo.db.Save(&skillCategoryObj).Error; err != nil {
		return err
	}
	fmt.Println("Skill category has been updated")
	return nil
}
func (repo *skillRepositoryPostgres) fetchAllSkillCategories(limit, offset int, orderBy string) ([]SkillCategory, int64, error) {
	var skillCategoryList []SkillCategory
	var totalRecords int64
	if err := repo.db.Model(&SkillCategory{}).Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}
	if err := repo.db.Order(orderBy).Offset(offset).Limit(limit).Find(&skillCategoryList).Error; err != nil {
		return nil, 0, err
	}

	return skillCategoryList, totalRecords, nil
}
