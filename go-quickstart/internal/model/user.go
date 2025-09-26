package model

import (
	"time"
)

// User represents a user entity
type User struct {
	Id       uint      `json:"id" orm:"id"       description:"User ID"`
	Username string    `json:"username" orm:"username" description:"Username"`
	Email    string    `json:"email" orm:"email"     description:"Email address"`
	Password string    `json:"-" orm:"password"    description:"Password (hidden in JSON)"`
	Status   int       `json:"status" orm:"status"    description:"User status (1:active, 0:inactive)"`
	CreateAt time.Time `json:"create_at" orm:"create_at" description:"Creation time"`
	UpdateAt time.Time `json:"update_at" orm:"update_at" description:"Update time"`
}

// UserCreateReq represents the request for creating a user
type UserCreateReq struct {
	Username string `json:"username" v:"required|length:3,20#Username is required|Username length should be 3-20"`
	Email    string `json:"email" v:"required|email#Email is required|Email format is invalid"`
	Password string `json:"password" v:"required|length:6,20#Password is required|Password length should be 6-20"`
}

// UserUpdateReq represents the request for updating a user
type UserUpdateReq struct {
	Id       uint   `json:"id" v:"required#User ID is required"`
	Username string `json:"username" v:"length:3,20#Username length should be 3-20"`
	Email    string `json:"email" v:"email#Email format is invalid"`
	Status   int    `json:"status" v:"in:0,1#Status should be 0 or 1"`
}

// UserGetReq represents the request for getting a user
type UserGetReq struct {
	Id uint `json:"id" v:"required#User ID is required"`
}

// UserDeleteReq represents the request for deleting a user
type UserDeleteReq struct {
	Id uint `json:"id" v:"required#User ID is required"`
}

// UserListReq represents the request for listing users
type UserListReq struct {
	Page   int    `json:"page" d:"1" v:"min:1#Page should be at least 1"`
	Size   int    `json:"size" d:"10" v:"min:1|max:100#Size should be 1-100"`
	Search string `json:"search" d:""`
	Status int    `json:"status" d:"-1" v:"in:-1,0,1#Status should be -1, 0, or 1"`
}

// UserRes represents the response for user data
type UserRes struct {
	Id       uint      `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Status   int       `json:"status"`
	CreateAt time.Time `json:"create_at"`
}

// UserListRes represents the response for user list
type UserListRes struct {
	List  []UserRes `json:"list"`
	Total int       `json:"total"`
	Page  int       `json:"page"`
	Size  int       `json:"size"`
}

// UserDeleteRes represents the response for user deletion
type UserDeleteRes struct {
	Success bool `json:"success"`
}
