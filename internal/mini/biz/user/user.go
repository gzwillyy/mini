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
	"github.com/gzwillyy/mini/pkg/api/mini/v1"
	"regexp"

	"github.com/jinzhu/copier"

	"github.com/gzwillyy/mini/internal/mini/store"
	"github.com/gzwillyy/mini/internal/pkg/errno"
	"github.com/gzwillyy/mini/internal/pkg/model"
)

//将 UserBiz 的实现放在了 internal/mini/biz/user 目录中
//原因是考虑到以后业务逻辑层代码量会比较大
//按 REST 资源保存在不同的目录中，后期阅读和维护都会比较简单

// UserBiz 定义了 user 模块在 biz 层所实现的方法.
type UserBiz interface {
	Create(ctx context.Context, r *v1.CreateUserRequest) error
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

//接受来自 Controller 层的入参：context.Context、*v1.CreateUserRequest；
//根据 Store 层 Users().Create() 的入参要求，构建 UserM 结构体；
//调用 Users().Create() 创建用户，并检查返回结果。这里检查了报错是否是用户已存在。如果是，就返回指定的业务错误（ErrUserAlreadyExist 是我们新增的业务错误）。
//为了提高开发效率，简化代码量，我在从 v1.CreateUserRequest 构造 model.UserM 的时候，使用了 github.com/jinzhu/copier 包的 Copy 函数
