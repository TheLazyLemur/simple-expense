package api

type createUserResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type getUserResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
