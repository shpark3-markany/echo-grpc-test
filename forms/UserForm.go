package forms

type UserForm struct {
	Email    string `json:"email" gorm:"unique;not null"`
	UserName string `json:"user_name"`
	Password string `json:"password" gorm:"not null"`
	Age      uint64 `json:"age"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

type LoginForm struct {
	Email string `json:"email"`
	Password string `json:"password"`
}