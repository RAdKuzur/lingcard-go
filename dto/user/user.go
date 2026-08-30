package user

type UserDTO struct {
	Name             string `json:"name"`
	Password         string `json:"password"`
	Email            string `json:"email"`
	Role             int    `json:"role"`
	TargetLanguageID int    `json:"target_language_id"`
	BaseLanguageID   int    `json:"base_language_id"`
	IsBanned         bool   `json:"is_banned"`
}
