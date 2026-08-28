package word

type WordProgressDTO struct {
	ID            int    `json:"id"`
	Text          string `json:"text"`
	Translation   string `json:"translation"`
	Transcription string `json:"transcription"`
	Level         string `json:"level"`
	RepeatTime    string `json:"repeat_time"`
}
