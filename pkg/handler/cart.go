package handler

import (
	"gorm.io/gorm"
)

//var (
//	ErrCantFindProduct    = errors.New("Can't Find The Product ")
//	ErrCantDecodeProducts = errors.New("Can't Find The Product ")
//	ErrUserIdIsNotValid   = errors.New("This User Is Not Valid ")
//	ErrCantUpdateUser     = errors.New("Can't Add This Product To The Cart ")
//	ErrCantRemoveItemCart = errors.New("Can't Remove This Item From The Cart ")
//	ErrCantGetItem        = errors.New("Wasn't Able To Get The Item From The Cart ")
//	ErrCantBuyCartItem    = errors.New("Can't Update The Purchase ")
//)

type CartHandler struct {
	DbConnection *gorm.DB
}

func NewCartHandler(db *gorm.DB) CartHandler {
	return CartHandler{
		DbConnection: db,
	}
}

// AddProductToCart

//func (h *CartHandler) AddProductToCart(productID string, quantity int) error {
//	product := model.Product{}
//	if err := h.DbConnection.First(&product, productID).Error; err != nil {
//		return err
//	}
//	cartItem := model.CartItem{
//		ProductUserID: ,
//		Quantity:      quantity,
//	}
//	// Save cart item to database
//	if err := h.DbConnection.Create(&cartItem).Error; err != nil {
//		return err
//	}
//	return nil
//}

// phai co handleruser ( handleruser ID ) add product ( product ID ) to Cart. at the same time, check product is available,
// add item to handleruser cart

// RemoveCartItem

func RemoveCartItem() {}

// BuyItemFromCart

func BuyItemFromCart() {}

// InstantBuyer

func InstantBuyer() {}
