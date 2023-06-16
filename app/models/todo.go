package models

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/goravel/framework/database/orm"
	"gorm.io/gorm"
)

type Todo struct {
	ID uuid.UUID

	Title       string
	CompletedAt sql.NullTime
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
