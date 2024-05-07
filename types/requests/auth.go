package requests

type LoginRequest struct {
	Username string `json:"username" binding:"required,email,gte=2,lte=40"`
	Password string `json:"password" binding:"required,gt=2,lte=40"`
}
