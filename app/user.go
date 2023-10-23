package app

type UserLogin struct {
	Email string `json:"email" valid:"email,required"`
	Password string `json:"password" valid:"required, stringlength(6|255)"`
}

type UserRegister struct {
	Username string `json:"username" valid:"required"`
	Email string `json:"email" valid:"email,required"`
	Password string `json:"password" valid:"required, stringlength(6|255)"`
}

type UserUpdate struct {
	Username string `json:"username" valid:"required"`
	Email string `json:"email" valid:"email,required"`
	Password string `json:"password" valid:"required, stringlength(6|255)"`
}