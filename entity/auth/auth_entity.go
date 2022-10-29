package auth

type Auth struct {
	Token string `json:"refreshToken" binding:"required"`
}

type Credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
