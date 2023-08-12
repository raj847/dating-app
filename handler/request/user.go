package request

type User struct {
	Username string `json:"username" validate:"required,min=3,max=11,lowercase"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,passwd"`
	Name     string `json:"name" validate:"required"`
	Gender   string `json:"gender" validate:"required,gendr"`
	DOB      string `json:"date_of_birth" validate:"required"`
	Nickname string `json:"nickname"`
	Domicile string `json:"domicile"`
	Photo    string `json:"photo"`
	Job      string `json:"job"`
	Interest string `json:"interest"`
}

type Login struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
