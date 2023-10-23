package models

import (
	"time"
)

type User struct {
	ID			uint      	`gorm:"primaryKey;autoIncrement" json:"id"`
	Username 	string 		`gorm:"type:varchar(255);not null" json:"username" `
	Email 		string 		`gorm:"type:varchar(255);not null;unique" json:"email" `
	Password 	string 		`gorm:"type:varchar(255);not null" json:"password" `
	Photos 		[]Photo 	`gorm:"foreignKey:UserID;contraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"photos"`
	CreatedAt 	time.Time 	`gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt 	time.Time 	`gorm:"autoUpdateTime" json:"updated_at"`
}