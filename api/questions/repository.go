package questions

// QuestionRepository Used to store and retrieve questions
type QuestionRepository interface {
	GetQuestionById(id uint) (*Question, error)
	UpdateQuestionByID(question *Question) error
	GetAllQuestions(limitInt, offsetInt int, orderBy string) ([]Question, int64, error)
	DeleteQuestionByID(questionId uint) error
	AddQuestion(question AddQuestionRequestBulk) error
	//createCategories(jsonData []Category) error
}
