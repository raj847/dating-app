package request

type User struct {
	Username string `form:"username" validate:"required,min=3,max=11,lowercase"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=8,passwd"`
	Name     string `form:"name" validate:"required"`
	Gender   string `form:"gender" validate:"required,gendr"`
	DOB      string `form:"date_of_birth" validate:"required"`
	Nickname string `form:"nickname"`
	Domicile string `form:"domicile"`
	Photo    string `form:"photo"`
	Job      string `form:"job"`
	Interest string `form:"interest"`
}

type Login struct {
	Chars    string `json:"username/email"`
	Password string `json:"password"`
}
