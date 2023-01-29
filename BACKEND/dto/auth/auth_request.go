package authdto

type RegisterRequest struct {
	Fullname   string `gorm:"type: varchar(255)" json:"fullname" validate:"required"`
	Username   string `json:"username" gorm:"type: varchar(255)" validate:"required"`
	Email      string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password   string `gorm:"type: varchar(255)" json:"password" validate:"required"`
	ListAsRole string `json:"listAsRole" gorm:"type: varchar(255)" validate:"required" `
	Gender     string `json:"gender" gorm:"type: varchar(255)" validate:"required"`
	Phone      string `json:"phone" gorm:"type: varchar(255)" validate:"required"`
	Address    string `json:"address" gorm:"type: text" validate:"required"`
	Image      string `json:"image" gorm:"type: varchar(255)"`
}

type LoginRequest struct {
	Username string `json:"username" gorm:"type: varchar(255)" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}
