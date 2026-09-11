package profile

type ProfileDTO struct {
	Username         string `json:"username"`
	Role             string `json:"role"`
	BaseLanguageID   int    `json:"base_language_id"`
	TargetLanguageID int    `json:"target_language_id"`
	NoneWords        int    `json:"none_words"`
	LearningWords    int    `json:"learning_words"`
	LearnedWords     int    `json:"learned_words"`
}
