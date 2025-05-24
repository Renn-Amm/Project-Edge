package form

type LoginRequest struct {
	Email string `gorm:"uniqure;not null" json:"email" binding:"required,email"`
	Password string `gorm:"not null" json:"-" binding:"required,min=6"`
}