package skills

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
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

func (repo *skillRepositoryPostgres) createSkill(skillObj *Skill, userID uint, skillLevel string) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&skillObj).Error; err != nil {
			return err
		}
		userSkillObj := UserSkill{
			SkillLevel: skillLevel,
			SkillID:    skillObj.ID,
			UserID:     userID,
		}
		if err := tx.Create(&userSkillObj).Error; err != nil {
			return err
		}
		fmt.Println("Skill and UserSkill objects have been stored")
		return nil
	})
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

func (repo *skillRepositoryPostgres) getSkillById(id uint) (Skill, error) {
	var skillObj Skill
	if err := repo.db.Where("id = ?", id).Preload("SkillCategory").Preload("Bookings").First(&skillObj).Error; err != nil {
		return Skill{}, err
	}

	fmt.Printf("Skill by id %d has been fetched\n", id)
	return skillObj, nil
}

func (repo *skillRepositoryPostgres) updateSkill(skillObj Skill) error {
	if err := repo.db.Save(&skillObj).Error; err != nil {
		return err
	}
	fmt.Println("Skill has been updated")
	return nil
}

func (repo *skillRepositoryPostgres) deleteSkillById(id uint) error {
	result := repo.db.Where("id = ?", id).Delete(&Skill{})
	if result.RowsAffected == 0 {
		return fmt.Errorf("skill with id %d not found", id)
	}
	if result.Error != nil {
		return result.Error
	}

	fmt.Printf("Skill by id %d has been deleted\n", id)
	return nil
}

func (repo *skillRepositoryPostgres) fetchAllSkill(limit, offset int, orderBy, keyword string) ([]Skill, int64, error) {
	var skillList []Skill
	var totalRecords int64

	query := repo.db.Model(&Skill{}).Where("deleted_at IS NULL")
	if keyword != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(keyword)+"%")
	}

	err := query.Count(&totalRecords).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Order(orderBy).Limit(limit).Offset(offset).Preload("SkillCategory").Preload("Bookings").Find(&skillList).Error
	if err != nil {
		return nil, totalRecords, err
	}

	return skillList, totalRecords, nil
}
