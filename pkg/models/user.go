package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"unique"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Password  string `json:"password,omitempty"`
	Token     string `json:"token,omitempty"`
}

//type Success struct {
//	ResponseCode      int    `json:"code"`
//	Message           string `json:"message"`
//	Data              byte   `json:"data"`
//	ExternalReference string `json:"ext_ref"`
//}
//
//type Error struct {
//	ResponseCode      int    `json:"code"`
//	Message           string `json:"message"`
//	Detail            string `json:"detail"`
//	ExternalReference string `json:"ext_ref"`
//}
