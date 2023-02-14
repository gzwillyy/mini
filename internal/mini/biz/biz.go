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

package biz

import (
	"github.com/gzwillyy/mini/internal/mini/biz/user"
	"github.com/gzwillyy/mini/internal/mini/store"
)

//使用工厂模式开发了用以创建 Biz 层实例的代码

// IBiz 定义了 Biz 层需要实现的方法.

type IBiz interface {
	Users() user.UserBiz
}

// biz 是 IBiz 的一个具体实现.

type biz struct {
	ds store.IStore
}

// 确保 biz 实现了 IBiz 接口.
var _ IBiz = (*biz)(nil)

// NewBiz 创建一个 IBiz 类型的实例. 在创建 *biz 实例的时候，我们传入了其所依赖的 Store 层实例 ds，并用 ds 创建 IBiz 的子类

func NewBiz(ds store.IStore) *biz {
	return &biz{ds: ds}
}

// Users 返回一个实现了 UserBiz 接口的实例.
func (b *biz) Users() user.UserBiz {
	return user.New(b.ds)
}
