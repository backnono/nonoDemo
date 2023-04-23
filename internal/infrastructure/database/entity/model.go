// Package entity @Author nono.he 2023/4/23 21:10:00
package entity

import (
	"xorm.io/xorm"

	"nonoDemo/pkg/framework"
)

func Sync(dbEngine *xorm.Engine, logger framework.Logger) {
	beans := []interface{}{
		new(BlogManager),
		// ...
	}
	SyncWithXorm(dbEngine, logger, beans...)
	return
}

func SyncWithXorm(
	dbEngine *xorm.Engine, logger framework.Logger, beans ...interface{},
) {
	// db engine 初始化所有需要的字段的 zh name 映射
	// sturctFieldNameMapZhName := make(map[string]string)
	// dbFieldNameMapStructFieldName := make(map[string]string)

	/*for i := range beans {
		// currMPtr := beans[i]
		TraverseSingleStructPtr(beans[i], &sturctFieldNameMapZhName, &dbFieldNameMapStructFieldName)
	}*/

	err := dbEngine.Sync2(beans...)
	if err != nil {
		logger.Error("sync database", err)
		panic("数据库同步失败！")
	}
	return
}
