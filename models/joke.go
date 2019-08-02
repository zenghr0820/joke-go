package models

import (
	"time"
)

// 糗事 实体
type Joke struct {
	Id         string    `gorm:"column:id; type: varchar(50); primary_key" json:"id" form:"id"`
	Type       string    `gorm:"column:type; type: varchar(50)" json:"type" form:"type"`
	Title      string    `gorm:"column:title" json:"title" form:"title"`
	Content    string    `gorm:"column:content; type: text" json:"content" form:"content"`
	AuthorName string    `gorm:"column:author_name; type: varchar(50)" json:"author_name" form:"author_name"`
	AuthorUrl  string    `gorm:"column:author_url" json:"author_url" form:"author_url"`
	AuthorImg  string    `gorm:"column:author_img" json:"author_img" form:"author_img"`
	ImageUrl   string    `gorm:"column:image_url; type: text" json:"image_url" form:"image_url"`
	VideoUrl   string    `gorm:"column:video_url" json:"video_url" form:"video_url"`
	Vote       string    `gorm:"column:vote; type: varchar(11)" json:"vote" form:"vote"`
	Time       time.Time `gorm:"column:time" json:"time" form:"time"`
}

type Jokes [] Joke
