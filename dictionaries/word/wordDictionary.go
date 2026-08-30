package word

type StatusWordDictionary struct{}

const (
	NONE     = 1
	LEARNING = 2
	LEARNED  = 3
)

func (r StatusWordDictionary) GetList() map[int]string {
	return map[int]string{
		NONE:     "Не выучено",
		LEARNING: "В процессе обучения",
		LEARNED:  "Изучено",
	}
}

func (r StatusWordDictionary) Get(index int) string {
	name, _ := r.GetList()[index]
	return name
}
