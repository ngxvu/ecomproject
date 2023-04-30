package model

import "time"

type Address struct {
	Address_ID string `bson:"id" gorm:"primaryKey"`
	House      string `json:"house" bson:"house"`
	Street     string `json:"street" bson:"street"`
	City       string `json:"city" bson:"city"`
	Zipcode    string `json:"zipcode" bson:"zipcode"`
}

type ProductUser struct {
	Product_ID   string  `bson:"product_id" gorm:"primaryKey"`
	Product_Name string  `json:"product_name" bson:"product_name"`
	Price        float64 `json:"price" bson:"price"`
	Rating       int     `json:"rating" bson:"rating"`
	Image        string  `json:"image" bson:"image"`
}

//type Product struct {
//	Product_ID   string  `bson:"product_id" gorm:"primaryKey"`
//	Product_Name string  `json:"product_name" `
//	Price        float64 `json:"price"`
//	Rating       int     `json:"rating"`
//	Image        string  `json:"image"`
//}
//

type Order struct {
	Order_ID       string        `bson:"order_id" gorm:"primaryKey"`
	Order_Cart     []ProductUser `json:"order_cart"  bson:"order_cart" `
	Order_At       time.Time     `json:"order_at" bson:"order_at" `
	Price          float64       `json:"price" bson:"price" `
	Discount       float64       `json:"discount" bson:"discount" `
	Payment_Method Payment       `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool
	Cod     bool
}
