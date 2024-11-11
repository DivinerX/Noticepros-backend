package requests

type LoginRuquest struct {
	Type     uint8
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}
