// Package database @Author nono.he 2023/4/20 15:33:00
// Domain 层 gateway 中接口方法的实现
package database

import (
	"context"

	"nonoDemo/internal/domain/model"
	"nonoDemo/internal/infrastructure/database/entity"
	"nonoDemo/pkg/framework"
	"xorm.io/xorm"
)

type BlogRepository struct {
	db     *xorm.Engine
	logger framework.Logger
}

func NewBlogRepository(db *xorm.Engine, logger framework.Logger) *BlogRepository {
	return &BlogRepository{db: db, logger: logger}
}

func (b *BlogRepository) Load(ctx context.Context, id string) (res model.BlogManager, err error) {
	var blog entity.BlogManager
	query := b.db.Where("code", id).And("deleted = ?", false).OrderBy("code asc")
	_, err = query.Get(&blog)
	if err != nil {
		b.logger.Error("BlogRepository:load", err)
		return model.BlogManager{}, framework.ErrorCommonDB.Wrap(err, "failed to get blog by code")
	}
	return blog.ToModel(), err // 将 DO（数据对象）转成 Domain 层 mode
}

func (b *BlogRepository) FindBlogByOptions(
	options ...entity.Option,
) ([]entity.BlogManager, int64, error) {
	var blogs []entity.BlogManager
	session := b.db.NewSession()
	for _, option := range options {
		session = option(session)
	}
	count, err := session.FindAndCount(&blogs)
	if err != nil {
		return nil, 0, err
	}
	if err != nil {
		return nil, 0, err
	}
	return blogs, count, nil
}
