package level

type LevelDictionary struct{}

const (
	Beginner          = 1
	Elementary        = 2
	Intermediate      = 3
	UpperIntermediate = 4
	Advanced          = 5
	Proficiency       = 6
)

func (r LevelDictionary) GetList() map[int]string {
	return map[int]string{
		Beginner:          "Начальный",
		Elementary:        "Базовый",
		Intermediate:      "Средний",
		UpperIntermediate: "Выше среднего",
		Advanced:          "Продвинутый",
		Proficiency:       "Профессиональный",
	}
}

func (r LevelDictionary) Get(index int) (string, bool) {
	name, exists := r.GetList()[index]
	return name, exists
}
