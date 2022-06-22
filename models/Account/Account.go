package account

type Account struct {
	User_id  int    `json:"user_id"`
	Username string `json:"username,omitempty" binding:"required,min=1,max=50"`
	Email    string `json:"email,omitempty" binding:"required,min=1,max=100"`
	Password string `json:"password" binding:"required,min=8,max=16"`
}

type RegisterResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
