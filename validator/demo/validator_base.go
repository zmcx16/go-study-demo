package demo

type RegisterAxisClut struct {
	Username       string `validate:"gt=0"`
	PasswordNew    string `validate:"gt=0"`
	PasswordRepeat string `validate:"eqfield=PasswordNew"`
	Email          string `validate:"email"`
}
