package models

import (
	"github.com/google/uuid"
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/support/carbon"
	"gorm.io/gorm"
)

type Todo struct {
	ID uuid.UUID

	Title       string
	CompletedAt *carbon.Carbon
	UserID      uuid.UUID

	orm.Timestamps
	orm.SoftDeletes

	// Relations
	User User `gorm:"foreignKey:UserID"`
}

func (t *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	// sets uuid
	t.ID = uuid.New()
	return
}
