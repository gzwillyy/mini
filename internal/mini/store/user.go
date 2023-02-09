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

package store

import (
	"context"

	"github.com/gzwillyy/mini/internal/pkg/model"
	"gorm.io/gorm"
)

// UserStore 定义了 user 模块在 store 层所实现的方法.
type UserStore interface {
	Create(ctx context.Context, user *model.UserM) error
}

// UserStore 接口的实现.
type users struct {
	db *gorm.DB
}

func newUsers(db *gorm.DB) *users {
	return &users{db}
}

// 确保 users 实现了 UserStore 接口.
var _ UserStore = (*users)(nil)

// Create users具体实现 UserStore 接口的 Create 方法
func (u users) Create(ctx context.Context, user *model.UserM) error {
	return u.db.Create(&user).Error
}
