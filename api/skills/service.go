package skills

type SkillService struct {
	skillRepository SkillRepository
}

func NewService(r SkillRepository) SkillService {
	return SkillService{skillRepository: r}
}

func (svc *SkillService) CreateCategories(jsonData []SkillCategory) error {
	return svc.skillRepository.createCategories(jsonData)
}
func (svc *SkillService) createSkill(userSkillData *UserSkillRequest) error {
	return svc.skillRepository.createSkill(userSkillData)
}
func (svc *SkillService) CreateSkillCategories(skillCategoryObj []SkillCategory) error {
	return svc.skillRepository.createSkillCategories(skillCategoryObj)
}

func (svc *SkillService) GetSkillCategoryById(id uint) (SkillCategory, error) {
	return svc.skillRepository.getSkillCategoryById(id)
}
func (svc *SkillService) DeleteSkillCategoryById(id uint) error {
	return svc.skillRepository.deleteSkillCategoryById(id)
}
func (svc *SkillService) UpdateSkillCategory(skillCategoryObj SkillCategory) error {
	return svc.skillRepository.updateSkillCategory(skillCategoryObj)
}
func (svc *SkillService) FetchAllSkillCategories(limit, offset int, orderBy string) ([]SkillCategory, int64, error) {
	return svc.skillRepository.fetchAllSkillCategories(limit, offset, orderBy)
}
func (svc *SkillService) UpdateSkill(skillObj Skill) error {
	return svc.skillRepository.updateSkill(skillObj)
}
func (svc *SkillService) GetSkillById(id uint) (Skill, error) {
	return svc.skillRepository.getSkillById(id)
}
func (svc *SkillService) DeleteSkillById(id uint) error {
	return svc.skillRepository.deleteSkillById(id)
}
func (svc *SkillService) FetchAllSkill(limit, offset int, orderBy, keyword string) ([]Skill, int64, error) {
	return svc.skillRepository.fetchAllSkill(limit, offset, orderBy, keyword)
}
