package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	FirstName   string    `json:"first_name" validate:"required, min=2,max =30"`
	LastName    string    `json:"last_name" validate:"required, min=2,max =30"`
	Password    string    `json:"password" validate:"required, min=6"`
	Email       string    `json:"email" validate:"email, required"`
	Phone       string    `json:"phone" validate:"required"`
	Token       string    `json:"token"`
	RefeshToken string    `json:"refesh_token"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      string    `json:"user_id"`
	//Product        []ProductUser `json:"user_cart" bson:"user_cart" gorm:"foreignKey:ProductUserID"`
	AddressDetails []Address `json:"address"  bson:"address" gorm:"foreignKey:AddressID"`
	OrderStatus    []Order   `json:"orders" bson:"orders" gorm:"foreignKey:OrderID"`
}

type Address struct {
	AddressID string `json:"id" gorm:"primaryKey"`
	House     string `json:"house" bson:"house"`
	Street    string `json:"street" bson:"street"`
	City      string `json:"city" bson:"city"`
	Zipcode   string `json:"zipcode" bson:"zipcode"`
}

type Product struct {
	ProductID   uuid.UUID `json:"product_id" gorm:"primaryKey;default:uuid_generate_v4()"`
	ProductName string    `json:"product_name" `
	Price       int       `json:"price"`
	Rating      int       `json:"rating"`
	Image       string    `json:"image"`
}

type ProductUser struct {
	ProductUserID uuid.UUID `json:"product_user_id" gorm:"primaryKey"`
	ProductName   string    `json:"product_name" bson:"product_name"`
	Price         int       `json:"price" bson:"price"`
	Rating        int       `json:"rating" bson:"rating"`
	Image         string    `json:"image" bson:"image"`
}

type Order struct {
	OrderID string `json:"order_id" gorm:"primaryKey"`
	//OrderCart     []ProductUser `json:"order_cart"  bson:"order_cart" gorm:"foreignKey:ProductUserID" `
	OrderAt       time.Time `json:"order_at" bson:"order_at" `
	Price         float64   `json:"price" bson:"price" `
	Discount      float64   `json:"discount" bson:"discount" `
	PaymentMethod Payment   `json:"payment_method" bson:"payment_method" gorm:"foreignKey:PaymentID"`
}

type Payment struct {
	PaymentID string `gorm:"primaryKey"`
	Digital   bool
	Cod       bool
}

type CartItem struct {
	CartItemID    string `gorm:"primaryKey"`
	ProductUserID string `gorm:"foreignKey:ProductUserID"`
	Quantity      int
}
