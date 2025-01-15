package todoList

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"user_name"`
	Password string `json:"password"`
}
