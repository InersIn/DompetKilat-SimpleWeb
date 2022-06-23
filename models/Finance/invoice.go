package finance

// Models For Conventional Invoice and Productive Invoice
type Invoce struct {
	Name   string `json:"name"`
	Amount int64  `json:"amount"`
	Tenor  int    `json:"tenor"`
	Grade  string `json:"grade"`
	Rate   int    `json:"rate"`
	Type   string `json:"type"`
}

// Models For SBN
type Sbn struct {
	Name    string `json:"name"`
	Amount  int64  `json:"amount"`
	Tenor   int    `json:"tenor"`
	Rate    int    `json:"rate"`
	Type    string `json:"type"`
	User_id string `json:"user_id"`
}
