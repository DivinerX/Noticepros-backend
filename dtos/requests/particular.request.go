package requests

type ParticularRequest struct {
	RentFrom    string `binding:"required"`
	RentThrough string `binding:"required"`
	Dollars     uint   `binding:"required,gte=0"`
	Cents       uint   `binding:"required,gte=0,lte=99"`
	Written     string `binding:"required"`
	PayToFirst  string `binding:"required"`
	PayToLast   string `binding:"required"`
	Telephone   string `binding:"required"`
	Address     string `binding:"required"`
	City        string `binding:"required"`
	Unit        string
	State       string `binding:"required"`
	ZipCode     string `binding:"required"`
	County      string `binding:"required"`
	OpenHours   string `binding:"required"`
	OpenDays    string `binding:"required"`
	PID         string `binding:"required"`
}
