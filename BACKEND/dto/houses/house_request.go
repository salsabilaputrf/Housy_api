package housesdto

import (
	// "housy/models"

	"gorm.io/datatypes"
)

type HouseRequest struct {
	Name      string         `json:"name" gorm:"type: varchar(255)" validate:"required"`
	CityName  string         `json:"city_name" gorm:"type: varchar(255)" validate:"required"`
	Address   string         `json:"address" gorm:"type: varchar(255)" validate:"required"`
	Price     int            `json:"price" gorm:"type:int" validate:"required"`
	TypeRent  string         `json:"type_rent" gorm:"type: varchar(255)" validate:"required"`
	Amenities datatypes.JSON `json:"amenities" gorm:"type: json" validate:"required"`
	Bedroom   int            `json:"bedroom" gorm:"type: int" validate:"required"`
	Bathroom  int            `json:"bathroom" gorm:"type: int" validate:"required"`
	Image     datatypes.JSON `json:"image" gorm:"type: json" validate:"required"`
}
