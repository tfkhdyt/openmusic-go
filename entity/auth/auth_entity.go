package auth

type Auth struct {
	Token string `json:"token" binding:"required"`
}

type Credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type JWTPayload struct {
	UserId string `json:"userId"`
}
