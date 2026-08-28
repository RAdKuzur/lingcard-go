package post

type StatusPostDictionary struct{}

const (
	WAITING  = 1
	APPROVED = 2
	REJECTED = 3
)

func (r StatusPostDictionary) GetList() map[int]string {
	return map[int]string{
		WAITING:  "В ожидании рассмотрения",
		APPROVED: "Одобрено к публикации",
		REJECTED: "Отказано в публикации",
	}
}

func (r StatusPostDictionary) Get(index int) (string, bool) {
	name, exists := r.GetList()[index]
	return name, exists
}
