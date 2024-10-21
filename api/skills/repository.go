package skills

// SkillRepository Used to store and retrieve skills based on experience and bookings
type SkillRepository interface {
	createCategories(jsonData []SkillCategory) error
	createSkill(userSkillData *UserSkillRequest) error
	createSkillCategories(skillCategories []SkillCategory) error
	getSkillCategoryById(id uint) (SkillCategory, error)
	deleteSkillCategoryById(id uint) error
	updateSkillCategory(skillCategoryObj SkillCategory) error
	fetchAllSkillCategories(limit, offset int, orderBy string) ([]SkillCategory, int64, error)
	getSkillById(id uint) (Skill, error)
	updateSkill(skillObj Skill) error
	deleteSkillById(id uint) error
	fetchAllSkill(limit, offset int, orderBy, keyword string) ([]Skill, int64, error)
}
