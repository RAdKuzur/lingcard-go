package role

type RoleDictionary struct{}

const (
	RoleUser  = 1
	RoleAdmin = 2
)

func (r RoleDictionary) GetList() map[int]string {
	return map[int]string{
		RoleUser:  "Пользователь",
		RoleAdmin: "Администратор",
	}
}

func (r RoleDictionary) Get(index int) (string, bool) {
	name, exists := r.GetList()[index]
	return name, exists
}
