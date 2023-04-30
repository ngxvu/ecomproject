package model

import (
	"time"
)

type User struct {
	ID              string        `json:"id" bson:"id" gorm:"primaryKey"`
	First_Name      string        `json:"first_name" validate:"required, min=2,max =30"`
	Last_Name       string        `json:"last_name" validate:"required, min=2,max =30"`
	Password        string        `json:"password" validate:"required, min=6"`
	Email           string        `json:"email" validate:"email, required"`
	Phone           string        `json:"phone" validate:"required"`
	Token           string        `json:"token"`
	Refesh_Token    string        `json:"refesh_token"`
	Created_At      time.Time     `json:"created_at"`
	Updated_At      time.Time     `json:"updated_at"`
	User_ID         string        `json:"user_id"`
	User_Cart       []ProductUser `json:"user_cart" bson:"user_cart"  gorm:"foreignKey:Product_ID"`
	Address_Details []Address     `json:"address"  bson:"address" gorm:"foreignKey:Address_ID"`
	//Order_Status    []Order       `json:"orders" bson:"orders" gorm:"foreignKey:Order_ID "`
}
