package api

// The response that is sent to the client when a user is created.
type createUserResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// The response that is sent to the client when user information is requested.
type getUserResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type loginUserResponse struct {
	Token string `json:"token"`
}

type createExpenseResponse struct {
	Amount int64 `json:"amount"`
}
