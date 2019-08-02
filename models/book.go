package models

import "github.com/gofrs/uuid"

type Book struct {
	Id uuid.UUID `gorm:"column:id" json:"id" form:"id"`

	AuthorId uuid.UUID `json:"author_id" form:"author_id"`
}
