// Copyright 2023 The Mini Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package user

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"regexp"

	"github.com/jinzhu/copier"

	"github.com/gzwillyy/mini/internal/mini/store"
	"github.com/gzwillyy/mini/internal/pkg/errno"
	"github.com/gzwillyy/mini/internal/pkg/model"
	"github.com/gzwillyy/mini/pkg/api/mini/v1"
	"github.com/gzwillyy/mini/pkg/auth"
	"github.com/gzwillyy/mini/pkg/token"
)

//将 UserBiz 的实现放在了 internal/mini/biz/user 目录中
//原因是考虑到以后业务逻辑层代码量会比较大
//按 REST 资源保存在不同的目录中，后期阅读和维护都会比较简单

// UserBiz 定义了 user 模块在 biz 层所实现的方法.
type UserBiz interface {
	ChangePassword(ctx context.Context, username string, r *v1.ChangePasswordRequest) error
	Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error)
	Create(ctx context.Context, r *v1.CreateUserRequest) error
	Get(ctx context.Context, username string) (*v1.GetUserResponse, error)
}

// userBiz UserBiz接口的实现.
type userBiz struct {
	ds store.IStore
}

// NewUserBiz New 创建一个实现了 UserBiz 接口的实例.
func NewUserBiz(ds store.IStore) *userBiz {
	return &userBiz{ds: ds}
}

// 确保 userBiz 实现了 UserBiz 接口.
var _ UserBiz = (*userBiz)(nil)

// ChangePassword 是 UserBiz 接口中 `ChangePassword` 方法的实现.
func (b *userBiz) ChangePassword(ctx context.Context, username string, r *v1.ChangePasswordRequest) error {
	userM, err := b.ds.Users().Get(ctx, username)
	if err != nil {
		return err
	}

	if err := auth.Compare(userM.Password, r.OldPassword); err != nil {
		return errno.ErrPasswordIncorrect
	}

	userM.Password, _ = auth.Encrypt(r.NewPassword)
	if err := b.ds.Users().Update(ctx, userM); err != nil {
		return err
	}

	return nil
}

// Login 是 UserBiz 接口中 `Login` 方法的实现.
func (b *userBiz) Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error) {
	// 获取登录用户的所有信息
	user, err := b.ds.Users().Get(ctx, r.Username)
	if err != nil {
		return nil, errno.ErrUserNotFound
	}

	// 对比传入的明文密码和数据库中已加密过的密码是否匹配
	if err := auth.Compare(user.Password, r.Password); err != nil {
		return nil, errno.ErrPasswordIncorrect
	}

	// 如果匹配成功，说明登录成功，签发 token 并返回
	t, err := token.Sign(r.Username)
	if err != nil {
		return nil, errno.ErrSignToken
	}

	return &v1.LoginResponse{Token: t}, nil
}

// Create 是 UserBiz 接口中 `Create` 方法的实现.
func (b *userBiz) Create(ctx context.Context, r *v1.CreateUserRequest) error {
	var userM model.UserM
	_ = copier.Copy(&userM, r)
	if err := b.ds.Users().Create(ctx, &userM); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error()); match {
			return errno.ErrUserAlreadyExist
		}

		return err
	}

	return nil
}

// Get 是 UserBiz 接口中 `Get` 方法的实现.
func (b *userBiz) Get(ctx context.Context, username string) (*v1.GetUserResponse, error) {
	user, err := b.ds.Users().Get(ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrUserNotFound
		}

		return nil, err
	}

	var resp v1.GetUserResponse
	_ = copier.Copy(&resp, user)

	resp.CreatedAt = user.CreatedAt.Format("2006-01-02 15:04:05")
	resp.UpdatedAt = user.UpdatedAt.Format("2006-01-02 15:04:05")

	return &resp, nil
}

//接受来自 Controller 层的入参：context.Context、*v1.CreateUserRequest；
//根据 Store 层 Users().Create() 的入参要求，构建 UserM 结构体；
//调用 Users().Create() 创建用户，并检查返回结果。这里检查了报错是否是用户已存在。如果是，就返回指定的业务错误（ErrUserAlreadyExist 是我们新增的业务错误）。
//为了提高开发效率，简化代码量，我在从 v1.CreateUserRequest 构造 model.UserM 的时候，使用了 github.com/jinzhu/copier 包的 Copy 函数
