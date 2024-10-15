package skills

// SkillRepository Used to store and retrieve skills based on experience and bookings
type SkillRepository interface {
	createCategories(jsonData []SkillCategory) error
}
