package requests

type LoginRuquest struct {
	Type     uint8  `binding:"required,number"`
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}
