package auth

type RegisterRequest struct {
	Name     string `form:"name" json:"name" binding:"required,min=4,max=100"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=20"`
}
