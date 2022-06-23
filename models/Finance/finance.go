package finance

type Finance struct {
	Finance_id int64  `json:"finance_id"`
	Name       string `json:"name"`
	Count      int    `json:"count"`
	Sub        string `json:"sub"`
	User_id    int    `json:"user_id"`
}
