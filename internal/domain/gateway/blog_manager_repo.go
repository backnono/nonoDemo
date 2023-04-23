// Package gateway @Author nono.he 2023/4/20 15:23:00
// 定义核心业务逻辑的接口方法
package gateway

import (
	"context"
	"nonoDemo/internal/domain/model"
	"nonoDemo/internal/infrastructure/database/entity"
)

type IBlogManager interface {
	Load(ctx context.Context, id string) (res model.BlogManager, err error)
	//  Save(...) ...

	IBase
}

//IBase 基础查询option
type IBase interface {
	// FindBlogByOptions
	// @Description blog的公共筛选查询方法
	// @Author nono.he 2023-03-13 15:50:54
	// @Return find and count
	FindBlogByOptions(
		options ...entity.Option,
	) ([]entity.BlogManager, int64, error)

	// TODO 还可以提取出公共的get类型的方法 GetCamByOptions
}
