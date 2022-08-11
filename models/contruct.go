package models

import "time"

type Contruct struct {
	ID       uint `json:"id" gorm:"primaryKey"`
	CreateAt time.Time
	Body     string `json:"body"`
	IDpage   string `json:idpage`
}
