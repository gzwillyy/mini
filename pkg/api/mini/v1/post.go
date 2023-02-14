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

package v1

// CreatePostRequest 指定了 `POST /v1/posts` 接口的请求参数.
type CreatePostRequest struct {
	Title   string `json:"title" valid:"required,stringlength(1|256)"`
	Content string `json:"content" valid:"required,stringlength(1|10240)"`
}

// CreatePostResponse 指定了 `POST /v1/posts` 接口的返回参数.
type CreatePostResponse struct {
	PostID string `json:"postID"`
}

// GetPostResponse 指定了 `GET /v1/posts/{postID}` 接口的返回参数.
type GetPostResponse PostInfo

// UpdatePostRequest 指定了 `PUT /v1/posts` 接口的请求参数.
type UpdatePostRequest struct {
	Title   *string `json:"title" valid:"stringlength(1|256)"`
	Content *string `json:"content" valid:"stringlength(1|10240)"`
}

// PostInfo 指定了博客的详细信息.
type PostInfo struct {
	Username  string `json:"username,omitempty"`
	PostID    string `json:"postID,omitempty"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// ListPostRequest 指定了 `GET /v1/posts` 接口的请求参数.
type ListPostRequest struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

// ListPostResponse 指定了 `GET /v1/posts` 接口的返回参数.
type ListPostResponse struct {
	TotalCount int64       `json:"totalCount"`
	Posts      []*PostInfo `json:"posts"`
}
