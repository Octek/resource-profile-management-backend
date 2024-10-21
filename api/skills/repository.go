package skills

// SkillRepository Used to store and retrieve skills based on experience and bookings
type SkillRepository interface {
	createCategories(jsonData []SkillCategory) error
	createSkill(skillObj *Skill) error
	createSkillCategories(skillCategories []SkillCategory) error
	getSkillCategoryById(id uint) (SkillCategory, error)
	deleteSkillCategoryById(id uint) error
	updateSkillCategory(skillCategoryObj SkillCategory) error
	fetchAllSkillCategories(limit, offset int, orderBy string) ([]SkillCategory, int64, error)
}
