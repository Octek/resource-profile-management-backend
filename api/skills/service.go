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
