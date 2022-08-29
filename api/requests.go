package api

type createUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type loginUserRequest struct {
	ID       int32  `json:"id"`
	Password string `json:"password"`
}
