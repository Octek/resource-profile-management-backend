package user

// UserRepository Used to store and retrieve user details
type UserRepository interface {
	createCategories(jsonData []UserCategory) error
}
