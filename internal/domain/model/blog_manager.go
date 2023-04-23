// Package model @Author nono.he 2023/4/20 15:51:00
package model

type BlogManager struct {
	ID      int    `json:"id"`
	BlogID  string `json:"blog_id"`
	Name    string `json:"name"`
	Deleted bool   `json:"deleted"`
}
