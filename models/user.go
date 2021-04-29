package models

import (
	"time"

	"gorm.io/gorm"
)

// DB Users
type User struct {
	ID        uint           `json:"id", form:"id", gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt", form:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt", form:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt", form:"deletedAt", gorm:"index"`
	Name      string         `json:"name", form:"name"`
	Email     string         `json:"email", form:"email"`
	Password  string         `json:"password", form:"password"`
}

type TokenResponse struct {
	ID    uint   `json:"id", form:"id", gorm:"primarykey"`
	Name  string `json:"name", form:"name"`
	Email string `json:"email", form:"email"`
	Token string `json:"token", form:"token"`
}

type UserRequest struct {
	Name     string `json:"name", form:"name", validate:"required"`
	Email    string `json:"email", form:"email", validate:"required, email"`
	Password string `json:"password", form:"password", validate:"required"`
}

type UserTokenResponseAny struct {
	Code    int           `json:"code", form:"code"`
	Message string        `json:"message", form:"message"`
	Status  string        `json:"status", form:"status"`
	Data    TokenResponse `json:"data", form:"data"`
}

type UserTokenResponseMany struct {
	Code    int             `json:"code", form:"code"`
	Message string          `json:"message", form:"message"`
	Status  string          `json:"status", form:"status"`
	Data    []TokenResponse `json:"data", form:"data"`
}

type UserResponseAny struct {
	Code    int    `json:"code", form:"code"`
	Message string `json:"message", form:"message"`
	Status  string `json:"status", form:"status"`
	Data    User   `json:"data", form:"data"`
}

type UserResponseMany struct {
	Code    int    `json:"code", form:"code"`
	Message string `json:"message", form:"message"`
	Status  string `json:"status", form:"status"`
	Data    []User `json:"data", form:"data"`
}
