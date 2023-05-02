package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID              uuid.UUID     `json:"id" bson:"id" gorm:"primaryKey"`
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
	Order_Status    []Order       `json:"orders" bson:"orders" gorm:"foreignKey:Order_ID "`
}

type Address struct {
	Address_ID string `bson:"id" gorm:"primaryKey"`
	House      string `json:"house" bson:"house"`
	Street     string `json:"street" bson:"street"`
	City       string `json:"city" bson:"city"`
	Zipcode    string `json:"zipcode" bson:"zipcode"`
}

type ProductUser struct {
	Product_ID   string  `bson:"product_id" gorm:"foreignKey:Product_ID"`
	Product_Name string  `json:"product_name" bson:"product_name"`
	Price        float64 `json:"price" bson:"price"`
	Rating       int     `json:"rating" bson:"rating"`
	Image        string  `json:"image" bson:"image"`
}

type Product struct {
	Product_ID   string  `bson:"product_id" gorm:"primaryKey"`
	Product_Name string  `json:"product_name" `
	Price        float64 `json:"price"`
	Rating       int     `json:"rating"`
	Image        string  `json:"image"`
}

type Order struct {
	Order_ID       string        `bson:"order_id" gorm:"primaryKey"`
	Order_Cart     []ProductUser `json:"order_cart"  bson:"order_cart" gorm:"foreignKey:Product_ID" `
	Order_At       time.Time     `json:"order_at" bson:"order_at" `
	Price          float64       `json:"price" bson:"price" `
	Discount       float64       `json:"discount" bson:"discount" `
	Payment_Method Payment       `json:"payment_method" bson:"payment_method" gorm:"foreignKey:Payment_ID"`
}

type Payment struct {
	Payment_ID string `gorm:"primaryKey"`
	Digital    bool
	Cod        bool
}
