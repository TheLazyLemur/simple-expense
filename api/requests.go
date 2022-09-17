package api

// The request that is sent to the server to create a new user.
type createUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

// The request that is sent to the server to login.
type loginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type createExpenseRequest struct {
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
}
