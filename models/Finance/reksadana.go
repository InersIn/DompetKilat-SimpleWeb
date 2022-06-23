package finance

// Models For Reksadana
type Reksadana struct {
	Name    string `json:"name"`
	Amount  int64  `json:"amount"`
	Return  int    `json:"return"`
	User_id int    `json:"user_id"`
}
