package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"go-quickstart/internal/dao"
	"go-quickstart/internal/model"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type IUser interface {
	Create(ctx context.Context, req *model.UserCreateReq) (*model.User, error)
	GetById(ctx context.Context, id uint) (*model.User, error)
	GetList(ctx context.Context, req *model.UserListReq) ([]*model.User, int, error)
	Update(ctx context.Context, req *model.UserUpdateReq) (*model.User, error)
	Delete(ctx context.Context, id uint) error
}

type userImpl struct{}

var userService = userImpl{}

func User() IUser {
	return &userService
}

// Create creates a new user
func (s *userImpl) Create(ctx context.Context, req *model.UserCreateReq) (*model.User, error) {
	// Check if username already exists
	count, err := dao.User.Ctx(ctx).Where("username", req.Username).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "failed to check username existence")
	}
	if count > 0 {
		return nil, gerror.New("username already exists")
	}

	// Check if email already exists
	count, err = dao.User.Ctx(ctx).Where("email", req.Email).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "failed to check email existence")
	}
	if count > 0 {
		return nil, gerror.New("email already exists")
	}

	// Hash password (in real application, use proper hashing)
	hashedPassword := fmt.Sprintf("%x", md5.Sum([]byte(req.Password)))

	// Create user
	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Status:   1, // Active by default
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	id, err := dao.User.Ctx(ctx).Data(user).InsertAndGetId()
	if err != nil {
		return nil, gerror.Wrap(err, "failed to create user")
	}

	user.Id = gconv.Uint(id)
	return user, nil
}

// GetById retrieves a user by ID
func (s *userImpl) GetById(ctx context.Context, id uint) (*model.User, error) {
	user := &model.User{}
	err := dao.User.Ctx(ctx).Where("id", id).Scan(user)
	if err != nil {
		return nil, gerror.Wrap(err, "failed to get user")
	}
	if user.Id == 0 {
		return nil, gerror.New("user not found")
	}
	return user, nil
}

// GetList retrieves a list of users with pagination
func (s *userImpl) GetList(ctx context.Context, req *model.UserListReq) ([]*model.User, int, error) {
	// Build query
	query := dao.User.Ctx(ctx)
	
	// Add search condition
	if req.Search != "" {
		query = query.Where("username LIKE ? OR email LIKE ?", "%"+req.Search+"%", "%"+req.Search+"%")
	}
	
	// Add status filter
	if req.Status != -1 {
		query = query.Where("status", req.Status)
	}

	// Get total count
	total, err := query.Count()
	if err != nil {
		return nil, 0, gerror.Wrap(err, "failed to get user count")
	}

	// Get paginated results
	var users []*model.User
	err = query.Page(req.Page, req.Size).Order("create_at DESC").Scan(&users)
	if err != nil {
		return nil, 0, gerror.Wrap(err, "failed to get user list")
	}

	return users, total, nil
}

// Update updates a user
func (s *userImpl) Update(ctx context.Context, req *model.UserUpdateReq) (*model.User, error) {
	// Check if user exists
	user, err := s.GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	// Build update data
	updateData := gdb.Map{
		"update_at": time.Now(),
	}

	if req.Username != "" && req.Username != user.Username {
		// Check if new username already exists
		count, err := dao.User.Ctx(ctx).Where("username", req.Username).Where("id !=", req.Id).Count()
		if err != nil {
			return nil, gerror.Wrap(err, "failed to check username existence")
		}
		if count > 0 {
			return nil, gerror.New("username already exists")
		}
		updateData["username"] = req.Username
	}

	if req.Email != "" && req.Email != user.Email {
		// Check if new email already exists
		count, err := dao.User.Ctx(ctx).Where("email", req.Email).Where("id !=", req.Id).Count()
		if err != nil {
			return nil, gerror.Wrap(err, "failed to check email existence")
		}
		if count > 0 {
			return nil, gerror.New("email already exists")
		}
		updateData["email"] = req.Email
	}

	if req.Status != -1 {
		updateData["status"] = req.Status
	}

	// Update user
	_, err = dao.User.Ctx(ctx).Where("id", req.Id).Data(updateData).Update()
	if err != nil {
		return nil, gerror.Wrap(err, "failed to update user")
	}

	// Return updated user
	return s.GetById(ctx, req.Id)
}

// Delete deletes a user
func (s *userImpl) Delete(ctx context.Context, id uint) error {
	// Check if user exists
	_, err := s.GetById(ctx, id)
	if err != nil {
		return err
	}

	// Delete user
	_, err = dao.User.Ctx(ctx).Where("id", id).Delete()
	if err != nil {
		return gerror.Wrap(err, "failed to delete user")
	}

	return nil
}
