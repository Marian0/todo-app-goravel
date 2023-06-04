package models

import (
	"github.com/google/uuid"
	"github.com/goravel/framework/database/orm"
	"gorm.io/gorm"
)

type User struct {
	ID uuid.UUID
	orm.Timestamps
	Name     string
	Email    string
	Password string
	orm.SoftDeletes
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// sets uuid
	u.ID = uuid.New()
	return
}
