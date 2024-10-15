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
