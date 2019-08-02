package models

import (
	"github.com/gofrs/uuid"
)

type Author struct {
	Id    uuid.UUID `gorm:"column:id" json:"id" form:"id"`
	Name  string    `gorm:"column:name" json:"name" form:"name"`
	Sex   string    `gorm:"column:sex" json:"sex" form:"sex"`
	Books []Book    `gorm:"ForeignKey:AuthorId;AssociationForeignKey:Refer"`
}
