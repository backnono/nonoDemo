// Package dto @Author nono.he 2023/4/20 15:27:00
package dto

import "nonoDemo/internal/domain/model"

type Blog struct {
	blogID   string
	BlogName string
	//...
}

func BlogFromModel(input model.BlogManager) (blog Blog) {
	//TODO
	return Blog{
		blogID:   input.BlogID,
		BlogName: input.Name,
	}
}
