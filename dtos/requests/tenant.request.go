package requests

type TenantRequest []struct {
	FirstName     string `binding:"required"`
	LastName      string `binding:"required"`
	TelePhone     string `binding:"required,phone"`
	TelePhoneCell string `binding:"required,phone"`
	TelePhoneFax  string `binding:"required,phone"`
	Email1        string `binding:"required,email"`
	Email2        string `binding:"omitempty,email"`
	PID           string `binding:"required"`
}
