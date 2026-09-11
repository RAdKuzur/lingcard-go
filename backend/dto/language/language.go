package language

type LangDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type LangMapDTO struct {
	Code           string   `json:"code"`
	Label          string   `json:"label"`
	AvailableCodes []string `json:"available_codes"`
}
