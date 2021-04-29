package controllers

import (
	"net/http"
	"project/configs"
	"project/middleware"
	"project/models"

	"github.com/labstack/echo"
	"gorm.io/gorm/clause"
)

func CreateCheckoutController(c echo.Context) error {
	//data user
	userId := middleware.ExtractUserIdFromJWT(c)
	var userDB models.User
	err_userDB := configs.DB.Find(&userDB, userId).Error
	if err_userDB != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err_userDB.Error(),
			Status:  "error",
		})
	}

	//data cart
	var input models.CheckoutRequest
	c.Bind(&input)
	var cartDB models.Cart
	row_cartDB := configs.DB.Find(&cartDB, input.IDCart).RowsAffected
	if row_cartDB == 0 {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: "Invalid id cart",
			Status:  "error",
		})
	}
	err_cartDB := configs.DB.Find(&cartDB, input.IDCart).Error
	if err_cartDB != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err_cartDB.Error(),
			Status:  "error",
		})
	}
	productQuantity := cartDB.Quantity

	//data product
	var productDB models.Product
	err_productDB := configs.DB.Find(&productDB, cartDB.IDProduct).Error
	if err_productDB != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err_productDB.Error(),
			Status:  "error",
		})
	}
	productPrice := productDB.Price
	productId := productDB.ID

	//save data transaction
	var transactionDB models.Transaction
	STATUS := "checkout"
	TOTAL := productPrice * productQuantity
	transactionDB.IDUser = uint(userId)
	transactionDB.Status = STATUS
	transactionDB.Total = TOTAL
	transactionDB.User = userDB
	// transactionDB.Detail = detTrxDB

	if err_trx := configs.DB.Save(&transactionDB).Error; err_trx != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err_trx.Error(),
			Status:  "error",
		})
	}
	//save detail transaksi
	var detailTrxDB models.DetailTransaction
	detailTrxDB.IDTransaction = transactionDB.ID
	detailTrxDB.IDProduct = productId
	detailTrxDB.Quantity = productQuantity
	detailTrxDB.Transactions = transactionDB
	detailTrxDB.Product = productDB
	if err_detailTrxDB := configs.DB.Save(&detailTrxDB).Error; err_detailTrxDB != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err_detailTrxDB.Error(),
			Status:  "error",
		})
	}

	//load data transaksi untuk di tampilkan
	var loadTransactionDB models.Transaction
	if err := configs.DB.Preload("User").Preload("Detail").Find(&loadTransactionDB, transactionDB.ID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error",
		})
	}

	//pengurangan stock
	configs.DB.Find(&productDB, cartDB.IDProduct)
	productDB.Stock = productDB.Stock - cartDB.Quantity
	configs.DB.Save(&productDB)
	//end pengurangan stok

	return c.JSON(http.StatusOK, models.TransactionResponseAny{
		Code:    http.StatusOK,
		Message: "Success checkout",
		Status:  "success",
		Data:    loadTransactionDB,
	})
}
func GetCheckoutController(c echo.Context) error {
	userId := middleware.ExtractUserIdFromJWT(c)
	STATUS := "checkout"
	var transactionDB []models.Transaction
	c.Bind(&transactionDB)
	if err := configs.DB.Preload(clause.Associations).Where("id_user = ? AND status = ?", userId, STATUS).Find(&transactionDB).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "success",
		})
	}
	return c.JSON(http.StatusOK, models.TransactionResponseMany{
		Code:    http.StatusOK,
		Message: "List checkout",
		Status:  "success",
		Data:    transactionDB,
	})
}
