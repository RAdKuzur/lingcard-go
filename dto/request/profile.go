package request

type ProfileUpdateDTO struct {
	BaseLanguageID   int `json:"base_language_id"`
	TargetLanguageID int `json:"target_language_id"`
}
