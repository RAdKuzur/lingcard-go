package reaction

type StatusReactionDictionary struct{}

const (
	Like    = 1
	Dislike = 2
)

func (r StatusReactionDictionary) GetList() map[int]string {
	return map[int]string{
		Like:    "Лайк",
		Dislike: "Дизлайк",
	}
}

func (r StatusReactionDictionary) Get(index int) (string, bool) {
	name, exists := r.GetList()[index]
	return name, exists
}
