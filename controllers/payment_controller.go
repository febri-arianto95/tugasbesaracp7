package controllers

import (
	"net/http"
	"project/configs"
	"project/middleware"
	"project/models"

	"github.com/labstack/echo"
	"gorm.io/gorm/clause"
)

func PaymentController(c echo.Context) error {
	userId := middleware.ExtractUserIdFromJWT(c)
	STATUS := "checkout"
	var payment models.PaymentRequest
	c.Bind(&payment)
	var transactionDB models.Transaction
	if row := configs.DB.Where("id_user = ? AND status = ?", userId, STATUS).Find(&transactionDB, payment.IDTransaction).RowsAffected; row == 0 {
		return c.JSON(http.StatusOK, models.ResponseNotif{
			Code:    http.StatusOK,
			Message: "Nothing checkout to paid",
			Status:  "success",
		})
	}
	if err := configs.DB.Preload(clause.Associations).Where("id_user = ? AND status = ?", userId, STATUS).Find(&transactionDB, payment.IDTransaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error",
		})
	}
	transactionDB.Status = "paid"
	if err := configs.DB.Save(&transactionDB).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error",
		})
	}
	return c.JSON(http.StatusOK, models.TransactionResponseAny{
		Code:    http.StatusOK,
		Message: "Success checkout",
		Status:  "success",
		Data:    transactionDB,
	})
}
func GetPaymentController(c echo.Context) error {
	userId := middleware.ExtractUserIdFromJWT(c)
	STATUS := "paid"
	var transactionDB []models.Transaction
	c.Bind(&transactionDB)
	if err := configs.DB.Preload(clause.Associations).Where("id_user = ? AND status = ?", userId, STATUS).Find(&transactionDB).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error",
		})
	}
	return c.JSON(http.StatusOK, models.TransactionResponseMany{
		Code:    http.StatusOK,
		Message: "List paid",
		Status:  "success",
		Data:    transactionDB,
	})
}
