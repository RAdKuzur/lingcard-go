package reaction

type StatusReactionDictionary struct{}

const (
	LIKE    = 1
	DISLIKE = 2
)

func (r StatusReactionDictionary) GetList() map[int]string {
	return map[int]string{
		LIKE:    "Лайк",
		DISLIKE: "Дизлайк",
	}
}

func (r StatusReactionDictionary) Get(index int) string {
	name, _ := r.GetList()[index]
	return name
}
