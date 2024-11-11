package experience

// ExperienceRepository Used to store and retrieve user experince
type ExperienceRepository interface {
	AddExperienceWithUserAndSkills(userID, skillId uint, experience *Experience) (*Experience, error)
	GetExperienceById(id uint) (*Experience, error)
	GetUserExperienceByUserIdAndExperienceId(userId, experienceId uint) (*UserExperience, error)
	UpdateExperience(experience *Experience) error
	DeleteUserExperienceByUserID(userId, id uint) error
	GetAllUserExperience(userId, id uint, limit int, offset int, orderBy string) ([]Experience, uint, error)
	//createCategories(jsonData []Category) error
}
