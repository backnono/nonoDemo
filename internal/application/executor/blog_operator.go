// Package executor @Author nono.he 2023/4/20 15:20:00
package executor

import (
	"context"
	"nonoDemo/pkg/framework"

	"nonoDemo/internal/application/dto"
	"nonoDemo/internal/domain/gateway"
)

type BlogOperator struct {
	blogManager gateway.IBlogManager // 字段 type 是接口类型，通过 Infra 层具体实现进行依赖注入
}

func NewBlogOperator(blogManager gateway.IBlogManager) *BlogOperator {
	return &BlogOperator{
		blogManager: blogManager,
	}
}

func (b *BlogOperator) GetBlog(ctx context.Context, blogID string) (dto.Blog, error) {
	blog, err := b.blogManager.Load(ctx, blogID)
	if err != nil {
		return dto.Blog{}, framework.Wrap(err, "获取博客信息失败")
	}

	// TODO
	return dto.BlogFromModel(blog), nil // 通过 DTO 传递数据到外层
}
