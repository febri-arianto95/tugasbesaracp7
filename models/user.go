package models

import (
	"time"

	"gorm.io/gorm"
)

// DB Users
type User struct {
	ID        uint           `JSON:"id", form:"id", gorm:"primarykey"`
	CreatedAt time.Time      `JSON:"createdAt", form:"createdAt"`
	UpdatedAt time.Time      `JSON:"updatedAt", form:"updatedAt"`
	DeletedAt gorm.DeletedAt `JSON:"deletedAt", form:"deletedAt", gorm:"index"`
	Name      string         `JSON:"name", form:"name"`
	Email     string         `JSON:"email", form:"email"`
	Password  string         `JSON:"password", form:"password"`
}

type TokenResponse struct {
	ID    uint   `json:"id", form:"id", gorm:"primarykey"`
	Name  string `json:"name", form:"name"`
	Email string `json:"email", form:"email"`
	Token string `json:"token", form:"token"`
}

type UserRequestRegister struct {
	Name     string `JSON:"name", form:"name"`
	Email    string `JSON:"email", form:"email"`
	Password string `JSON:"password", form:"password"`
}
type UserRequestLogin struct {
	Email    string `JSON:"email", form:"email"`
	Password string `JSON:"password", form:"password"`
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
