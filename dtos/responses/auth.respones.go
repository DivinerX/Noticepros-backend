package responses

type Sub struct {
	ID   string `binding:"required" json:"id"`
	Type uint8  `binding:"required" json:"type"`
}
