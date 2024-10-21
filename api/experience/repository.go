package experience

// ExperienceRepository Used to store and retrieve user experince
type ExperienceRepository interface {
	AddExperienceWithUserAndSkills(userID, skillId uint, experience Experience) (ExperienceResponse, error)
	GetExperienceById(id uint) (*Experience, error)
	GetUserExperienceByUserIdAndExperienceId(userId, experienceId uint) (*UserExperience, error)
	UpdateExperience(experience *Experience) error
	GetAllUserExperienceList(expID, userID uint) ([]Experience, error)
	DeleteUserExperienceByID(id uint) error
	//createCategories(jsonData []Category) error
}
