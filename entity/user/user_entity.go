package user

type User struct {
	ID       string `json:"id"`
	Username string `json:"username" binding:"required" gorm:"uniqueIndex"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
}
