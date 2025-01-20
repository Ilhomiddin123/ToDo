package todoList

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
