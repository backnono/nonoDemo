// Package entity @Author nono.he 2023/4/20 15:40:00
package entity

import "nonoDemo/internal/domain/model"

type BlogManager struct {
	ID      int    `xorm:"id int autoincr pk"`
	BlogID  string `xorm:"blog_id varchar(20) notnull"`
	Name    string `xorm:"name varchar(50) notnull unique"`
	Deleted bool   `xorm:"'deleted' bool notnull default false"`
}

func (b BlogManager) ToModel() model.BlogManager {
	return model.BlogManager{
		ID:      b.ID,
		BlogID:  b.BlogID,
		Name:    b.Name,
		Deleted: false,
	}
}
