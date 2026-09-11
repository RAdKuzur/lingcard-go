package level

type LevelDictionary struct{}

const (
	BEGINNER          = 1
	ELEMENTARY        = 2
	INTERMEDIATE      = 3
	UPPERINTERMEDIATE = 4
	ADVANCED          = 5
	PROFICIENCY       = 6
)

func (r LevelDictionary) GetList() map[int]string {
	return map[int]string{
		BEGINNER:          "Начальный",
		ELEMENTARY:        "Базовый",
		INTERMEDIATE:      "Средний",
		UPPERINTERMEDIATE: "Выше среднего",
		ADVANCED:          "Продвинутый",
		PROFICIENCY:       "Профессиональный",
	}
}

func (r LevelDictionary) Get(index int) string {
	name, _ := r.GetList()[index]
	return name
}
