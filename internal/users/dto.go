package users

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=8,lte=100"`
}

type UserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
