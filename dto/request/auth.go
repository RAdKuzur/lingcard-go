package request

type LoginDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AuthUserDTO struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

type RegisterDTO struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	BaseLangId int    `json:"base_language_id"`
	TargetLang int    `json:"target_language_id"`
	Role       int    `json:"role"`
	IsBanned   bool   `json:"is_banned"`
}
