package questions

// QuestionRepository Used to store and retrieve questions
type QuestionRepository interface {
	AddQuestion(question *Question, name []string) (*Question, error)
	GetQuestionById(id uint) (*Question, error)
	UpdateQuestionByID(question *Question) error
	GetAllQuestions(limitInt, offsetInt int, orderBy string) ([]Question, int64, error)
	DeleteQuestionByID(questionId uint) error
	//createCategories(jsonData []Category) error
}
