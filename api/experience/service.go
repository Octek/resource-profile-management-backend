package experience

type ExperienceService struct {
	experienceRepository ExperienceRepository
}

func NewService(r ExperienceRepository) ExperienceService {
	return ExperienceService{experienceRepository: r}
}

//func (svc *Experience) CreateCategories(jsonData []Category) error {
//	return svc.userRepository.createCategories(jsonData)
//}

func (svc *ExperienceService) AddExperienceWithUserAndSkills(userID, skillId uint, experience *Experience) (*Experience, error) {
	return svc.experienceRepository.AddExperienceWithUserAndSkills(userID, skillId, experience)
}

func (svc *ExperienceService) GetExperienceById(id uint) (*Experience, error) {
	return svc.experienceRepository.GetExperienceById(id)
}

func (svc *ExperienceService) GetUserExperienceByUserIdAndExperienceId(userId, experienceId uint) (*UserExperience, error) {
	return svc.experienceRepository.GetUserExperienceByUserIdAndExperienceId(userId, experienceId)
}

func (svc *ExperienceService) UpdateExperience(experience *Experience) error {
	return svc.experienceRepository.UpdateExperience(experience)
}

func (svc *ExperienceService) DeleteUserExperienceByUserID(userId, id uint) error {
	return svc.experienceRepository.DeleteUserExperienceByUserID(userId, id)
}
func (svc *ExperienceService) GetAllUserExperience(userId, id uint, limit int, offset int, orderBy string) ([]Experience, uint, error) {
	return svc.experienceRepository.GetAllUserExperience(userId, id, limit, offset, orderBy)
}
