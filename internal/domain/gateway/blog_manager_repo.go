// Package gateway @Author nono.he 2023/4/20 15:23:00
// 定义核心业务逻辑的接口方法
package gateway

import (
	"context"
	"nonoDemo/internal/domain/model"
)

type IBlogManager interface {
	Load(ctx context.Context, id string) (res model.BlogManager, err error)
	//Save(...) ...
}
