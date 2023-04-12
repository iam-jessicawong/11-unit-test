package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	Title       string `gorm:"not null;type:varchar(100)" json:"title" form:"title" valid:"required~Product title is required"`
	Description string `gorm:"not null" json:"description" form:"description" valid:"required~Product description is required"`
	UserID      uint
	User        *User
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
