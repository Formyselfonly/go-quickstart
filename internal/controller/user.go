package controller

import (
	"context"
	"go-quickstart/internal/model"
	"go-quickstart/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// UserController handles user-related HTTP requests
type UserController struct{}

// Create creates a new user
func (c *UserController) Create(ctx context.Context, req *model.UserCreateReq) (res *model.UserRes, err error) {
	user, err := service.User().Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return &model.UserRes{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Status:   user.Status,
		CreateAt: user.CreateAt,
	}, nil
}

// Get retrieves a user by ID
func (c *UserController) Get(ctx context.Context, req *model.UserGetReq) (res *model.UserRes, err error) {
	user, err := service.User().GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &model.UserRes{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Status:   user.Status,
		CreateAt: user.CreateAt,
	}, nil
}

// List retrieves a list of users with pagination
func (c *UserController) List(ctx context.Context, req *model.UserListReq) (res *model.UserListRes, err error) {
	users, total, err := service.User().GetList(ctx, req)
	if err != nil {
		return nil, err
	}
	
	var userList []model.UserRes
	for _, user := range users {
		userList = append(userList, model.UserRes{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email,
			Status:   user.Status,
			CreateAt: user.CreateAt,
		})
	}
	
	return &model.UserListRes{
		List:  userList,
		Total: total,
		Page:  req.Page,
		Size:  req.Size,
	}, nil
}

// Update updates a user
func (c *UserController) Update(ctx context.Context, req *model.UserUpdateReq) (res *model.UserRes, err error) {
	user, err := service.User().Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return &model.UserRes{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Status:   user.Status,
		CreateAt: user.CreateAt,
	}, nil
}

// Delete deletes a user
func (c *UserController) Delete(ctx context.Context, req *model.UserDeleteReq) (res *model.UserDeleteRes, err error) {
	err = service.User().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &model.UserDeleteRes{Success: true}, nil
}
