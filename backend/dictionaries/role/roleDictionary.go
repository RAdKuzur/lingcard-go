package role

type RoleDictionary struct{}

const (
	ROLEUSER  = 1
	ROLEADMIN = 2
)

func (r RoleDictionary) GetList() map[int]string {
	return map[int]string{
		ROLEUSER:  "Пользователь",
		ROLEADMIN: "Администратор",
	}
}

func (r RoleDictionary) Get(index int) string {
	name, _ := r.GetList()[index]
	return name
}
